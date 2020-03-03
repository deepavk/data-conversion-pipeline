package writers

import (
	"os"
	"pipeline/models"
	"testing"
)

var jw *JsonWriter

func setUp(t *testing.T, perm os.FileMode) {
	filename := "../writers/testdata/valid.json"
	fh, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, perm)
	if err != nil {
		t.Errorf("JsonWriter.WriteData() error = %v", err)
	}
	jw = NewJsonDataWriter()
	jw.FileHandle = fh
	jw.fileName = filename
}

func TestJsonWriter_WriteData_Valid(t *testing.T) {
	setUp(t, 0777)
	restaurant := models.Restaurant{Name: "abc", Address: "addr1", Rating: 5,
		Contact: "def", Phone: "123-444", Url: "http://abc.com"}
	restaurants := []*models.Restaurant{}
	restaurants = append(restaurants, &restaurant)

	allRestaurants := models.Restaurants{RestaurantData: restaurants}
	if err := jw.WriteData(&allRestaurants); (err != nil) != false {
		t.Errorf("Valid json: JsonWriter.WriteData() have error = %v, wantErr %v", err, false)
	}
	tearDown(t)
}

func TestJsonWriter_WriteData_PermissionDenied(t *testing.T) {
	setUp(t, 0000)
	restaurant := models.Restaurant{"abc", "addr1", 5,
		"def", "123-444", "http://abc.com"}
	restaurants := []*models.Restaurant{}
	restaurants = append(restaurants, &restaurant)

	allRestaurants := models.Restaurants{RestaurantData: restaurants}
	if err := jw.WriteData(&allRestaurants); err == nil {
		t.Errorf("Expected permission denied error for file write")
	}
	tearDown(t)
}

func tearDown(t *testing.T) {
	err := os.Remove(jw.fileName)
	if err != nil {
		t.Errorf("teardown() error = %v", err)
	}
	jw.Close()
}
