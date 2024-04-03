package main

import (
	"bufio"        // scanning csv: scanner := bufio.NewScanner(file)
	"encoding/csv" // reading, writing csv: reader := csv.NewReader(file)
	"errors"       // error handling
	"fmt"          // printing
	"io"           // csv
	"os"           // interface, arguments, file operations
	"regexp"       // regex patterns
	"strings"      // string manipulation: contentStr := strings.Join(content, "\n")
	"time"         // measuring and displaying time
)

////////////
// Variables
////////////

type AirportLookup map[string]string // key : value pairs

var line string // each line read from input

/////////////////////////////////////////////////////////////////////////////////////////

////////////
// Functions
////////////

func main() {
	// if too few or too many arguments, print usage
	if len(os.Args) < 4 || len(os.Args) > 5 {
		displayUsage()
		return // exits the program
	}

	// if only "go run . -h"
	if len(os.Args) > 1 && os.Args[1] == "-h" {
		displayUsage()
		return
	}

	// command line arguments
	inputFile := os.Args[1]         // input.txt
	outputFile := os.Args[2]        // output.txt
	airportLookupFile := os.Args[3] // airport-lookup.csv

	// load data from csv file, call loadAirportLookup()
	airportLookup, err := loadAirportLookup(airportLookupFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// call prettifyItinerary() with command line arguments
	err = prettifyItinerary(inputFile, outputFile, airportLookup)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Output generated successfully!") // everything ok!
	}
} // main() END
/////////////////////////////////////////////////////////////////////////////////////////

// instructions for using the program
func displayUsage() {
	fmt.Println("Itinerary usage:")
	fmt.Println("go run . <flag: -h> ./input.txt ./output.txt ./airport-lookup.csv")
} // displayUsage() END
/////////////////////////////////////////////////////////////////////////////////////////

// loads airport data from csv file into airportLookup map
func loadAirportLookup(filename string) (AirportLookup, error) {
	airportLookup := make(AirportLookup)

	// open csv file
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("airport lookup not found")
	}
	defer file.Close() // close csv!

	// create csv reader object
	reader := csv.NewReader(file)

	// read header row created by reader, returns slice of strings
	header, err := reader.Read()
	if err != nil {
		// Return an error if the error is not EOF (end of file)
		if err != io.EOF {
			return nil, err
		}
	}
	// standard header
	expectedHeader := []string{"name", "iso_country", "municipality", "icao_code", "iata_code", "coordinates"}

	if len(header) != len(expectedHeader) { // if not enough columns
		return nil, errors.New("airport lookup malformed: incorrect number of columns")
	}

	for i, col := range header { // iterate over header
		if col != expectedHeader[i] { // if header order do not match expected order
			return nil, errors.New("airport lookup malformed: incorrect header")
		}
	}

	for {
		record, err := reader.Read() // read every part of csv, return record
		if err == io.EOF {           // end of file error
			break
		}
		if err != nil { // other error
			return nil, err
		}
		// check every column, is empty cells or no data
		for _, col := range record {
			if col == "" || len(strings.TrimSpace(col)) == 0 || len(col) == 0 {
				return nil, errors.New("airport lookup malformed: missing or blank data")
			}
		}
		airportLookup[record[3]] = record[0] // ICAO code to airport name
		airportLookup[record[4]] = record[0] // IATA code to airport name
	}
	return airportLookup, nil
} //loadAirportLookup() END
/////////////////////////////////////////////////////////////////////////////////////////

// formats input data, writes prettified output file
func prettifyItinerary(inputFile, outputFile string, airportLookup AirportLookup) error {
	/* <- You can remove this
	// Check if output file already exists, won't let create same output file twice
	if _, err := os.Stat(outputFile); err == nil {
		return errors.New("output file already exists, please choose a different file name")
	} You can remove this -> */

	// read input file
	file, err := os.Open(inputFile)
	if err != nil {
		return errors.New("input not found")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) // read input line by line
	var content []string              // empty slice for modified lines

	// iterate over input file
	for scanner.Scan() {
		line = scanner.Text()              // every line in input
		pattern := `(#|##)([A-Z0-9]{3,4})` // regex pattern for airport codes
		// # or ##, three to four characters A-Z or 0-9
		re := regexp.MustCompile(pattern)

		// find airport codes in input file
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			// Extract the airport code and its prefix
			code := match[2]   // airpor code
			prefix := match[1] // # or ##

			if prefix == "#" || prefix == "##" {
				// if airport code exists in the lookup map
				if name, ok := airportLookup[code]; ok {
					// replace airport code with airport name
					line = strings.ReplaceAll(line, prefix+code, name)
				}
			}
		}

		// calls parseDateTime() with lines from input
		line = parseDateTime(line)

		// replace whitespace characters with newlines
		line = strings.ReplaceAll(line, ` (\v)`, "\n")
		line = strings.ReplaceAll(line, ` (\f)`, "\n")
		line = strings.ReplaceAll(line, ` (\r)`, "\n")

		// adds modified line to content
		content = append(content, line)
	}

	if err := scanner.Err(); err != nil {
		return err // return error if errors during scanning
	}

	// joins lines separated by newlines
	contentStr := strings.Join(content, "\n")
	// calls trimWhiteSpace() with content line
	contentStr = trimWhiteSpace(contentStr)

	// create output file if not exist, clear it if exist
	outputFileHandle, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outputFileHandle.Close() // closing file!

	_, err = outputFileHandle.WriteString(contentStr) // write content to output.txt
	if err != nil {
		return err
	}
	return nil
} // prettifyItinerary() END
/////////////////////////////////////////////////////////////////////////////////////////

