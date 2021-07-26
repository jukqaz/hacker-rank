package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'activityNotifications' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY expenditure
 *  2. INTEGER d
 */

func activityNotifications(expenditure []int32, d int32) int32 {
	counting_array := make([]int, 201)
	for i := 0; int32(i) < d; i++ {
		counting_array[expenditure[i]]++
	}

	notification := int32(0)
	for i := d; int(i) < len(expenditure); i++ {
		median := findMedian(counting_array, d)

		if int32(median*2) <= expenditure[i] {
			notification++
		}

		counting_array[expenditure[i]]++
		counting_array[expenditure[i-int32(d)]]--
	}
	return notification
}

func findMedian(counting_array []int, d int32) float64 {
	median := .0

	countSum := int32(0)
	first := 0
	second := 0
	for i, count := range counting_array {
		countSum += int32(count)
		if d%2 == 0 {
			if first == 0 && countSum > d/2-1 {
				first = i
			}
			if second == 0 && countSum > d/2 {
				second = i

				median = float64(first+second) / 2.0
				break
			}
		} else {
			if countSum > d/2 {
				median = float64(i)
				break
			}
		}
	}
	return median
}

func main() {
	file, err := os.Open("input/input.txt")
	checkError(err)

	reader := bufio.NewReaderSize(file, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	dTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	expenditureTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var expenditure []int32

	for i := 0; i < int(n); i++ {
		expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
		checkError(err)
		expenditureItem := int32(expenditureItemTemp)
		expenditure = append(expenditure, expenditureItem)
	}

	result := activityNotifications(expenditure, d)

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
