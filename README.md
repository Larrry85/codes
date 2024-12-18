# My Project

This is my first task: Itinerary.
This tool converts airplane booking data into more user friendly form.

## Description

1. Program prints usage if given too few arguments or flag: -h.
2. Program needs three arguments saved in same directory: input file, output file and csv file.
3. If CSV file does not exist or is malformed error message is shown.
4. If input file does not exist error message is shown.
5. Program creates output file if it doesn't exist, and overwrites it every time program is run.
6. (Option in lines 129 - 133: program does not accept output file that already exist.)
7. Program converts IATA and ICAO codes into airport names.
8.  Program converts ISO times into user friendly form.
9.  Program formats output.txt in a way there are no whitespaces or too many newlines.

## Usage

Print usage with flag: -h
```
go run . -h ./input.txt ./output.txt ./airport-lookup.csv

```
How to use:
```
go run . ./input.txt ./output.txt ./airport-lookup.csv

```


## Coder

Laura Levistö 4/24