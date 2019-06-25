package elasticsearch

import (
	"context"
	"encoding/json"
)

// IndexRequest indexes a request document
func (s *Service) IndexRequest(identifiable DocIdentifiable, doc interface{}) error {
	// TODO: accept ctx as param
	return s.Index(s.RequestsIndex, identifiable.DocID(), "_doc", doc)
}

// DeleteRequest drops an indexed request document
func (s *Service) DeleteRequest(identifiable DocIdentifiable) error {
	// TODO: accept ctx as param
	return s.Delete(s.RequestsIndex, identifiable.DocID(), "_doc")
}

// GetRequest searches an indexed request document
func (s *Service) GetRequest(identifiable DocIdentifiable) (*json.RawMessage, error) {
	// TODO: accept ctx as param
	get, err := s.Client.Get().
		Index(s.RequestsIndex).
		Type("_doc").
		Id(identifiable.DocID()).
		Do(context.TODO())

	if err != nil {
		return nil, err
	}
	return get.Source, nil
}
