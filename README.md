## intern-supply-scraper

__Update (Jan. 11, 2017):__ It looks like Intern Supply isn't being updated as actively as I'd hoped it'd be when developing this. For this reason, I'm going to decrease the frequency of dumps to once/week (instead of daily).

A scraper for [Intern Supply](http://www.intern.supply).

- [Data](#data)
- [Installation](#installation)
- [Usage](#usage)
- [Contribute](#contribute)

### Data

New data is scraped and committed daily. Find the set of scraped data [here](./data).

### Installation

  - __Requirements__:

    + [Go](https://golang.org/)
    + [Git](https://git-scm.com/)

  - Clone repository
  
    ```
    $ git clone https://github.com/kshvmdn/intern-supply-scraper.git
    $ cd intern-supply-scraper
    ```

  - Install dependencies

    ```
    $ go get
    ```

### Usage

  - The barebone scraper will generate a CSV file on every run and commit that once every 24 hours.

  - Run the scraper:

    ```sh
    $ go run scraper.go [OPTIONS]
    ```

  - If you choose, you can build first, then run:

  ```sh
  $ go build scraper.go
  $ ./scraper [OPTIONS]
  ```

  - Command line flags:

  ```
  --no-commit     Omit the git pull/add/commit/push step.
  --verbose       Get detailed logging output.
  --diff          Run git diff on the 2 latest dump files. Requires there to be at least 2 files present.
  ```

### Contribute

The project is completely open source. Feel free to [open a issue](https://github.com/kshvmdn/intern-supply-scraper/issues) or [submit a pull request](https://github.com/kshvmdn/intern-supply-scraper/pulls)! :smile:
