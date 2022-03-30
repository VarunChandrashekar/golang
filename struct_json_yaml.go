package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type Details struct {
	Name   string
	Age    int
	Gender string
	Sport  []string
	Height float32
	Weight int
}

func main() {
	//slice := make([]Details, 0, 3)
	var filename = "in.txt"
	slicecomplete := make([]Details, 0)
	//fmt.Println("Opening the in.txt")
	file, error := os.Open(filename)
	if error != nil {
		fmt.Println("file not found ", error)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var info Details
		//var text Details
		splitsquare := strings.Split(scanner.Text(), "[")
		splitsquare1 := splitsquare[0]
		splitsquare2 := splitsquare[1]
		//fmt.Println(splitsquare[1])
		//fmt.Println(splitsquare1)
		part1 := strings.Split(splitsquare1, ",")
		//fmt.Println(part1)

		info.Name = part1[0]
		//fmt.Println(info.name)
		info.Gender = part1[2]
		//fmt.Println(info.gender)
		ageint, err := strconv.Atoi(part1[1])
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		info.Age = ageint
		//fmt.Println(info.age)

		part2 := strings.Split(splitsquare2, "]")
		//fmt.Println(part2[1])
		//info.sport = part2[0]

		part22 := strings.Split(part2[0], ",")
		info.Sport = part22

		//fmt.Println(info.sport)

		part3 := strings.Split(part2[1], ",")
		//fmt.Println(part3[2])
		weightint, err := strconv.Atoi(part3[2])
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		info.Weight = weightint
		//fmt.Println(info.weight)

		heightfloat, err := strconv.ParseFloat(part3[1], 32)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		info.Height = float32(heightfloat)

		slicecomplete = append(slicecomplete, info)

	}
	//fmt.Println(slicecomplete)

	//fmt.Println(string(yamlData))

	//Export the environment variable using viper. Optional in this case.
	var format int
	_ = viper.BindEnv("FORMAT")

	convertType := viper.Get("FORMAT")
	if !(convertType == "YAML" || convertType == "JSON") {
		fmt.Println("No environment variable set for FORMAT")
		return
	}

	//Take Environment variable input from the user
	fmt.Println("Enter 1 for JSON and 2 for YAML")
	fmt.Scan(&format)
	if format > 2 {
		fmt.Println("Invalid input. Enter a valid number")
		os.Exit(2)
	}
	//fmt.Println(format)
	if format == 1 {
		os.Setenv("FORMAT", "JSON")
		jsonData, err := json.Marshal(&slicecomplete)

		if err != nil {
			panic("Unable to write data into the file")
		}
		jsonfileName := "test.json"
		err = ioutil.WriteFile(jsonfileName, jsonData, 0644)
		if err != nil {
			panic("Unable to write data into the file")
		}
	}
	if format == 2 {
		os.Setenv("FORMAT", "YAML")
		yamlData, err := yaml.Marshal(&slicecomplete)
		if err != nil {
			panic("Unable to write data into the file")
		}
		yamlfileName := "test.yaml"
		err = ioutil.WriteFile(yamlfileName, yamlData, 0644)
		if err != nil {
			panic("Unable to write data into the file")
		}
	}
	//for _, env := range os.Environ() {
	//	fmt.Println(env)
	//}

}
