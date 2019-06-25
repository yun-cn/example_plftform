package instabase

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/yanshiyason/noonde_platform/instabase/constants"
)

// SpaceDetailService provides access to the SpaceDetail stay endpoint
type SpaceDetailService service

// SpaceDetail SpaceDetail rooms for night stay
func (s *SpaceDetailService) SpaceDetail(ctx context.Context, listingID string) (*SpaceDetailResponsePayload, *Response, error) {
	req, err := s.NewRequest(listingID)

	if err != nil {
		return nil, nil, err
	}

	payload := &SpaceDetailResponsePayload{}
	resp, err := s.client.Do(ctx, req, payload)

	if err != nil {
		return nil, resp, err
	}

	return payload, resp, nil
}

// NewRequest request for SpaceDetail
func (s *SpaceDetailService) NewRequest(listingID string) (*http.Request, error) {
	payload := s.requestPayload(&SpaceDetailQueryVariables{ID: listingID})
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
func (s *SpaceDetailService) requestPayload(p *SpaceDetailQueryVariables) *SpaceDetailRequestPayload {
	return &SpaceDetailRequestPayload{
		Variables: *p,
		Query:     constants.SpaceDetailQuery,
	}
}

// SpaceDetailResponsePayload shape of the SpaceDetail response
type SpaceDetailResponsePayload struct {
	Data struct {
		Rooms []struct {
			Typename               string `json:"__typename"`
			UID                    string `json:"uid"`
			BookingBeforeLimitDays int    `json:"bookingBeforeLimitDays"`
			MinBookingHours        int    `json:"minBookingHours"`
			SiteURL                string `json:"siteUrl"`
			SpaceCategory          struct {
				Typename string `json:"__typename"`
				Title    string `json:"title"`
			} `json:"spaceCategory"`
			SeoDescription string `json:"seoDescription"`
			Usages         []struct {
				Typename string `json:"__typename"`
				ID       string `json:"id"`
				Title    string `json:"title"`
			} `json:"usages"`
			SummaryBusinessDays []struct {
				Typename      string `json:"__typename"`
				Days          string `json:"days"`
				BusinessHours string `json:"businessHours"`
			} `json:"summaryBusinessDays"`
			SummaryPrice      string `json:"summaryPrice"`
			SummaryPriceTable []struct {
				Typename  string `json:"__typename"`
				Timetable []struct {
					Typename string `json:"__typename"`
					Time     string `json:"time"`
					Price    string `json:"price"`
				} `json:"timetable"`
				Days string `json:"days"`
			} `json:"summaryPriceTable"`
			Attentions []struct {
				Typename    string `json:"__typename"`
				ID          string `json:"id"`
				IsOfficial  bool   `json:"isOfficial"`
				Title       string `json:"title"`
				Priority    int    `json:"priority"`
				Description string `json:"description"`
			} `json:"attentions"`
			Notice   string `json:"notice"`
			Building struct {
				Typename   string  `json:"__typename"`
				Address    string  `json:"address"`
				ID         string  `json:"id"`
				Lat        float64 `json:"lat"`
				Lon        float64 `json:"lon"`
				Prefecture struct {
					Typename string `json:"__typename"`
					ID       string `json:"id"`
				} `json:"prefecture"`
				Title           string `json:"title"`
				SummaryAccesses []struct {
					Typename string `json:"__typename"`
					Access   string `json:"access"`
					Line     string `json:"line"`
					Station  string `json:"station"`
				} `json:"summaryAccesses"`
				ParentArea interface{} `json:"parentArea"`
			} `json:"building"`
			Review struct {
				Typename  string `json:"__typename"`
				ID        string `json:"id"`
				Title     string `json:"title"`
				Comment   string `json:"comment"`
				Point     int    `json:"point"`
				Usage     string `json:"usage"`
				Age       int    `json:"age"`
				Gender    int    `json:"gender"`
				CreatedAt string `json:"createdAt"`
			} `json:"review"`
			CancelPolicies []struct {
				Typename   string `json:"__typename"`
				Title      string `json:"title"`
				Percentage string `json:"percentage"`
			} `json:"cancelPolicies"`
			OtherEquipmentNames []string `json:"otherEquipmentNames"`
			FreeEquipments      []struct {
				Typename         string `json:"__typename"`
				Title            string `json:"title"`
				Description      string `json:"description"`
				CountDescription string `json:"countDescription"`
			} `json:"freeEquipments"`
			ChargedEquipments []struct {
				Typename         string      `json:"__typename"`
				SummaryPrice     string      `json:"summaryPrice"`
				Title            string      `json:"title"`
				Description      string      `json:"description"`
				CountDescription interface{} `json:"countDescription"`
			} `json:"chargedEquipments"`
			Images []struct {
				Typename string `json:"__typename"`
				SpaceID  string `json:"spaceId"`
				ID       string `json:"id"`
				FilePath string `json:"filePath"`
			} `json:"images"`
			ID              string  `json:"id"`
			Title           string  `json:"title"`
			FriendlyTitle   string  `json:"friendlyTitle"`
			IsAnyAvailable  bool    `json:"isAnyAvailable"`
			IsOrderApprove  bool    `json:"isOrderApprove"`
			Square          int     `json:"square"`
			ReviewCount     int     `json:"reviewCount"`
			AveragePoint    float64 `json:"averagePoint"`
			SpaceType       int     `json:"spaceType"`
			Capacity        int     `json:"capacity"`
			SummaryMinPrice int     `json:"summaryMinPrice"`
			SummaryMaxPrice int     `json:"summaryMaxPrice"`
			SpaceURL        string  `json:"spaceUrl"`
		} `json:"rooms"`
	} `json:"data"`
}

// SpaceDetailRequestPayload payload for SpaceDetail graphQL endpoint
type SpaceDetailRequestPayload struct {
	Query     string                    `json:"query"`
	Variables SpaceDetailQueryVariables `json:"variables"`
}

// SpaceDetailQueryVariables allowed parameters
type SpaceDetailQueryVariables struct {
	ID string `json:"id"`
}
