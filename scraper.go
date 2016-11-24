package main

import (
	"encoding/csv"
	"fmt"
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

		fmt.Printf("Done, %s.\n", filename)
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
		panic(err)
	}

	// Close f on exit.
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	w := csv.NewWriter(f)

	if err := w.WriteAll(records); err != nil {
		panic(err)
	}

	w.Flush()

	if err := w.Error(); err != nil {
		panic(err)
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

	var cmd exec.Cmd

	switch command {
	case "pull":
		cmd = exec.Command(git, "pull", remote, branch)
	case "add":
		cmd = exec.Command(git, "add", ".")
	case "commit":
		cmd = exec.Command(git, "commit", "-am", getDateString())
	case "push":
		cmd = exec.Command(git, "push", remote, branch)
	default:
		cmd = exec.Command(git, "status")
	}

	out, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	fmt.Println(out)
}
