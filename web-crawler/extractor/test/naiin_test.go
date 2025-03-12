package extractor_test

import (
	// "encoding/json"
	// "fmt"
	"io"

	"web-crawler/config"
	"web-crawler/extractor"

	"net/http"
	"net/http/cookiejar"

	// "os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNaiinExtractor_IsValidBookPage(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want bool
	}{
		{
			name: "Valid 1",
			url:  "https://www.naiin.com/product/detail/639369",
			want: true,
		},
		{
			name: "Valid 2",
			url:  "https://www.naiin.com/product/detail/508064",
			want: true,
		},
		{
			name: "Invalid (main page)",
			url:  "https://www.naiin.com/books/",
			want: false,
		},
		{
			name: "Invalid (category page)",
			url:  "https://www.naiin.com/category?category_1_code=28&product_type_id=1",
			want: false,
		},
		{
			name: "Invalid (toy page)",
			url:  "https://www.naiin.com/product/detail/603593",
			want: false,
		},
	}

	jar, _ := cookiejar.New(nil)
	client := http.Client{
		Jar: jar,
	}
	n := extractor.NaiinExtractor{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, tt.url, nil)
			req.Header.Set("User-Agent", config.GetRandomUserAgents())

			resp, err := client.Do(req)
			if err != nil {
				t.Errorf("Failed to get URL: %v", err)
			}

			defer resp.Body.Close()
			html, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("Failed to read body: %v", err)
			}

			// os.WriteFile(fmt.Sprintf("%s.html", tt.name), []byte(html), 0644)

			if got := n.IsValidBookPage(tt.url, string(html)); got != tt.want {
				t.Errorf("NaiinExtractor.IsValidBookPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNaiinExtractor_Extract(t *testing.T) {
	tests := []struct {
		name string
		url  string
	}{
		{
			name: "Book 1",
			url:  "https://www.naiin.com/product/detail/639369",
		},
		{
			name: "Book 2",
			url:  "https://www.naiin.com/product/detail/508064",
		},

		{
			name: "Book 3",
			url:  "https://www.naiin.com/product/detail/485046",
		},
	}

	jar, _ := cookiejar.New(nil)
	client := http.Client{
		Jar: jar,
	}
	n := extractor.NaiinExtractor{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, tt.url, nil)
			req.Header.Set("User-Agent", config.GetRandomUserAgents())

			resp, err := client.Do(req)
			if err != nil {
				t.Errorf("Failed to get URL: %v", err)
			}

			defer resp.Body.Close()
			html, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("Failed to read body: %v", err)
			}

			// os.WriteFile(fmt.Sprintf("%s.html", tt.name), []byte(html), 0644)

			bookWithAuthors, errr := n.Extract(string(html))
			book := bookWithAuthors.Book
			if errr != nil {
				t.Errorf("NaiinExtractor.Extract() error = %v", errr)
			}

			assert.NotEmpty(t, book.URL)
			assert.NotEmpty(t, book.ImageURL)
			assert.NotEmpty(t, book.Title)
			assert.NotEmpty(t, bookWithAuthors.Authors)
			assert.NotEmpty(t, book.ISBN)
			assert.NotEmpty(t, book.Description)

			// jsonData, _ := json.MarshalIndent(map[string]interface{}{
			// 	"product_url": book.ProductURL.String(),
			// 	"image_url":   book.ImageURL.String(),
			// 	"title":       book.Title,
			// 	"authors":     book.Authors,
			// 	"isbn":        book.ISBN,
			// 	"description": book.Description,
			// }, "", "  ")
			// os.WriteFile(fmt.Sprintf("%s.json", tt.name), jsonData, 0644)
		})
	}
}
