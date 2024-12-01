package main

import (
	"bufio"
	"fmt"
	"os"
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

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	fl := map[int]int{}
	sl := map[int]int{}

	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "   ")

		one, err := strconv.Atoi(arr[0])
		if err != nil {
			panic(err)
		}

		if _, ok := fl[one]; ok {
			fl[one] += 1
		} else {
			fl[one] = 1
		}

		two, err := strconv.Atoi(arr[1])
		if err != nil {
			panic(err)
		}

		if _, ok := sl[two]; ok {
			sl[two] += 1
		} else {
			sl[two] = 1
		}
	}

	result := 0

	for key, val := range fl {
		if rp, ok := sl[key]; ok {
			result += (key * val * rp)
		}
	}

	fmt.Println(result)

}
