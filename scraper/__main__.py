#!/usr/bin/env python3

import glob
import os
import sys

from collections import OrderedDict
from datetime import datetime

from bs4 import BeautifulSoup

from helpers import diff, fetch, log, write

BASE_URL = 'http://intern.supply/'


def scrape(url):
    html = fetch.fetch(url)
    soup = BeautifulSoup(html, 'html.parser')

    companies = soup.find('ul', id='companies')

    if not companies:
        return None

    data = []

    for li in companies.find_all('li'):
        name = li.find(text=True, recursive=False)

        is_open, link = False, None

        if li.find('a'):
            is_open = True
            link = li.find('a')['href']

        data.append(OrderedDict([
            ('name', name),
            ('open', is_open),
            ('link', link)
        ]))

    return data


def main():
    if '--help' in sys.argv:
        with open('help.txt', 'r') as f:
            print('\n%s' % f.read())
        sys.exit(0)

    companies = scrape(BASE_URL)

    if not companies:
        log.log('Failed to scrape data, try again later.', force=True)
        sys.exit(1)

    if not os.path.exists('data'):
        os.makedirs('data')

    now = datetime.now()

    if '--no-csv' not in sys.argv:
        if write._csv('data/%s.csv' % now, companies):
            log.log('CSV done, %s.' % now)

    if '--json' in sys.argv:
        if write._json('data/%s.json' % now, companies):
            log.log('JSON done, %s.' % now)

    if '--diff' in sys.argv:
        files = sorted(glob.iglob('data/*.csv'), key=os.path.getctime)

        if len(files) < 2:
            log.log('Can\'t diff < 2 files. Try running again.', force=True)
        else:
            if not diff.diff(*files[-2:]):
                log.log('No changes since last run.')

if __name__ == '__main__':
    main()
