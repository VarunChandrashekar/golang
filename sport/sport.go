package sport

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sport/logger"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

type Details struct {
	Name   string
	Age    int
	Gender string
	Sport  []string
	Height float32
	Weight int
}

type DetailWriter interface {
	WriteToFile(file string) error
}

type DetailWriterJson struct {
	ListJson []Details
}

type DetailWriterYaml struct {
	ListYaml []Details
}

func (a DetailWriterJson) WriteToFile(file string) error {
	jsonData, err := json.Marshal(a.ListJson)

	if err != nil {
		//panic("Unable to write data into the file")
		logger.Fatal("Unable to write data into the file")
	}
	jsonfileName := "test.json"
	err = ioutil.WriteFile(jsonfileName, jsonData, 0644)
	if err != nil {
		//panic("Unable to write data into the file")
		logger.Fatal("Unable to write data into the file")
	}
	return nil
}

func (a DetailWriterYaml) WriteToFile(file string) error {
	yamlData, err := yaml.Marshal(a.ListYaml)

	if err != nil {
		//panic("Unable to write data into the file")
		logger.Fatal("Unable to write data into the file")
	}
	yamlfileName := "test.yaml"
	err = ioutil.WriteFile(yamlfileName, yamlData, 0644)
	if err != nil {
		//panic("Unable to write data into the file")
		logger.Fatal("Unable to write data into the file")
	}
	return nil
}

func Outputfile(d DetailWriter, file string) {

	err := d.WriteToFile(file)
	if err != nil {
		logger.Error("In function:Outputfile. Error Writing to the file")
		return
	}
}
func PrintHello(name string) string {

	fmt.Println("hello ", name)
	fmt.Println("Hello, Modules! This is mypackage speaking!")
	return name
}

func Filestructconv(filename string) []Details {

	Slicecomplete := make([]Details, 0)
	//fmt.Println("Opening the in.txt")
	file, error := os.Open(filename)
	if error != nil {
		//fmt.Println("file not found ", error)
		logger.Error("file not found ")
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

		Slicecomplete = append(Slicecomplete, info)

	}
	return Slicecomplete
}
