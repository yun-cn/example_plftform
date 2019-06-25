package spacemarket

import (
	"context"
	"net/http"
	"time"

	"github.com/yanshiyason/noonde_platform/spacemarket/types"
)

// RoomsDayService provides access to the RoomsDay endpoint
type RoomsDayService service

// NewRequest prepares a request
func (s *RoomsDayService) NewRequest(id int) (*http.Request, error) {
	url, err := URLFor(id, types.RTDay)
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
func (s *RoomsDayService) Room(ctx context.Context, id int) (*DayRoomResponsePayload, *Response, error) {
	req, err := s.NewRequest(id)

	if err != nil {
		return nil, nil, err
	}

	payload := &DayRoomResponsePayload{}
	resp, err := s.client.Do(ctx, req, payload)

	return payload, resp, nil
}

// DayRoomResponsePayload payload for requests on day time rooms
type DayRoomResponsePayload struct {
	ID         int    `json:"id" mapstructure:"id"`
	UID        string `json:"uid" mapstructure:"uid"`
	Status     int    `json:"status" mapstructure:"status"`
	RentType   int    `json:"rent_type" mapstructure:"rent_type"`
	Space      Space  `json:"space" mapstructure:"space"`
	Owner      Owner  `json:"owner" mapstructure:"owner"`
	Thumbnails []struct {
		Image       string `json:"image" mapstructure:"image"`
		Description string `json:"description" mapstructure:"description"`
	} `json:"thumbnails" mapstructure:"thumbnails"`
	PriceTextList []struct {
		MinPriceText     string      `json:"min_price_text" mapstructure:"min_price_text"`
		MinPriceUnitText interface{} `json:"min_price_unit_text" mapstructure:"min_price_unit_text"`
		MaxPriceText     string      `json:"max_price_text" mapstructure:"max_price_text"`
		MaxPriceUnitText string      `json:"max_price_unit_text" mapstructure:"max_price_unit_text"`
	} `json:"price_text_list" mapstructure:"price_text_list"`
	RoomsDaycore               string           `json:"reputation_score" mapstructure:"reputation_score"`
	ReputationCount            int              `json:"reputation_count" mapstructure:"reputation_count"`
	Name                       string           `json:"name" mapstructure:"name"`
	Capacity                   int              `json:"capacity" mapstructure:"capacity"`
	HasDirectReservationPlans  bool             `json:"has_direct_reservation_plans" mapstructure:"has_direct_reservation_plans"`
	HasLastMinuteDiscountPlans bool             `json:"has_last_minute_discount_plans" mapstructure:"has_last_minute_discount_plans"`
	IsFavorite                 bool             `json:"is_favorite" mapstructure:"is_favorite"`
	InquiryOnly                bool             `json:"inquiry_only" mapstructure:"inquiry_only"`
	InquiryText                string           `json:"inquiry_text" mapstructure:"inquiry_text"`
	MinPriceText               string           `json:"min_price_text" mapstructure:"min_price_text"`
	MinPriceUnitText           string           `json:"min_price_unit_text" mapstructure:"min_price_unit_text"`
	SpaceID                    int              `json:"space_id" mapstructure:"space_id"`
	OwnerID                    int              `json:"owner_id" mapstructure:"owner_id"`
	CountryID                  interface{}      `json:"country_id" mapstructure:"country_id"`
	AllowRentTypes             []types.RentType `json:"allow_rent_types" mapstructure:"allow_rent_types"`
	ReservationMethod          int              `json:"reservation_method" mapstructure:"reservation_method"`
	PriceDisplayType           int              `json:"price_display_type" mapstructure:"price_display_type"`
	CreatedAt                  time.Time        `json:"created_at" mapstructure:"created_at"`
	UpdatedAt                  time.Time        `json:"updated_at" mapstructure:"updated_at"`
	RoomDisplayType            int              `json:"room_display_type" mapstructure:"room_display_type"`
	TaxType                    int              `json:"tax_type" mapstructure:"tax_type"`
	CanDisplay                 bool             `json:"can_display" mapstructure:"can_display"`
	StatusText                 string           `json:"status_text" mapstructure:"status_text"`
	IsAvailable                bool             `json:"is_available" mapstructure:"is_available"`
	Memo                       string           `json:"memo" mapstructure:"memo"`
	AttachFiles                []interface{}    `json:"attach_files" mapstructure:"attach_files"`
	EventTypes                 []struct {
		Name     string `json:"name" mapstructure:"name"`
		NameText string `json:"name_text" mapstructure:"name_text"`
	} `json:"event_types" mapstructure:"event_types"`
	Description          string        `json:"description" mapstructure:"description"`
	Area                 float64       `json:"area" mapstructure:"area"`
	EmbedVideoURL        interface{}   `json:"embed_video_url" mapstructure:"embed_video_url"`
	EmbedVrURL           interface{}   `json:"embed_vr_url" mapstructure:"embed_vr_url"`
	IsTemairazu          bool          `json:"is_temairazu" mapstructure:"is_temairazu"`
	FoodDescription      string        `json:"food_description" mapstructure:"food_description"`
	TrashDescription     string        `json:"trash_description" mapstructure:"trash_description"`
	EquipmentDescription string        `json:"equipment_description" mapstructure:"equipment_description"`
	RoomType             interface{}   `json:"room_type" mapstructure:"room_type"`
	KeyExchangeType      int           `json:"key_exchange_type" mapstructure:"key_exchange_type"`
	SeatedCapacity       interface{}   `json:"seated_capacity" mapstructure:"seated_capacity"`
	StandingCapacity     interface{}   `json:"standing_capacity" mapstructure:"standing_capacity"`
	Bedrooms             interface{}   `json:"bedrooms" mapstructure:"bedrooms"`
	Beds                 interface{}   `json:"beds" mapstructure:"beds"`
	Bathrooms            interface{}   `json:"bathrooms" mapstructure:"bathrooms"`
	Toilets              interface{}   `json:"toilets" mapstructure:"toilets"`
	Policy               interface{}   `json:"policy" mapstructure:"policy"`
	PolicyType           int           `json:"policy_type" mapstructure:"policy_type"`
	PolicyTypeTitle      string        `json:"policy_type_title" mapstructure:"policy_type_title"`
	Terms                interface{}   `json:"terms" mapstructure:"terms"`
	Plans                []Plan        `json:"plans" mapstructure:"plans"`
	OptionItems          []interface{} `json:"option_items" mapstructure:"option_items"`
	Amenities            []struct {
		Name     string `json:"name" mapstructure:"name"`
		NameText string `json:"name_text" mapstructure:"name_text"`
	} `json:"amenities" mapstructure:"amenities"`
	RelatedEntries []interface{} `json:"related_entries" mapstructure:"related_entries"`
	FavoritesCount int           `json:"favorites_count" mapstructure:"favorites_count"`
	UsageRates     []struct {
		ID            int     `json:"id" mapstructure:"id"`
		RoomID        int     `json:"room_id" mapstructure:"room_id"`
		EventTypeID   int     `json:"event_type_id" mapstructure:"event_type_id"`
		Rate          float64 `json:"rate" mapstructure:"rate"`
		EventTypeText string  `json:"event_type_text" mapstructure:"event_type_text"`
	} `json:"usage_rates" mapstructure:"usage_rates"`
	SiblingRooms []interface{} `json:"sibling_rooms" mapstructure:"sibling_rooms"`
	RoomsDay     []Reputation  `json:"RoomsDay" mapstructure:"RoomsDay"`
}
