package binview

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func BinView(filePath string, perLine int) string {
	input, err := readInput(filePath)
	if err != nil {
		panic(err)
	}
	return formatBytes(input, perLine)
}

func formatBytes(input []byte, perLine int) string {
	// TODO: optimize this
	inputLen := len(input)

	var builder strings.Builder
	// 8 from byte plus 1 from the separator, times the number of bytes
	builder.Grow(9 * inputLen)

	for i := 0; i < inputLen; i++ {
		builder.WriteString(fmt.Sprintf("%08b", input[i]))
		if (i+1)%perLine == 0 { // FIXME: prove this algorithm
			builder.WriteRune('\n')
		} else {
			builder.WriteRune(' ')
		}
	}

	return builder.String()
}

func readInput(filePath string) (input []byte, err error) {
	if filePath == "" {
		input, err = ioutil.ReadAll(os.Stdin)
	} else {
		input, err = ioutil.ReadFile(filePath)
	}
	return input, err
}
