package main

import (
	"fmt"
	"log"

	colorable "github.com/mattn/go-colorable"
	"github.com/mmcdole/gofeed"
	"github.com/urfave/cli/v2"
)

type RSSFeed struct {
	URL   string
	Theme string
}

func fetchNews(c *cli.Context, feedURLs map[string][]RSSFeed) error {
	theme := c.String("theme")
	feeds, ok := feedURLs[theme]
	if !ok {
		return fmt.Errorf("invalid theme")
	}
	fp := gofeed.NewParser()
	for _, feed := range feeds {
		rssFeed, err := fp.ParseURL(feed.URL)
		if err != nil {
			log.Printf("error fetching RSS feed: %v", err)
			continue
		}
		fmt.Printf("theme: %s\n\n", feed.Theme)
		writer := colorable.NewColorableStdout()
		for _, item := range rssFeed.Items {
			fmt.Fprintf(writer, "%s%s%s\n", "\x1b[31m", "title: "+item.Title, "\x1b[0m")
			// fmt.Println("title:", item.Title)
			fmt.Println("description:", item.Description)
			fmt.Println("link:", item.Link)
			fmt.Println()
		}
	}
	return nil
}
