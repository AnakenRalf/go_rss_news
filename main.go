package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	feedURLs, err := readFeedURLs("feeds.yaml")
	if err != nil {
		log.Fatalf("Error reading feed URLs: %v", err)
	}
	app := &cli.App{
		Name:  "RSS Feed News cli",
		Usage: "RSSFeedReader.exe [options]",
		Commands: []*cli.Command{
			{
				Name:    "news",
				Aliases: []string{"n"},
				Usage:   "Fetch and print RSS news",
				Action: func(c *cli.Context) error {
					return fetchNews(c, feedURLs)
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "theme",
						Usage: "Select a news theme: AI, ML, Computers, World news",
					},
				},
			},
		},
	}
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
