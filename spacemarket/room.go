package spacemarket

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/yanshiyason/noonde_platform/spacemarket/types"
)

// BaseURL for the spacemarket api
const BaseURL = "https://api.spacemarket.com/1"

// RoomHeaders default headers for rooms endpoint
func RoomHeaders() http.Header {
	hdrs := DefaultHeaders()
	hdrs.Set("host", "api.spacemarket.com")
	hdrs.Set("Connection", "keep-alive")
	hdrs.Set("Authorization", "Bearer (null)")
	return hdrs
}

// URLFor returns correct URL depending on RentType
func URLFor(roomID int, rt types.RentType) (*url.URL, error) {
	s := fmt.Sprintf("%s/rooms/%d.json?rent_type=%d", BaseURL, roomID, rt)
	return url.Parse(s)
}

// Space space
type Space struct {
	ID                      int           `json:"id"`
	Name                    string        `json:"name"`
	Username                string        `json:"username"`
	Access                  string        `json:"access"`
	SpaceType               string        `json:"space_type"`
	Capacity                int           `json:"capacity"`
	Latitude                string        `json:"latitude"`
	Longitude               string        `json:"longitude"`
	City                    string        `json:"city"`
	State                   string        `json:"state"`
	StateText               string        `json:"state_text"`
	SpaceTypeText           string        `json:"space_type_text"`
	CountryAlpha2           string        `json:"country_alpha2"`
	OwnerID                 int           `json:"owner_id"`
	CountryID               int           `json:"country_id"`
	HasMapView              bool          `json:"has_map_view"`
	HasStreetView           bool          `json:"has_street_view"`
	Heading                 string        `json:"heading"`
	Pitch                   string        `json:"pitch"`
	Fax                     string        `json:"fax"`
	Homepage                string        `json:"homepage"`
	BusinessStartHour       interface{}   `json:"business_start_hour"`
	BusinessEndHour         interface{}   `json:"business_end_hour"`
	BusinessHourDescription string        `json:"business_hour_description"`
	HasExtraBusinessHours   interface{}   `json:"has_extra_business_hours"`
	ExtraBusinessHours      interface{}   `json:"extra_business_hours"`
	ReputationCount         int           `json:"reputation_count"`
	ReputationScore         string        `json:"reputation_score"`
	AvailableRoomCount      int           `json:"available_room_count"`
	RegistrationType        string        `json:"registration_type"`
	Status                  int           `json:"status"`
	CreatedAt               time.Time     `json:"created_at"`
	UpdatedAt               time.Time     `json:"updated_at"`
	CanDisplay              bool          `json:"can_display"`
	EnableBankPayment       bool          `json:"enable_bank_payment"`
	Description             string        `json:"description"`
	PostalCode              string        `json:"postal_code"`
	Address1                string        `json:"address_1"`
	Address2                string        `json:"address_2"`
	Tel                     string        `json:"tel"`
	Policy                  string        `json:"policy"`
	Terms                   string        `json:"terms"`
	ReservableStart         int           `json:"reservable_start"`
	ReservableDuration      int           `json:"reservable_duration"`
	StatusText              string        `json:"status_text"`
	UtcOffset               int           `json:"utc_offset"`
	CoverPhoto              interface{}   `json:"cover_photo"`
	Thumbnails              []interface{} `json:"thumbnails"`
	IsFavorite              bool          `json:"is_favorite"`
	IsAvailable             bool          `json:"is_available"`
}

