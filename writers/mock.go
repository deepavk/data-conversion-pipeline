package writers

import (
	"errors"
	"os"
	"pipeline/models"
)

type MockDataWriter struct {
	FileHandle *os.File
	fileName   string
	Data       *models.Restaurants
}

func GetMockWriter() DataWriter {
	return new(MockDataWriter)
}

func (m *MockDataWriter) Setup() error {
	return nil
}

func (m *MockDataWriter) WriteData(restaurants *models.Restaurants) error {
	if len(restaurants.RestaurantData) == 0 {
		return errors.New("Empty data")
	}
	m.Data = restaurants
	return nil
}

func (m *MockDataWriter) Close() {
	return
}

func (m MockDataWriter) GetData() *models.Restaurants {
	return m.Data
}
