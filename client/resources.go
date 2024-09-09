package yadisk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/context"
)

type ResourcesService service

func (s *ResourcesService) Get(
	ctx context.Context, opts ResourcesGetOpts,
) (*Resource, *http.Response, error) {
	url := s.client.BaseURL.JoinPath(s.URL)
	req, err := s.client.NewRequest(ctx, "GET", url, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var result = new(Resource)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, nil, err
	}

	return result, resp, nil
}

func (s *ResourcesService) Create(
	ctx context.Context, opts ResourcesCreateOpts,
) (*Link, *http.Response, error) {
	url := s.client.BaseURL.JoinPath(s.URL)
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

func (s *ResourcesService) Update(
	ctx context.Context, body *ResourcesUpdateBody, opts ResourcesUpdateOpts,
) (*Resource, *http.Response, error) {
	requestBody, err := json.Marshal(body)
	if err != nil {
		return nil, nil, err
	}

	url := s.client.BaseURL.JoinPath(s.URL)
	req, err := s.client.NewRequest(ctx, "PATCH", url, bytes.NewBuffer(requestBody), opts)
	if err != nil {
		return nil, nil, err
	}

	var result = new(Resource)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, nil, err
	}

	return result, resp, nil
}

func (s *ResourcesService) Delete(
	ctx context.Context, opts ResourcesDeleteOpts,
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

func (s *ResourcesService) Copy(
	ctx context.Context, opts ResourcesCopyOpts,
) (*Link, *http.Response, error) {
	url := s.client.BaseURL.JoinPath(s.URL + "/copy")
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

func (s *ResourcesService) GetDownloadLink(
	ctx context.Context, opts ResourcesGetDownloadLinkOpts,
) (*Link, *http.Response, error) {
	url := s.client.BaseURL.JoinPath(s.URL + "/download")
	req, err := s.client.NewRequest(
		ctx, "GET", url, nil, opts)
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

func (s *ResourcesService) GetFiles(
	ctx context.Context, opts *ResourcesGetFilesOpts,
) (*FilesResourceList, *http.Response, error) {
	url := s.client.BaseURL.JoinPath(s.URL + "/files")
	req, err := s.client.NewRequest(
		ctx, "GET", url, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var result = new(FilesResourceList)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, nil, err
	}

	return result, resp, nil
}

func (s *ResourcesService) GetLastUploaded(
	ctx context.Context, opts *ResourcesGetLastUploadedOpts,
) (*LastUploadedResourceList, *http.Response, error) {
	url := s.client.BaseURL.JoinPath(s.URL + "/last-uploaded")
	req, err := s.client.NewRequest(
		ctx, "GET", url, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var result = new(LastUploadedResourceList)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, nil, err
	}

	return result, resp, nil
}

func (s *ResourcesService) Move(
	ctx context.Context, opts ResourcesMoveOpts,
) (*Link, *http.Response, error) {
	url := s.client.BaseURL.JoinPath(s.URL + "/move")
	req, err := s.client.NewRequest(
		ctx, "POST", url, nil, opts)
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

func (s *ResourcesService) GetPublic(
	ctx context.Context, opts *ResourcesGetPublicOpts,
) (*PublicResourceList, *http.Response, error) {
	url := s.client.BaseURL.JoinPath(s.URL + "/public")
	req, err := s.client.NewRequest(
		ctx, "GET", url, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var result = new(PublicResourceList)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, nil, err
	}

	return result, resp, nil
}

func (s *ResourcesService) Publish(
	ctx context.Context, body *ResourcesPublishBody, opts ResourcesPublishOpts,
) (*Link, *http.Response, error) {
	requestBody, err := json.Marshal(body)
	if err != nil {
		return nil, nil, err
	}

	url := s.client.BaseURL.JoinPath(s.URL + "/publish")
	req, err := s.client.NewRequest(ctx, "PUT", url, bytes.NewBuffer(requestBody), opts)
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

func (s *ResourcesService) Unpublish(
	ctx context.Context, opts ResourcesUnpublishOpts,
) (*Link, *http.Response, error) {
	url := s.client.BaseURL.JoinPath(s.URL + "/unpublish")
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

func (s *ResourcesService) GetUploadLink(
	ctx context.Context, opts ResourcesGetUploadLinkOpts,
) (*ResourceUploadLink, *http.Response, error) {
	url := s.client.BaseURL.JoinPath(s.URL + "/upload")
	req, err := s.client.NewRequest(ctx, "GET", url, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var result = new(ResourceUploadLink)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, nil, err
	}

	return result, resp, nil
}

func (s *ResourcesService) UploadFile(
	ctx context.Context, uploadUrl string, file io.Reader,
) (*Link, *http.Response, error) {
	url, err := url.Parse(uploadUrl)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, "PUT", url, file, nil)
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

func (s *ResourcesService) GetUploadedBytes(
	ctx context.Context, uploadUrl string,
) (int64, error) {
	url, err := url.Parse(uploadUrl)
	if err != nil {
		return 0, err
	}

	req, err := s.client.NewRequest(ctx, "HEAD", url, nil, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return 0, err
	}

	var result int64
	if resp.StatusCode == http.StatusOK {
		contentRange := resp.Header.Get("Range")
		if strings.HasPrefix(contentRange, "bytes ") {
			parts := strings.Split(contentRange, "-")
			if len(parts) == 2 {
				fmt.Sscanf(parts[1], "%d", result)
			}
		}
	}

	return result, nil
}
