package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func extractValues(str string) int {
	if strings.Contains(str, "mul(") {
		if before, after, found := strings.Cut(str, "mul("); found {
			if strings.Contains(after, "mul(") {
				return extractValues(after)
			}
			if before, found = strings.CutSuffix(after, ")"); found {
				nums := strings.Split(before, ",")
				if len(nums) == 2 {
					num1, err := strconv.Atoi(nums[0])
					if err != nil {
						return 0
					}
					num2, err := strconv.Atoi(nums[1])
					if err != nil {
						return 0
					}
					return (num1 * num2)
				}
			}
		}
	}
	return 0
}

func main() {

	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	result := 0

	for {
		str, err := reader.ReadString(')')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		values := extractValues(str)
		result += values
	}

	fmt.Println(result)
}