// parses and formats date and time
func parseDateTime(input string) string {
	// regex patterns for D, T12 and T24 ISO times
	// D(YYYY-MM-DD T any)
	datePattern := `D\((\d{4}-\d{2}-\d{2})T.*\)`
	// T12(YYYY-MM-DD T HH:MM, (+/-, HH:MM or Z))
	time12Pattern := `T12\(\d{4}-\d{2}-\d{2}T(\d{2}:\d{2})([-+]\d{2}:\d{2}|Z)\)`
	// T24(YYYY-MM_DD T HH:MM, (+/-, HH:MM or Z))
	time24Pattern := `T24\(\d{4}-\d{2}-\d{2}T(\d{2}:\d{2})([-+]\d{2}:\d{2}|Z)\)`

	// patterns to variables
	dateRegex := regexp.MustCompile(datePattern)
	time12Regex := regexp.MustCompile(time12Pattern)
	time24Regex := regexp.MustCompile(time24Pattern)

	// find D time in input
	input = dateRegex.ReplaceAllStringFunc(input, func(match string) string {
		dTime := dateRegex.FindStringSubmatch(match)
		dateString := dTime[1]                            // first captured group: date
		date, err := time.Parse("2006-01-02", dateString) // input string to right format
		if err != nil {
			return match // leave unchanged if malformed
		}
		// date as a string
		return date.Format("02 Jan 2006")
	})

	// find T12 time in input
	input = time12Regex.ReplaceAllStringFunc(input, func(match string) string {
		t12time := time12Regex.FindStringSubmatch(match)
		timeString := t12time[1]   // first captured group: time
		offsetString := t12time[2] // second captured group: offset
		if offsetString == "Z" {   // if Z in T12 time
			offsetString = "+00:00" // change offset to Zulu time
		}
		time, err := time.Parse("15:04", timeString) // input string to right format
		if err != nil {
			return match // leave unchanged if malformed
		}
		// format with AM/PM
		timeStringFormatted := time.Format("03:04PM")
		// time and offset as a string
		return fmt.Sprintf("%s (%s)", timeStringFormatted, offsetString)
	})

	// find T24 time in input
	input = time24Regex.ReplaceAllStringFunc(input, func(match string) string {
		t24time := time24Regex.FindStringSubmatch(match)
		timeString := t24time[1]   // first captured group: time
		offsetString := t24time[2] // second captured group: offset
		if offsetString == "Z" {   // if Z in T24 time
			offsetString = "+00:00" // change offset to Zulu time
		}
		time, err := time.Parse("15:04", timeString) // input string to right format
		if err != nil {
			return match // leave unchanged if malformed
		}
		// format with AM/PM
		timeStringFormatted := time.Format("15:04")
		// time and offset as a string
		return fmt.Sprintf("%s (%s)", timeStringFormatted, offsetString)
	})

	return input // return modified input
} // parseDateTime() END
// ///////////////////////////////////////////////////////////////////////////////////////

// removes empty space from input, max two newlines in a row
func trimWhiteSpace(text string) string {
	lines := strings.Split(text, "\n") // split input into lines separated ny newline
	var result []string                // store non empty lines
	prevEmpty := false                 // if previous line was empty
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line) // remove empty space from start/end
		empty := trimmedLine == ""             // if current line is empty
		if empty && prevEmpty {                // if both previous and current lines are empty
			// If more than two consecutive empty lines, skip adding this line
			continue
		}
		result = append(result, trimmedLine) // trimmed line to result slice
		prevEmpty = empty                    // update prevEmpty
	}
	return strings.Join(result, "\n") // join non empty lines into string, separated by newlines
} //trimWhiteSpace() END
/////////////////////////////////////////////////////////////////////////////////////////
