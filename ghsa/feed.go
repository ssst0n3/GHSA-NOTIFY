package ghsa

import (
	"encoding/xml"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/feeds"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"io/ioutil"
	"time"
)

func ParseRss(filepath string) (rss feeds.RssFeedXml, err error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	err = xml.Unmarshal(content, &rss)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func CompareFeed(newFeed feeds.Feed, oldFeedFile string) (equal bool, err error) {
	oldRss, _ := ParseRss(oldFeedFile)
	newRss := (&feeds.Rss{Feed: &newFeed}).FeedXml().(*feeds.RssFeedXml)
	oldRss.Channel.PubDate = newRss.Channel.PubDate
	spew.Dump(oldRss)
	spew.Dump(newRss)
	equal = *newRss == oldRss
	return
}

func GenerateNewFeed(securityVulnerabilities []SecurityVulnerability) (feed *feeds.Feed, err error) {
	now := time.Now()
	feed = &feeds.Feed{
		Title:       "Container Software GHSA Feeds",
		Link:        &feeds.Link{Href: "https://raw.githubusercontent.com/ssst0n3/GHSA-NOTIFY/main/output/feed.xml"},
		Description: "",
		Author:      &feeds.Author{Name: "ssst0n3", Email: "ssst0n3@gmail.com"},
		Created:     now,
		//Updated:     now,
	}

	feed.Items = []*feeds.Item{}
	for _, securityVulnerability := range securityVulnerabilities {
		feed.Items = append(feed.Items, &feeds.Item{
			Title:       securityVulnerability.Advisory.Summary,
			Link:        &feeds.Link{Href: securityVulnerability.Advisory.Permalink},
			Description: securityVulnerability.Advisory.Description,
			Created:     securityVulnerability.Advisory.UpdatedAt,
		})
	}
	return
}

func WriteRss(feed *feeds.Feed, filepath string) (err error) {
	rss, err := feed.ToRss()
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	err = ioutil.WriteFile(filepath, []byte(rss), 0644)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}
