# jsony
Simple package to read JSON data from a file

[![Go Reference](https://pkg.go.dev/badge/github.com/wsw70/jsony.svg)](https://pkg.go.dev/github.com/wsw70/jsony)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
![GitHub branch checks state](https://img.shields.io/github/checks-status/wsw70/jsony/main)
## Why `jsony`

Reading and writing JSON files requires lines and lines of the same code (reading, marshaling, `if err != nil` ...). It is time to simplify this by having a few helper functions for these common tasks.

The idea of `jsony` was inspired by [`resty`](https://github.com/go-resty/resty), a wonderful package which immensly simplifies HTTP calls ❤️. 

## Usage

### Importing

    import "github.com/wsw70/jsony"

### `Write()` JSON data to a file

    err := jsony.Write(data, filename)

where `data` is the `struct` (or other serializable type) you want to write to file `filename` (a `string`).


### `Read()` JSON data from a file

    err := jsony.Read(filename, &data)

where `filename` (a `string`) is the file you want to read JSON from and put them in `data` (a `struct`, or other container).
