package noonde

import (
	"time"
)

// ListingIndexer indexes data somewhere
type ListingIndexer interface {
	Index(index string, id string, docType string, doc ListingIndexable) error
	Delete(index string, id string, docType string) error
}

// ListingIndexable implement this to be indexable
type ListingIndexable interface {
	// TODO: Maybe remove all these methods. Maybe it's easier to have 1 method interface.
	// Type() ListingSearchType
	// SearchQuery() string
	// ID() string
	// IDOnPlatform() string
	// RefreshedAt() *time.Time
	// EventTypes() []EventType
	// Amenities() []Amenity
	// Location() Location
	// Platform() Platform
	// ReviewCount() int
	// TODO: Consider if I need the methods above or just this 1 method.
	ToListingSearchDocument() *ListingSearchDocument
}

// ListingSearchType is either "room" or "space".
type ListingSearchType string

const (
	// ListingSearchTypeRoom means the user wants to stay overnight.
	ListingSearchTypeRoom ListingSearchType = "room"
	// ListingSearchTypeSpace means the user wants to rent the space during daytime.
	ListingSearchTypeSpace = "space"
)

// ListingSearchDocument is the indexed document schema in our search system.
type ListingSearchDocument struct {
	// Type is a ListingSearchType.
	Type ListingSearchType `json:"type"`
	// Search is the search query.
	Search string `json:"search"`
	// ID is the ID of the record in a storage system.
	ID string `json:"id"`
	// IDOnPlatform is the ID of the listing on it's platform (spacemarket, spacee, airbnb, etc).
	IDOnPlatform string `json:"id_on_platform"`
	// RefreshedAt is the last time at which the data was updated from the platform.
	RefreshedAt *time.Time `json:"refreshed_at"`
	// EventTypes is the types of event the listing can be used for.
	EventTypes []EventType `json:"event_type"`
	// Amenities is a list of amenities included in the rental.
	Amenities []Amenity `json:"amenities"`
	// Location is the geo_point for the listing.
	Location Location `json:"location"`
	// DocID is the ID of the document in the search system.
	DocID string `json:"-"`
	// Platform is the platform this listing is on.
	Platform Platform `json:"platform"`
	// ReviewCount is the number of reviews this listing has on it's platform.
	ReviewCount int `json:"review_count"`

	// TODO: Adapt or Remove.
	// Tags I don't know what this will be used for yet.
	Tags []string `json:"tags"`
}

//
// Current elasticsearch mappings:
//
//   "mappings": {
//     "_doc": {
//       "properties": {
//         "type": {
//           "type": "keyword"
//         },
//         "search": {
//           "type": "text",
//           "analyzer": "my_analyzer"
//         },
//         "id": {
//           "type": "long"
//         },
//         "id_on_platform": {
//           "type": "keyword"
//         },
//         "refreshed_at": {
//           "type": "date"
//         },
//         "date": {
//           "type": "date"
//         },
//         "review_count": {
//            "type": "integer"
//         },
//         "event_type": {
//            "type": "keyword"
//         },
//         "amenities": {
//            "type": "keyword"
//         },
//         "location": {
//           "type": "geo_point"
//         },
//         "platform": {
//           "type": "keyword"
//         }
//       }
//     }
//   },
