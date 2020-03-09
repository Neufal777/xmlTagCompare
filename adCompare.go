package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

//Ad type description
type Ad struct {
	URL string `xml:"url"`
}

func (a *Ad) filesProcess(files []string, adTag string) map[string]int {

	/*
		We analyze the files and get the tags
		to process them later and find matches.
	*/

	//Show processing message
	fmt.Println("Processing..")

	totalUrls := map[string]int{}

	var (
		urlprocessed int
		inElement    string
		urls         []string
	)

	for _, file := range files {

		data, _ := ioutil.ReadFile(file)

		xml.Unmarshal(data, &a)

		r := bytes.NewReader(data)

		decoder := xml.NewDecoder(r)

		for {
			t, _ := decoder.Token()
			if t == nil {
				break
			}

			switch se := t.(type) {
			case xml.StartElement:
				inElement = se.Name.Local
				if inElement == adTag {
					decoder.DecodeElement(&a, &se)
					urls = append(urls, a.URL)
					urlprocessed++
				}
			}
		}
	}
	fmt.Println("url processed", urlprocessed)

	//check if the URL's are duplicated
	for _, tagUrl := range urls {

		if totalUrls[tagUrl] > 0 {
			totalUrls[tagUrl]++
		} else {
			totalUrls[tagUrl] = 1
		}
	}

	return totalUrls
}

func (a *Ad) showDuplicates(tagmap map[string]int) {

	var sumTotalDup int
	for tag, times := range tagmap {
		if times > 1 {
			fmt.Println(tag, "Duplicated", times, "Times")
			sumTotalDup += times
		}
	}

	fmt.Println("Total duplicates: ", sumTotalDup)
}

//func (a *Ad) formatUrl(urls []string) []string {}
