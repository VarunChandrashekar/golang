package main

import (
	"fmt"
	"os"
	"sport/logger"
	"sport/sport"

	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Hello, Modules!")
	sport.PrintHello("Varun")
	var outfile = "test"
	//Export the environment variable using viper. Optional in this case.
	filename := "in.txt"
	var format int
	_ = viper.BindEnv("FORMAT")

	convertType := viper.Get("FORMAT")
	fmt.Println(convertType)
	if !(convertType == "YAML" || convertType == "JSON") {
		//fmt.Println("No environment variable set for FORMAT. Please set an env variable")
		logger.Error("No environment variable set for FORMAT. Please set an env variable")
		os.Exit(2)
	}

	//Take Environment variable input from the user
	fmt.Println("Enter 1 for JSON and 2 for YAML")
	fmt.Scan(&format)
	if format > 2 {
		//fmt.Println("Invalid input. Enter a valid number")
		logger.Error("Invalid input. Enter a valid number")
		os.Exit(2)
	}

	processedlist := sport.Filestructconv(filename)
	//fmt.Println(processedlist)

	if format == 1 {
		os.Setenv("FORMAT", "JSON")
		//fmt.Println("Converting to JSON and storing to test.json")
		logger.Info("Converting to JSON and storing to test.json")
		inputslice := sport.DetailWriterJson{processedlist}
		sport.Outputfile(inputslice, outfile)

	}
	if format == 2 {
		os.Setenv("FORMAT", "YAML")
		//fmt.Println("Converting to YAML and storing to test.yaml")
		logger.Info("Converting to YAML and storing to test.yaml")
		inputslice := sport.DetailWriterYaml{processedlist}
		sport.Outputfile(inputslice, outfile)
	}
}
