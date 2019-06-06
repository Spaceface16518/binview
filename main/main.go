package main

import (
	"bufio"
	"flag"
	"github.com/Spaceface16518/binview"
	"os"
)

var perLine int
var filePath string

func init() {
	flag.IntVar(&perLine, "line", 4, "The number of bytes per line")
	flag.StringVar(&filePath, "path", "", "A path to a binary file. If left blank, input will be taken standard input")
}

func main() {
	flag.Parse()

	writer := bufio.NewWriter(os.Stdout)
	_, err := writer.WriteString(binview.BinView(filePath, perLine))
	if err != nil {
		panic(err)
	}
}
