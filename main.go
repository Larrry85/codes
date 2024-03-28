package main

import (
	//"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
)

////////////
// variables
////////////

type AirportLookup map[string]string

// Define flags for bonus features
var (
	colorOutput    bool
	reverseColumns bool
	airportNames   bool
)

////////////
// functions
////////////

func main() {

	// if too few, over or five, or -h: usage
	if len(os.Args) < 4 || len(os.Args) > 5 || (len(os.Args) == 5 && os.Args[4] == "-h") {
		displayUsage()
		return
	}

	// Parse command line arguments
	inputFile := os.Args[1]         // input.txt
	outputFile := os.Args[2]        //output.txt
	airportLookupFile := os.Args[3] //airport csv

	// Parse optional flags, bonus features
	if len(os.Args) > 4 {
		parseFlags(os.Args[4:])
	}

	// Load data, call loadAirportLookup()
	airportLookup, err := loadAirportLookup(airportLookupFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Call prettifyItinerary() with command line arguments
	err = prettifyItinerary(inputFile, outputFile, airportLookup)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Output generated successfully!")
	}
} // main() END
/////////////////////////////////////////////////////////////////////////////////////////

func displayUsage() { // instructions for using the program
	fmt.Println("itinerary usage:")
	fmt.Println("go run . ./input.txt ./output.txt ./airports_lookup.csv [flags]")
	fmt.Println("Flags:")
	fmt.Println("-color: Enable color in output")
	fmt.Println("-reverse: Enable reverse order of columns in CSV lookup")
	fmt.Println("-airport: Enable conversion of airport names to IATA/ICAO codes")
} // displayUsage() END
/////////////////////////////////////////////////////////////////////////////////////////

func parseFlags(flags []string) {
	for _, flag := range flags {
		switch flag {
		case "-color":
			colorOutput = true // bonus feature: colors
		case "-reverse":
			reverseColumns = true // bonus feature: different order
		case "-airport":
			airportNames = true //bonus feature: reversed
		default:
			fmt.Println("Unknown flag:", flag)
		}
	}
} //parseFlags() /
/////////////////////////////////////////////////////////////////////////////////////////

func loadAirportLookup(filename string) (AirportLookup, error) {
	airportLookup := make(AirportLookup)

	// open file, input.txt
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("Airport lookup not found")
	}
	defer file.Close() // closing the file!

	// read file
	reader := csv.NewReader(file)
	_, err = reader.Read() // Skip header row
	if err != nil {
		return nil, errors.New("Airport lookup malformed: missing header row")
	}

	// Parse the CSV data and populate the airport lookup map.
	// (if different order in columns, still works, column name as key)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		airportLookup[record[4]] = record[0] // iata column
		airportLookup[record[3]] = record[0] // icao column
	}

	return airportLookup, nil
} //loadAirportLookup() END
/////////////////////////////////////////////////////////////////////////////////////////

// function to format strings, string manipulating
func prettifyItinerary(inputFile, outputFile string, airportLookup AirportLookup) error {
	/* Check if the output file already exists
	if _, err := os.Stat(outputFile); err == nil {
		return errors.New("Output file already exists. Please choose a different file name.")
	}*/

	// read file
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return errors.New("Input not found")
	}

	// convert to string
	contentStr := string(content)

	// replace \v, \f, and \r with \n
	re := regexp.MustCompile(`\\[vfr]`)
	contentStr = re.ReplaceAllString(contentStr, "\n")

	// split content by \n
	lines := strings.Split(contentStr, "\n")

	// parseDateTime() in each line
	for i, line := range lines {
		lines[i] = parseDateTime(line)
	}

	// join the modified lines back into string
	contentStr = strings.Join(lines, "\n")

	// replace icao and iata with  airport names
	for code, name := range airportLookup {
		contentStr = strings.ReplaceAll(contentStr, "##"+code, "TEMP_PLACEHOLDER"+name)
		contentStr = strings.ReplaceAll(contentStr, "#"+code, name)
		contentStr = strings.ReplaceAll(contentStr, "TEMP_PLACEHOLDER"+name, name)
	}

	// if airport code not found
	re = regexp.MustCompile(`##?[A-Z]{3}`)
	missingCodes := re.FindAllString(contentStr, -1)
	if len(missingCodes) > 0 {
		//fmt.Println("Missing airport codes:", missingCodes)
		//return fmt.Errorf("Airport codes not found in the lookup: %v", missingCodes)
		return fmt.Errorf("Airport lookup malformed")
	}

	// calls parseDateTime()
	contentStr = parseDateTime(contentStr)

	// calls trimWhitespace()
	contentStr = trimWhiteSpace(contentStr)

	// call bonus features
	//bonus()
	//colorOutput
	//reverseColumns
	//airportNames

	// Another way to write into file
	/*
		file, err := os.OpenFile(outputFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = file.Write([]byte(contentStr))
		if err != nil {
			return err
		}*/
	file, err := os.Create(outputFile) // creates output txt file
	if err != nil {
		return err
	}
	defer file.Close() // closing file!

	_, err = file.WriteString(contentStr) // write content into output file
	if err != nil {
		return err
	}
	return nil
} // prettifyItinerary() END
/////////////////////////////////////////////////////////////////////////////////////////

func parseDateTime(text string) string {
	// regular expression pattern to match ISO 8601 date-time formats
	re := regexp.MustCompile(`(\w)\(([^)]+)\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		mode := parts[1]        // start of the date string: D, or T12 or T24
		datetimeStr := parts[2] // the rest: time

		// Parse the datetime string ????????????????????????
		datetime, err := time.Parse(time.RFC3339, datetimeStr)
		if err != nil {
			// Unable to parse, return unchanged
			return match
		}

		var formattedDateTime string
		switch mode { //???????????????????????
		case "D":
			// Format date as DD-Mmm-YYYY
			formattedDateTime = datetime.Format("02 Jan 2006")
		case "T12":
			// Format time in 12-hour format
			formattedDateTime = datetime.Format("03:04PM (-07:00)")
		case "T24":
			// Format time in 24-hour format
			formattedDateTime = datetime.Format("15:04 (-07:00)")
		}

		return formattedDateTime
	})
}

// parseDateTime() END
/////////////////////////////////////////////////////////////////////////////////////////

func trimWhiteSpace(text string) string {
	lines := strings.Split(text, "\n")
	var result []string
	prevEmpty := false // Flag to track if the previous line was empty
	for _, line := range lines {
		// Trim leading and trailing spaces from the line
		trimmedLine := strings.TrimSpace(line)
		// Check if the current line is empty
		empty := trimmedLine == ""
		// Check if the previous line was empty and the current line is also empty
		if prevEmpty && empty {
			// Skip adding the current line if the previous line was empty
			continue
		}
		// Keep track of the number of consecutive empty lines
		if empty && prevEmpty {
			// If there are more than two consecutive empty lines, skip adding this line
			continue
		}
		result = append(result, trimmedLine)
		prevEmpty = empty
	}
	return strings.Join(result, "\n")
}

//trimWhiteSpace() END
/////////////////////////////////////////////////////////////////////////////////////////

//////////////////
// Bonus functions
//////////////////

func bonus() {
	//colorOutput
	//reverseColumns
	//airportNames

	return
} //bonus() END
/////////////////////////////////////////////////////////////////////////////////////////

/* README

To run this Go code, save it in a file named itinerary_prettifier.go and execute it using:

go run . ./input.txt ./output.txt ./airport-lookup.csv [flags]
*/
