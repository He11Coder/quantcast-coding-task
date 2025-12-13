# quantcast-coding-task

A command-line tool designed to process cookie logs and identify the most active cookies for a specific date.

## Usage

Package `main` is located within `cmd` directory. The source file `most_active_cookie.go` contains `main()` function.

### Build and Run

To start using the tool, you may build the project into a binary using `go build`:
```bash
go build -o cookie-tool ./cmd
```

Or use the following `make` command from `Makefile`:
```bash
make build
```

Then simply run the binary and specify arguments:
```bash
./cookie-tool -f cookie_log.csv -d 2018-12-09
```

### Arguments

* `-f`: path to a cookie log file (CSV format is supported)

* `-d`: date to search in YYYY-MM-DD format (UTC)

## Testing

### Run all tests

```bash
go test ./...
```

or

```bash
make test
```

### Check Code Coverage

To see the total coverage excluding the main entry point:

```bash
make coverage
```

## Project Layout

* `cmd/`: application entry point

* `pkg/`: library code

    * `argparse/`: command-line flag parsing and validation

    * `logparse/`: CSV reading and required data structure definitions

    * `stats/`: core business logic (finding the most frequent cookies)

    * `utils/`: helper functions

* `Makefile`: shortcut commands