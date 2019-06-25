package elasticsearch

import (
	"context"
	"encoding/json"

	"github.com/olivere/elastic"
)

// IndexReview indexes a Review document
func (s *Service) IndexReview(identifiable DocIdentifiable, doc interface{}) error {
	// TODO: accept ctx as param
	return s.Index(s.ReviewsIndex, identifiable.DocID(), "_doc", doc)
}

// DeleteReview drops an indexed Review document
func (s *Service) DeleteReview(identifiable DocIdentifiable) error {
	// TODO: accept ctx as param
	return s.Delete(s.ReviewsIndex, identifiable.DocID(), "_doc")
}

// SearchReviews searches an indexed Review document
func (s *Service) SearchReviews(query *Query) (interface{}, error) {
	// TODO: accept ctx as param
	return s.Search(s.ReviewsIndex, query)
}

// GetReview searches an indexed request document
func (s *Service) GetReview(identifiable DocIdentifiable) (*json.RawMessage, error) {
	// TODO: accept ctx as param
	get, err := s.Client.Get().
		Index(s.ReviewsIndex).
		Type("_doc").
		Id(identifiable.DocID()).
		Do(context.TODO())

	if err != nil {
		return nil, err
	}
	return get.Source, nil
}

// GetReviews gets multiple documents by their IDS
func (s *Service) GetReviews(indentifiables []DocIdentifiable) ([]*json.RawMessage, error) {
	// TODO: accept ctx as param

	mg := s.Client.MultiGet()
	for _, identifiable := range indentifiables {
		mg = mg.Add(elastic.NewMultiGetItem().Index(s.ReviewsIndex).Type("_doc").Id(identifiable.DocID()))
	}

	get, err := mg.Do(context.TODO())

	if err != nil {
		return nil, err
	}
	results := []*json.RawMessage{}

	for _, doc := range get.Docs {
		results = append(results, doc.Source)
	}
	return results, nil
}
