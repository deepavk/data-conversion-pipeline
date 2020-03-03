package models

import "sort"

const (
	name     = "name"
	rating   = "rating"
	JsonData = "json"
	XmlData  = "xml"
)

var RestaurantFieldsCount = 6

type Restaurant struct {
	Name    string `json:"name" xml:"name"`
	Address string `json:"address" xml:"address"`
	Rating  int    `json:"rating" xml:"rating"`
	Contact string `json:"contact" xml:"contact"`
	Phone   string `json:"phone" xml:"phone"`
	Url     string `json:"url" xml:"url"`
}

type RestaurantData []*Restaurant

type Restaurants struct {
	RestaurantData `json:"restaurants" xml:"restaurant"`
}

func (h RestaurantData) Sort(field string, asc bool) {
	if len(field) == 0 {
		return
	}
	if asc {
		if field == name {
			sort.Slice(h[:], func(i, j int) bool {
				return h[i].Name < h[j].Name
			})
		}
		if field == rating {
			sort.Slice(h[:], func(i, j int) bool {
				return h[i].Rating < h[j].Rating
			})
		}
	} else {
		if field == name {
			sort.Slice(h[:], func(i, j int) bool {
				return h[i].Name > h[j].Name
			})
		}
		if field == rating {
			sort.Slice(h[:], func(i, j int) bool {
				return h[i].Rating > h[j].Rating
			})
		}
	}
}
