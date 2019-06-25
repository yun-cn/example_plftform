package types

type PriceType string

const (
	PTHourly PriceType = "HOURLY"
)

// EventType represent renting purpose
type EventType int

const (
	ETParty         EventType = 1
	ETClassRoom     EventType = 2
	ETPhotoShoot    EventType = 3
	ETFilmShoot     EventType = 4
	ETSocialEvent   EventType = 5
	ETPerformance   EventType = 7
	ETStudio        EventType = 8
	ETSports        EventType = 9
	ETOffice        EventType = 10
	ETWedding       EventType = 11
	ETOther         EventType = 12
	ETStayBuisiness EventType = 23
	ETStayParty     EventType = 24
	ETStayTrip      EventType = 25
	ETStayGroup     EventType = 26
	ETStayVacation  EventType = 27
)

func (e EventType) String() string {
	return map[EventType]string{
		1:  "Party",
		2:  "ClassRoom",
		3:  "PhotoShoot",
		4:  "FilmShoot",
		5:  "SocialEvent",
		7:  "Performance",
		8:  "Studio",
		9:  "Sports",
		10: "Office",
		11: "Wedding",
		12: "Other",
		23: "StayBuisiness",
		24: "StayParty",
		25: "StayTrip",
		26: "StayGroup",
		27: "StayVacation",
	}[e]
}

// DeduceRentType finds rentType from eventType
func (e EventType) DeduceRentType() RentType {
	switch e {
	case ETParty, ETClassRoom, ETPhotoShoot, ETFilmShoot, ETSocialEvent, ETPerformance, ETStudio, ETSports, ETOffice, ETWedding, ETOther:
		return RTDay
	case ETStayBuisiness, ETStayTrip, ETStayParty, ETStayGroup, ETStayVacation:
		return RTNight
	default:
		panic("unknown type")
	}
}

// RentType represents day time usage or night stay
type RentType int

const (
	_ RentType = iota
	RTDay
	RTNight
)

func (s RentType) String() string {
	return [...]string{"", "day_time", "night_stay"}[s]
}

// StayRoomType shared, private, entire
type StayRoomType int

const (
	_ StayRoomType = iota
	SharedRoom
	PrivateRoom
	EntirePlace
)

func (s StayRoomType) String() string {
	return [...]string{"shared_room", "private_room", "entire_place"}[s]
}

// RoomRentType is the rentType represented as a string enum
type RoomRentType string

const (
	DAYTIME         RoomRentType = "DAY_TIME"
	STAY            RoomRentType = "STAY"
	RoomRentTypeALL RoomRentType = ""
)

// remaining types to decipher:
//
// "key_exchange_type": 1,
// "policy_type": 2,
// "policy_type_title": "普通：48時間キャンセル無料",
// "price_calc_type": 0,
// "price_display_type": 0,
// "price_display_type": 0,
// "registration_type": "normal_entry",
// "room_display_type": 0,
// "space_type": "residential",
// "space_type_text": "住宅",
// "tax_type": 0,
// "user_tel_display_type": 0,
//

// data from spacemarket:
//
// EventTypes
// [{
//     "id": 1,
//     "rent_type": 1,
//     "name": "party",
//     "name_text": "パーティー"
//   },{
//     "id": 2,
//     "rent_type": 1,
//     "name": "class",
//     "name_text": "会議・研修"
//   },{
//     "id": 3,
//     "rent_type": 1,
//     "name": "photo_shoot",
//     "name_text": "写真撮影"
//   },{
//     "id": 4,
//     "rent_type": 1,
//     "name": "film_shoot",
//     "name_text": "ロケ撮影"
//   },{
//     "id": 5,
//     "rent_type": 1,
//     "name": "event",
//     "name_text": "イベント"
//   },{
//     "id": 7,
//     "rent_type": 1,
//     "name": "performance",
//     "name_text": "演奏・パフォーマンス"
//   },{
//     "id": 8,
//     "rent_type": 1,
//     "name": "studio",
//     "name_text": "個展・展示会"
//   },{
//     "id": 9,
//     "rent_type": 1,
//     "name": "sports",
//     "name_text": "スポーツ・フィットネス"
//   },{
//     "id": 10,
//     "rent_type": 1,
//     "name": "office",
//     "name_text": "オフィス"
//   },{
//     "id": 11,
//     "rent_type": 1,
//     "name": "wedding",
//     "name_text": "結婚式"
//   },{
//     "id": 12,
//     "rent_type": 1,
//     "name": "other",
//     "name_text": "その他"
//   },{
//     "id": 23,
//     "rent_type": 2,
//     "name": "stay-business",
//     "name_text": "出張・ビジネス"
//   },{
//     "id": 24,
//     "rent_type": 2,
//     "name": "stay-party",
//     "name_text": "パーティー"
//   },{
//     "id": 25,
//     "rent_type": 2,
//     "name": "stay-trip",
//     "name_text": "旅行"
//  },{"id": 26,
//     "rent_type": 2,
//     "name": "stay-group",
//     "name_text": "合宿・グループ"
// },{ "id": 27,
//     "rent_type": 2,
//     "name": "stay-vacation",
//     "name_text": "バケーションレンタル"}]
//
// StayRoomTypeField
//
// [{ "id": "1",
//    "name": "SHARED_ROOM",
//    "nameText": "shared_room"},
//   {"id": "2",
//    "name": "PRIVATE_ROOM",
//    "nameText": "private_room"},
//   {"id": "3",
//    "name": "ENTIRE_PLACE",
//    "nameText": "entire_place"}]
