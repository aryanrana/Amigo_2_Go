package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func unique(s []string) []string {
	for i := 0; i < len(s); i++ {
		if contains(s[i+1:], s[i]) {
			s = remove(s, i)
			i--
		}
	}
	return s
}

func main() {

	webPage := "https://en.wikipedia.org/wiki/Synthetic_diamond"
	resp, err := http.Get(webPage)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	map1 := map[string][]string{}
	text := doc.Find("p").Text()
	textchange := regexp.MustCompile("<*>")
	result := textchange.ReplaceAllString(text, " ")
	textch := regexp.MustCompile(`\[([^\[\]]*)\]`)
	final := textch.ReplaceAllString(result, "")
	t1 := strings.Replace(final, ",", "", -1)
	t2 := strings.Replace(t1, ";", "", -1)
	t3 := strings.Replace(t2, "(", "", -1)
	t4 := strings.Replace(t3, ")", "", -1)
	t5 := strings.Replace(t4, "'", "", -1)
	t6 := strings.Replace(t5, "\"", "", -1)
	finale := strings.Replace(t6, "^", "", -1)
	sentences := strings.Split(finale, ".")
	for i := 0; i < len(sentences); i++ {
		word := strings.Split(sentences[i], " ")
		for x := 0; x < len(word); x++ {
			if contains(word[x+1:], word[x]) {
				word = remove(word, x)
				x--
			}
		}
		for j := 0; j < len(word); j++ {
			map1[(word[j])] = append(map1[word[j]], sentences[i])
		}
	}
	word := os.Args[1]
	if value, ok := map1[word]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("\n\nNo instances found")
	}
}
