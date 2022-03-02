package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type EksiTopic struct {
	Title string
	URL   string
	Count int
}

func (e EksiTopic) String() string {
	return e.Title
}

type Topics []EksiTopic

func (t Topics) toStringArray() []string {
	var result []string
	for _, topic := range t {
		result = append(result, topic.String())
	}
	return result
}

func GetEksiAgenda() Topics {
	var (
		url        = "https://eksisozluk.com"
		topicClass = ".topic-list"
		result     Topics
	)

	c := colly.NewCollector()

	c.OnHTML(topicClass, func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, h *colly.HTMLElement) {
			topicUrl := h.ChildAttr("a", "href")
			topicUrl = fmt.Sprintf("%s%s", url, topicUrl)
			title, count := getTitleAndNumberFromText(h.ChildText("a"))
			if title != "" {
				result = append(result, EksiTopic{
					Title: title,
					URL:   topicUrl,
					Count: count,
				})
			}
		})

	})

	if err := c.Visit(url); err != nil {
		log.Fatalf("error visiting %s: %v\n", url, err)
	}

	return result
}

// this is the title  5 -> (this is the title, 5)
func getTitleAndNumberFromText(text string) (string, int) {
	a := strings.Split(text, " ")
	if len(a) == 0 {
		return text, 0
	}
	count, _ := strconv.Atoi(a[len(a)-1])
	return strings.Join(a[:len(a)-1], " "), count
}
