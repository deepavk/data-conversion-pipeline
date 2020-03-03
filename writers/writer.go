package writers

import (
	"errors"
	"pipeline/models"
)

type DataWriter interface {
	Setup() error
	WriteData(restaurants *models.Restaurants) error
	Close()
}

func GetDataWriter(outputDataType string) (DataWriter, error) {
	switch outputDataType {
	case models.JsonData:
		return NewJsonDataWriter(), nil
	case models.XmlData:
		return NewXmlDataWriter(), nil
	default:
		return nil, errors.New("unsupported output file format")
	}
}
