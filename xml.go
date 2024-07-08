package main

import (
	"encoding/xml"
	"errors"
	"io"
	"net/http"
)

type XML struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}
type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchXML(url string) (XML, error) {
	resp, err := http.Get(url)
	var content XML
	if err != nil {
		return XML{}, errors.New("COULD NOT ACCESS URL")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return XML{}, errors.New("COULD NOT READ CONTENT")
	}
	xml.Unmarshal(body, &content)
	return content, nil
}
