package yadisk

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

const API_VERSION = "v1"

type YandexDiskClient struct {
	Token      string
	BaseURL    *url.URL
	HTTPClient *http.Client
	// services
	Disk      *DiskService
	Resources *ResourcesService
	Public    *PublicService
}

type service struct {
	client *YandexDiskClient
	URL    string
}

func NewYandexDiskClient(token string) *YandexDiskClient {
	baseUrl, _ := url.Parse("https://cloud-api.yandex.net")
	baseUrl = baseUrl.JoinPath(API_VERSION)

	c := &YandexDiskClient{
		Token:      token,
		BaseURL:    baseUrl,
		HTTPClient: http.DefaultClient,
	}

	c.Disk = &DiskService{c, "disk"}
	c.Resources = &ResourcesService{c, "disk/resources"}
	c.Public = &PublicService{c, "disk/public/resources"}

	return c
}

func (c *YandexDiskClient) Do(req *http.Request, mod interface{}) (*http.Response, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_, _ = io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}()

	if err = c.Validate(resp); err != nil {
		return resp, err
	}

	if mod != nil {
		if w, ok := mod.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				return nil, err
			}
		} else {
			body, err := io.ReadAll(resp.Body)
			if err != nil && err != io.EOF {
				return nil, err
			}
			if err := json.Unmarshal(body, &mod); err != nil {
				return nil, err
			}
		}
	}

	return resp, err
}

func (c *YandexDiskClient) NewRequest(
	ctx context.Context,
	method string,
	url *url.URL,
	body io.Reader,
	opts interface{},
) (*http.Request, error) {
	v, _ := query.Values(opts)
	url.RawQuery = v.Encode()

	req, err := http.NewRequestWithContext(ctx, method, url.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "OAuth "+c.Token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *YandexDiskClient) Validate(resp *http.Response) error {

	// skip HEAD request
	if resp.Request.Method == http.MethodHead {
		return nil
	}

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return nil
	}

	var errResp = new(DiskError)
	if err := json.NewDecoder(resp.Body).Decode(errResp); err != nil {
		return errResp
	}

	return fmt.Errorf("error: %d %s", resp.StatusCode, errResp.Message)
}
