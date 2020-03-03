package restaurants

import (
	"pipeline/csv_utility"
	"pipeline/models"
	"pipeline/writers"
	"reflect"
	"testing"
)

func TestRestaurantWriter_writeValidRestaurants(t *testing.T) {
	hw := &RestaurantWriter{
		writers: []writers.DataWriter{writers.GetMockWriter()},
	}
	hw.validRestaurants = append(hw.validRestaurants, &models.Restaurant{})
	err := hw.writeValidRestaurants()
	if err != nil {
		t.Errorf("writeValidRestaurants() error %+v", err)
	}
	writer := hw.writers[0].(*writers.MockDataWriter)
	dataWritten := writer.GetData()
	if len(dataWritten.RestaurantData) != len(hw.validRestaurants) {
		t.Errorf("writeValidRestaurants() error %+v", err)
	}
}

func TestRestaurantWriter_writeInvalidRestaurants(t *testing.T) {
	rows := [][]string{{"name1", "address1", "contact1", "phone1"}}
	hw := &RestaurantWriter{
		writers:        []writers.DataWriter{writers.GetMockWriter()},
		csvFileHandler: csv_utility.GetMockCsv(rows),
	}
	hw.invalidRestaurants = rows
	hw.writeInvalidRestaurants()
	csvh := hw.csvFileHandler.(*csv_utility.MockCsvHandler)
	dataWritten := csvh.GetDataWritten()
	if !reflect.DeepEqual(dataWritten, hw.invalidRestaurants) {
		t.Errorf("writeInValidRestaurants() error")
	}
}

func TestRestaurantWriter_getValidatedData(t *testing.T) {
	hw := new(RestaurantWriter)
	restaurant := &models.Restaurant{"name1", "address1", 4,
		"contact-p1", "123-456", "http://test.com"}
	tests := []struct {
		name    string
		content []string
		isValid bool
		restaurant   *models.Restaurant
	}{
		{"invalid_url", []string{"name1", "address1", "4", "contact-p1", "123-456", "http//url.in"},
			false, nil},
		{"invalid_rating", []string{"name1", "address1", "rating", "contact-p1", "123-456", "http//url.in"},
			false, nil},
		{"invalid_len", []string{"name1", "address1", "4"}, false, nil},
		{"valid", []string{"name1", "address1", "4", "contact-p1", "123-456", "http://test.com"},
			true, restaurant},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid, restaurant := hw.getValidatedData(tt.content)
			if isValid != tt.isValid {
				t.Errorf("RestaurantWriter.getValidatedData() got = %v, want %v", isValid, tt.isValid)
			}
			if !reflect.DeepEqual(restaurant, tt.restaurant) {
				t.Errorf("writeInValidRestaurants() error")
			}
		})

	}
}
