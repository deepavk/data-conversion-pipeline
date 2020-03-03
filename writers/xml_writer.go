package writers

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"pipeline/models"
)

var xmlFilePath = "../pipeline/output_files/restaurant.xml"

type XmlWriter struct {
	FileHandle *os.File
	fileName   string
}

func NewXmlDataWriter() *XmlWriter {
	return new(XmlWriter)
}

func (xw *XmlWriter) Setup() error {
	fh, err := os.OpenFile(xmlFilePath, os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Printf("error in opening file to write xml: %s", err)
		return err
	}
	xw.FileHandle = fh
	xw.fileName = xmlFilePath
	return nil
}

func (xw *XmlWriter) WriteData(restaurants *models.Restaurants) error {
	output, err := xml.MarshalIndent(restaurants, "", "\t")
	if err != nil {
		log.Println("error marshalling to xml:", err)
		return err
	}

	err = ioutil.WriteFile(xw.fileName, output, 0755)
	if err != nil {
		log.Println("error writing xml to file:", err)
		return err
	}
	log.Printf("data written to xml file %s successfully", xw.fileName)
	return nil
}

func (xw *XmlWriter) Close() {
	xw.FileHandle.Close()
}
