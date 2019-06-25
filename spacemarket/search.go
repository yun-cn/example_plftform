package spacemarket

import (
	"math"
	"net/http"
)

// TotalPages returns total amount of pages concidering number of items requested per page
func (r *SearchStayRoomsResponsePayload) TotalPages(perPage int) int {
	totalResults := r.Data.SearchStayRooms.PageInfo.TotalCount
	return int(math.Ceil(float64(totalResults) / float64(perPage)))
}

// TotalPages returns total amount of pages concidering number of items requested per page
func (r *SearchRoomsResponsePayload) TotalPages(perPage int) int {
	totalResults := r.Data.SearchRooms.PageInfo.TotalCount
	return int(math.Ceil(float64(totalResults) / float64(perPage)))
}

// DefaultHeaders returns default headers
func DefaultHeaders() http.Header {
	hdrs := http.Header{}
	hdrs.Set("host", "v3api.spacemarket.com")
	hdrs.Set("Content-Type", "application/json")
	hdrs.Set("Accept", "*/*")
	hdrs.Set("X-Sm-Ios-App-Version", "4.2.1")
	hdrs.Set("Accept-Timezone", "Asia/Tokyo")
	hdrs.Set("X-Api-Key", apiKey)
	hdrs.Set("Accept-Language", "en_JP")
	hdrs.Set("Accept-Encoding", "br, gzip, deflate")
	hdrs.Set("User-Agent", userAgent)
	hdrs.Set("Spacemarket-Version", "2018-11-05")
	return hdrs
}
