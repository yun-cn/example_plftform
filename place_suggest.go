package noonde

import (
	"context"
	"fmt"

	"github.com/yanshiyason/noonde_platform/instabase"
)

// searchService search
var instabaseClient *instabase.Client

func init() {
	instabaseClient = instabase.NewClient(nil)
}

// Suggestions collection of Suggestion
type Suggestions []*Suggestion

// Suggestion is a potential place supported by instabase
type Suggestion struct {
	Name         string `json:"name"`
	WardID       int    `json:"ward_id"`
	AreaID       int    `json:"area_id"`
	PrefectureID int    `json:"prefecture_id"`
}

// WithAreaID filter those with area ids
func (ss Suggestions) WithAreaID() []*Suggestion {
	var results []*Suggestion
	for _, s := range ss {
		if s.AreaID != 0 {
			results = append(results, s)
		}
	}
	return results
}

// WithPrefectureID filter those with prefecture ids
func (ss Suggestions) WithPrefectureID() []*Suggestion {
	var results []*Suggestion
	for _, s := range ss {
		if s.PrefectureID != 0 {
			results = append(results, s)
		}
	}
	return results
}

// WithWardID filter those with ward ids
func (ss Suggestions) WithWardID() []*Suggestion {
	var results []*Suggestion
	for _, s := range ss {
		if s.WardID != 0 {
			results = append(results, s)
		}
	}
	return results
}

// Names return suggestion Names
func (ss Suggestions) Names() []string {
	var results []string
	for _, s := range ss {
		results = append(results, s.Name)
	}
	return results
}

// PlaceSuggest returns possible locations in instabase.
func PlaceSuggest(keyword string) Suggestions {
	payload, _, err := instabaseClient.PlaceSuggest.PlaceSuggest(context.Background(), keyword)

	if err != nil {
		fmt.Printf("error searching place suggest: %+v\n", err)
		return nil
	}

	// Find all listing ids
	var suggestions = []*Suggestion{}
	for _, doc := range payload.Data.PlaceSearch {
		s := &Suggestion{Name: doc.Name}
		switch doc.ModelType {
		case "s":
			s.AreaID = doc.ModelID
		case "wardId":
			s.WardID = doc.ModelID
		case "prefectureId":
			s.PrefectureID = doc.ModelID
		default:
			panic(fmt.Errorf("UNKNOWN modelType for instabase place suggest: %s", doc.ModelType))
		}

		suggestions = append(suggestions, s)
	}

	return suggestions
}
