package noonde

import (
	"fmt"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/yanshiyason/noonde_platform/instabase"
)

// ListingsSearchResults todo
type ListingsSearchResults struct {
	ListingIds  []string    `json:"listing_ids"`
	Results     interface{} `json:"results"`
	Platform    Platform    `json:"platform"`
	RefreshedAt *time.Time  `json:"refreshed_at"`
}

// SearchInstabase will either retrived cached results or fetch new results from Instabase
func SearchInstabase(query *SearchQuery) []*Listing {
	listingIds := instabaseQueryIndex(query)
	listingDetails := instabaseRefreshListingDetails(listingIds)

	listings := []*Listing{}
	for _, d := range listingDetails {
		payload := instabase.SpaceDetailResponsePayload{}

		// TODO: CLEANUP THIS MESS
		data := d.Data.(map[string]interface{})
		stringToDateTimeHook := func(
			f reflect.Type,
			t reflect.Type,
			data interface{}) (interface{}, error) {
			if t == reflect.TypeOf(time.Time{}) && f == reflect.TypeOf("") {
				return time.Parse(time.RFC3339, data.(string))
			}

			return data, nil
		}

		config := mapstructure.DecoderConfig{
			DecodeHook: stringToDateTimeHook,
			Result:     &payload,
		}

		decoder, err := mapstructure.NewDecoder(&config)
		if err != nil {
			panic(err)
		}

		err = decoder.Decode(data)
		if err != nil {
			panic(err)
		}
		// TODO: CLEANUP THIS MESS -- END...

		room := payload.Data.Rooms[0]

		amenities := []Amenity{}
		for _, a := range room.OtherEquipmentNames {
			amenities = append(amenities, InstabaseToNoondeAmenity(a))
		}
		for _, a := range room.FreeEquipments {
			amenities = append(amenities, InstabaseToNoondeAmenity(a.Title))
		}

		min := room.SummaryMinPrice
		max := room.SummaryMaxPrice
		unit := "時間"
		priceRange := fmt.Sprintf("%d ~ %d / %s", min, max, unit)

		thumbnails := []*ListingThumbnail{}
		for _, t := range room.Images {
			thumbnails = append(thumbnails, &ListingThumbnail{URL: t.FilePath, Description: ""})
		}

		listings = append(listings, &Listing{
			Platform:    PlatformInstabase,
			Name:        room.FriendlyTitle,
			Description: room.SeoDescription,
			Amenities:   amenities,
			Lat:         fmt.Sprintf("%f", room.Building.Lat),
			Lon:         fmt.Sprintf("%f", room.Building.Lon),
			ReviewCount: room.ReviewCount,
			Capacity:    room.Capacity,
			PriceRange:  priceRange,
			Thumbnails:  thumbnails,
		})
	}
	return listings
}
