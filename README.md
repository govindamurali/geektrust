[![Go](https://github.com/govindamurali/geektrust/actions/workflows/go.yml/badge.svg)](https://github.com/govindamurali/geektrust/actions/workflows/go.yml)

# Pre-requisites
* Go 1.17
* go tool

# How to run the code

There are scripts to execute the code. 

Use `run.sh` if you are Linux/Unix/macOS Operating systems and `run.bat` if you are on Windows. It takes input file path from console and executes the command
Internally both the scripts run the following commands 


 * `go build .` - This will build an executable by the name geektrust in the directory $GOPATH/src/geektrust besides the main.go file .
 * `go run main.go`


 # How to execute the unit tests

The unit tests are ran and the coverage is calculated using the library `gotestsum`. You might need to install to run these. 

We execute the unit tests by running the following command from the directory $GOPATH/src/geektrust

`gotestsum --hide-summary=all ./...`
We check for the coverage of unit tests by executing the following command. from the directory $GOPATH/src/geektrust

`gotestsum --hide-summary=all -- -coverprofile=cover.out ./...`
