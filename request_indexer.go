package noonde

import "time"

// RequestIndexer indexes data somewhere
type RequestIndexer interface {
	Index(index string, id string, docType string, doc RequestIndexable) error
	Delete(index string, id string, docType string) error
}

// RequestIndexable implement this to be an indexable request
type RequestIndexable interface {
	ToRequestSearchDocument() *RequestSearchDocument
}

// RequestSearchDocument this represents a request made to the search endpoint of a platform.
// it includes the search results, the aggregated IDs of the listings, and the refreshed_at timestamp.
type RequestSearchDocument struct {
	// Key is a string representing this request
	// something like
	//   "spacemarket://location:Tokyo,guests:2,page:1"
	//        "spacee://location:Tokyo,guests:2,page:1"
	//     "instabase://location:Tokyo,guests:2,page:1"
	Key         string        `json:"key"`
	Results     []interface{} `json:"results"`
	ListingIds  []string      `json:"listing_ids"`
	RefreshedAt *time.Time    `json:"refreshed_at"`
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
