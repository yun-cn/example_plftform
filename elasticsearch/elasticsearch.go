package elasticsearch

import (
	"context"

	"github.com/olivere/elastic"
)

// implements the SearchIndexer interface

// DocIdentifiable to be able to be used as a document ID.
type DocIdentifiable interface {
	DocID() string
}

// Service ..
type Service struct {
	Client *elastic.Client
	// ListingsIndex name of the index for listings
	ListingsIndex string
	// RequestsIndex name of the index for requests
	RequestsIndex string
	// PlaceSuggestionsIndex name of the index for instabase place suggestions
	PlaceSuggestionsIndex string
	// UsersIndex name of the index for users
	UsersIndex string
	// ReviewsIndex name of the index for Reviews
	ReviewsIndex string
}

// NewDefaultService with default client
func NewDefaultService() *Service {
	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}

	return &Service{
		Client:                client,
		ListingsIndex:         "noonde_listings",
		RequestsIndex:         "noonde_search_requests",
		PlaceSuggestionsIndex: "noonde_place_suggestions",
		UsersIndex:            "noonde_users",
		ReviewsIndex:          "noonde_reviews",
	}
}

// Index indexes a document
func (s *Service) Index(indexName string, docID string, docType string, doc interface{}) error {
	// TODO: accept ctx as param
	ctx := context.TODO()
	_, err := s.Client.Index().Index(indexName).Type(docType).Id(docID).BodyJson(doc).Do(ctx)
	return err
}

// Delete drops an indexed document
func (s *Service) Delete(index string, id string, docType string) error {
	// TODO: accept ctx as param
	ctx := context.TODO()
	_, err := s.Client.Delete().Index(index).Type(docType).Id(id).Do(ctx)
	return err
}

// Search implements the Searcher interface
func (s *Service) Search(index string, query *Query) (interface{}, error) {
	// TODO: accept ctx as param

	bquery := elastic.NewBoolQuery()

	// if query.Filter != "" && query.Start != nil && query.End != nil {
	// 	bquery.Filter(
	// 		elastic.NewRangeQuery(query.Filter).
	// 			From(query.Start).
	// 			To(query.End),
	// 	)
	// }

	// if query.Type != "" {
	// 	bquery = bquery.Must(
	// 		elastic.NewSimpleQueryStringQuery(query.Type).
	// 			Field("type").
	// 			DefaultOperator("AND"))
	// }

	// if query.IDsQuery != "" {
	// 	bquery = bquery.Must(
	// 		elastic.NewSimpleQueryStringQuery(query.IDsQuery).
	// 			Field("id").
	// 			DefaultOperator("OR"))
	// }

	// if query.SearchQuery != "" {
	// 	bquery = bquery.Must(
	// 		elastic.NewSimpleQueryStringQuery(query.SearchQuery).
	// 			Field("search").
	// 			DefaultOperator("AND"))
	// }

	// if query.TagsQuery != "" {
	// 	bquery = bquery.Must(
	// 		elastic.NewSimpleQueryStringQuery(query.TagsQuery).
	// 			Field("tags").
	// 			DefaultOperator("AND"))
	// }

	if query.IDQuery != "" {
		bquery = bquery.Must(
			elastic.NewIdsQuery(index, query.IDQuery),
		)
	}

	src := elastic.
		NewSearchSource().
		Query(bquery)

	// if query.SortField != "" {
	// 	src = src.Sort(query.SortField, query.SortAsc)
	// }

	// Debug ..
	// sc, _ := src.Source()
	// js, _ := json.Marshal(sc)
	// fmt.Printf("\n\n\n\n\n\n%s\n\n\n\n\n\n", js)

	// TODO implement context
	ctx := context.TODO()

	res, err := s.Client.Search(index).
		SearchSource(src).
		Type("_doc").
		Do(ctx)

	if err != nil {
		// return nil, 0, err
		return nil, err
	}

	// var ids = []int64{}
	// if res.Hits.TotalHits > 0 {
	// 	for _, h := range res.Hits.Hits {
	// 		bb, err := h.Source.MarshalJSON()
	// 		if err != nil {
	// 			return nil, 0, err
	// 		}
	// 		id := gjson.Get(string(bb), "id").Int()
	// 		ids = append(ids, id)
	// 	}
	// }

	// return res.Hits.TotalHits
	return res.Hits.Hits, nil
}

// curl -X GET "localhost:9200/my_index/_search" -H 'Content-Type: application/json' -d'
// {
//   "query": {
//     "terms": {
//       "_uid": [ "_doc#1", "_doc#2" ]
//     }
//   },
//   "aggs": {
//     "UIDs": {
//       "terms": {
//         "field": "_uid",
//         "size": 10
//       }
//     }
//   },
//   "sort": [
//     {
//       "_uid": {
//         "order": "desc"
//       }
//     }
//   ],
//   "script_fields": {
//     "UID": {
//       "script": {
//          "lang": "painless",
//          "source": "doc[\u0027_uid\u0027]"
//       }
//     }
//   }
// }
// '
