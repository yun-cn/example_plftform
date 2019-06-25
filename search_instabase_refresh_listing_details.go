package noonde

import (
	"encoding/json"
	"fmt"
	"sync"
)

// instabaseRefreshListingDetails query each listing individually.
func instabaseRefreshListingDetails(listingIDs []string) []*IndexedListingDetails {
	fmt.Printf("Instabase: Got ids %+v\n", listingIDs)

	var wg sync.WaitGroup

	for _, id := range listingIDs {
		wg.Add(1)

		go func(listingID string) {
			defer wg.Done()

			instabaseQueryListingDetails(listingID)
		}(id)
	}

	// Waiting for all listing details requests to complete.
	wg.Wait()

	listingDocIDs := make([]interface{}, len(listingIDs))
	for i, listingID := range listingIDs {
		listingDocIDs[i] = ListingDocID{
			ListingID: listingID,
			Platform:  PlatformInstabase,
		}
	}

	listings, err := searchService.GetListings(listingDocIDs)
	if err != nil {
		fmt.Printf("error multigetting listings: %+v\n", err)
	}

	listingDetails := []*IndexedListingDetails{}
	for _, listing := range listings {
		if listing == nil {
			continue
		}
		d := IndexedListingDetails{}
		json.Unmarshal(*listing, &d)
		listingDetails = append(listingDetails, &d)
	}

	return listingDetails
}
