The in.txt file is the input file with contents to be processed.
The in.txt is processed into a slice using the Filestructconv() function in the sport package.
Based on the environment variables("YAML" or "JSON"), an output file (test.json or test.yaml) is created.
The make file can be tested by issuing the command "make build". It will create a directory named bin.
The bin directory contains the main executable.
The logger for the program is done with the help of uber/zap package present in the logger package
The testing program is the sport_test.go. It can be run by issuing the command "go test"

