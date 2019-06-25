package noonde

import (
	"fmt"
	"sort"
)

//Amenity is an amenity included in the rental.
type Amenity string

// Amenities an array of amenity
type Amenities []Amenity

func (aa Amenities) String() string {
	sorted := aa.AsStringArray()
	sort.Strings(sorted)
	return fmt.Sprintf("%s", sorted)
}

// AsStringArray convert each to string.
func (aa Amenities) AsStringArray() []string {
	strs := []string{}
	for _, a := range aa {
		strs = append(strs, string(a))
	}

	return strs
}

func (a Amenity) String() string {
	return string(a)
}

//TODO:...
const (
	AmenityAccessibility        Amenity = "accessibility"
	AmenityAirConditioner               = "air_conditioner"
	AmenityBar                          = "bar"
	AmenityBathtub                      = "bathtub"
	AmenityCableHdmi                    = "cable_hdmi"
	AmenityCateringServices             = "catering_services"
	AmenityChairs                       = "chairs"
	AmenityChangingRoom                 = "changing_room"
	AmenityChild                        = "child"
	AmenityCommercialShoot              = "commercial_shoot"
	AmenityCooking                      = "cooking"
	AmenityDisposal                     = "disposal"
	AmenityDrinking                     = "drinking"
	AmenityDvdBrPlayer                  = "dvd_br_player"
	AmenityElevator                     = "elevator"
	AmenityExtensionCord                = "extension_cord"
	AmenityFilmEquipment                = "film_equipment"
	AmenityFlooring                     = "flooring"
	AmenityFoodAllowed                  = "food_allowed"
	AmenityFullLengthMirror             = "full_length_mirror"
	AmenityHangerRack                   = "hanger_rack"
	AmenityHotplate                     = "hotplate"
	AmenityInternetWifiMobile           = "internet_wifi_mobile"
	AmenityKitchenEquipment             = "kitchen_equipment"
	AmenityKitchenFacilities            = "kitchen_facilities"
	AmenityLedLight                     = "led_light"
	AmenityLightingEquipment            = "lighting_equipment"
	AmenityLockers                      = "lockers"
	AmenityMarket                       = "market"
	AmenityMicrowave                    = "microwave"
	AmenityMirrors                      = "mirrors"
	AmenityMonitor                      = "monitor"
	AmenityNaturalLight                 = "natural_light"
	AmenityNearStation10min             = "near_station_10min"
	AmenityOverPm10                     = "over_pm10"
	AmenityParking                      = "parking"
	AmenityPet                          = "pet"
	AmenityPhotographyEquipment         = "photography_equipment"
	AmenityPot                          = "pot"
	AmenityPowerEquipment               = "power_equipment"
	AmenityPrinter                      = "printer"
	AmenityPrivateRoom                  = "private_room"
	AmenityProjector                    = "projector"
	AmenityPublicTransportation         = "public_transportation"
	AmenityRefrigerator                 = "refrigerator"
	AmenityRestaurant                   = "restaurant"
	AmenityShowers                      = "showers"
	AmenitySmoking                      = "smoking"
	AmenitySmokingArea                  = "smoking_area"
	AmenitySofa                         = "sofa"
	AmenitySoundSystem                  = "sound_system"
	AmenitySurveillanceCamera           = "surveillance_camera"
	AmenityTables                       = "tables"
	AmenityTakoyaki                     = "takoyaki"
	AmenityToilet                       = "toilet"
	AmenityTv                           = "tv"
	AmenityWhiteboard                   = "whiteboard"
	AmenityWifi                         = "wifi"
)

// StringToAmenity convert string back to amenity
var StringToAmenity map[string]Amenity

func init() {
	StringToAmenity = map[string]Amenity{
		"accessibility":         AmenityAccessibility,
		"air_conditioner":       AmenityAirConditioner,
		"bar":                   AmenityBar,
		"bathtub":               AmenityBathtub,
		"cable_hdmi":            AmenityCableHdmi,
		"catering_services":     AmenityCateringServices,
		"chairs":                AmenityChairs,
		"changing_room":         AmenityChangingRoom,
		"child":                 AmenityChild,
		"commercial_shoot":      AmenityCommercialShoot,
		"cooking":               AmenityCooking,
		"disposal":              AmenityDisposal,
		"drinking":              AmenityDrinking,
		"dvd_br_player":         AmenityDvdBrPlayer,
		"elevator":              AmenityElevator,
		"extension_cord":        AmenityExtensionCord,
		"film_equipment":        AmenityFilmEquipment,
		"flooring":              AmenityFlooring,
		"food_allowed":          AmenityFoodAllowed,
		"full_length_mirror":    AmenityFullLengthMirror,
		"hanger_rack":           AmenityHangerRack,
		"hotplate":              AmenityHotplate,
		"internet_wifi_mobile":  AmenityInternetWifiMobile,
		"kitchen_equipment":     AmenityKitchenEquipment,
		"kitchen_facilities":    AmenityKitchenFacilities,
		"led_light":             AmenityLedLight,
		"lighting_equipment":    AmenityLightingEquipment,
		"lockers":               AmenityLockers,
		"market":                AmenityMarket,
		"microwave":             AmenityMicrowave,
		"mirrors":               AmenityMirrors,
		"monitor":               AmenityMonitor,
		"natural_light":         AmenityNaturalLight,
		"near_station_10min":    AmenityNearStation10min,
		"over_pm10":             AmenityOverPm10,
		"parking":               AmenityParking,
		"pet":                   AmenityPet,
		"photography_equipment": AmenityPhotographyEquipment,
		"pot":                   AmenityPot,
		"power_equipment":       AmenityPowerEquipment,
		"printer":               AmenityPrinter,
		"private_room":          AmenityPrivateRoom,
		"projector":             AmenityProjector,
		"public_transportation": AmenityPublicTransportation,
		"refrigerator":          AmenityRefrigerator,
		"restaurant":            AmenityRestaurant,
		"showers":               AmenityShowers,
		"smoking":               AmenitySmoking,
		"smoking_area":          AmenitySmokingArea,
		"sofa":                  AmenitySofa,
		"sound_system":          AmenitySoundSystem,
		"surveillance_camera":   AmenitySurveillanceCamera,
		"tables":                AmenityTables,
		"takoyaki":              AmenityTakoyaki,
		"toilet":                AmenityToilet,
		"tv":                    AmenityTv,
		"whiteboard":            AmenityWhiteboard,
		"wifi":                  AmenityWifi,
	}
}
