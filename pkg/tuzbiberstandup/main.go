package tuzbiberstandup

import (
	"github.com/gocolly/colly/v2"

	"github.com/egegunes/radar/pkg/item"
)

func Scrape(c *colly.Collector) (items []*item.Item) {
	c.OnHTML("div.tribe-events-calendar-list", func(e *colly.HTMLElement) {
		e.ForEach("div.tribe-events-calendar-list__event-row", func(_ int, c *colly.HTMLElement) {
			date := c.ChildText("span.tribe-event-date-start")
			title := c.ChildText("a.tribe-events-calendar-list__event-title-link")
			link := c.ChildAttr("a.tribe-events-calendar-list__event-title-link", "href")
			venue := c.ChildText("span.tribe-events-calendar-list__event-venue-title")
			items = append(items, &item.Item{Title: title + " " + date + " @" + venue, Link: link})
		})
	})

	c.Visit("https://tuzbiberstandup.com/event-dates/")

	return items
}
