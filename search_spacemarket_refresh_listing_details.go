package noonde

import (
	"encoding/json"
	"fmt"
	"sync"
)

func spacemarketRefreshListingDetails(listingIDs []int) []*IndexedListingDetails {
	fmt.Printf("Spacemarket: Got ids %+v\n", listingIDs)

	var wg sync.WaitGroup

	for _, id := range listingIDs {
		wg.Add(1)

		go func(listingID int) {
			defer wg.Done()

			spacemarketQueryListingDetails(listingID)
		}(id)
	}

	// Waiting for all listing details requests to complete.
	wg.Wait()

	listingDocIDs := make([]interface{}, len(listingIDs))
	for i, listingID := range listingIDs {
		listingDocIDs[i] = ListingDocID{
			ListingID: fmt.Sprintf("%d", listingID),
			Platform:  PlatformSpacemarket,
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
