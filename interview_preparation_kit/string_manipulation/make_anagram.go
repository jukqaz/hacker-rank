package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'makeAnagram' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING a
 *  2. STRING b
 */

func makeAnagram(a string, b string) int32 {
	count := int32(0)

	countMap := make(map[string]int32)
	for _, r := range a {
		countMap[string(r)]++
	}

	for _, r := range b {
		countMap[string(r)]--
	}

	for _, v := range countMap {
		count += abs(v)
	}

	return count
}

func abs(i int32) int32 {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	a := readLine(reader)

	b := readLine(reader)

	res := makeAnagram(a, b)

	fmt.Fprintf(writer, "%d\n", res)

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
