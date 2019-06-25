package elasticsearch

import (
	"context"
	"encoding/json"

	"github.com/olivere/elastic"
)

// IndexUser indexes a User document
func (s *Service) IndexUser(identifiable DocIdentifiable, doc interface{}) error {
	// TODO: accept ctx as param
	return s.Index(s.UsersIndex, identifiable.DocID(), "_doc", doc)
}

// DeleteUser drops an indexed User document
func (s *Service) DeleteUser(identifiable DocIdentifiable) error {
	// TODO: accept ctx as param
	return s.Delete(s.UsersIndex, identifiable.DocID(), "_doc")
}

// SearchUsers searches an indexed User document
func (s *Service) SearchUsers(query *Query) (interface{}, error) {
	// TODO: accept ctx as param
	return s.Search(s.UsersIndex, query)
}

// GetUser searches an indexed request document
func (s *Service) GetUser(identifiable DocIdentifiable) (*json.RawMessage, error) {
	// TODO: accept ctx as param
	get, err := s.Client.Get().
		Index(s.UsersIndex).
		Type("_doc").
		Id(identifiable.DocID()).
		Do(context.TODO())

	if err != nil {
		return nil, err
	}
	return get.Source, nil
}

// GetUsers gets multiple documents by their IDS
func (s *Service) GetUsers(indentifiables []DocIdentifiable) ([]*json.RawMessage, error) {
	// TODO: accept ctx as param

	mg := s.Client.MultiGet()
	for _, identifiable := range indentifiables {
		mg = mg.Add(elastic.NewMultiGetItem().Index(s.UsersIndex).Type("_doc").Id(identifiable.DocID()))
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
