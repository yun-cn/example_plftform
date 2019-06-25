package noonde

import (
	"context"
	"fmt"

	"github.com/yanshiyason/noonde_platform/spacemarket"
)

// spacemarketQueryReviews fetch reviews for the listing. (and seed users.)
func spacemarketQueryReviews(listingID int) {
	fmt.Printf("Spacemarket: Querying Reviews for %d\n", listingID)

	listingIDString := fmt.Sprintf("%d", listingID)
	// first query has no "cursor"
	payload, _, err := spacemarketClient.Reputations.List(context.TODO(), &spacemarket.ReputationsParams{
		PerPage: 20,
		Page:    1,
		RoomID:  listingIDString,
	})

	if err != nil {
		fmt.Printf("Spacemarket: error querying reviews: %+v\n", err)
		return
	}

	// TODO: think if we need to recursively fetch next pages.
	// if need, must make use of the perPage, and page params.
	var reviewsDocuments []*IndexedReviewDoc

	for _, r := range payload.Data.Reputations.Results {
		reviewsDocuments = append(reviewsDocuments, &IndexedReviewDoc{
			Description:         r.Description,
			IDOnPlatform:        fmt.Sprintf("listingID:%d,usersID:%s", listingID, r.From.ID),
			ListingIDOnPlatform: listingIDString,
			Platform:            PlatformSpacemarket,
			UserIDOnPlatform:    r.From.ID,
			UserName:            r.From.Name,
			UserAvatarURL:       r.From.ProfileImage,
		})
	}

	// Create an indexed review and an indexed user.
	for _, rDoc := range reviewsDocuments {
		userDocID := UserDocID{
			Platform: PlatformSpacemarket,
			ID:       rDoc.UserIDOnPlatform,
		}

		uDoc := IndexedUserDoc{
			Platform:     PlatformSpacemarket,
			IDOnPlatform: rDoc.UserIDOnPlatform,
			Name:         rDoc.UserName,
			AvatarURL:    rDoc.UserAvatarURL,
		}

		searchService.IndexUser(userDocID, uDoc)

		reviewDocID := ReviewDocID{
			Platform:  PlatformSpacemarket,
			ListingID: rDoc.ListingIDOnPlatform,
			UserID:    rDoc.UserIDOnPlatform,
			ReviewID:  rDoc.IDOnPlatform,
		}

		searchService.IndexReview(reviewDocID, rDoc)
	}
}
