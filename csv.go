package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadCsv() ([]*CsvRecord, error) {

	f, err := os.Open(csvFile)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("unable to read csv file. Error: %s", err.Error()))
	}

	defer f.Close()

	csvReader := csv.NewReader(f)

	csvReader.Comma = ';'

	dateIndex := 1
	shopIndex := 5
	priceIndex := 8

	csvRecords := make([]*CsvRecord, 0)
	first := true

	for {
		record, err := csvReader.Read()

		// skip header
		if first {
			first = false
			continue
		}

		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, errors.New(fmt.Sprintf("error reading csv record. Error: %s\n", err.Error()))
		}

		var date time.Time
		var shop string
		var price float64

		for idx, value := range record {

			if idx == dateIndex {
				date, err = time.Parse("02.01.06", value)
				if err != nil {
					return nil, errors.New(fmt.Sprintf("unable to parse date %s. Error: %s\n", value, err.Error()))
				}
			}

			if idx == shopIndex {
				shop = value
			}

			if idx == priceIndex {

				value := strings.Replace(value, ",", ".", -1)

				parsed, err := strconv.ParseFloat(value, 64)
				if err != nil {
					return nil, errors.New(fmt.Sprintf("unable to parse price %s. Error: %s\n", value, err.Error()))
				}
				price = parsed * -1
			}

		}

		csvRecord := &CsvRecord{
			Date:  date,
			Shop:  shop,
			Price: price,
		}

		csvRecords = append(csvRecords, csvRecord)

	}

	return csvRecords, nil
}

type CsvRecord struct {
	Date     time.Time
	Shop     string
	Price    float64
	Category string
}
