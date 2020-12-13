package item

import "github.com/gorilla/feeds"

type Item struct {
	Title       string
	Link        string
	Description string
}

func (i *Item) FeedItem() *feeds.Item {
	return &feeds.Item{Title: i.Title, Link: &feeds.Link{Href: i.Link}, Description: i.Description}
}
