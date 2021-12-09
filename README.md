# Milstone 1 Scanner Part
#### Team: GolblinSheep
#### Date: Oct 28, 2021

## Run .golite files
To run the scanner on samples, please
```
$ cd proj/golite/
$ go run golite.go -lex simple.golite
```

We have provided multiple samples in the golite folder, named as "simple1.golite", "simple2.golite", "simple3.golite" ... Change filename in step 2 and run on other golite files.

## Run testcases
In the meantime, we have written all above sample tests in /proj/scanner/scanner_test.go. To run all tests, please
```
$ cd proj/scanner/
$ go test -v
```

# Milstone 2 Parser Part
#### Team: GolblinSheep
#### Date: Nov 11, 2021

1. Parse and build AST

## Run .golite files
To run the parser on samples, please
```
$ cd proj/golite/
$ go run golite.go -ast simple.golite
```

## Run testcases
```
$ cd proj/parser/
$ go test -v
```

2. Semantic Analysis: Build up Symbol Table and Type Check

## Run testcases
```
$ cd proj/semantic/
$ go test -v
```