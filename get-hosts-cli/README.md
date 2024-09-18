# Get Hosts

## Summary

- [About](#about)
- [Usage](#usage)
- [Tests](#tests)

## About <a name = "about"></a>

Get Hosts is CLI application developed in Golang with the proposes to get a server name or IP address of a specific host inputted as a flag params in shell terminal

## Usage <a name = "usage"></a>

In the root folder, run this command in the terminal to get the IP address of a host

```shell
go run main.go ip --host amazon.com.br
```

In the root folder, run this command in the terminal to get the server name of a host

```shell
go run main.go server --host amazon.com.br
```

## Tests <a name = "tests"></a>

For run all the tests files in a project

```shell
go test ./...
```

For run all the tests in verbose mode

```shell
go test ./.. --v
```

For run all the tests and see the percentage of coverage

```shell
go test ./.. --cover
```

Generate the file containing the coverage information

```shell
go test ./... --coverprofile coverage.txt
```

Tool to read the coverage information file and output in terminal

```shell
go tool cover --func=coverage.txt
```

Generate an html file containing which lines are not covered and which are

```shell
go tool cover --func=coverage.txt
```
