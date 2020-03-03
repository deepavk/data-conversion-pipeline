package csv_utility

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

type Csv interface {
	Read() ([][]string, error)
	Write(filename string, headers []string, rows [][]string) error
}

type CsvHandler struct {
	fileName string
}

func NewCsvReader(filename string) Csv {
	return &CsvHandler{fileName: filename}
}

func (csvh *CsvHandler) Read() ([][]string, error) {
	// Skip first row of headers
	f, err := os.Open(csvh.fileName)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	row1, err := bufio.NewReader(f).ReadSlice('\n')
	if err != nil {
		log.Printf("error in reading csv headers %s", err)
		return nil, err
	}
	_, err = f.Seek(int64(len(row1)), io.SeekStart)
	if err != nil {
		log.Printf("error in reading csv  %s", err)
		return nil, err
	}

	// Read remaining rows
	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		log.Printf("error in reading csv rows %s", err)
		return nil, err
	}

	return rows, nil
}

func (csvh *CsvHandler) Write(filename string, headers []string, rows [][]string) error {
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("error in writing csv %s", err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range rows {
		err := writer.Write(value)
		if err != nil {
			log.Printf("error in writing csv %s", err)
			return err
		}
	}
	return nil
}
