package spacemarket

import (
	"bytes"
	"context"
	"encoding/json"
	"math"
	"net/http"
	"time"

	"github.com/yanshiyason/noonde_platform/spacemarket/constants"
	"github.com/yanshiyason/noonde_platform/spacemarket/types"
)

// ReputationsService provides access to the Reputations endpoint
type ReputationsService service

// ReputationsResponsePayload response payload
type ReputationsResponsePayload struct {
	Data struct {
		Reputations struct {
			Typename string `json:"__typename"`
			PageInfo struct {
				Typename   string `json:"__typename"`
				TotalCount int    `json:"totalCount"`
			} `json:"pageInfo"`
			Results []struct {
				Typename    string `json:"__typename"`
				Description string `json:"description"`
				From        struct {
					Typename     string `json:"__typename"`
					ID           string `json:"id"`
					Name         string `json:"name"`
					ProfileImage string `json:"profileImage"`
					Username     string `json:"username"`
				} `json:"from"`
				Reservation struct {
					Typename       string      `json:"__typename"`
					StartedAt      time.Time   `json:"startedAt"`
					EventTypeText  string      `json:"eventTypeText"`
					NumberOfGuests interface{} `json:"numberOfGuests"`
				} `json:"reservation"`
			} `json:"results"`
		} `json:"reputations"`
	} `json:"data"`
}

// ReputationsParams params
type ReputationsParams struct {
	Page int `json:"page"`
	// TODO: think about how to handle null strings...
	RentType types.RoomRentType `json:"rentType,omitempty"`
	PerPage  int                `json:"perPage"`
	RoomID   string             `json:"roomId"`
}

// NewRequest prepares a request
func (s *ReputationsService) NewRequest(p *ReputationsParams) (*http.Request, error) {
	payloadBytes, err := json.Marshal(s.newRequestPayload(p.toQueryVariable()))
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", graphQLEndpoint, body)

	if err != nil {
		return nil, err
	}

	req.Header = DefaultHeaders()

	return req, nil
}

// List returns the reviews for a room
func (s *ReputationsService) List(ctx context.Context, p *ReputationsParams) (*ReputationsResponsePayload, *Response, error) {
	req, err := s.NewRequest(p)

	if err != nil {
		return nil, nil, err
	}

	payload := &ReputationsResponsePayload{}
	resp, err := s.client.Do(ctx, req, payload)

	if err != nil {
		return nil, resp, err
	}

	return payload, resp, nil
}

// TotalPages returns total amount of pages concidering number of items requested per page
func (r *ReputationsResponsePayload) TotalPages(perPage int) int {
	totalResults := r.Data.Reputations.PageInfo.TotalCount
	return int(math.Ceil(float64(totalResults) / float64(perPage)))
}

// toQueryVariable TODO: remove once figured out how to deal with null values
func (p *ReputationsParams) toQueryVariable() *reputationsQueryVariables {
	// Must transform string into pointer because json marshaling must return null rather than ""
	// TODO, think about solution to json marshaling
	var rt *types.RoomRentType
	if p.RentType == types.RoomRentTypeALL {
		rt = nil
	} else {
		rt = &p.RentType
	}

	return &reputationsQueryVariables{
		PerPage:  p.PerPage,
		Page:     p.Page,
		RentType: rt,
		RoomID:   p.RoomID,
	}
}

// private

// newRequestPayload builds the payload
func (s *ReputationsService) newRequestPayload(v *reputationsQueryVariables) *reputationsRequestPayload {
	return &reputationsRequestPayload{
		Variables: *v,
		Query:     constants.ReputationsQuery,
	}
}

// reputationsRequestPayload payload for searchStayRooms graphQL endpoint
type reputationsRequestPayload struct {
	Variables reputationsQueryVariables `json:"variables"`
	Query     string                    `json:"query"`
}

// reputationsQueryVariables payload for Reputations endpoint
type reputationsQueryVariables struct {
	Page int `json:"page"`
	// TODO: think about how to handle null strings...
	RentType *types.RoomRentType `json:"rentType"`
	PerPage  int                 `json:"perPage"`
	RoomID   string              `json:"roomId"`
}
