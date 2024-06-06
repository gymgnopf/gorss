package rss

import (
	"encoding/xml"
	"unicode/utf8"
)

// Feed represents the structure of an Atom feed
type RSS struct {
	XMLName xml.Name `xml:"feed"`
	ID      string   `xml:"id"`
	Link    Link     `xml:"link"`
	Title   string   `xml:"title"`
	Updated string   `xml:"updated"`
	Entries []Entry  `xml:"entry"`
}

// Link represents the structure of a link in an Atom feed
type Link struct {
	Href string `xml:"href,attr"`
}

// Entry represents the structure of an entry in an Atom feed
type Entry struct {
	Title   string  `xml:"title"`
	Link    Link    `xml:"link"`
	ID      string  `xml:"id"`
	Author  Author  `xml:"author"`
	Summary Summary `xml:"summary"`
}

// Author represents the structure of an author in an Atom feed entry
type Author struct {
	Name string `xml:"name"`
}

// Summary represents the structure of a summary in an Atom feed entry
type Summary struct {
	Type string `xml:"type,attr"`
	Body string `xml:",chardata"`
}

// ParseRSS parst das RSS-Feed aus einem Byte-Slice
func ParseAtom(data []byte) (*RSS, error) {
	var rss RSS
	sanitzedData := removeInvalidUTF8(data)

	err := xml.Unmarshal(sanitzedData, &rss)
	if err != nil {
		return nil, err
	}
	return &rss, nil
}

func removeInvalidUTF8(data []byte) []byte {
	valid := make([]byte, 0, len(data))
	for len(data) > 0 {
		r, size := utf8.DecodeRune(data)
		if r == utf8.RuneError && size == 1 {
			data = data[size:]
			continue
		}
		valid = append(valid, data[:size]...)
		data = data[size:]
	}
	return valid
}
