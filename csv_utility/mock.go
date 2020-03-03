package csv_utility

type MockCsvHandler struct {
	RowsRead    [][]string
	outputFile  string
	RowsWritten [][]string
}

func GetMockCsv(rows [][]string) Csv {
	return &MockCsvHandler{RowsRead: rows}
}

func (m *MockCsvHandler) Read() ([][]string, error) {
	return m.RowsRead, nil
}

func (m *MockCsvHandler) Write(filename string, headers []string, rows [][]string) error {
	m.RowsWritten = rows
	return nil
}

func (m *MockCsvHandler) GetDataWritten() [][]string {
	return m.RowsWritten
}
