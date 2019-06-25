package spacemarket

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/yanshiyason/noonde_platform/spacemarket/constants"
	"github.com/yanshiyason/noonde_platform/spacemarket/types"
)

// SearchNightService provides access to the search stay endpoint
type SearchNightService service

// SearchStayRoomsParams params for request
type SearchStayRoomsParams = searchStayRoomsQueryVariables

// Search search rooms for night stay
func (s *SearchNightService) Search(ctx context.Context, params *SearchStayRoomsParams) (*SearchStayRoomsResponsePayload, *Response, error) {
	req, err := s.NewRequest(params)

	if err != nil {
		return nil, nil, err
	}

	payload := &SearchStayRoomsResponsePayload{}
	resp, err := s.client.Do(ctx, req, payload)

	if err != nil {
		return nil, resp, err
	}

	return payload, resp, nil
}

// NewRequest request for SearchStayRooms
func (s *SearchNightService) NewRequest(p *SearchStayRoomsParams) (*http.Request, error) {
	payload := s.requestPayload(p)
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

// RequestPayload builds the payload
func (s *SearchNightService) requestPayload(p *SearchStayRoomsParams) *searchStayRoomsRequestPayload {
	return &searchStayRoomsRequestPayload{
		Variables: *p,
		Query:     constants.SearchStayRoomsQuery,
	}
}

// SearchStayRoomsResponsePayload shape of the searchStayRooms response
type SearchStayRoomsResponsePayload struct {
	Data struct {
		SearchStayRooms struct {
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
				StayCapacity               int     `json:"stayCapacity"`
				TotalReputationCount       int     `json:"totalReputationCount"`
				TotalReputationScore       float32 `json:"totalReputationScore"`
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
		} `json:"searchStayRooms"`
	} `json:"data"`
}

// searchStayRoomsRequestPayload payload for searchStayRooms graphQL endpoint
type searchStayRoomsRequestPayload struct {
	Variables searchStayRoomsQueryVariables `json:"variables"`
	Query     string                        `json:"query"`
}

// searchStayRoomsQueryVariables allowed parameters
type searchStayRoomsQueryVariables struct {
	Amenities                  interface{}     `json:"amenities"`
	EndedAt                    interface{}     `json:"endedAt"`
	EventType                  types.EventType `json:"eventType"`
	Geocode                    interface{}     `json:"geocode"`
	HasDirectReservationPlans  interface{}     `json:"hasDirectReservationPlans"`
	HasLastMinuteDiscountPlans interface{}     `json:"hasLastMinuteDiscountPlans"`
	HasMonthlyDiscountPlans    interface{}     `json:"hasMonthlyDiscountPlans"`
	HasTodayReservationPlans   interface{}     `json:"hasTodayReservationPlans"`
	HasWeeklyDiscountPlans     interface{}     `json:"hasWeeklyDiscountPlans"`
	Keyword                    interface{}     `json:"keyword"`
	Location                   string          `json:"location"`
	MaxPrice                   interface{}     `json:"maxPrice"`
	MinCapacity                interface{}     `json:"minCapacity"`
	MinPrice                   interface{}     `json:"minPrice"`
	Page                       int             `json:"page"`
	PerPage                    int             `json:"perPage"`
	SponsoredPromotionIds      interface{}     `json:"sponsoredPromotionIds"`
	StartedAt                  interface{}     `json:"startedAt"`
	State                      interface{}     `json:"state"`
	StayRoomTypes              interface{}     `json:"stayRoomTypes"`
	WithRecommend              interface{}     `json:"withRecommend"`
}
