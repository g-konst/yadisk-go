package yadisk

import (
	"context"
	"net/http"
)

type OperationsService service

func (s *OperationsService) Get(
	ctx context.Context, opts OperationsGetOpts,
) (*OperationStatus, *http.Response, error) {
	url := s.client.BaseURL.JoinPath(s.URL + opts.OperationId)
	req, err := s.client.NewRequest(ctx, "GET", url, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var result = new(OperationStatus)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, nil, err
	}

	return result, resp, nil
}
