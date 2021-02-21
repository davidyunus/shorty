package url

import (
	"context"
	"time"

	"github.com/davidyunus/shorty/internal/data/url"
	"github.com/davidyunus/shorty/pkg/generator"
)

// Service implement url service
type Service struct {
	url url.IStorage
}

// NewService create new url service
func NewService(u url.IStorage) *Service {
	return &Service{
		url: u,
	}
}

// CreateURL create url service
func (s *Service) CreateURL(ctx context.Context, u, short string) (*url.Url, error) {
	stringSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	if short == "" {
		short, _ = generator.RandomStringSet(6, stringSet)
	}
	pURL := &url.Url{
		URL:          u,
		Shortcode:    short,
		StartDate:    time.Now(),
		LastSeenDate: time.Now(),
	}

	err := s.url.Create(ctx, pURL)
	if err != nil {
		return nil, err
	}

	query := `"shortcode" = $1`
	uri, err := s.url.Single(ctx, query, short)
	if err != nil {
		return nil, err
	}
	return uri, nil
}

// GetURL get url service
func (s *Service) GetURL(ctx context.Context, short string) (*url.Url, error) {
	query := `"shortcode" = $1`
	uri, err := s.url.Single(ctx, query, short)
	if err != nil {
		return nil, err
	}

	return uri, err
}

// GetURLandAddCount add redirect url count
func (s *Service) GetURLandAddCount(ctx context.Context, short string) (*url.Url, error) {
	query := `"shortcode" = $1`
	uri, err := s.url.Single(ctx, query, short)
	if err != nil {
		return nil, err
	}
	if uri == nil {
		return nil, err
	}
	uri.RedirectCount++
	uri.LastSeenDate = time.Now()
	err = s.url.Update(ctx, uri)
	if err != nil {
		return nil, err
	}

	return s.GetURL(ctx, short)
}
