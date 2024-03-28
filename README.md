# My Project

This is my first task: itinerary.

## To Do:

1. print usage if too few arguments or -h  
2. if csv file not exist -> error msg "Airport lookup not found"  
3. malformed csv -> error msg "Airport lookup malformed"  
4. if input file not exist -> error msg "Input not found"  
5. no creating or overwriting the output.txt in case of error  
6. luo kopio csv tiedostosta, jolla testaat viallisen csv tiedoston  
7. test: if missing (name) column -> no creating output.txt  
8. test: create output.txt, test it is not overwritten  
9. luo testi input.txt, jolla testaat nämä kaikki kohdat:  
10. IATA, ICAO to airport names  
11. ISO to user friendly  
12. no whitespace, max two newlines  
13. \v \f \r into: \n
14. Bonus: alkuvalikkoon näille vaihtoehto:  
15. IATA, ICAO to city names -feature  
16. different order in columns -feature  
17. formatting, highlight, bold, underline, italic, colors etc -feature  
18. kaikille funktio tai muu, jolla ne menee päälle  
19. other features?

## Usage

```
Itinerary usage:
go run . ./input.txt ./output.txt ./airport-lookup.csv [flags]

Flags:
-color: Enable color in outpu
-reverse: Enable reverse order of columns in CSV lookup"
-airport: Enable conversion of airport names to IATA/ICAO codes

```

