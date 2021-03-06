package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const url string = "http://www.intern.supply/"

func main() {
	for {
		dateString := getDateString()
		filename := fmt.Sprintf("data/%s.csv", dateString)

		companies := scrape()
		writeToCsv(filename, companies)

		if !isArg("--no-commit") {
			doGit("pull")
			doGit("add")
			doGit("commit")
			doGit("push")
		}

		if isArg("--diff") {
			doGit("diff")
		}

		doLog(fmt.Sprintf("Done, %s.", filename), true)
		time.Sleep(time.Duration(24) * time.Hour)
	}
}

func getDateString() string {
	y, m, d := time.Now().Date()
	return fmt.Sprintf("%04d-%02d-%02d", y, m, d)
}

func writeToCsv(filename string, records [][]string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err.Error())
	}

	// Close f on exit.
	defer func() {
		if err := f.Close(); err != nil {
			panic(err.Error())
		}
	}()

	w := csv.NewWriter(f)

	if err := w.WriteAll(records); err != nil {
		panic(err.Error())
	}

	w.Flush()

	if err := w.Error(); err != nil {
		panic(err.Error())
	}
}

func scrape() [][]string {
	var doc *goquery.Document
	var e error

	if doc, e = goquery.NewDocument(url); e != nil {
		panic(e.Error())
	}

	var companies [][]string
	companies = append(companies, []string{"name", "open", "link"})

	doc.Find("ul#companies li").Each(func(i int, s *goquery.Selection) {
		status := s.Find("a")
		if status.Text() == "" {
			status = s.Find("h1")
		}

		companyName := strings.TrimRight(s.Text(), status.Text())
		applicationHref, _ := status.Attr("href")
		isOpen := status.Text() == "Apply"

		companies = append(companies, []string{
			companyName, strconv.FormatBool(isOpen), applicationHref})
	})

	return companies
}

func doGit(command string) {
	git := "git"
	remote := "origin"
	branch := "master"

	var cmd *exec.Cmd

	switch command {
	case "pull":
		cmd = exec.Command(git, "pull", remote, branch)
	case "add":
		cmd = exec.Command(git, "add", "./data")
	case "commit":
		cmd = exec.Command(git, "commit", "-m", fmt.Sprintf("Data dump, %s", getDateString()))
	case "push":
		cmd = exec.Command(git, "push", remote, branch)
	case "diff":
		files, err := ioutil.ReadDir("./data")
		if err != nil {
			panic(err.Error())
		}

		if len(files) < 2 {
			fmt.Println("Can't diff < 2 files. Try again tomorrow.")
			return
		}

		fileA := fmt.Sprintf("data/%s", files[len(files)-2].Name())
		fileB := fmt.Sprintf("data/%s", files[len(files)-1].Name())

		cmd = exec.Command(git, "diff", "--no-index", "--color", fileA, fileB)
	default:
		cmd = exec.Command(git, "status")
	}

	out, err := cmd.Output()

	if err != nil {
		panic(err.Error())
	}

	output := string(out)

	if output == "" && command == "diff" {
		output = "No changes since last run."
	}

	doLog(output, false)
}

func doLog(message string, important bool) {
	if important || isArg("--verbose") {
		fmt.Println(message)
	}
}

func isArg(arg string) bool {
	for _, current := range os.Args {
		if current == arg {
			return true
		}
	}

	return false
}
