package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	//url to scrape
	url := "https://techcrunch.com/"

	//get the response from the url
	res, err := http.Get(url)
	checkError(err)
	defer res.Body.Close()

	if res.StatusCode > 400 {
		fmt.Println("Status code: ", res.StatusCode)
	}

	//convert response body to a goquery document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

	//create a file for our posts to populate
	file, err := os.Create("posts.csv")
	checkError(err)

	//create a writer to write our posts to our file
	writer := csv.NewWriter(file)

	//find the articles container and loop through each item
	doc.Find("div.river").Find("div.post-block").Each(func(index int, item *goquery.Selection) {
		//find the title text, the href of the story, and the excerpt of the story
		h2 := item.Find("h2")
		title := strings.TrimSpace(h2.Text())
		url, _ := h2.Find("a").Attr("href")
		excerpt := strings.TrimSpace(item.Find("div.post-block__content").Text())

		//posts slice
		posts := []string{title, url, excerpt}

		//write posts to file
		writer.Write(posts)

	})

	writer.Flush()

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
