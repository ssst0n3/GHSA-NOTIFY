package ghsa

import (
	"github.com/gorilla/feeds"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"io/ioutil"
	"time"
)

func UpdateFeed(securityVulnerabilities []SecurityVulnerability) (feed *feeds.Feed, err error) {
	now := time.Now()
	feed = &feeds.Feed{
		Title:       "Container Software GHSA Feeds",
		Link:        &feeds.Link{Href: "https://raw.githubusercontent.com/ssst0n3/GHSA-NOTIFY/main/output/feed.xml"},
		Description: "",
		Author:      &feeds.Author{Name: "ssst0n3", Email: "ssst0n3@gmail.com"},
		Created:     now,
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
