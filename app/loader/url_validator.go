package loader

import (
	"net/url"
)

type URLValidator interface {
	IsValid(string) bool
}

type urlValidator struct{}

func NewURLValidator() URLValidator {
	return &urlValidator{}
}

func (*urlValidator) IsValid(u string) bool {
	parsedURL, err := url.ParseRequestURI(u)
	return err == nil && len(parsedURL.Scheme) != 0 && len(parsedURL.Host) != 0
}
