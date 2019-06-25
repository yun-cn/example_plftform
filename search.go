package noonde

import (
	"fmt"
	"time"

	"github.com/yanshiyason/noonde_platform/elasticsearch"

	"github.com/yanshiyason/noonde_platform/spacemarket"
)

// CacheDuration is how long we cache search results
var CacheDuration = 10 * time.Minute

// searchService search
var searchService *elasticsearch.Service
var spacemarketClient *spacemarket.Client

func init() {
	searchService = elasticsearch.NewDefaultService()
	spacemarketClient = spacemarket.NewClient(nil)
}

// SearchQuery is the base search query. This will be translated to each platform.
// TODO: include all possible parameters
type SearchQuery struct {
	SearchTerm        string
	Location          string
	CheckIn           *time.Time
	DurationInMinutes int
	NumberOfGuests    int
	EventType         EventType
	Page              int
	Amenities         Amenities
}

// Hash returns a unique key for the query
// TODO: include all possible parameters
func (s *SearchQuery) Hash() string {
	return fmt.Sprintf(
		"location:%s,check_in:%s,duration_in_minutes:%d,number_of_guests:%d,page:%d,amenities:%s",
		s.Location,
		s.CheckIn,
		s.DurationInMinutes,
		s.NumberOfGuests,
		s.Page,
		s.Amenities.String(),
	)
}

func (s *SearchQuery) String() string {
	return s.Hash()
}

// Listing is a search result
type Listing struct {
	Platform    Platform            `json:"platform"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Amenities   []Amenity           `json:"amenities"`
	Lat         string              `json:"lat"`
	Lon         string              `json:"lon"`
	ReviewCount int                 `json:"review_count"`
	Capacity    int                 `json:"capacity"`
	PriceRange  string              `json:"price_range"`
	Thumbnails  []*ListingThumbnail `json:"thumbnails"`
}

// ListingThumbnail img url and description.
type ListingThumbnail struct {
	URL         string `json:"url"`
	Description string `json:"description"`
}

// InstabaseRefreshListingDetails

// Search returns results form multiple platforms
func Search(query *SearchQuery) []*Listing {
	spacemarketListings := SearchSpacemarket(query)
	instabaseListings := SearchInstabase(query)

	fmt.Printf(`
	got %d listings for spacemarket
	got %d listings for instabase
	`, len(spacemarketListings), len(instabaseListings))

	var result []*Listing

	result = append(result, spacemarketListings...)
	result = append(result, instabaseListings...)

	return result
}

// TODO benchmark when executed concurrently.
// func Search(query *SearchQuery) []*Listing {
// 	var wg sync.WaitGroup
// 	var spacemarketListings []*Listing
// 	var instabaseListings []*Listing

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		spacemarketListings = SearchSpacemarket(query)
// 	}()

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		instabaseListings = SearchInstabase(query)
// 	}()

// 	wg.Wait()

// 	fmt.Printf(`
// 	got %d listings for spacemarket
// 	got %d listings for instabase
// 	`, len(spacemarketListings), len(instabaseListings))

// 	var result []*Listing

// 	result = append(result, spacemarketListings...)
// 	result = append(result, instabaseListings...)

// 	return result
// }

// The noonde server will proxy search requests to all platforms, aggregate the results, and return them to the client.
//
//
// The client will send us:
//      Search:
//			Location: 新宿
//			check-in: 2019/03/20
//          type:     space
//
// Nonde server will translate that to the vocabulary each platform uses:
//
//		SpacemarketSearch:
//			Place: 新宿
//          Start: 2019/03/20
//          RoomType: 2
//
//		Spacee:
//			Location: 新宿
//          First: 2019/03/20
//
//		Instabase:
//			In: 新宿
//          Start: 2019/03/20
//
// For each param group that noonde sends to a platform, it will hash the values of the query to gain a key,
// and associate the room_ids returned by the platform to it.
// Next time the same query is made, the hashed value will match, and we return the room ids right away (if cache is busted).
//
//

// TODO:
// // SearchSpacee will either retrived cached results or fetch new results from spacee
// func SearchSpacee(query *SearchQuery) {
// }
