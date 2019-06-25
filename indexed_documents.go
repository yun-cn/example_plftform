package noonde

import "time"

// IndexedListingDetails todo
type IndexedListingDetails struct {
	IDOnPlatform string      `json:"id_on_platform"`
	Data         interface{} `json:"data"`
	RefreshedAt  *time.Time  `json:"refreshed_at"`
}

// IndexedPlaceSuggestions todo
type IndexedPlaceSuggestions struct {
	Results []*Suggestion `json:"results"`
}

// IndexedUserDoc ..
type IndexedUserDoc struct {
	Platform     Platform `json:"platform"`
	IDOnPlatform string   `json:"id_on_platform"`
	Name         string   `json:"name"`
	AvatarURL    string   `json:"avatar_url"`
}

// IndexedReviewDoc ..
type IndexedReviewDoc struct {
	IDOnPlatform        string   `json:"id_on_platform"`
	UserIDOnPlatform    string   `json:"user_id_on_platform"`
	ListingIDOnPlatform string   `json:"listing_id_on_platform"`
	Platform            Platform `json:"platform"`
	Description         string   `json:"description"`
	UserName            string   `json:"user_name"`
	UserAvatarURL       string   `json:"user_avatar_url"`
}