// Owner room owner data
type Owner struct {
	ID                     int         `json:"id"`
	Username               string      `json:"username"`
	Profile                string      `json:"profile"`
	IsCorp                 bool        `json:"is_corp"`
	CorpName               interface{} `json:"corp_name"`
	CorpNameKana           interface{} `json:"corp_name_kana"`
	Status                 int         `json:"status"`
	Rank                   int         `json:"rank"`
	ProfileImage           string      `json:"profile_image"`
	Name                   string      `json:"name"`
	NameKana               string      `json:"name_kana"`
	IsVerifiedEmail        bool        `json:"is_verified_email"`
	IsVerifiedTel          bool        `json:"is_verified_tel"`
	CountryID              int         `json:"country_id"`
	Homepage               string      `json:"homepage"`
	CancelCount            int         `json:"cancel_count"`
	CardStatus             string      `json:"card_status"`
	CorpNameEn             interface{} `json:"corp_name_en"`
	RoomDisplayType        int         `json:"room_display_type"`
	PriceDisplayType       int         `json:"price_display_type"`
	ReplyTimeAvg           float64     `json:"reply_time_avg"`
	ReplyRate              float64     `json:"reply_rate"`
	ConfirmRate            float64     `json:"confirm_rate"`
	ReputationCount        int         `json:"reputation_count"`
	ReputationScore        string      `json:"reputation_score"`
	DeletedAt              interface{} `json:"deleted_at"`
	CreatedAt              time.Time   `json:"created_at"`
	UpdatedAt              time.Time   `json:"updated_at"`
	FriendlyID             string      `json:"friendly_id"`
	IsListedGuesthouse     bool        `json:"is_listed_guesthouse"`
	UserTelDisplayType     int         `json:"user_tel_display_type"`
	EnableCard             int         `json:"enable_card"`
	State                  string      `json:"state"`
	StateText              string      `json:"state_text"`
	StatusText             string      `json:"status_text"`
	ReplyTimeAvgText       string      `json:"reply_time_avg_text"`
	ReplyTimeAvgEvaluation int         `json:"reply_time_avg_evaluation"`
	ReplyRateText          string      `json:"reply_rate_text"`
	ReplyRateEvaluation    int         `json:"reply_rate_evaluation"`
	ConfirmRateText        string      `json:"confirm_rate_text"`
	ConfirmRateEvaluation  int         `json:"confirm_rate_evaluation"`
	EnableJcbCards         bool        `json:"enable_jcb_cards"`
}

