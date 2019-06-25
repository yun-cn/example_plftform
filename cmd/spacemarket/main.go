package main

import (
	"context"
	"fmt"
	"io/ioutil"

	client "github.com/yanshiyason/noonde_platform/spacemarket"
	"github.com/yanshiyason/noonde_platform/spacemarket/types"
)

func main() {

	c := client.NewClient(nil)

	// payload, resp, err := c.SearchDay.Search(context.Background(), &client.SearchRoomsParams{
	// 	Location:  "tokyo",
	// 	PerPage:   20,
	// 	Page:      1,
	// 	PriceType: types.PTHourly,
	// 	EventType: types.ETParty,
	// })

	// payload, resp, err := c.RoomsDay.Room(context.Background(), 36598)

	// c.RoomsDay.Room(context.TODO(), )

	// payload, resp, err := c.SearchNight.Search(context.TODO(), &client.SearchStayRoomsParams{
	// 	Location:  "tokyo",
	// 	PerPage:   20,
	// 	Page:      1,
	// 	EventType: types.ETStayBuisiness,
	// })

	payload, resp, err := c.Reputations.List(context.Background(), &client.ReputationsParams{
		Page:     1,
		PerPage:  20,
		RentType: types.RoomRentTypeALL,
		RoomID:   "43485",
	})

	bytes, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("Error: %s\n", err)
	fmt.Printf("response: %v\n", string(bytes))
	fmt.Printf("payload: %v\n", payload)
}
