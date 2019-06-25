package noonde

import "fmt"

// SearchQueryDocID doc id in elasticsearch
type SearchQueryDocID struct {
	Query *SearchQuery
	Platform
}

// ListingDocID doc id in elasticsearch
type ListingDocID struct {
	ListingID string
	Platform
}

// DocID doc id in elasticsearch
func (doc ListingDocID) DocID() string {
	return doc.String()
}

func (doc ListingDocID) String() string {
	return fmt.Sprintf("%s://%s", doc.Platform, doc.ListingID)
}

// DocID prefixed with platform
func (doc SearchQueryDocID) DocID() string {
	return doc.String()
}

// String stringer
func (doc SearchQueryDocID) String() string {
	return fmt.Sprintf("%s://%s", doc.Platform, doc.Query)
}

// UserDocID key indexable in elasticsearch
type UserDocID struct {
	Platform Platform
	ID       string
}

// DocID key in elasticsearch
func (doc UserDocID) DocID() string {
	return fmt.Sprintf("%s://%s", doc.Platform, doc.ID)
}

// ReviewDocID key indexable in elasticsearch
type ReviewDocID struct {
	Platform Platform
	// ListingID on platform
	ListingID string
	// UserID on platform
	UserID string
	// ReviewID on platform
	ReviewID string
}

// DocID key in elasticsearch
func (doc ReviewDocID) DocID() string {
	return fmt.Sprintf("%s://user:%s,listing:%s,review:%s", doc.Platform, doc.UserID, doc.ListingID, doc.ReviewID)
}
