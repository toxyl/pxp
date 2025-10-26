package language

import (
	"strconv"
	"strings"

	"github.com/toxyl/math"
)

// @Name: load-csv
// @Desc: Loads a CSV file
// @Param:      path    	- -   -   	Path to the CSV file
// @Param:      sep     	- -   "\t"  The separator to split columns with
// @Param:      hasHeader         true  Whether the first row is a header row
// @Returns:    result  	- -   -   	A 2D slice with the data
func loadCSV(path, sep string, hasHeader bool) ([][]float64, error) {
	_, data, err := loadFile(path)
	if err != nil {
		return nil, err
	}
	sep = normalizeSep(sep)
	lines := strings.Split(string(data), "\n")
	if hasHeader {
		lines = lines[1:]
	}
	res := [][]float64{}
	for _, line := range lines {
		line = strings.TrimRight(line, "\r")
		if line == "" {
			continue
		}
		row := []float64{}
		for _, col := range strings.Split(line, sep) {
			if f, err := strconv.ParseFloat(strings.TrimSpace(col), 64); err == nil {
				row = append(row, f)
			} else {
				row = append(row, math.NaN[float64]())
			}
		}
		res = append(res, row)
	}
	return res, nil
}

// @Name: load-csv-column
// @Desc: Loads a column from a CSV file
// @Param:      path    	- -   -   	Path to the CSV file
// @Param:      index     	- -   -  	The index of the column to retrieve
// @Param:      sep     	- -   "\t"  The separator to split columns with
// @Param:      hasHeader         true  Whether the first row is a header row
// @Returns:    result  	- -   -   	A slice with the data
func loadCSVColumn(path string, index int, sep string, hasHeader bool) ([]float64, error) {
	_, data, err := loadFile(path)
	if err != nil {
		return nil, err
	}

	sep = normalizeSep(sep)
	lines := strings.Split(string(data), "\n")
	if hasHeader {
		lines = lines[1:]
	}
	res := []float64{}
	for _, line := range lines {
		line = strings.TrimRight(line, "\r")
		if line == "" {
			continue
		}
		val := math.NaN[float64]()
		parts := strings.Split(line, sep)
		if index >= 0 && index < len(parts) {
			if f, err := strconv.ParseFloat(strings.TrimSpace(parts[index]), 64); err == nil {
				val = f
			}
		}
		res = append(res, val)
	}
	return res, nil
}

// @Name: load-csv-row
// @Desc: Loads a row from a CSV file
// @Param:      path    	- -   -   	Path to the CSV file
// @Param:      index     	- -   -  	The index of the row to retrieve
// @Param:      sep     	- -   "\t"  The separator to split columns with
// @Param:      hasHeader         true  Whether the first row is a header row
// @Returns:    result  	- -   -   	A slice with the data
func loadCSVRow(path string, index int, sep string, hasHeader bool) ([]float64, error) {
	_, data, err := loadFile(path)
	if err != nil {
		return nil, err
	}

	sep = normalizeSep(sep)
	lines := strings.Split(string(data), "\n")
	if hasHeader {
		lines = lines[1:]
	}
	if index < 0 || index >= len(lines) {
		return nil, nil
	}
	row := strings.TrimRight(lines[index], "\r")
	res := []float64{}
	for _, col := range strings.Split(row, sep) {
		if f, err := strconv.ParseFloat(strings.TrimSpace(col), 64); err == nil {
			res = append(res, f)
		} else {
			res = append(res, math.NaN[float64]())
		}
	}

	return res, nil
}

func normalizeSep(sep string) string {
	switch sep {
	case "\\t":
		return "\t"
	case "\\n":
		return "\n"
	case "\\r":
		return "\r"
	}
	return sep
}
