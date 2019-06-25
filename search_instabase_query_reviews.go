package noonde

import (
	"context"
	"fmt"
)

// instabaseQueryReviews fetch reviews for the listing. (and seed users.)
func instabaseQueryReviews(listingID string) {
	fmt.Printf("Instabase: Querying Reviews for %s\n", listingID)
	// first query has no "cursor"
	payload, _, err := instabaseClient.Reviews.Reviews(context.TODO(), listingID, "")

	if err != nil {
		fmt.Printf("Instabase: error querying reviews: %+v\n", err)
		return
	}

	// TODO: think if we need to recursively fetch next pages.
	// if need, must take last cursor from edges.last.cursor.
	// payload.Data.Reviews.PageInfo

	// shape of the review data from instabase:
	//
	// "id": "15961",
	// "title": "アクセスが良く使いやすいスペース",
	// "comment": "ゆっくりとした空間で、会議や商談等に必要な備品も完備されており、使い勝手が良いスペースです。",
	// "point": 5, <--- rating
	// "usage": "電話会議",
	// "age": 6,
	// "gender": 1,
	// "createdAt": "2019-03-07 07:35:18 +0900"
	//
	type instabaseReview struct {
		ID      string
		Title   string
		Comment string
		Rating  int
		// Instabase doesn't give us usernames nor userid...
		// UserName string
	}
	var reviews []*instabaseReview

	for _, e := range payload.Data.Reviews.Edges {
		reviews = append(reviews, &instabaseReview{
			ID:      e.Node.ID,
			Title:   e.Node.Title,
			Comment: e.Node.Comment,
			Rating:  e.Node.Point,
		})
	}

	// Create an indexed review and an indexed user.
	for _, review := range reviews {
		userDocID := UserDocID{
			Platform: PlatformInstabase,
			ID:       review.ID,
		}

		// Using review.ID here because instabase doesn't give us the user ids...
		// this is for creating our "fake" users.
		// Instabase doesn't give us usernames neither ... think how to handle that...
		uDoc := IndexedUserDoc{
			Platform:     PlatformInstabase,
			IDOnPlatform: review.ID,
			Name:         "",
			AvatarURL:    "",
		}

		searchService.IndexUser(userDocID, uDoc)

		reviewDocID := ReviewDocID{
			Platform:  PlatformInstabase,
			ListingID: listingID,
			UserID:    review.ID,
			ReviewID:  review.ID,
		}

		rDoc := &IndexedReviewDoc{
			IDOnPlatform:        review.ID,
			Platform:            PlatformInstabase,
			Description:         fmt.Sprintf("%s. %s", review.Title, review.Comment),
			UserName:            "",
			UserIDOnPlatform:    review.ID,
			ListingIDOnPlatform: listingID,
		}

		searchService.IndexReview(reviewDocID, rDoc)
	}
}
