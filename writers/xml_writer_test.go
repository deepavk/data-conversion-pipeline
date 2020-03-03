package writers

import (
	"os"
	"pipeline/models"
	"testing"
)

var xw *XmlWriter

func setUpData(t *testing.T, perm os.FileMode) {
	filename := "../writers/testdata/valid.xml"
	fh, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, perm)
	if err != nil {
		t.Errorf("XmlWriter.WriteData() error = %v", err)
	}
	xw = NewXmlDataWriter()
	xw.FileHandle = fh
	xw.fileName = filename
}

func TestXmlWriter_WriteData_Valid(t *testing.T) {
	setUpData(t, 0777)
	restaurant := models.Restaurant{"abc", "addr1", 5,
		"def", "123-444", "http://abc.com"}
	restaurants := []*models.Restaurant{}
	restaurants = append(restaurants, &restaurant)

	allRestaurants := models.Restaurants{RestaurantData: restaurants}
	if err := xw.WriteData(&allRestaurants); (err != nil) != false {
		t.Errorf("Valid Xml: XmlWriter.WriteData() have error = %v, wantErr %v", err, false)
	}
	tearDownData(t)
}

func TestXmlWriter_WriteData_PermissionDenied(t *testing.T) {
	setUpData(t, 0000)
	restaurant := models.Restaurant{"abc", "addr1", 5,
		"def", "123-444", "http://abc.com"}
	restaurants := []*models.Restaurant{}
	restaurants = append(restaurants, &restaurant)

	allRestaurants := models.Restaurants{RestaurantData: restaurants}
	if err := xw.WriteData(&allRestaurants); err == nil {
		t.Errorf("Expected permission denied error for file write")
	}
	tearDownData(t)
}

func tearDownData(t *testing.T) {
	err := os.Remove(xw.fileName)
	if err != nil {
		t.Errorf("teardown() error = %v", err)
	}
	xw.Close()
}
