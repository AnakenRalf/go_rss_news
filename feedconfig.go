package main

import (
	"io/ioutil"

	"github.com/goccy/go-yaml"
)

type FeedConfig struct {
	Theme string `yaml:"theme"`
	Feeds []struct {
		URL   string `yaml:"url"`
		Title string `yaml:"title"`
	} `yaml:"feeds"`
}

// readFeedURLs reads a YAML file containing FeedConfig structs and returns a map
// of RSSFeeds based on the theme of each FeedConfig. The function takes in a
// string filename representing the name of the YAML file to read and returns a
// map[string][]RSSFeed and an error if encountered.

func readFeedURLs(filename string) (map[string][]RSSFeed, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var config []FeedConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	feedURLs := make(map[string][]RSSFeed)
	for _, feedConfig := range config {
		rssFeeds := make([]RSSFeed, 0)
		for _, feed := range feedConfig.Feeds {
			rssFeeds = append(rssFeeds, RSSFeed{URL: feed.URL, Theme: feedConfig.Theme})
		}
		feedURLs[feedConfig.Theme] = rssFeeds
	}
	return feedURLs, nil
}
