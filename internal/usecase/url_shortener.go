package usecase

import (
	"LinkShortener/internal/domain"
	"context"
	"crypto/rand"
	"net/url"
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

// validateURL checks if the provided URL is valid.
func (u *UrlShortener) validateURL(originalURL string) error {
	// Parse the URL to check its validity
	if originalURL == "" {
		return domain.ErrEmptyURL
	}

	parsedURL, err := url.Parse(originalURL)
	if err != nil {
		return domain.ErrInvalidURL
	}

	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return domain.ErrInvalidURL
	}

	return nil //all is good
}

func (u *UrlShortener) generateShortCode() string {

	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	//length := 6 // You can adjust the length of the short code as needed

	length := 6

	// Create a byte slice to hold the short code
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		// Генерируем безопасное случайное число
		b := make([]byte, 1)
		_, err := rand.Read(b)
		if err != nil {
			panic(err) // или лог, если хочешь
		}
		// Преобразуем байт в индекс символа
		result[i] = chars[int(b[0])%len(chars)]
	}
	// Кодируем в base64 и обрезаем до нужной длины
	return string(result)
}
