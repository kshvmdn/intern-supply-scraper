## intern-supply-scraper

A scraper for [Intern Supply](http://www.intern.supply).

- [Data](#data)
- [Installation](#installation)
- [Usage](#usage)
- [Contribute](#contribute)

### Data

New data is scraped and committed daily. Find the set of scraped data [here](./data).

### Installation

  - __Requirements__:

    + [Python 3+](https://www.python.org/)
    + [Git](https://git-scm.com/)

  - Clone repository
  
    ```
    $ git clone https://github.com/kshvmdn/intern-supply-scraper.git
    $ cd intern-supply-scraper
    ```

  - Install dependencies

    ```
    $ pip install -r ./requirements.txt
    ```

### Usage

  - The barebone scraper will generate a CSV file on every run. You can optionally also generate a JSON file (`--json`) and omit the CSV file (`--no-csv`).

  - You can run a `git diff` on the **two most recent CSV** files by adding `--diff`. This'll show you the changes between the 2 most recent scrapes.

  - Run the scraper with `scraper/__main__.py`.

  ```
  $ ./scraper/__main__.py --help
  usage: ./scraper/__main__.py [--help] [--no-csv] [--json] [--diff] [--verbose]

  optional arguments:
    --help      show this help message and exit
    --no-csv    skip generating a csv file
    --json      generate json document
    --diff      run git diff on the two most recent *csv* files
    --verbose   show detailed output
  ```

### Contribute

The project is completely open source. Feel free to [open a issue](https://github.com/kshvmdn/intern-supply-scraper/issues) or [submit a pull request](https://github.com/kshvmdn/intern-supply-scraper/pulls)! :smile:
