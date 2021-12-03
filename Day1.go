package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("Day1Input.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()

	var data []int

	for _, eachline := range txtlines {
		var num int
		num, err := strconv.Atoi(eachline)
		if err != nil {
			fmt.Println(err)
		}
		data = append(data, num)
	}
	//PART 1
	var count int = 0
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			count++
		}
	}

	fmt.Println(count)

	//PART 2
	var c int = 0
	for i := 0; i < len(data)-3; i++ {
		if (data[i] + data[i+1] + data[i+2]) < (data[i+1] + data[i+2] + data[i+3]) {
			c++
		}
	}
	fmt.Println(c)
}
