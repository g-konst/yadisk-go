package yadisk

import (
	"context"
	"net/http"
)

type TrashService service

func (s *TrashService) Get(
	ctx context.Context, opts TrashGetOpts,
) (*TrashResource, *http.Response, error) {
	url := s.client.BaseURL.JoinPath(s.URL)
	req, err := s.client.NewRequest(ctx, "GET", url, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var result = new(TrashResource)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, nil, err
	}

	return result, resp, nil
}

func (s *TrashService) Clear(
	ctx context.Context, opts *TrashClearOpts,
) (*Link, *http.Response, error) {
	url := s.client.BaseURL.JoinPath(s.URL)
	req, err := s.client.NewRequest(ctx, "DELETE", url, nil, opts)
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

func (s *TrashService) Restore(
	ctx context.Context, opts TrashRestoreOpts,
) (*Link, *http.Response, error) {
	url := s.client.BaseURL.JoinPath(s.URL + "/restore")
	req, err := s.client.NewRequest(ctx, "PUT", url, nil, opts)
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
