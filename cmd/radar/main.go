package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gorilla/feeds"

	"github.com/egegunes/radar/pkg/tuzbiberstandup"
)

func main() {
	c := colly.NewCollector()

	now := time.Now()
	feed := &feeds.Feed{
		Title:       "radar.egeinternal.xyz",
		Link:        &feeds.Link{Href: "http://radar.egeinternal.xyz"},
		Description: "Radar",
		Author:      &feeds.Author{Name: "Ege Güneş", Email: "egegunes@gmail.com"},
		Created:     now,
	}

	go func() {
		var items []*feeds.Item

		for {
			for _, item := range tuzbiberstandup.Scrape(c) {
				items = append(items, item.FeedItem())
			}

			feed.Items = items

			time.Sleep(5 * time.Minute)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rss, err := feed.ToRss()
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		fmt.Fprintln(w, rss)
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
