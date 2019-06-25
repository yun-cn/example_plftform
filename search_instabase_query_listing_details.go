package noonde

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func instabaseQueryListingDetails(listingID string) {
	docID := ListingDocID{listingID, PlatformInstabase}

	if doc, _ := searchService.GetListing(docID); doc != nil {
		fmt.Printf("Instabase: Found existing doc %s\n", listingID)

		result := IndexedListingDetails{}
		json.Unmarshal(*doc, &result)

		if time.Now().After(result.RefreshedAt.Add(CacheDuration)) {
			fmt.Printf("Instabase: Refreshing cached results for listing:  %s\n", listingID)
		} else {
			fmt.Printf("Instabase: Cached results still valid for listing: %s\n", listingID)
			return
		}
	}

	fmt.Printf("Instabase: Fetching doc %s\n", docID)

	// Fetch reviews async
	JobQueue.InstabaseFetchListingReviews(listingID)

	payload, _, err := instabaseClient.SpaceDetail.SpaceDetail(context.TODO(), listingID)
	if err != nil {
		fmt.Printf("error getting room details from Instabase: %+v\n", err)
		return
	}

	now := time.Now()
	results := &IndexedListingDetails{
		IDOnPlatform: listingID,
		Data:         payload,
		RefreshedAt:  &now,
	}

	err = searchService.IndexListing(docID, results)
	if err != nil {
		fmt.Printf("error indexing search results: %+v\n", err)
		return
	}
}
