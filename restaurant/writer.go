package restaurants

import (
	"log"
	"net/url"
	"pipeline/csv_utility"
	"pipeline/models"
	"pipeline/writers"
	"strconv"
	"unicode/utf8"
)

var headers = []string{"name", "address", "rating", "contact", "phone", "url"}

const invalidRestaurantsFile = "../pipeline/output_files/invalid.csv"

type SortOption struct {
	Key string
	Asc bool
}

type RestaurantWriter struct {
	outputDataFormat []string
	invalidRestaurants    [][]string
	writers          []writers.DataWriter
	validRestaurants      models.RestaurantData
	sort             SortOption
	csvFileHandler   csv_utility.Csv
}

func NewRestaurantWriter(filename string, outputFormats []string, sort SortOption) (*RestaurantWriter, error) {
	hw := &RestaurantWriter{outputDataFormat: outputFormats, sort: sort}
	hw.csvFileHandler = csv_utility.NewCsvReader(filename)
	err := hw.initializeWriters()
	if err != nil {
		return nil, err
	}
	return hw, nil
}

func (hw *RestaurantWriter) WriteRestaurantData() error {
	rowsRead, err := hw.csvFileHandler.Read()
	if err != nil {
		return err
	}
	totalRowCount := len(rowsRead)
	validCount := 0
	for _, content := range rowsRead {
		isValid, restaurant := hw.getValidatedData(content)
		if !isValid {
			hw.invalidRestaurants = append(hw.invalidRestaurants, content)
			continue
		}
		hw.validRestaurants = append(hw.validRestaurants, restaurant)
		validCount += 1
	}
	err = hw.writeValidRestaurants()
	if err != nil {
		return err
	}
	hw.writeInvalidRestaurants()
	hw.closeWriters()
	log.Printf("%d valid rows and %d invalid rows out of total rows: %d",
		validCount, totalRowCount-validCount, totalRowCount)
	return nil
}

func (hw *RestaurantWriter) writeValidRestaurants() error {
	for _, writer := range hw.writers {
		hw.validRestaurants.Sort(hw.sort.Key, hw.sort.Asc)
		allRestaurants := &models.Restaurants{RestaurantData: hw.validRestaurants}
		err := writer.WriteData(allRestaurants)
		if err != nil {
			log.Printf("error writing restaurants %s", err)
			return err
		}
	}
	return nil
}

func (hw *RestaurantWriter) writeInvalidRestaurants() {
	if len(hw.invalidRestaurants) == 0 {
		return
	}
	err := hw.csvFileHandler.Write(invalidRestaurantsFile, headers, hw.invalidRestaurants)
	if err != nil {
		log.Printf("error writing invalid restaurants to csv file %s", err)
	}
}

func (hw *RestaurantWriter) getValidatedData(content []string) (bool, *models.Restaurant) {
	// validate length of each row before accessing the index
	if len(content) != models.RestaurantFieldsCount {
		return false, nil
	}
	name := content[0]
	if !utf8.ValidString(name) {
		return false, nil
	}
	rating, err := strconv.Atoi(content[2])
	if err != nil {
		return false, nil
	}
	uri, err := url.ParseRequestURI(content[5])
	if err != nil {
		return false, nil
	}

	restaurant := &models.Restaurant{
		Name:    name,
		Address: content[1],
		Rating:  rating,
		Contact: content[3],
		Phone:   content[4],
		Url:     uri.String(),
	}
	return true, restaurant
}

func (hw *RestaurantWriter) initializeWriters() error {
	for _, format := range hw.outputDataFormat {
		dataWriter, err := writers.GetDataWriter(format)
		if err != nil {
			return err
		}
		err = dataWriter.Setup()
		if err != nil {
			return err
		}
		hw.writers = append(hw.writers, dataWriter)
	}
	return nil
}

func (hw *RestaurantWriter) closeWriters() {
	for _, w := range hw.writers {
		w.Close()
	}
}
