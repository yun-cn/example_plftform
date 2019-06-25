package noonde

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/yanshiyason/noonde_platform/spacemarket"
	"github.com/yanshiyason/noonde_platform/spacemarket/types"
)

// spacemarketQueryIndex will either retrived cached results or fetch new results from spacemarket
// returns the listingIds
func spacemarketQueryIndex(query *SearchQuery) []int {
	// (*client.SearchRoomsResponsePayload, *client.Response, error)

	// Get docID
	docID := SearchQueryDocID{query, PlatformSpacemarket}

	// Retrieve last results from storage
	if doc, _ := searchService.GetRequest(docID); doc != nil {
		results := SpacemarketListingsSearchResults{}
		json.Unmarshal(*doc, &results)

		if time.Now().After(results.RefreshedAt.Add(CacheDuration)) {
			fmt.Printf("Spacemarket: Refreshing cached results for:  %s\n", docID)
		} else {
			fmt.Printf("Spacemarket: Cached results still valid for: %s\n", docID)
			return results.ListingIds
		}
	}

	// Requesting new search results
	payload, _, err := spacemarketClient.SearchDay.Search(context.Background(), &spacemarket.SearchRoomsParams{
		Location:  query.Location,
		PerPage:   20,
		Page:      query.Page,
		PriceType: types.PTHourly,
		EventType: query.EventType.ToSpacemarket(),
		Amenities: query.Amenities.ToSpacemarket(),
	})

	if err != nil {
		fmt.Printf("error searching spacemarket: %+v\n", err)
		return nil
	}

	// Find all listing ids
	var listingIds = []int{}
	for _, doc := range payload.Data.SearchRooms.Results {
		listingIds = append(listingIds, doc.ID)
	}

	now := time.Now()
	results := &SpacemarketListingsSearchResults{
		ListingIds:  listingIds,
		Results:     payload,
		Platform:    PlatformSpacemarket,
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
