package noonde

import (
	spacemarket "github.com/yanshiyason/noonde_platform/spacemarket/types"
)

// EventType is the type of event the room will be used for.
type EventType string

// TODO ... skip
const (
	EventTypeParty         EventType = "party"
	EventTypeClassRoom               = "class_room"
	EventTypePhotoShoot              = "photo_shoot"
	EventTypeFilmShoot               = "film_shoot"
	EventTypeSocialEvent             = "social_event"
	EventTypePerformance             = "performance"
	EventTypeStudio                  = "studio"
	EventTypeSports                  = "sports"
	EventTypeOffice                  = "office"
	EventTypeWedding                 = "wedding"
	EventTypeOther                   = "other"
	EventTypeStayBuisiness           = "stay_buisiness"
	EventTypeStayParty               = "stay_party"
	EventTypeStayTrip                = "stay_trip"
	EventTypeStayGroup               = "stay_group"
	EventTypeStayVacation            = "stay_vacation"
)

// ToSpacemarket translate Noonde event types to spacemarket event types.
func (et EventType) ToSpacemarket() spacemarket.EventType {
	return map[EventType]spacemarket.EventType{
		EventTypeParty:         spacemarket.ETParty,
		EventTypeClassRoom:     spacemarket.ETClassRoom,
		EventTypePhotoShoot:    spacemarket.ETPhotoShoot,
		EventTypeFilmShoot:     spacemarket.ETFilmShoot,
		EventTypeSocialEvent:   spacemarket.ETSocialEvent,
		EventTypePerformance:   spacemarket.ETPerformance,
		EventTypeStudio:        spacemarket.ETStudio,
		EventTypeSports:        spacemarket.ETSports,
		EventTypeOffice:        spacemarket.ETOffice,
		EventTypeWedding:       spacemarket.ETWedding,
		EventTypeOther:         spacemarket.ETOther,
		EventTypeStayBuisiness: spacemarket.ETStayBuisiness,
		EventTypeStayParty:     spacemarket.ETStayParty,
		EventTypeStayTrip:      spacemarket.ETStayTrip,
		EventTypeStayGroup:     spacemarket.ETStayGroup,
		EventTypeStayVacation:  spacemarket.ETStayVacation,
	}[et]
}
