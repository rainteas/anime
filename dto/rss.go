package dto

import "encoding/xml"

// 定义 RSS 和相关结构
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
	Name    string
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Id          string
	Title       string      `xml:"title"`
	Description string      `xml:"description"`
	Link        string      `xml:"link"`
	PubDate     string      `xml:"pubDate"`
	Enclosures  []Enclosure `xml:"enclosure"`
	Download    int
}

type RssMeta struct {
	Id       int
	AnmeName string
	Season   string
	Url      string
}

type Enclosure struct {
	URL    string `xml:"url,attr"`
	Length int    `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}
