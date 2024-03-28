# My Project

This is my first task: itinerary.

## To Do:

<span style="color:red">1.</span> print usage if too few arguments or -h\n
<span style="color:red">2.</span> if csv file not exist -> error msg "Airport lookup not found"\n
<span style="color:red">3.</span> malformed csv -> error msg  "Airport lookup malformed"\n
<span style="color:red">4.</span> if input file not exist -> error msg "Input not found"\n
<span style="color:red">5.</span> no creating or overwriting the output.txt in case of error\n
<span style="color:red">6.</span> luo kopio csv tiedostosta, jolla testaat viallisen csv tiedoston\n
<span style="color:red">7.</span> test: if missing (name) column -> no creating output.txt\n
<span style="color:red">8.</span> test: create output.txt, test it is not overwritten\n
<span style="color:red">9.</span> luo testi input.txt, jolla testaat nämä kaikki kohdat:\n
<span style="color:red">10.</span> IATA, ICAO to airport names\n
<span style="color:red">11.</span> ISO to user friendly\n
<span style="color:red">12.</span> no whitespace, max two newlines\n
<span style="color:red">13.</span> Bonus: alkuvalikkoon näille vaihtoehto:\n
<span style="color:red">14.</span> IATA, ICAO  to city names -feature\n
<span style="color:red">15.</span> different order in columns -feature\n
<span style="color:red">16.</span> different order in columns -feature\n
<span style="color:red">17.</span> formatting, highlight, bold, underline, italic, colors etc -feature\n
<span style="color:red">18.</span> kaikille funktio tai muu, jolla ne menee päälle\n
<span style="color:red">19.</span> other features?\n

## Usage

```
Itinerary usage:
go run . ./input.txt ./output.txt ./airport-lookup.csv [flags]

Flags:
-color: Enable color in outpu
-reverse: Enable reverse order of columns in CSV lookup"
-airport: Enable conversion of airport names to IATA/ICAO codes

```

