# Fruit Price Calculator
This program allows to print the average price of each fruit given an input file where each line contains a fruit name and a price.
The application support absolute and relative path as a command line parameter.
If any line of the given file is not in the right format ```fruit,price``` the application will return an error.

## Build the application
```
    cd cmd/
    go build -gcflags="-N -l" -o fruit-average-calculator
```

## Run the application
```
    ./fruit-average-calculator ../test/fruits.txt
```

## Example file content
banana,1.95

apple,9.92

apple,9.09

orange,2.46 

apple,8.65

apple,5.01

## Example file content
Average price of a apple is $6.65

Average price of a banana is $5.07

Average price of a orange is $4.63
