package writers

import (
	"pipeline/models"
	"reflect"
	"testing"
)

func TestGetDataWriter(t *testing.T) {
	tests := []struct {
		name     string
		fileType string
		want     DataWriter
		wantErr  bool
	}{
		{"json_writer", models.JsonData, NewJsonDataWriter(), false},
		{"xml_writer", models.XmlData, NewXmlDataWriter(), false},
		{"invalid", "toml", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDataWriter(tt.fileType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDataWriter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDataWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}
