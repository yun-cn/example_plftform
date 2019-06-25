package instabase

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Client represents the http client for instabase
//
// c := client.NewClient()
//
// r, err := c.Search(&client.SearchQueryVariables{
// 	Location: "tokyo",
// 	PerPage:  20,
// 	Page:     1,
// })
//

// DefaultHeaders returns default headers
func DefaultHeaders() http.Header {
	hdrs := http.Header{}
	hdrs.Set("Accept-Encoding", "br, gzip, deflate")
	hdrs.Set("Accept-Language", "en-us")
	hdrs.Set("Accept", "*/*")
	hdrs.Set("app_platform", "IB-iOS")
	hdrs.Set("app_type", "user")
	hdrs.Set("app_version", "2.0.0")
	hdrs.Set("Authorization", "Bearer")
	hdrs.Set("Content-Type", "application/json")
	hdrs.Set("User-Agent", userAgent)
	// Think about auth && log in later!
	hdrs.Set("cookie", "_instabase_session=671451940b41a8d4f7e5ed3bf8ecfe96; _insttid=01dccc6ad6daf62972ae742278048a7abf517f5c")

	return hdrs
}

const (
	graphQLEndpoint = "https://instabase.jp/graphql"
	userAgent       = "Instabase/34 CFNetwork/976 Darwin/18.2.0"
)

// Client manages communication with instabase api
type Client struct {
	client *http.Client

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different endpoints of the instabase API.
	Search       *SearchService
	PlaceSuggest *PlaceSuggestService
	SpaceDetail  *SpaceDetailService
	Reviews      *ReviewsService
}

type service struct {
	client *Client
}

// NewClient returns a new instabase API client. If a nil httpClient is
// provided, http.DefaultClient will be used.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := &Client{client: httpClient}
	c.common.client = c
	c.Search = (*SearchService)(&c.common)
	c.PlaceSuggest = (*PlaceSuggestService)(&c.common)
	c.SpaceDetail = (*SpaceDetailService)(&c.common)
	c.Reviews = (*ReviewsService)(&c.common)

	return c
}

// Response represents a reponse from an detail page
type Response struct {
	*http.Response
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by p, or returned as an
// error if an API error has occurred. If payload implements the io.Writer
// interface, the raw response body will be written to p, without attempting to
// first decode it.
//
// The provided ctx must be non-nil. If it is canceled or times out,
// ctx.Err() will be returned.
//
// adapted from go-github
//
func (c *Client) Do(ctx context.Context, req *http.Request, p interface{}) (*Response, error) {
	req.WithContext(ctx)

	resp, err := c.client.Do(req)

	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}
	defer resp.Body.Close()

	// Response is gziped
	reader, err := gzip.NewReader(resp.Body)

	bodyBytes, _ := ioutil.ReadAll(reader)

	var decodeErr error
	if p != nil {
		if w, ok := p.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decodeErr = json.NewDecoder(bytes.NewBuffer(bodyBytes)).Decode(p)
			if decodeErr != nil {
				fmt.Printf("DECODE ERROR: %s\n", decodeErr)
			}
		}
	}

	//reset the response body to the original unread state
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return &Response{resp}, decodeErr
}
