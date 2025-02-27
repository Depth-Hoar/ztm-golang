// to run code run go run ./demo/readers
// next run 1 1 1
// then hit Ctrl d
// should return Sum: 3

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)

	sum := 0

	for {
		input, inputErr := scanner.ReadString(' ')

		num, convErr := strconv.Atoi(strings.TrimSpace(input))
		if convErr != nil {
			fmt.Println(convErr)
		} else {
			sum += num
		}

		if inputErr == io.EOF {
			break
		}
		if inputErr != nil {
			fmt.Println("Error reading Stdin:", inputErr)
		}
	}
	fmt.Printf("Sum: %v\n", sum)
}
