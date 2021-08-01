package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'isValid' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func isValid(s string) string {
	frequencyMap := make(map[string]int32)
	for _, r := range s {
		frequencyMap[string(r)]++
	}

	countMap := make(map[int32]int)
	for _, frequency := range frequencyMap {
		countMap[frequency]++
	}

	if len(countMap) == 1 {
		return "YES"
	}

	if len(countMap) > 2 {
		return "NO"
	}

	frequencies := make([]int32, 0, len(countMap))
	for k := range countMap {
		frequencies = append(frequencies, k)
	}
	maxFrequency, minFrequency := sortFrequency(frequencies[0], frequencies[1])

	if maxFrequency-minFrequency == 1 && countMap[maxFrequency] == 1 {
		return "YES"
	}
	if countMap[minFrequency] == 1 {
		return "YES"
	}

	return "NO"
}

func sortFrequency(frequencyA, frequencyB int32) (int32, int32) {
	if frequencyA < frequencyB {
		return frequencyB, frequencyA
	}
	return frequencyA, frequencyB
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	s := readLine(reader)

	result := isValid(s)

	fmt.Fprintf(writer, "%s\n", result)

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
