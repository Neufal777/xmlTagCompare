package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Ad struct {
	Url string `xml:"url"`
}

func filesProcess(files []string, adTag string) []string {

	/*
		Analizamos los archivos y obtenemos las urls del tag <url>
		para procesarlas mas tarde y encontrar coincidencias.
	*/
	urls := []string{}

	for _, file := range files {

		data, _ := ioutil.ReadFile(file)

		ad := &Ad{}

		xml.Unmarshal(data, &ad)
		r := bytes.NewReader(data)

		var inElement string

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
					decoder.DecodeElement(&ad, &se)
					urls = append(urls, ad.Url)
				}
			default:
			}
		}

	}

	return urls
}

func UrlsCompare(urlSlice []string) {

	/*
		Comparar las urls de los distintos enlaces
		y obtener enlaces duplicados entre los archivos
	*/

	total := map[string]int{}

	for _, urls := range urlSlice {

		if total[urls] > 0 {
			total[urls] += 1

		} else {

			total[urls] = 1
		}
	}

	sumaTotal := 0
	for url, veces := range total {

		if veces > 1 {

			fmt.Println("La url:", url, "esta repetida", veces, "veces")
			sumaTotal += veces

		}
	}

	fmt.Println("Hay ", sumaTotal, "Duplicados en total")
}

func main() {

	files := []string{
		"file1.xml",
		"file2.xml",
		//"file3.xml",
		//"file4.xml",
	}

	UrlsCompare(filesProcess(files, "ad"))
}
