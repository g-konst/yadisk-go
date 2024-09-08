package yadisk

import (
	"context"
	"net/http"
)

type PublicService service

func (s *PublicService) Get(
	ctx context.Context, opts PublicGetOpts,
) (*PublicResource, *http.Response, error) {
	url := s.client.BaseURL.JoinPath(s.URL)
	req, err := s.client.NewRequest(ctx, "GET", url, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var result = new(PublicResource)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, nil, err
	}

	return result, resp, nil
}

func (s *PublicService) GetDownloadLink(
	ctx context.Context, opts PublicGetDownloadLinkOpts,
) (*Link, *http.Response, error) {
	url := s.client.BaseURL.JoinPath(s.URL + "/download")
	req, err := s.client.NewRequest(ctx, "GET", url, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var result = new(Link)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, nil, err
	}

	return result, resp, nil
}

func (s *PublicService) SaveToDisk(
	ctx context.Context, opts PublicSaveToDiskOpts,
) (*Link, *http.Response, error) {
	url := s.client.BaseURL.JoinPath(s.URL + "/save-to-disk")
	req, err := s.client.NewRequest(ctx, "POST", url, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var result = new(Link)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, nil, err
	}

	return result, resp, nil
}
