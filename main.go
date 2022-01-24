package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/gocolly/colly"
)

func main() {

	url := "https://klasika.kompas.id/baca/inspiraksi-kemilau-perayaan-hut-ke-50-kompas/"
	fileName := "data.txt"

	// delete exist file
	os.Remove(fileName)

	os.Create(fileName)

	getData(url, fileName)

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}

	showCountAlphabet(string(data))

}

func cleanData(data string, e *colly.HTMLElement) string {

	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}

	data = reg.ReplaceAllString(e.Text, "")
	data = strings.ToUpper(data)

	return data

}

func getData(url, fileName string) {

	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("p", func(e *colly.HTMLElement) {
		f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		alphaData := cleanData(e.Text, e)

		if _, err = f.WriteString(alphaData); err != nil {
			panic(err)
		}
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		alphaData := cleanData(e.Text, e)

		if _, err = f.WriteString(alphaData); err != nil {
			panic(err)
		}

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(url)
}

func showCountAlphabet(data string) {
	counter := make(map[string]int)
	for _, c := range data {
		if string(c) != " " {
			counter[string(c)]++
		}
	}

	chars := make([]string, 0, len(counter))
	for chr := range counter {
		chars = append(chars, chr)
	}

	sort.Slice(chars, func(i, j int) bool {
		return counter[chars[i]] > counter[chars[j]]
	})

	for _, chr := range chars {
		fmt.Printf("%v: %v\n", chr, counter[chr])
	}

}
