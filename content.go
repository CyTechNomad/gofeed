package main

import (
    "time"
    "net/http"
    "net/url"
    "github.com/go-shiori/go-readability"
    "github.com/SlyMarbo/rss"
)

const (
    defaultRefreshRate = 15 * time.Minute
)

var (
    feeds = map[string]*rss.Feed{
        "https://tailscale.com/blog/index.xml": nil,
        "https://blog.golang.org/feed.atom": nil,
        "https://kubernetes.io/feed.xml": nil,
    }
)

func loadFeeds(subs map[string]*rss.Feed) {
    for sub := range subs {
        feed, err := rss.Fetch(sub)
        if err != nil {
            panic(err)
        }

        feeds[sub] = feed
    }
}

func scrapeTextContent(l string) string {
    
    resp, err := http.Get(l)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    
    u, err := url.Parse(l)

    article, err := readability.FromReader(resp.Body, u)
    if err != nil {
        panic(err)
    }

    return article.TextContent
}

func RefreshFeeds() {
    for _, feed := range feeds {
        err := feed.Update() 
        if err != nil {
            panic(err)
        }
    }
}
