package noonde_test

import (
	"testing"

	noonde "github.com/yanshiyason/noonde_platform"
)

func Test_SearchQuery_Hash(t *testing.T) {
	cases := map[string]struct {
		in  *noonde.SearchQuery
		out string
	}{
		"empty": {
			&noonde.SearchQuery{},
			"location:,check_in:<nil>,duration_in_minutes:0,number_of_guests:0,page:0,amenities:[]",
		},
		"amenities": {
			&noonde.SearchQuery{
				Amenities: noonde.Amenities{noonde.AmenityAirConditioner, noonde.AmenityBathtub},
			},
			"location:,check_in:<nil>,duration_in_minutes:0,number_of_guests:0,page:0,amenities:[air_conditioner bathtub]",
		},
		"amenities order don't matter": {
			&noonde.SearchQuery{
				Amenities: noonde.Amenities{noonde.AmenityBathtub, noonde.AmenityAirConditioner},
			},
			"location:,check_in:<nil>,duration_in_minutes:0,number_of_guests:0,page:0,amenities:[air_conditioner bathtub]",
		},
	}

	for caseName, c := range cases {
		result := c.in.Hash()
		if result != c.out {
			t.Errorf(`
			case name: %s
			expected:  %s
			actual:    %s`, caseName, c.out, result)
		}
	}
}

func Test_Amenities__ToSpacemarketQuery(t *testing.T) {
	cases := map[string]struct {
		in  noonde.Amenities
		out string
	}{
		"amenities are concatenated": {
			noonde.Amenities{noonde.AmenityAirConditioner, noonde.AmenityBathtub},
			"air_conditioner,bathtub",
		},
		"amenities are concatenated and non sorted": {
			noonde.Amenities{noonde.AmenityBathtub, noonde.AmenityAirConditioner, noonde.AmenityCableHdmi},
			"bathtub,air_conditioner,cable_hdmi",
		},
		"empty is empty string": {
			noonde.Amenities{},
			"",
		},
	}

	for caseName, c := range cases {
		result := c.in.ToSpacemarket()
		if result != c.out {
			t.Errorf(`
			case name: %s
			expected:  %s
			actual:    %s`, caseName, c.out, result)
		}
	}
}

func Test_SpacemarketToNoondeAmenity(t *testing.T) {
	cases := []struct {
		in  string
		out noonde.Amenity
	}{
		{"air_conditioner", noonde.AmenityAirConditioner},
		{"bathtub", noonde.AmenityBathtub},
		{"cable_hdmi", noonde.AmenityCableHdmi},
	}

	for _, c := range cases {
		result := noonde.SpacemarketToNoondeAmenity(c.in)
		if result != c.out {
			t.Errorf(`
			expected:  %s
			actual:    %s`, c.out, result)
		}
	}
}
