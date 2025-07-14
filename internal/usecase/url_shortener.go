package usecase

import (
	"LinkShortener/internal/domain"
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/url"
	"strings"
	"time"
)

type UrlShortener struct {

	//urlRepo its like a database repository for URLs
	urlRepo domain.URLRepository

	//baseURL is the base URL for the short links
	baseURL string
}

func NewUrlShortener(repository domain.URLRepository, siteBaseURL string) *UrlShortener {
	return &UrlShortener{
		urlRepo: repository,
		baseURL: siteBaseURL,
	}
}

// ShortenURL generates a short code for a given URL and saves it in the repository.
func (u *UrlShortener) ShortenURL(ctx context.Context, originalURL string) (*domain.URL, error) {
	//validate the original URL
	if err := u.validateURL(originalURL); err != nil {
		return nil, err
	}
	//generate a unique short code
	shortCode := u.generateShortCode()

	// Create a new URL object with the original URL and generated short code
	urlObj := &domain.URL{
		Original:  originalURL,
		ShortCode: shortCode,
		CreatedAt: time.Now(),
	}
	// Save the URL object in the repository
	err := u.urlRepo.Save(ctx, urlObj)
	if err != nil {
		return nil, err
	}
	// Return the URL object with the full short URL
	return urlObj, nil
}
