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

func main() {
	data := readDataFromFile()

	count := 0

	for i := 0; i < len(data); i++ {
		lineArr := data[i]
		safe := true
		desc := lineArr[0] > lineArr[1]
		for j := 0; j < (len(lineArr) - 1); j++ {
			if desc {
				if lineArr[j] <= lineArr[j+1] || (lineArr[j]-lineArr[j+1]) > 3 {
					safe = false
					break
				}
			} else {
				if lineArr[j] >= lineArr[j+1] || (lineArr[j+1]-lineArr[j]) > 3 {
					safe = false
					break
				}

			}
		}
		if safe {
			count += 1
		}
	}

	fmt.Println(count)
}
