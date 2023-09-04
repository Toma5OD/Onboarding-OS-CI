
# Go Sorting Program

## Description

This Go program reads lines from an input file, sorts them in ascending alphabetical order, and then writes the sorted lines to an output file.

## How It Works

- Reads lines from `input.txt` using Go's `bufio` package.
- Sorts the lines using Go's `sort` package.
- Writes the sorted lines to `sorted.txt`.

## How to Use

1. Ensure you have Go installed on your machine.
2. Place an `input.txt` file in the same directory as the Go program. This file should contain the lines you want to sort.
3. Run the program with the command `go run main.go`.
4. Check the generated `sorted.txt` for the sorted lines.

## Testing

### Unit Testing with `main_test.go`

To run the unit tests, use the following command:

```
go test
```

This will run all test cases defined in `main_test.go` to ensure the sorting logic is working as expected.

### End-to-End Testing with `e2e_test.go`

To run end-to-end tests, use the following command:

```
go test -v
```

This will execute the `TestEndToEnd` function in `e2e_test.go`, which tests the program's functionality from start to finish.

### Benchmark Testing

To run the benchmark test, use the following command:

```
go test -bench .
```

This will execute the benchmark defined in `main_test.go`, measuring the performance of the sorting logic.

