package instabase

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/yanshiyason/noonde_platform/instabase/constants"
)

// SearchService provides access to the search stay endpoint
type SearchService service

// SearchParams params for request
type SearchParams = searchQueryVariables

// Search search rooms for night stay
func (s *SearchService) Search(ctx context.Context, params *SearchParams) (*SearchResponsePayload, *Response, error) {
	req, err := s.NewRequest(params)

	if err != nil {
		return nil, nil, err
	}

	payload := &SearchResponsePayload{}
	resp, err := s.client.Do(ctx, req, payload)

	if err != nil {
		return nil, resp, err
	}

	return payload, resp, nil
}

// NewRequest request for Search
func (s *SearchService) NewRequest(p *SearchParams) (*http.Request, error) {
	payload := s.requestPayload(p)
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
func (s *SearchService) requestPayload(p *SearchParams) *searchRequestPayload {
	return &searchRequestPayload{
		Variables: *p,
		Query:     constants.SearchResultQuery,
	}
}

// SearchResponsePayload shape of the search response
type SearchResponsePayload struct {
	Data struct {
		Spaces []struct {
			Typename string `json:"__typename"`
			Images   []struct {
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
			SeoDescription  string  `json:"seoDescription"`
			Square          int     `json:"square"`
			ReviewCount     int     `json:"reviewCount"`
			AveragePoint    float64 `json:"averagePoint"`
			SpaceType       int     `json:"spaceType"`
			Capacity        int     `json:"capacity"`
			SummaryPrice    string  `json:"summaryPrice"`
			SummaryMinPrice int     `json:"summaryMinPrice"`
			SummaryMaxPrice int     `json:"summaryMaxPrice"`
			SpaceURL        string  `json:"spaceUrl"`
			Building        struct {
				Typename        string `json:"__typename"`
				SummaryAccesses []struct {
					Typename string `json:"__typename"`
					Access   string `json:"access"`
					Line     string `json:"line"`
					Station  string `json:"station"`
				} `json:"summaryAccesses"`
				Lat        float64     `json:"lat"`
				Lon        float64     `json:"lon"`
				ParentArea interface{} `json:"parentArea"`
			} `json:"building"`
		} `json:"spaces"`
	} `json:"data"`
}

// searchRequestPayload payload for search graphQL endpoint
type searchRequestPayload struct {
	Query     string               `json:"query"`
	Variables searchQueryVariables `json:"variables"`
}

// searchQueryVariables allowed parameters
type searchQueryVariables struct {
	AreaID         interface{} `json:"areaId"`
	BottomRightLat interface{} `json:"bottomRightLat"`
	BottomRightLon interface{} `json:"bottomRightLon"`
	CapacityIds    []int       `json:"capacityIds"` //
	CategoryIds    []int       `json:"categoryIds"` //
	ConditionIds   []int       `json:"conditionIds"`
	EquipmentIds   []int       `json:"equipmentIds"`  // todo figure out the numbers
	FromDateDay    int         `json:"fromDateDay"`   // 15
	FromDateMonth  int         `json:"fromDateMonth"` // 3
	FromDateYear   int         `json:"fromDateYear"`  // 2019
	FromTime       string      `json:"fromTime"`      // "12:00"
	ToTime         string      `json:"toTime"`        // "13:00"
	OrderBy        interface{} `json:"orderBy"`
	Page           int         `json:"page"`    // 1
	PerPage        interface{} `json:"perPage"` // null
	PrefectureID   interface{} `json:"prefectureId"`
	StationID      int         `json:"stationId"` // 2600316 get from suggest api
	TopLeftLat     interface{} `json:"topLeftLat"`
	TopLeftLon     interface{} `json:"topLeftLon"`
	UsageIds       []int       `json:"usageIds"` // EventType
	WardID         interface{} `json:"wardId"`
}
