package types

// type PriceType string

// const (
// 	PTHourly PriceType = "HOURLY"
// )

// EventType represent renting purpose
type EventType int

// TODO ...
const (
	ETMeeting             EventType = 1  // 打ち合わせ
	ETLesson              EventType = 2  // レッスン・講座
	ETSeminar             EventType = 3  // セミナー・研究
	ETParty               EventType = 4  // パーティー
	ETStudyGroup          EventType = 5  // 勉強会
	ETTherapy             EventType = 6  // セラピー
	ETYoga                EventType = 7  // ヨガ・ダンス
	ETWorkplace           EventType = 8  // 作業所
	ETPhotoShoot          EventType = 10 // 撮影・収録
	ETInterviewOrTest     EventType = 11 // 面接・試験
	ETBoardgame           EventType = 12 // ボードゲーム
	ETCounseling          EventType = 13 // カウンセリング
	ETMothersOrGirlsParty EventType = 14 // 女子会・ママ会
	ETWorkshop            EventType = 15 // ワークショップ
	ETShowingOfAMovie     EventType = 16 // 上映会・映画鑑賞
)

// func (e EventType) String() string {
// 	return map[EventType]string{
// 		1:  "Party",
// 		2:  "ClassRoom",
// 		3:  "PhotoShoot",
// 		4:  "FilmShoot",
// 		5:  "SocialEvent",
// 		7:  "Performance",
// 		8:  "Studio",
// 		9:  "Sports",
// 		10: "Office",
// 		11: "Wedding",
// 		12: "Other",
// 		23: "StayBuisiness",
// 		24: "StayParty",
// 		25: "StayTrip",
// 		26: "StayGroup",
// 		27: "StayVacation",
// 	}[e]
// }

// // DeduceRentType finds rentType from eventType
// func (e EventType) DeduceRentType() RentType {
// 	switch e {
// 	case ETParty, ETClassRoom, ETPhotoShoot, ETFilmShoot, ETSocialEvent, ETPerformance, ETStudio, ETSports, ETOffice, ETWedding, ETOther:
// 		return RTDay
// 	case ETStayBuisiness, ETStayTrip, ETStayParty, ETStayGroup, ETStayVacation:
// 		return RTNight
// 	default:
// 		panic("unknown type")
// 	}
// }

// RentType represents day time usage or night stay
// type RentType int

// const (
// 	_ RentType = iota
// 	RTDay
// 	RTNight
// )

// func (s RentType) String() string {
// 	return [...]string{"", "day_time", "night_stay"}[s]
// }

// StayRoomType shared, private, entire
// type StayRoomType int

// const (
// 	_ StayRoomType = iota
// 	SharedRoom
// 	PrivateRoom
// 	EntirePlace
// )

// func (s StayRoomType) String() string {
// 	return [...]string{"shared_room", "private_room", "entire_place"}[s]
// }

// RoomRentType is the rentType represented as a ...
// type RoomRentType string

// const (
// 	DAYTIME         RoomRentType = "DAY_TIME"
// 	STAY            RoomRentType = "STAY"
// 	RoomRentTypeALL RoomRentType = ""
// )

// data from instabase:
//
// EventTypes
// usage ids:
// uchiawase: 1,打ち合わせ
// lecon: 2,レッスン・講座
// seminar 3,セミナー・研究
// party: 4,パーティー
// studygroup 5,勉強会
// therapy 6,セラピー
// yoga 7,ヨガ・ダンス
// workplace 8,作業所
// photoshoot: 10,撮影・収録
// interview/test 11,面接・試験
// boardgame 12,ボードゲーム
// counseling 13,カウンセリング
// mom party 14,女子会・ママ会
// workshop 15,ワークショップ
// showing film 16,上映会・映画鑑賞
