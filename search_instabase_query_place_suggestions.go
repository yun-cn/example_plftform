package noonde

import (
	"encoding/json"
	"fmt"
)

// PlaceSuggestionDocID key indexable in elasticsearch
type PlaceSuggestionDocID string

// DocID key in elasticsearch
func (doc PlaceSuggestionDocID) DocID() string {
	return string(doc)
}

func (doc PlaceSuggestionDocID) String() string {
	return string(doc)
}

func instabaseQueryPlaceSuggestions(location string) *Suggestion {
	var suggestions Suggestions
	docID := PlaceSuggestionDocID(location)
	// Try get place suggestions from storage
	// Else query instabase, and index the results
	if doc, _ := searchService.GetPlaceSuggestion(docID); doc != nil {
		indexedSuggestions := IndexedPlaceSuggestions{}
		json.Unmarshal(*doc, &indexedSuggestions)
		suggestions = indexedSuggestions.Results
		fmt.Printf("Found place suggestions for %s: %+v\n", docID, suggestions.Names())

	} else {
		// TODO: pass in a context.
		fmt.Printf("Didn't find place suggestions for %s, querying...\n", docID)
		suggestions = PlaceSuggest(location)
		indexedSuggestions := IndexedPlaceSuggestions{Results: suggestions}
		searchService.IndexPlaceSuggestion(docID, indexedSuggestions)
	}

	// invalid location.
	if len(suggestions) == 0 {
		return nil
	}

	withArea := suggestions.WithAreaID()
	withPref := suggestions.WithPrefectureID()
	withWard := suggestions.WithWardID()

	var suggestion *Suggestion
	if len(withArea) != 0 {
		suggestion = withArea[0]
	} else if len(withPref) != 0 {
		suggestion = withPref[0]
	} else if len(withWard) != 0 {
		suggestion = withWard[0]
	}

	return suggestion
}
