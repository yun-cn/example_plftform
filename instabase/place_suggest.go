package instabase

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/yanshiyason/noonde_platform/instabase/constants"
)

// PlaceSuggestService provides access to the PlaceSuggest stay endpoint
type PlaceSuggestService service

// PlaceSuggest PlaceSuggest rooms for night stay
func (s *PlaceSuggestService) PlaceSuggest(ctx context.Context, query string) (*PlaceSuggestResponsePayload, *Response, error) {
	req, err := s.NewRequest(query)

	if err != nil {
		return nil, nil, err
	}

	payload := &PlaceSuggestResponsePayload{}
	resp, err := s.client.Do(ctx, req, payload)

	if err != nil {
		return nil, resp, err
	}

	return payload, resp, nil
}

// NewRequest request for PlaceSuggest
func (s *PlaceSuggestService) NewRequest(query string) (*http.Request, error) {
	payload := s.requestPayload(&PlaceSuggestQueryVariables{Keyword: query})
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", graphQLEndpoint, body)
	req.Host = "www.instabase.jp"

	if err != nil {
		return nil, err
	}

	req.Header = DefaultHeaders()

	return req, nil
}

// RequestPayload builds the payload
func (s *PlaceSuggestService) requestPayload(p *PlaceSuggestQueryVariables) *PlaceSuggestRequestPayload {
	return &PlaceSuggestRequestPayload{
		Variables: *p,
		Query:     constants.PlaceSuggestQuery,
	}
}

// PlaceSuggestResponsePayload shape of the PlaceSuggest response
type PlaceSuggestResponsePayload struct {
	Data struct {
		PlaceSearch []struct {
			Typename  string `json:"__typename"`
			Name      string `json:"name"`
			ModelType string `json:"modelType"`
			ModelID   int    `json:"modelId"`
		} `json:"placeSearch"`
	} `json:"data"`
}

// PlaceSuggestRequestPayload payload for PlaceSuggest graphQL endpoint
type PlaceSuggestRequestPayload struct {
	Query     string                     `json:"query"`
	Variables PlaceSuggestQueryVariables `json:"variables"`
}

// PlaceSuggestQueryVariables allowed parameters
type PlaceSuggestQueryVariables struct {
	Keyword string `json:"keyword"`
}
