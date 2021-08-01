package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'commonChild' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING s1
 *  2. STRING s2
 */

func commonChild(s1 string, s2 string) int32 {
	count := int32(0)

	start := 0
	for i := 0; i < len(s1); i++ {
		for j := start; j < len(s2); j++ {
			if !strings.Contains(s2, string(s1[i])) {
				break
			}
			if s1[i] == s2[j] {
				if start == j {
					break
				}
				start = j
				count++
				break
			}
		}
	}

	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	s1 := readLine(reader)

	s2 := readLine(reader)

	result := commonChild(s1, s2)

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
