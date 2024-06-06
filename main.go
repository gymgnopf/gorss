package main

import (
	"fmt"
	"gorss/http"
	"gorss/rss"
	"log"
)

func main() {
	// URL des RSS-Feeds
	url := "https://stitcher.io/rss"

	// Abrufen und Parsen des RSS-Feeds
	body, err := http.FetchRSS(url)
	if err != nil {
		log.Fatalf("Fehler beim Abrufen des Feeds: %v", err)
	}

	feed, err := rss.ParseAtom(body)
	if err != nil {
		log.Fatalf("Fehler beim Parsen des Feeds: %v", err)
	}

	// Ausgabe der Titel und Links der Artikel
	for _, item := range feed.Entries {
		fmt.Printf("Titel: %s\nLink: %s\n\n", item.Title, item.Link)
	}
}
