import requests

HEADERS = {
    'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71 Safari/537.36',
    'Referer': 'http://intern.supply'
}


def fetch(url, headers={}):
    return requests.get(url, {**HEADERS, **headers}).text


if __name__ == '__main__':
    x = fetch('http://google.com')
    print(x)
