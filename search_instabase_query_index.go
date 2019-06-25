package noonde

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/yanshiyason/noonde_platform/instabase"
)

// instabaseQueryIndex will either retrived cached results or fetch new results from Instabase
// returns the listingIds
func instabaseQueryIndex(query *SearchQuery) []string {
	suggestion := instabaseQueryPlaceSuggestions(query.Location)

	// invalid location
	if suggestion == nil {
		fmt.Printf("No suggesitions found for location: %s\n", query.Location)
		return nil
	}

	// Get query key
	docID := SearchQueryDocID{query, PlatformInstabase}

	// Retrieve last results from storage
	if doc, _ := searchService.GetRequest(docID); doc != nil {
		results := ListingsSearchResults{}
		json.Unmarshal(*doc, &results)

		if time.Now().After(results.RefreshedAt.Add(CacheDuration)) {
			fmt.Printf("Instabase: Refreshing cached results for:  %s\n", docID)
		} else {
			fmt.Printf("Instabase: Cached results still valid for: %s\n", docID)
			return results.ListingIds
		}
	}

	params := &instabase.SearchParams{
		PerPage: 20,
		Page:    query.Page,
		// TODO. Noonde EventType!
		// UsageIds
		// Figure out what these mean: CategoryIds, CapacityIds, ConditionIds:
		//
	}
	if suggestion.AreaID != 0 {
		params.AreaID = suggestion.AreaID
	}
	if suggestion.WardID != 0 {
		params.WardID = suggestion.WardID
	}
	if suggestion.PrefectureID != 0 {
		params.PrefectureID = suggestion.PrefectureID
	}

	if len(query.Amenities) != 0 {
		params.EquipmentIds = query.Amenities.ToInstabaseIDs()
	}

	// AreaID
	// PrefectureID
	// StationID
	// UsageIds
	// Requesting new search results
	payload, _, err := instabaseClient.Search.Search(context.Background(), params)

	if err != nil {
		fmt.Printf("error searching Instabase: %+v\n", err)
		return nil
	}

	// Find all listing ids
	var listingIds = []string{}
	for _, doc := range payload.Data.Spaces {
		listingIds = append(listingIds, doc.ID)
	}

	now := time.Now()
	results := &ListingsSearchResults{
		ListingIds:  listingIds,
		Results:     payload,
		Platform:    PlatformInstabase,
		RefreshedAt: &now,
	}

	// Index the search results
	err = searchService.IndexRequest(docID, results)
	if err != nil {
		fmt.Printf("error indexing search results: %+v\n", err)
		return nil
	}

	return listingIds
}
