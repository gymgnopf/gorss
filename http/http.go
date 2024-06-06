package http

import (
	"io"
	"net/http"
)

// FetchRSS lädt den RSS-Feed von der angegebenen URL
func FetchRSS(url string) ([]byte, error) {
	// HTTP GET-Anfrage an die URL des RSS-Feeds
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Lesen des HTTP-Antwortkörpers
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
