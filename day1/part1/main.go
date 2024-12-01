package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func ascSort(a, b int) int {
	return a - b
}

func main() {

	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	result := 0

	flList := []int{}
	slList := []int{}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "   ")

		one, err := strconv.Atoi(arr[0])
		if err != nil {
			panic(err)
		}
		flList = append(flList, one)

		two, err := strconv.Atoi(arr[1])
		if err != nil {
			panic(err)
		}
		slList = append(slList, two)
	}

	slices.SortStableFunc(flList, ascSort)
	slices.SortStableFunc(slList, ascSort)

	sort.Sort(sort.IntSlice(flList))
	sort.Sort(sort.IntSlice(slList))

	for i := 0; i < len(flList); i++ {
		one := flList[i]
		two := slList[i]

		if one > two {
			result += (one - two)
		} else {
			result += (two - one)
		}
	}

	fmt.Println(result)
}
