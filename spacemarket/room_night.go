package spacemarket

import (
	"context"
	"net/http"
	"time"

	"github.com/yanshiyason/noonde_platform/spacemarket/types"
)

// RoomsNightService provides access to the RoomsDay endpoint
type RoomsNightService service

// NewRequest prepares a request
func (s *RoomsNightService) NewRequest(id int) (*http.Request, error) {
	url, err := URLFor(id, types.RTNight)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url.String(), nil)

	if err != nil {
		return nil, err
	}

	req.Header = DefaultHeaders()

	return req, nil
}

// Room returns the details of a specific room
func (s *RoomsNightService) Room(ctx context.Context, id int) (*NightRoomResponsePayload, *Response, error) {
	req, err := s.NewRequest(id)

	if err != nil {
		return nil, nil, err
	}

	payload := &NightRoomResponsePayload{}
	resp, err := s.client.Do(ctx, req, payload)

	return payload, resp, nil
}

// NightRoomResponsePayload payload for requests on night time rooms
type NightRoomResponsePayload struct {
	ID         int            `json:"id"`
	UID        string         `json:"uid"`
	Status     int            `json:"status"`
	RentType   types.RentType `json:"rent_type"`
	Space      Space          `json:"space"`
	Owner      Owner          `json:"owner"`
	Thumbnails []struct {
		Image       string      `json:"image"`
		Description interface{} `json:"description"`
	} `json:"thumbnails"`
	PriceTextList []struct {
		MinPriceText     string      `json:"min_price_text"`
		MinPriceUnitText string      `json:"min_price_unit_text"`
		MaxPriceText     interface{} `json:"max_price_text"`
		MaxPriceUnitText interface{} `json:"max_price_unit_text"`
	} `json:"price_text_list"`
	ReputationScore            string           `json:"reputation_score"`
	ReputationCount            int              `json:"reputation_count"`
	Name                       string           `json:"name"`
	Capacity                   int              `json:"capacity"`
	HasDirectReservationPlans  bool             `json:"has_direct_reservation_plans"`
	HasLastMinuteDiscountPlans bool             `json:"has_last_minute_discount_plans"`
	IsFavorite                 bool             `json:"is_favorite"`
	InquiryOnly                bool             `json:"inquiry_only"`
	InquiryText                string           `json:"inquiry_text"`
	MinPriceText               string           `json:"min_price_text"`
	MinPriceUnitText           string           `json:"min_price_unit_text"`
	SpaceID                    int              `json:"space_id"`
	OwnerID                    int              `json:"owner_id"`
	CountryID                  interface{}      `json:"country_id"`
	AllowRentTypes             []types.RentType `json:"allow_rent_types"`
	ReservationMethod          int              `json:"reservation_method"`
	PriceDisplayType           int              `json:"price_display_type"`
	CreatedAt                  time.Time        `json:"created_at"`
	UpdatedAt                  time.Time        `json:"updated_at"`
	RoomDisplayType            int              `json:"room_display_type"`
	TaxType                    int              `json:"tax_type"`
	CanDisplay                 bool             `json:"can_display"`
	StatusText                 string           `json:"status_text"`
	IsAvailable                bool             `json:"is_available"`
	Memo                       string           `json:"memo"`
	AttachFiles                []interface{}    `json:"attach_files"`
	EventTypes                 []struct {
		Name     string `json:"name"`
		NameText string `json:"name_text"`
	} `json:"event_types"`
	Description          string      `json:"description"`
	Area                 float64     `json:"area"`
	EmbedVideoURL        interface{} `json:"embed_video_url"`
	EmbedVrURL           interface{} `json:"embed_vr_url"`
	IsTemairazu          bool        `json:"is_temairazu"`
	FoodDescription      interface{} `json:"food_description"`
	TrashDescription     interface{} `json:"trash_description"`
	EquipmentDescription string      `json:"equipment_description"`
	RoomType             int         `json:"room_type"`
	KeyExchangeType      int         `json:"key_exchange_type"`
	SeatedCapacity       interface{} `json:"seated_capacity"`
	StandingCapacity     interface{} `json:"standing_capacity"`
	Bedrooms             int         `json:"bedrooms"`
	Beds                 int         `json:"beds"`
	Bathrooms            int         `json:"bathrooms"`
	Toilets              int         `json:"toilets"`
	Policy               interface{} `json:"policy"`
	PolicyType           int         `json:"policy_type"`
	PolicyTypeTitle      string      `json:"policy_type_title"`
	Terms                interface{} `json:"terms"`
	Plans                []Plan      `json:"plans"`
	OptionItems          []struct {
		ID               int           `json:"id"`
		Name             string        `json:"name"`
		Description      string        `json:"description"`
		HasDailyPrice    bool          `json:"has_daily_price"`
		DailyPrice       interface{}   `json:"daily_price"`
		HasHourlyPrice   bool          `json:"has_hourly_price"`
		HourlyPrice      interface{}   `json:"hourly_price"`
		Countable        bool          `json:"countable"`
		MaxCount         int           `json:"max_count"`
		HasSessionPrice  bool          `json:"has_session_price"`
		SessionPrice     string        `json:"session_price"`
		SessionPriceUnit string        `json:"session_price_unit"`
		Status           int           `json:"status"`
		SortOrder        int           `json:"sort_order"`
		CreatedAt        time.Time     `json:"created_at"`
		UpdatedAt        time.Time     `json:"updated_at"`
		PriceText        string        `json:"price_text"`
		PriceUnitText    string        `json:"price_unit_text"`
		Attachments      []interface{} `json:"attachments"`
	} `json:"option_items"`
	Amenities      []interface{} `json:"amenities"`
	RelatedEntries []interface{} `json:"related_entries"`
	FavoritesCount int           `json:"favorites_count"`
	UsageRates     []struct {
		ID            int             `json:"id"`
		RoomID        int             `json:"room_id"`
		EventTypeID   types.EventType `json:"event_type_id"`
		Rate          float64         `json:"rate"`
		EventTypeText string          `json:"event_type_text"`
	} `json:"usage_rates"`
	SiblingRooms []struct {
		ID       int            `json:"id"`
		UID      string         `json:"uid"`
		Status   int            `json:"status"`
		RentType types.RentType `json:"rent_type"`
		Space    struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Username      string `json:"username"`
			Access        string `json:"access"`
			SpaceType     string `json:"space_type"`
			Capacity      int    `json:"capacity"`
			Latitude      string `json:"latitude"`
			Longitude     string `json:"longitude"`
			City          string `json:"city"`
			State         string `json:"state"`
			StateText     string `json:"state_text"`
			SpaceTypeText string `json:"space_type_text"`
			CountryAlpha2 string `json:"country_alpha2"`
		} `json:"space"`
		Owner struct {
			ID              int         `json:"id"`
			Username        string      `json:"username"`
			Profile         string      `json:"profile"`
			IsCorp          bool        `json:"is_corp"`
			CorpName        interface{} `json:"corp_name"`
			CorpNameKana    interface{} `json:"corp_name_kana"`
			Status          int         `json:"status"`
			Rank            int         `json:"rank"`
			ProfileImage    string      `json:"profile_image"`
			Name            string      `json:"name"`
			NameKana        string      `json:"name_kana"`
			IsVerifiedEmail bool        `json:"is_verified_email"`
			IsVerifiedTel   bool        `json:"is_verified_tel"`
		} `json:"owner"`
		Thumbnails                 []interface{} `json:"thumbnails"`
		PriceTextList              []interface{} `json:"price_text_list"`
		ReputationScore            string        `json:"reputation_score"`
		ReputationCount            int           `json:"reputation_count"`
		Name                       interface{}   `json:"name"`
		Capacity                   interface{}   `json:"capacity"`
		HasDirectReservationPlans  bool          `json:"has_direct_reservation_plans"`
		HasLastMinuteDiscountPlans bool          `json:"has_last_minute_discount_plans"`
		IsFavorite                 bool          `json:"is_favorite"`
		InquiryOnly                bool          `json:"inquiry_only"`
		InquiryText                string        `json:"inquiry_text"`
		MinPriceText               string        `json:"min_price_text"`
		MinPriceUnitText           string        `json:"min_price_unit_text"`
	} `json:"sibling_rooms"`
	Reputations []Reputation `json:"reputations"`
}
