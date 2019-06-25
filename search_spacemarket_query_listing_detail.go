package noonde

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func spacemarketQueryListingDetails(listingID int) {
	docID := ListingDocID{
		ListingID: fmt.Sprintf("%d", listingID),
		Platform:  PlatformSpacemarket,
	}

	if doc, _ := searchService.GetListing(docID); doc != nil {
		fmt.Printf("Found existing doc %s\n", docID)

		result := IndexedListingDetails{}
		json.Unmarshal(*doc, &result)

		if time.Now().After(result.RefreshedAt.Add(CacheDuration)) {
			fmt.Printf("Spacemarket: Refreshing cached results for listing:  %s\n", docID)
		} else {
			fmt.Printf("Spacemarket: Cached results still valid for listing: %s\n", docID)
			return
		}
	}

	fmt.Printf("Spacemarket: Fetching doc %s\n", docID)

	// Fetch reviews async
	JobQueue.SpacemarketFetchListingReviews(listingID)

	payload, _, err := spacemarketClient.RoomsDay.Room(context.TODO(), listingID)
	if err != nil {
		fmt.Printf("error getting room details from spacemarket: %+v\n", err)
		return
	}

	now := time.Now()
	results := &IndexedListingDetails{
		IDOnPlatform: fmt.Sprintf("%d", listingID),
		Data:         payload,
		RefreshedAt:  &now,
	}

	err = searchService.IndexListing(docID, results)
	if err != nil {
		fmt.Printf("error indexing search results: %+v\n", err)
		return
	}
}
