package urlbuilder_test

import (
	"github.com/dzahdev/urlbuilder"
	"testing"
)

func TestURLBuilder_WithOptions(t *testing.T) {
	builder := urlbuilder.NewURLBuilder(
		urlbuilder.WithScheme("https"),
		urlbuilder.WithHost("example.com"),
		urlbuilder.WithPath("/test"),
		urlbuilder.WithQueryParams(map[string]string{
			"param1": "value1",
			"param2": "value2",
		}),
		urlbuilder.WithAnchor("section1"),
	)

	url, err := builder.Build()
	if err != nil {
		t.Fatalf("Failed to build URL: %v", err)
	}

	expected := "https://example.com/test?param1=value1&param2=value2#section1"
	if url != expected {
		t.Errorf("Expected URL %q, got %q", expected, url)
	}
}

func TestURLBuilder_Setters(t *testing.T) {
	builder := urlbuilder.NewURLBuilder()
	builder.SetScheme("http")
	builder.SetHost("localhost")
	builder.SetPort(8080)
	builder.SetPath("/dashboard")
	builder.SetQueryParams(map[string]string{
		"lang": "en",
		"mode": "dark",
	})
	builder.SetAnchor("top")

	url, err := builder.Build()
	if err != nil {
		t.Fatalf("Failed to build URL: %v", err)
	}

	expected := "http://localhost:8080/dashboard?lang=en&mode=dark#top"
	if url != expected {
		t.Errorf("Expected URL %q, got %q", expected, url)
	}
}

func TestURLBuilder_DefaultPorts(t *testing.T) {
	tests := []struct {
		scheme   string
		host     string
		port     int
		expected string
	}{
		{"https", "secure.com", 443, "https://secure.com"},
		{"http", "example.com", 80, "http://example.com"},
		{"https", "secure.com", 8443, "https://secure.com:8443"},
		{"http", "example.com", 8080, "http://example.com:8080"},
	}

	for _, test := range tests {
		builder := urlbuilder.NewURLBuilder(
			urlbuilder.WithScheme(test.scheme),
			urlbuilder.WithHost(test.host),
			urlbuilder.WithPort(test.port),
		)

		url, err := builder.Build()
		if err != nil {
			t.Fatalf("Failed to build URL: %v", err)
		}

		if url != test.expected {
			t.Errorf("For scheme %q, host %q, port %d: expected %q, got %q", test.scheme, test.host, test.port, test.expected, url)
		}
	}
}
