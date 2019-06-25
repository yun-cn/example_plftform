package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic"
)

// IndexListing indexes a listing document
func (s *Service) IndexListing(identifiable DocIdentifiable, doc interface{}) error {
	// TODO: accept ctx as param
	return s.Index(s.ListingsIndex, identifiable.DocID(), "_doc", doc)
}

// DeleteListing drops an indexed listing document
func (s *Service) DeleteListing(identifiable DocIdentifiable) error {
	// TODO: accept ctx as param
	return s.Delete(s.ListingsIndex, identifiable.DocID(), "_doc")
}

// SearchListings searches an indexed listing document
func (s *Service) SearchListings(query *Query) (interface{}, error) {
	// TODO: accept ctx as param
	return s.Search(s.ListingsIndex, query)
}

// GetListing searches an indexed request document
func (s *Service) GetListing(identifiable DocIdentifiable) (*json.RawMessage, error) {
	// TODO: accept ctx as param
	get, err := s.Client.Get().
		Index(s.ListingsIndex).
		Type("_doc").
		Id(identifiable.DocID()).
		Do(context.TODO())

	if err != nil {
		return nil, err
	}
	return get.Source, nil
}

// GetListings gets multiple documents by their IDS
// ...DocIdentifiable
func (s *Service) GetListings(identifiables []interface{}) ([]*json.RawMessage, error) {
	// TODO: accept ctx as param

	mg := s.Client.MultiGet()
	for _, identifiable := range identifiables {
		if _, ok := identifiable.(DocIdentifiable); !ok {
			panic(fmt.Errorf("all elements must implement DocIdentifiable"))
		}
		id := identifiable.(DocIdentifiable)
		mg = mg.Add(elastic.NewMultiGetItem().Index(s.ListingsIndex).Type("_doc").Id(id.DocID()))
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
