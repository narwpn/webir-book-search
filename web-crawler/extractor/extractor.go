package extractor

import (
	"strings"

	"web-crawler/models"
)

type Extractor interface {
	// IsValidBookPage checks if the html is a valid book page. Valid book pages are pages that contain book information.
	IsValidBookPage(url string, html string) bool

	// Extract extracts the book and author information from the html and returns a BookWithAuthors struct.
	Extract(html string) (*models.BookWithAuthors, error)
}

func GetExtractor(hostUrl string) Extractor {
	if strings.Contains(hostUrl, "naiin.com") {
		return &NaiinExtractor{}
	}

	if strings.Contains(hostUrl, "chulabook.com") {
		return &ChulaExtractor{}
	}

	if strings.Contains(hostUrl, "booktopia.com.au") {
		return &BooktopiaExtractor{}
	}
	return nil
}
