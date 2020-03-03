package models

import (
	"reflect"
	"testing"
)

func TestRestaurantData_Sort(t *testing.T) {
	restaurant1 := &Restaurant{Name: "abbot", Rating: 5}
	restaurant2 := &Restaurant{Name: "darwin", Rating: 4}
	restaurant3 := &Restaurant{Name: "gibson", Rating: 3}

	noSortRestaurants := Restaurants{RestaurantData: []*Restaurant{restaurant1, restaurant2, restaurant3}}
	sortNameAsc := Restaurants{RestaurantData: []*Restaurant{restaurant1, restaurant2, restaurant3}}
	sortNameDesc := Restaurants{RestaurantData: []*Restaurant{restaurant3, restaurant2, restaurant1}}
	sortRatingDesc := Restaurants{RestaurantData: []*Restaurant{restaurant1, restaurant2, restaurant3}}
	sortRatingAsc := Restaurants{RestaurantData: []*Restaurant{restaurant3, restaurant2, restaurant1}}

	tests := []struct {
		name      string
		expected  Restaurants
		field     string
		ascending bool
	}{
		{"no_sort", noSortRestaurants, "", false},
		{"sort_name_asc", sortNameAsc, "name", true},
		{"sort_name_desc", sortNameDesc, "name", false},
		{"sort_rating_asc", sortRatingAsc, "rating", true},
		{"sort_rating_asc", sortRatingDesc, "rating", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			noSortRestaurants.Sort(tt.field, tt.ascending)
			if !reflect.DeepEqual(noSortRestaurants, tt.expected) {
				t.Errorf("Sorting failed, for key %s", tt.field)
			}
		})
	}
}
