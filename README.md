# My Project

This is my first task: itinerary.

## To Do:

1. print usage if too few arguments or -h
2. if csv file not exist -> error msg "Airport lookup not found"
<span style="color:red">3.</span>malformed csv -> error msg  "Airport lookup malformed"
<span style="color:red">4.</span>if input file not exist -> error msg "Input not found"
<span style="color:red">5.</span>no creating or overwriting the output.txt in case of error
<span style="color:red">6.</span>luo kopio csv tiedostosta, jolla testaat viallisen csv tiedoston
<span style="color:red">7.</span>test: if missing (name) column -> no creating output.txt
<span style="color:red">8.</span>test: create output.txt, test it is not overwritten
<span style="color:red">9.</span> luo testi input.txt, jolla testaat nämä kaikki kohdat:
<span style="color:red">10.</span>IATA, ICAO to airport names
<span style="color:red">11.</span>ISO to user friendly
<span style="color:red">12.</span>no whitespace, max two newlines
<span style="color:red">13.</span>Bonus: alkuvalikkoon näille vaihtoehto:
<span style="color:red">14.</span>IATA, ICAO  to city names -feature
<span style="color:red">15.</span>different order in columns -feature
<span style="color:red">16.</span>different order in columns -feature
<span style="color:red">17.</span>formatting, highlight, bold, underline, italic, colors etc -feature
<span style="color:red">18.</span>kaikille funktio tai muu, jolla ne menee päälle
<span style="color:red">19.</span>other features?

## Usage

```
Itinerary usage:
go run . ./input.txt ./output.txt ./airport-lookup.csv [flags]

Flags:
-color: Enable color in outpu
-reverse: Enable reverse order of columns in CSV lookup"
-airport: Enable conversion of airport names to IATA/ICAO codes

```