// Plan response json:plans object -- differant purchasing options
type Plan struct {
	ID                           int         `json:"id"`
	UID                          string      `json:"uid"`
	Name                         string      `json:"name"`
	Description                  string      `json:"description"`
	PriceCalcType                int         `json:"price_calc_type"`
	RentType                     int         `json:"rent_type"`
	HasMinRequiredHour           bool        `json:"has_min_required_hour"`
	HasMaxRequiredHour           bool        `json:"has_max_required_hour"`
	MaxRequiredHour              interface{} `json:"max_required_hour"`
	DailyPrice                   string      `json:"daily_price"`
	HourlyPrice                  string      `json:"hourly_price"`
	SessionPrice                 interface{} `json:"session_price"`
	IsRequiredFood               bool        `json:"is_required_food"`
	HasMinRequiredFoodPrice      bool        `json:"has_min_required_food_price"`
	MinRequiredFoodPrice         interface{} `json:"min_required_food_price"`
	MinHourlyPrice               interface{} `json:"min_hourly_price"`
	MinDailyPrice                interface{} `json:"min_daily_price"`
	MinSessionPrice              interface{} `json:"min_session_price"`
	MinRequiredHourUnit          int         `json:"min_required_hour_unit"`
	MaxRequiredHourUnit          int         `json:"max_required_hour_unit"`
	MinRequiredSessionHourUnit   int         `json:"min_required_session_hour_unit"`
	StayPrice                    interface{} `json:"stay_price"`
	ExtraGuestStayPrice          string      `json:"extra_guest_stay_price"`
	NumberOfStayPlanGuests       int         `json:"number_of_stay_plan_guests"`
	HasMinStayDays               bool        `json:"has_min_stay_days"`
	MinStayDays                  interface{} `json:"min_stay_days"`
	HasMaxStayDays               bool        `json:"has_max_stay_days"`
	MaxStayDays                  interface{} `json:"max_stay_days"`
	HasExtraStayCheckinoutHour   bool        `json:"has_extra_stay_checkinout_hour"`
	ExtraStayCheckinStartHour    interface{} `json:"extra_stay_checkin_start_hour"`
	ExtraStayCheckinEndHour      interface{} `json:"extra_stay_checkin_end_hour"`
	ExtraStayCheckoutHour        interface{} `json:"extra_stay_checkout_hour"`
	IsFoodPlan                   bool        `json:"is_food_plan"`
	IntervalHour                 interface{} `json:"interval_hour"`
	CreatedAt                    time.Time   `json:"created_at"`
	UpdatedAt                    time.Time   `json:"updated_at"`
	DirectReservationAccepted    bool        `json:"direct_reservation_accepted"`
	IsLastMinuteDiscount         bool        `json:"is_last_minute_discount"`
	LastMinuteDiscountPercentage interface{} `json:"last_minute_discount_percentage"`
	WeeklyDiscountPercentage     float64     `json:"weekly_discount_percentage"`
	MonthlyDiscountPercentage    float64     `json:"monthly_discount_percentage"`
	OptionCleaningPrice          interface{} `json:"option_cleaning_price"`
	StayCheckinStartHour         interface{} `json:"stay_checkin_start_hour"`
	StayCheckinEndHour           interface{} `json:"stay_checkin_end_hour"`
	StayCheckoutHour             interface{} `json:"stay_checkout_hour"`
	AvailableWeekday             string      `json:"available_weekday"`
	IsAvailableHoliday           bool        `json:"is_available_holiday"`
	Status                       int         `json:"status"`
	SortOrder                    int         `json:"sort_order"`
	IsPreview                    bool        `json:"is_preview"`
	HasSessionPrice              bool        `json:"has_session_price"`
	HasDailyPrice                bool        `json:"has_daily_price"`
	HasHourlyPrice               bool        `json:"has_hourly_price"`
	HasStayPrice                 bool        `json:"has_stay_price"`
	MinPriceText                 string      `json:"min_price_text"`
	MinPriceUnitText             string      `json:"min_price_unit_text"`
	DailyPriceText               string      `json:"daily_price_text"`
	DailyPriceUnitText           string      `json:"daily_price_unit_text"`
	HourlyPriceText              string      `json:"hourly_price_text"`
	HourlyPriceUnitText          string      `json:"hourly_price_unit_text"`
	StayPriceText                string      `json:"stay_price_text"`
	StayPriceUnitText            string      `json:"stay_price_unit_text"`
	SessionPriceText             string      `json:"session_price_text"`
	SessionPriceUnitText         string      `json:"session_price_unit_text"`
	OptionCleaningPriceText      interface{} `json:"option_cleaning_price_text"`
	MinRequiredSessionHour       interface{} `json:"min_required_session_hour"`
	MinRequiredHour              int         `json:"min_required_hour"`
	HasIntervalHour              bool        `json:"has_interval_hour"`
}

// Reputation response json:reputations object
type Reputation struct {
	ID              int       `json:"id"`
	Score           int       `json:"score"`
	Description     string    `json:"description"`
	ReputatableType string    `json:"reputatable_type"`
	FromType        string    `json:"from_type"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	From            struct {
		DeletedAt       interface{} `json:"deleted_at"`
		Name            string      `json:"name"`
		Username        string      `json:"username"`
		ProfileImage    string      `json:"profile_image"`
		CreatedAt       time.Time   `json:"created_at"`
		UpdatedAt       time.Time   `json:"updated_at"`
		ReputationCount float64     `json:"reputation_count"`
		ReputationScore float64     `json:"reputation_score"`
	} `json:"from"`
	Reputatable struct {
		StayCheckinAt time.Time      `json:"stay_checkin_at"`
		StartedAt     time.Time      `json:"started_at"`
		RentType      types.RentType `json:"rent_type"`
		EventType     string         `json:"event_type"`
		EventTypeText string         `json:"event_type_text"`
		Room          struct {
			Name       string        `json:"name"`
			UID        string        `json:"uid"`
			Thumbnails []interface{} `json:"thumbnails"`
			Space      struct {
				ID       int    `json:"id"`
				Name     string `json:"name"`
				Username string `json:"username"`
			} `json:"space"`
		} `json:"room"`
		NumberOfUsers int `json:"number_of_users"`
	} `json:"reputatable"`
	Comment interface{} `json:"comment"`
}
