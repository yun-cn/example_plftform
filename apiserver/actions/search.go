package actions

import (
	"strconv"
	"strings"

	"github.com/gobuffalo/buffalo"
	noonde "github.com/yanshiyason/noonde_platform"
)

// SearchHandler is a default handler to serve up
// a home page.
type searchParams struct {
	Page      int              `form:"page"`
	Location  string           `form:"location"`
	Amenities []noonde.Amenity `form:"amenities"`
}

func SearchHandler(c buffalo.Context) error {
	params := &searchParams{
		Location: c.Param("location"),
	}

	if c.Param("page") != "" {
		if page, err := strconv.Atoi(c.Param("page")); err == nil {
			params.Page = page
		}
	}

	if c.Param("amenities") != "" {
		var amenities []noonde.Amenity
		for _, a := range strings.Split(c.Param("amenities"), ",") {
			if amenity, ok := noonde.StringToAmenity[a]; ok {
				amenities = append(amenities, amenity)
			}
		}
		params.Amenities = amenities
	}

	q := &noonde.SearchQuery{
		Location:  params.Location,
		EventType: noonde.EventTypeOffice,
		Page:      params.Page,
		Amenities: params.Amenities,
	}

	searchResults := noonde.Search(q)

	// bb, _ := json.Marshal(searchResults)
	// w.Write(bb)

	// if dig == "" {
	// 	w.Write(searchResults)
	// 	return
	// }

	// "#.data.name"

	// data := gjson.GetBytes(searchResults, dig)

	// results := []string{}
	// for _, j := range data.Array() {
	// 	results = append(results, j.String())
	// }
	// bb, _ := json.Marshal(results)

	// w.Write(bb)

	return c.Render(200, r.JSON(&searchResults))
}

// func searchHandler(w http.ResponseWriter, r *http.Request) {
// 	params := r.URL.Query()
// 	location := params.Get("location")
// 	pageStr := params.Get("page")
// 	amenitiesStr := params.Get("amenities")
// 	// dig := params.Get("dig")

// 	var amenities []noonde.Amenity
// 	for _, a := range strings.Split(amenitiesStr, ",") {
// 		if amenity, ok := noonde.StringToAmenity[a]; ok {
// 			amenities = append(amenities, amenity)
// 			fmt.Printf("GOT AMENITY: %s\n", amenity)
// 		} else {
// 			fmt.Printf("UNKNOWN AMENITY: %s\n", a)
// 		}
// 	}

// 	page, err := strconv.Atoi(pageStr)
// 	if err != nil {
// 		w.Write([]byte("invalid page"))
// 		return
// 	}

// 	q := &noonde.SearchQuery{
// 		Location:  location,
// 		EventType: noonde.EventTypeOffice,
// 		Page:      page,
// 		Amenities: amenities,
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// searchResults := noonde.Search(q)

// bb, _ := json.Marshal(searchResults)
// w.Write(bb)

// // if dig == "" {
// // 	w.Write(searchResults)
// // 	return
// // }

// // // "#.data.name"

// // data := gjson.GetBytes(searchResults, dig)

// // results := []string{}
// // for _, j := range data.Array() {
// // 	results = append(results, j.String())
// // }
// // bb, _ := json.Marshal(results)

// // w.Write(bb)
// }
