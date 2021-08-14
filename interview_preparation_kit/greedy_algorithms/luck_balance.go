package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'luckBalance' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. 2D_INTEGER_ARRAY contests
 */

func luckBalance(k int32, contests [][]int32) int32 {
	sum := int32(0)

	sort.Slice(contests, func(i, j int) bool {
		return contests[i][0] > contests[j][0]
	})

	for _, c := range contests {
		if c[1] == 0 {
			sum += c[0]
		} else {
			if k > 0 {
				sum += c[0]
				k--
			} else {
				sum -= c[0]
			}
		}
	}

	return sum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	var contests [][]int32
	for i := 0; i < int(n); i++ {
		contestsRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var contestsRow []int32
		for _, contestsRowItem := range contestsRowTemp {
			contestsItemTemp, err := strconv.ParseInt(contestsRowItem, 10, 64)
			checkError(err)
			contestsItem := int32(contestsItemTemp)
			contestsRow = append(contestsRow, contestsItem)
		}

		if len(contestsRow) != 2 {
			panic("Bad input")
		}

		contests = append(contests, contestsRow)
	}

	result := luckBalance(k, contests)

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
