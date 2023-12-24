# FileCondenser
FileCondenser is a command-line utility written in Go that transforms the format of input files by collapsing line breaks. It's particularly useful for processing PEM files, where it replaces newline characters.

## Installation
To install FileCondenser, you need to have Go installed on your machine. Once you have Go installed, you can download and install FileCondenser by running:

```
git clone git@github.com:franklin83diaz/file-condenser.git
cd file-condenser
go build -o file-condenser main.go
```

## Usage
To use FileCondenser, you need to provide two arguments: the input file and the output file. Here's the syntax:

```
FileCondenser <input file> <output file>
```
For example, if you have a PEM file named private.pem and you want to output the result to private.pem.collapsed, you would run:
```
FileCondenser private.pem private.pem.collapsed

```



Contributing
Contributions to FileCondenser are welcome! Please submit a pull request or create an issue on GitHub.

License
FileCondenser is released under the MIT License. See the included LICENSE file for more details.