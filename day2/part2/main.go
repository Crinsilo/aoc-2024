package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readDataFromFile() [][]int {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	data := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()

		arr := strings.Split(line, " ")

		lineNums := []int{}

		for _, val := range arr {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}

			lineNums = append(lineNums, num)
		}

		data = append(data, lineNums)

	}

	return data
}

func safeOrNot(lineArr []int, exclude int) bool {
	safe := true

	a := 0
	b := 1
	if exclude == 0 {
		a = 1
		b = 2
	}
	if exclude == 1 {
		a = 0
		b = 2
	}
	desc := lineArr[a] > lineArr[b]
	for b < len(lineArr) {

		if desc {
			if lineArr[a] <= lineArr[b] || (lineArr[a]-lineArr[b]) > 3 {
				safe = false
				break
			}
		} else {
			if lineArr[a] >= lineArr[b] || (lineArr[b]-lineArr[a]) > 3 {
				safe = false
				break
			}

		}
		a++
		b++
		if exclude == a {
			a = exclude + 1
			b = exclude + 2
		}
		if exclude == b {
			b = exclude + 1
		}
	}
	return safe
}

func main() {
	data := readDataFromFile()

	count := 0

	for i := 0; i < len(data); i++ {
		lineArr := data[i]
		safe := false
		for j := 0; j < len(lineArr); j++ {
			safe = safe || safeOrNot(lineArr, j)
		}
		if safe {
			count += 1
		}
	}

	fmt.Println(count)
}
