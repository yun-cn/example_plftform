package spacemarket

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/yanshiyason/noonde_platform/spacemarket/constants"
	"github.com/yanshiyason/noonde_platform/spacemarket/types"
)

// SearchRoomsParams params for request
type SearchRoomsParams = searchRoomsQueryVariables

// SearchDayService provides access to the search endpoint
type SearchDayService service

// Search returns the reviews for a room
func (s *SearchDayService) Search(ctx context.Context, p *SearchRoomsParams) (*SearchRoomsResponsePayload, *Response, error) {
	req, err := s.NewRequest(p)

	if err != nil {
		return nil, nil, err
	}

	payload := &SearchRoomsResponsePayload{}
	resp, err := s.client.Do(ctx, req, payload)

	if err != nil {
		return nil, resp, err
	}

	return payload, resp, nil
}

// NewRequest request for SearchRooms
func (s *SearchDayService) NewRequest(params *SearchRoomsParams) (*http.Request, error) {
	payload := s.requestPayload(params)
	payloadBytes, err := json.Marshal(payload)
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

// newSearchRoomsRequestPayload builds the payload
func (s *SearchDayService) requestPayload(v *searchRoomsQueryVariables) *searchRoomsRequestPayload {
	return &searchRoomsRequestPayload{
		Variables: *v,
		Query:     constants.SearchRoomsQuery,
	}
}

// SearchRoomsResponsePayload shape of the searchRooms response
type SearchRoomsResponsePayload struct {
	Data struct {
		SearchRooms struct {
			Typename string `json:"__typename"`
			PageInfo struct {
				Typename   string `json:"__typename"`
				TotalCount int    `json:"totalCount"`
			} `json:"pageInfo"`
			Results []struct {
				Typename                   string  `json:"__typename"`
				ID                         int     `json:"id"`
				UID                        string  `json:"uid"`
				Name                       string  `json:"name"`
				HasLastMinuteDiscountPlans bool    `json:"hasLastMinuteDiscountPlans"`
				HasDirectReservationPlans  bool    `json:"hasDirectReservationPlans"`
				Capacity                   int     `json:"capacity"`
				TotalReputationCount       int     `json:"totalReputationCount"`
				TotalReputationScore       float64 `json:"totalReputationScore"`
				StateText                  string  `json:"stateText"`
				City                       string  `json:"city"`
				Access                     string  `json:"access"`
				Prices                     []struct {
					Typename    string `json:"__typename"`
					MinText     string `json:"minText"`
					MinUnitText string `json:"minUnitText"`
					MaxText     string `json:"maxText"`
					MaxUnitText string `json:"maxUnitText"`
				} `json:"prices"`
				Thumbnails []struct {
					Typename string `json:"__typename"`
					URL      string `json:"url"`
				} `json:"thumbnails"`
				IsInquiryOnly       bool          `json:"isInquiryOnly"`
				OwnerRank           int           `json:"ownerRank"`
				Latitude            float64       `json:"latitude"`
				Longitude           float64       `json:"longitude"`
				IsFavorite          bool          `json:"isFavorite"`
				SponsoredPromotions []interface{} `json:"sponsoredPromotions"`
				AvailablePlanCount  int           `json:"availablePlanCount"`
				Plans               interface{}   `json:"plans"`
				IsCancelFree        bool          `json:"isCancelFree"`
			} `json:"results"`
		} `json:"searchRooms"`
	} `json:"data"`
}

// searchRoomsRequestPayload payload for searchRooms graphQL endpoint
type searchRoomsRequestPayload struct {
	Variables searchRoomsQueryVariables `json:"variables"`
	Query     string                    `json:"query"`
}

// searchRoomsQueryVariables allowed parameters
type searchRoomsQueryVariables struct {
	Amenities                  interface{}     `json:"amenities"`
	EndedAt                    interface{}     `json:"endedAt"`
	EndedTime                  interface{}     `json:"endedTime"`
	EventType                  types.EventType `json:"eventType"`
	Geocode                    interface{}     `json:"geocode"`
	HasDirectReservationPlans  interface{}     `json:"hasDirectReservationPlans"`
	HasLastMinuteDiscountPlans interface{}     `json:"hasLastMinuteDiscountPlans"`
	HasTodayReservationPlans   interface{}     `json:"hasTodayReservationPlans"`
	Keyword                    interface{}     `json:"keyword"`
	Location                   interface{}     `json:"location"`
	MaxCapacity                interface{}     `json:"maxCapacity"`
	MaxPrice                   interface{}     `json:"maxPrice"`
	MinCapacity                interface{}     `json:"minCapacity"`
	MinPrice                   interface{}     `json:"minPrice"`
	Page                       int             `json:"page"`
	PerPage                    int             `json:"perPage"`
	PriceType                  types.PriceType `json:"priceType"`
	SponsoredPromotionIds      interface{}     `json:"sponsoredPromotionIds"`
	StartedAt                  interface{}     `json:"startedAt"`
	StartedTime                interface{}     `json:"startedTime"`
	State                      interface{}     `json:"state"`
	WithRecommend              interface{}     `json:"withRecommend"`
}
