package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkMAS(i, j int, data [][]string) int {
	if data[i][j] == "A" && i > 0 && j > 0 && i < len(data)-1 && j < len(data[i])-1 {
		if ((data[i-1][j-1] == "M" && data[i+1][j+1] == "S") ||
			(data[i-1][j-1] == "S" && data[i+1][j+1] == "M")) &&
			((data[i-1][j+1] == "M" && data[i+1][j-1] == "S") ||
				(data[i-1][j+1] == "S" && data[i+1][j-1] == "M")) {
			return 1
		}
	}
	return 0
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	data := [][]string{}

	for scanner.Scan() {
		text := scanner.Text()
		letters := strings.Split(text, "")
		data = append(data, letters)
	}

	result := 0

	for i, row := range data {
		for j, val := range row {
			if val == "A" {
				result += checkMAS(i, j, data)
			}
		}
	}

	fmt.Println(result)

}
