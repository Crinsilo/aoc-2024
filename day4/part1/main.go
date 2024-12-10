package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findXMAS(i, j int, data [][]string) int {
	count := 0
	if data[i][j] == "X" {
		if j+3 < len(data[i]) && data[i][j+1] == "M" && data[i][j+2] == "A" && data[i][j+3] == "S" {
			count += 1
		}
		if j-3 >= 0 && data[i][j-1] == "M" && data[i][j-2] == "A" && data[i][j-3] == "S" {
			count += 1
		}
		if i-3 >= 0 && data[i-1][j] == "M" && data[i-2][j] == "A" && data[i-3][j] == "S" {
			count += 1
		}
		if i+3 < len(data) && data[i+1][j] == "M" && data[i+2][j] == "A" && data[i+3][j] == "S" {
			count += 1
		}
		if i-3 >= 0 && j-3 >= 0 && data[i-1][j-1] == "M" && data[i-2][j-2] == "A" && data[i-3][j-3] == "S" {
			count += 1
		}
		if i+3 < len(data) && j+3 < len(data[i]) && data[i+1][j+1] == "M" && data[i+2][j+2] == "A" && data[i+3][j+3] == "S" {
			count += 1
		}
		if i+3 < len(data) && j-3 >= 0 && data[i+1][j-1] == "M" && data[i+2][j-2] == "A" && data[i+3][j-3] == "S" {
			count += 1
		}
		if i-3 >= 0 && j+3 < len(data[i]) && data[i-1][j+1] == "M" && data[i-2][j+2] == "A" && data[i-3][j+3] == "S" {
			count += 1
		}
	}
	return count
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
			if val == "X" {
				result += findXMAS(i, j, data)

			}
		}
	}

	fmt.Println(result)
}
