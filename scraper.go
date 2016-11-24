package main

import (
	"encoding/csv"
	"fmt"
	"os"
	_ "os/exec"
	"strconv"
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
	companies = append(companies, []string{"Name", "Open", "Link"})

	doc.Find("ul#companies li").Each(func(i int, s *goquery.Selection) {
		companyName := s.Text() // TODO
		applicationHref, isOpen := s.Find("a").Attr("href")

		companies = append(companies, []string{
			companyName, strconv.FormatBool(isOpen), applicationHref})
	})

	return companies
}
