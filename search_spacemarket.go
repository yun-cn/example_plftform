package noonde

import (
	"fmt"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"

	"github.com/yanshiyason/noonde_platform/spacemarket"
)

// SpacemarketListingsSearchResults todo
type SpacemarketListingsSearchResults struct {
	ListingIds  []int       `json:"listing_ids"`
	Results     interface{} `json:"results"`
	Platform    Platform    `json:"platform"`
	RefreshedAt *time.Time  `json:"refreshed_at"`
}

// SearchSpacemarket todo comment
func SearchSpacemarket(query *SearchQuery) []*Listing {
	listingIds := spacemarketQueryIndex(query)
	listingDetails := spacemarketRefreshListingDetails(listingIds)

	listings := []*Listing{}
	for _, d := range listingDetails {
		payload := spacemarket.DayRoomResponsePayload{}

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

		amenities := []Amenity{}
		for _, a := range payload.Amenities {
			amenities = append(amenities, SpacemarketToNoondeAmenity(a.Name))
		}

		min := payload.PriceTextList[0].MinPriceText
		max := payload.PriceTextList[0].MaxPriceText
		unit := payload.PriceTextList[0].MaxPriceUnitText
		priceRange := fmt.Sprintf("%s ~ %s / %s", min, max, unit)

		thumbnails := []*ListingThumbnail{}
		for _, t := range payload.Thumbnails {
			thumbnails = append(thumbnails, &ListingThumbnail{URL: t.Image, Description: t.Description})
		}

		listings = append(listings, &Listing{
			Platform:    PlatformSpacemarket,
			Name:        payload.Name,
			Description: payload.Description,
			Amenities:   amenities,
			Lat:         payload.Space.Latitude,
			Lon:         payload.Space.Longitude,
			ReviewCount: payload.Space.ReputationCount,
			Capacity:    payload.Capacity,
			PriceRange:  priceRange,
			Thumbnails:  thumbnails,
		})
	}
	return listings
}
