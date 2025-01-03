package urlbuilder

import (
	"fmt"
	"net/url"
)

type URLBuilder struct {
	scheme   string
	host     string
	port     int
	path     string
	params   url.Values
	user     string
	password string
	anchor   string
}

type Option func(*URLBuilder)

// WithScheme sets the URL scheme
func WithScheme(scheme string) Option {
	return func(b *URLBuilder) {
		b.scheme = scheme
	}
}

// WithHost sets the URL host
func WithHost(host string) Option {
	return func(b *URLBuilder) {
		b.host = host
	}
}

// WithPort sets the URL port
func WithPort(port int) Option {
	return func(b *URLBuilder) {
		b.port = port
	}
}

// WithPath sets the URL path
func WithPath(path string) Option {
	return func(b *URLBuilder) {
		b.path = path
	}
}

// WithQueryParams sets the URL query parameters
func WithQueryParams(params map[string]string) Option {
	return func(b *URLBuilder) {
		for key, value := range params {
			b.params.Add(key, value)
		}
	}
}

// WithBasicAuth sets the basic authentication for the URL
func WithBasicAuth(user, password string) Option {
	return func(b *URLBuilder) {
		b.user = user
		b.password = password
	}
}

// WithAnchor sets the URL anchor (fragment)
func WithAnchor(anchor string) Option {
	return func(b *URLBuilder) {
		b.anchor = anchor
	}
}

// SetScheme sets the URL scheme
func (b *URLBuilder) SetScheme(scheme string) {
	b.scheme = scheme
}

// SetHost sets the URL host
func (b *URLBuilder) SetHost(host string) {
	b.host = host
}

// SetPort sets the URL port
func (b *URLBuilder) SetPort(port int) {
	b.port = port
}

// SetPath sets the URL path
func (b *URLBuilder) SetPath(path string) {
	b.path = path
}

// SetQueryParams sets the URL query parameters
func (b *URLBuilder) SetQueryParams(params map[string]string) {
	for key, value := range params {
		b.params.Add(key, value)
	}
}

// SetBasicAuth sets the basic authentication for the URL
func (b *URLBuilder) SetBasicAuth(user, password string) {
	b.user = user
	b.password = password
}

// SetAnchor sets the URL anchor (fragment)
func (b *URLBuilder) SetAnchor(anchor string) {
	b.anchor = anchor
}

// NewURLBuilder creates a new instance of URLBuilder with options
func NewURLBuilder(options ...Option) *URLBuilder {
	builder := &URLBuilder{
		params: url.Values{},
	}

	for _, opt := range options {
		opt(builder)
	}

	return builder
}

// Build constructs the final URL
func (b *URLBuilder) Build() (string, error) {
	if b.scheme == "" {
		return "", fmt.Errorf("scheme is not set")
	}
	if b.host == "" {
		return "", fmt.Errorf("host is not set")
	}

	baseURL := &url.URL{
		Scheme: b.scheme,
		Host:   b.host,
		Path:   b.path,
	}

	if b.port != 0 && !((b.scheme == "https" && b.port == 443) || (b.scheme == "http" && b.port == 80)) {
		baseURL.Host = fmt.Sprintf("%s:%d", b.host, b.port)
	}

	baseURL.RawQuery = b.params.Encode()

	if b.user != "" && b.password != "" {
		baseURL.User = url.UserPassword(b.user, b.password)
	}

	if b.anchor != "" {
		baseURL.Fragment = b.anchor
	}

	return baseURL.String(), nil
}
