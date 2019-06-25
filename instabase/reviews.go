package instabase

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/yanshiyason/noonde_platform/instabase/constants"
)

// ReviewsService provides access to the Reviews stay endpoint
type ReviewsService service

// Reviews Reviews rooms for night stay
func (s *ReviewsService) Reviews(ctx context.Context, listingID string, after string) (*ReviewsResponsePayload, *Response, error) {
	req, err := s.NewRequest(listingID, after)

	if err != nil {
		return nil, nil, err
	}

	payload := &ReviewsResponsePayload{}
	resp, err := s.client.Do(ctx, req, payload)

	if err != nil {
		return nil, resp, err
	}

	return payload, resp, nil
}

// NewRequest request for Reviews
func (s *ReviewsService) NewRequest(listingID string, after string) (*http.Request, error) {
	v := &ReviewsQueryVariables{SpaceID: listingID}
	if after != "" {
		v.After = after
	}

	payload := s.requestPayload(v)
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", graphQLEndpoint, body)
	if err != nil {
		return nil, err
	}

	req.Host = "www.instabase.jp"
	req.Header = DefaultHeaders()

	return req, nil
}

// RequestPayload builds the payload
func (s *ReviewsService) requestPayload(p *ReviewsQueryVariables) *ReviewsRequestPayload {
	return &ReviewsRequestPayload{
		Variables: *p,
		Query:     constants.ReviewsQuery,
	}
}

// ReviewsResponsePayload shape of the Reviews response
type ReviewsResponsePayload struct {
	Data struct {
		Reviews struct {
			Typename string `json:"__typename"`
			Edges    []struct {
				Typename string `json:"__typename"`
				Node     struct {
					Typename  string `json:"__typename"`
					ID        string `json:"id"`
					Title     string `json:"title"`
					Comment   string `json:"comment"`
					Point     int    `json:"point"`
					Usage     string `json:"usage"`
					Age       int    `json:"age"`
					Gender    int    `json:"gender"`
					CreatedAt string `json:"createdAt"`
				} `json:"node"`
				Cursor string `json:"cursor"`
			} `json:"edges"`
			PageInfo struct {
				Typename        string `json:"__typename"`
				EndCursor       string `json:"endCursor"`
				HasNextPage     bool   `json:"hasNextPage"`
				HasPreviousPage bool   `json:"hasPreviousPage"`
			} `json:"pageInfo"`
		} `json:"reviews"`
	} `json:"data"`
}

// ReviewsRequestPayload payload for Reviews graphQL endpoint
type ReviewsRequestPayload struct {
	Query     string                `json:"query"`
	Variables ReviewsQueryVariables `json:"variables"`
}

// ReviewsQueryVariables allowed parameters
type ReviewsQueryVariables struct {
	// After is the value of the "cursor" of the last review of the previous request. Or null.
	After   string `json:"after,omitempty"`
	SpaceID string `json:"spaceId"`
}
