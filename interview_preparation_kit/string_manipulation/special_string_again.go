package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the substrCount function below.
func substrCount(n int32, s string) int64 {
	count := int64(n)

	frequencies := make(map[string]int64)

	for i := int32(2); i <= n; i++ {
		for j := int32(0); j <= n-i; j++ {
			frequencies[s[j:j+i]]++
		}
	}

	for str, frequency := range frequencies {
		length := len(str)
		pivot := length / 2
		left := str[:pivot]
		rightIndex := 0
		if length%2 == 0 {
			rightIndex = pivot
		} else {
			rightIndex = pivot + 1
		}
		right := str[rightIndex:]
		if !isReverse(left, right) {
			continue
		}
		if left == right {
			count += (frequency)
		}
	}

	return count
}

func isReverse(left, right string) bool {
	for i := 0; i < len(left); i++ {
		if left[0] != right[len(left)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	s := readLine(reader)

	result := substrCount(n, s)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
