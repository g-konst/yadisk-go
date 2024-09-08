package yadisk

import (
	"context"
	"net/http"
)

type DiskService service

func (s *DiskService) Get(
	ctx context.Context, opts *DiskGetOpts,
) (*Disk, *http.Response, error) {
	url := s.client.BaseURL.JoinPath(s.URL)
	req, err := s.client.NewRequest(ctx, "GET", url, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var result = new(Disk)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, nil, err
	}

	return result, resp, nil
}
