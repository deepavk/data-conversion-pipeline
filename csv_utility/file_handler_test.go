package csv_utility

import (
	"reflect"
	"testing"
)

func TestCsvHandler_ReadInvalid(t *testing.T) {
	csvh := &CsvHandler{}
	tests := []struct {
		name     string
		fileName string
		want     [][]string
		wantErr  bool
	}{
		{"file_not_found", "../csv_utility/testdata/not_present.csv", nil, true},
		{"invalid_cols", "../csv_utility/testdata/invalid_cols.csv", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csvh.fileName = tt.fileName
			got, err := csvh.Read()
			if err == nil && tt.wantErr {
				t.Errorf("CsvHandler.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CsvHandler.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCsvHandler_ReadValid(t *testing.T) {
	csvh := &CsvHandler{}
	csvh.fileName = "../csv_utility/testdata/valid.csv"
	rows, err := csvh.Read()
	validRows := [][]string{{"name1", "addr1", "1", "contact1", "phone1", "http://abc.com"}}
	if !reflect.DeepEqual(len(rows), len(validRows)) {
		t.Errorf("invalid  %d", len(rows))
	}
	if err != nil {
		t.Errorf("fail %s", err)
	}
}
