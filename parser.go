package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func openFile(filename string) (file *os.File) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func closeFile(fp *os.File) {
	fp.Close()
}

func getCommandAndNumber(line string) (command string, number int) {
	command = string(line[0])
	number = 0
	if len(line) > 1 {
		number = int(line[2])
	}
	return
}

func getInstructionFromCommandAndNumber(command string, number int) (instruction string) {
	switch command {
	case "P":
		instruction = "Select Pen " + string(number)
	case "D":
		instruction = "Pen Down"
	case "U":
		instruction = "Pen Up"
	case "N":
		instruction = "Draw North " + string(number)
	case "S":
		instruction = "Draw South " + string(number)
	case "E":
		instruction = "Draw East " + string(number)
	case "W":
		instruction = "Draw West " + string(number)
	default:
		instruction = "Command not Recognized: " + command + " " + string(number)

	}
	return
}

func main() {

	file := openFile("test_input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		command, number := getCommandAndNumber(line)
		instruction := getInstructionFromCommandAndNumber(command, number)
		fmt.Println(instruction)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	closeFile(file)
}
