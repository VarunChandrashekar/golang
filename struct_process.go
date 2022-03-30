package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Details struct {
	name   string
	age    int
	gender string
	sport  []string
	height float32
	weight int
}

func main() {
	//slice := make([]Details, 0, 3)
	var filename = "in.txt"
	//fmt.Println("Opening the in.txt")
	file, error := os.Open(filename)
	if error != nil {
		fmt.Println("file not found ", error)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		//var text Details
		splitsquare := strings.Split(scanner.Text(), "[")
		splitsquare1 := splitsquare[0]
		splitsquare2 := splitsquare[1]
		//fmt.Println(splitsquare[1])
		//fmt.Println(splitsquare1)
		part1 := strings.Split(splitsquare1, ",")
		//fmt.Println(part1)
		var info Details
		info.name = part1[0]
		//fmt.Println(info.name)
		info.gender = part1[2]
		//fmt.Println(info.gender)
		ageint, err := strconv.Atoi(part1[1])
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		info.age = ageint
		//fmt.Println(info.age)

		part2 := strings.Split(splitsquare2, "]")
		//fmt.Println(part2[1])
		//info.sport = part2[0]

		part22 := strings.Split(part2[0], ",")
		info.sport = part22

		//fmt.Println(info.sport)

		part3 := strings.Split(part2[1], ",")
		//fmt.Println(part3[2])
		weightint, err := strconv.Atoi(part3[2])
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		info.weight = weightint
		//fmt.Println(info.weight)

		heightfloat, err := strconv.ParseFloat(part3[1], 32)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		info.height = float32(heightfloat)
		//fmt.Println(info.height)

		fmt.Println(info.name, info.age, info.sport, info.height, info.weight)
	}
}
