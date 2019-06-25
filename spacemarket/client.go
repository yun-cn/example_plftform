package spacemarket

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Client represents the http client for spacemarket
//
// c := client.NewClient()
//
// r, err := c.SearchStayRooms(&client.SearchStayRoomsQueryVariables{
// 	Location: "tokyo",
// 	PerPage:  20,
// 	Page:     1,
// })
//
// r, err := c.SearchRooms(&client.SearchRoomsQueryVariables{
// 	Location:  "tokyo",
// 	PerPage:   20,
// 	Page:      1,
// 	PriceType: types.PTHourly,
// })
//
// r, err := c.RoomInitialize(&client.RoomInitializeQueryVariables{
// 	RoomRentType: client.DAYTIME,
// 	RoomUID:      "Jq1OGKhgsVbjR-B_",
// 	RoomID:       "32887",
// 	CanRentType:  client.STAY,
// })
//
// r, err := c.Reputations(&client.ReputationsQueryVariables{
// 	RoomID:  "32887",
// 	Page:    1,
// 	PerPage: 20,
// })
//

const (
	graphQLEndpoint = "https://v3api.spacemarket.com/graphql"
	userAgent       = "SpaceMarket/0 CFNetwork/976 Darwin/18.2.0"
	apiKey          = "ni3s06TLpB2iC02JC2BGj1rROh8H4Fqz7ufdyCc4"
)

// Client manages communication with spacemarket api
type Client struct {
	client *http.Client

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different endpoints of the spacemarket API.
	Reputations *ReputationsService
	RoomsDay    *RoomsDayService
	RoomsNight  *RoomsNightService
	SearchDay   *SearchDayService
	SearchNight *SearchNightService
}

type service struct {
	client *Client
}

// NewClient returns a new spacemarket API client. If a nil httpClient is
// provided, http.DefaultClient will be used.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := &Client{client: httpClient}
	c.common.client = c
	c.Reputations = (*ReputationsService)(&c.common)
	c.RoomsDay = (*RoomsDayService)(&c.common)
	c.RoomsNight = (*RoomsNightService)(&c.common)
	c.SearchDay = (*SearchDayService)(&c.common)
	c.SearchNight = (*SearchNightService)(&c.common)

	return c
}

// Response represents a reponse from an detail page
type Response struct {
	*http.Response

	// maybe stick this in here somehow
	// TotalPages int
}

// PaginatedResponse represents a reponse from an index page with paginated results
type PaginatedResponse struct {
	*http.Response

	TotalPages int
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
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

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
