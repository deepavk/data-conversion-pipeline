This utility supports the reading of csv files. The data can be written to different output formats such as json/xml. 

It can be extended to other formats by adding the specific logic for a different output formats. An implementation of the DataWriter can be added to the writers/ package to extend this utility to other formats.


To run the code:

```
$ cd pipeline
$ go build && ./pipeline

```

To include sorting in the output files:

```
$ cd pipeline 
$ go build && ./pipeline -sort=name -asc=true
$ go build && ./pipeline -sort=rating -asc=false

```

To run unit tests:

```
$ cd pipeline
$ go test -v ./...

```

Code structure 

```
pipeline/
|
|--- csv_utility    (utility to read and write csv files )
|
|--- csv_files      (contains the input file restaurants.csv )
|
|--- restaurants    (Co-ordinates the reading of csv files
|                    Initilizes output writers for different formats, 
|                    writing output files using writers which can be dynamically setup)
|
|--- models         (restaurant data defined here)
|
|--- output_files   (valid data from restaurants.csv is written to json and xml files and are stored here
|                   and invalid.csv if there are invalid restaurants in the csv file)
|
|--- writers        (contains DataWriter interface in writer.go to make the tool extensible to other formats, 
|                    json and xml writers are present here)
|
|---- main.go       (starts the processing of the csv)
|
                  
```
