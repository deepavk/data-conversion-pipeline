package writers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"pipeline/models"
)

var jsonFilePath = "../pipeline/output_files/restaurant.json"

type JsonWriter struct {
	FileHandle *os.File
	fileName   string
}

func NewJsonDataWriter() *JsonWriter {
	return new(JsonWriter)
}

func (jw *JsonWriter) Setup() error {
	fh, err := os.OpenFile(jsonFilePath, os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Printf("error in opening file to write json: %s", err)
		return err
	}
	jw.FileHandle = fh
	jw.fileName = jsonFilePath
	return nil
}

func (jw *JsonWriter) WriteData(restaurants *models.Restaurants) error {
	output, err := json.MarshalIndent(restaurants, "", "\t")
	if err != nil {
		log.Println("error marshalling to json:", err)
		return err
	}
	err = ioutil.WriteFile(jw.fileName, output, 0755)
	if err != nil {
		log.Println("error writing json to file:", err)
		return err
	}
	log.Printf("data written to json file %s successfully", jw.fileName)
	return nil
}

func (jw *JsonWriter) Close() {
	jw.FileHandle.Close()
}
