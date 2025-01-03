package main

import (
	"fmt"
	"github.com/your-repo/urlbuilder"
)

func main() {
	builder := urlbuilder.NewURLBuilder(
		urlbuilder.WithScheme("https"),
		urlbuilder.WithHost("example.com"),
		urlbuilder.WithPort(443),
		urlbuilder.WithPath("/api/v1/resource"),
		urlbuilder.WithQueryParams(map[string]string{
			"key":   "value",
			"token": "12345",
		}),
		urlbuilder.WithBasicAuth("user", "password"),
		urlbuilder.WithAnchor("section1"),
	)

	url, err := builder.Build()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Generated URL:", url)
}
