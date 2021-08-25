# jsony
Simple package to read JSON data from a file

## Why `jsony`

Reading and writing JSON files requires lines and lines of the same code (reading, marshaling, `if err != nil` ...). It is time to simplify this by having a few helper functions for these common tasks.

## Usage

### Importing

    import "github.com/wsw70/jsony"

### `Write()` JSON data to a file

    err := jsony.Write(data, filename)

where `data` is the `struct` (or other serializable type) you want to write to file `filename` (a `string`).


### `Read()` JSON data from a file

    err := jsony.Read(filename, &data)

where `filename` (a `string`) is the file you want to read JSON from and put them in `data` (a `struct`, or other container).
