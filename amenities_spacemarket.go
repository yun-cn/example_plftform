package noonde

import "strings"

// ToSpacemarket formats the amenities how spacemarket is expecting them.  (comma de) which spacemarket understands.
func (aa Amenities) ToSpacemarket() string {
	return strings.Join(aa.AsStringArray(), ",")
}

// SpacemarketToNoondeAmenity translates spacemarket language to noonde language.
func SpacemarketToNoondeAmenity(amenity string) Amenity {
	return Amenity(amenity)
}
