package main

import (
	"encoding/csv"
	"log"
	"os"
)

type AbbreviationRecord struct {
	Abbreviation string
	PlaceName    string
}

/*
 Given a 2 dimensional array of strings (say parsed from XML/JSON)
 return a list of abbreviation records
*/
func createAbbreviationRecords(input [][]string) []AbbreviationRecord {
	var abbrRecords []AbbreviationRecord

	for i, line := range input {
		if i > 0 { // omit header line
			var rec AbbreviationRecord
			for j, field := range line {
				if j == 0 {
					rec.Abbreviation = field
				} else if j == 1 {
					rec.PlaceName = field
				}
			}
			abbrRecords = append(abbrRecords, rec)
		}
	}
	return abbrRecords
}

/*
 Parse a list of abbreviation records from a .csv file.
*/
func parseAbbreviationsFromCsv(fileName string) []AbbreviationRecord {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// convert records to array of structs
	return createAbbreviationRecords(data)
}

/*
 Return, from a list of AbbreviationRecords, a list of strings of the abbreviations
 in that list.
*/
func abbreviationRecordsToStringLists(recs []AbbreviationRecord) ([]string, []string) {
	abbrStringList := make([]string, 0)
	placesStringList := make([]string, 0)
	for _, rec := range recs {
		abbrStringList = append(abbrStringList, rec.Abbreviation)
		placesStringList = append(placesStringList, rec.PlaceName)
	}

	return abbrStringList, placesStringList
}
