package elasticsearch

import (
	"context"
	"encoding/json"

	"github.com/olivere/elastic"
)

// IndexPlaceSuggestion indexes a PlaceSuggestion document
func (s *Service) IndexPlaceSuggestion(indentifiable DocIdentifiable, doc interface{}) error {
	// TODO: accept ctx as param
	return s.Index(s.PlaceSuggestionsIndex, indentifiable.DocID(), "_doc", doc)
}

// DeletePlaceSuggestion drops an indexed PlaceSuggestion document
func (s *Service) DeletePlaceSuggestion(indentifiable DocIdentifiable) error {
	// TODO: accept ctx as param
	return s.Delete(s.PlaceSuggestionsIndex, indentifiable.DocID(), "_doc")
}

// SearchPlaceSuggestions searches an indexed PlaceSuggestion document
func (s *Service) SearchPlaceSuggestions(query *Query) (interface{}, error) {
	// TODO: accept ctx as param
	return s.Search(s.PlaceSuggestionsIndex, query)
}

// GetPlaceSuggestion searches an indexed request document
func (s *Service) GetPlaceSuggestion(indentifiable DocIdentifiable) (*json.RawMessage, error) {
	// TODO: accept ctx as param
	get, err := s.Client.Get().
		Index(s.PlaceSuggestionsIndex).
		Type("_doc").
		Id(indentifiable.DocID()).
		Do(context.TODO())

	if err != nil {
		return nil, err
	}
	return get.Source, nil
}

// GetPlaceSuggestions gets multiple documents by their IDS
func (s *Service) GetPlaceSuggestions(indentifiables []DocIdentifiable) ([]*json.RawMessage, error) {
	// TODO: accept ctx as param

	mg := s.Client.MultiGet()
	for _, indentifiable := range indentifiables {
		mg = mg.Add(elastic.NewMultiGetItem().Index(s.PlaceSuggestionsIndex).Type("_doc").Id(indentifiable.DocID()))
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
