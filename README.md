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
  --no-commit     Omit the git pull/add/commit/push skip
  --verbose       Get detailed logging output.
  ```

### Contribute

The project is completely open source. Feel free to [open a issue](https://github.com/kshvmdn/intern-supply-scraper/issues) or [submit a pull request](https://github.com/kshvmdn/intern-supply-scraper/pulls)! :smile:
