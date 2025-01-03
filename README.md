# URL Builder for Go

This package provides a flexible and simple way to construct URLs in Go. It wraps the standard `net/url` package and supports setting various URL components such as scheme, host, port, path, query parameters, basic authentication, and anchor (fragment).

## Features

- Set URL scheme, host, port, path, and query parameters
- Support for basic authentication
- Include URL anchors (fragments)
- Fluent API with options or setters
- Automatic handling of default ports for `http` (80) and `https` (443)

## Installation

```bash
go get https://github.com/dzahdev/urlbuilder
```

## Usage

### Using Options

```go
package main

import (
	"fmt"
	"github.com/dzahdev/urlbuilder"
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
```

### Using Setters

```go
package main

import (
	"fmt"
	"github.com/dzahdev/urlbuilder"
)

func main() {
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
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Generated URL:", url)
}
```

## Output Examples

### Example 1: Full URL with HTTPS

```plaintext
https://example.com/api/v1/resource?key=value&token=12345#section1
```

### Example 2: Localhost with Custom Port

```plaintext
http://localhost:8080/dashboard?lang=en&mode=dark#top
```

## API Reference

### Constructor

- `NewURLBuilder(options ...Option) *URLBuilder`
    - Creates a new `URLBuilder` instance with optional parameters.

### Options

- `WithScheme(scheme string)`
- `WithHost(host string)`
- `WithPort(port int)`
- `WithPath(path string)`
- `WithQueryParams(params map[string]string)`
- `WithBasicAuth(user, password string)`
- `WithAnchor(anchor string)`

### Setters

- `SetScheme(scheme string)`
- `SetHost(host string)`
- `SetPort(port int)`
- `SetPath(path string)`
- `SetQueryParams(params map[string]string)`
- `SetBasicAuth(user, password string)`
- `SetAnchor(anchor string)`

### Builder Method

- `Build() (string, error)`
    - Constructs the URL based on the provided components.

## License

This package is open source and available under the MIT License.

