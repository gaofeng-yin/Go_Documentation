package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type Numbers struct {
	Numbers []Number `json:"Numbers"`
}

type Number struct {
	Name   string `json:"Name"`
	Number string `json:"Number"`
}

func main() {
	inputFile, err := os.Open("phoneNumber.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	fileData, _ := ioutil.ReadAll(inputFile)

	var numbers Numbers

	json.Unmarshal(fileData, &numbers)

	outputFile, err := os.Create("formatedNumbers.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer outputFile.Close()

	for _, value := range numbers.Numbers {
		writeToFile(value.Name, value.Number, outputFile)
	}

	fmt.Println("Success!")
}

func reformatPhoneNumber(phoneNumber string) string {
	if len(phoneNumber) < 2 {
		return "Unsupported number"
	}
	var cleanNumber string = removeSpacesAndDashes(phoneNumber)
	var cleanedNumber []string

	cleanedNumber = append(cleanedNumber, string(cleanNumber[0]))
	for i := 1; i < len(cleanNumber); i++ {
		if i%3 == 0 {
			cleanedNumber = append(cleanedNumber, "-")
		}
		cleanedNumber = append(cleanedNumber, string(cleanNumber[i]))
	}

	if len(cleanNumber)%3 == 1 {
		swapValue(&cleanedNumber[len(cleanedNumber)-3], &cleanedNumber[len(cleanedNumber)-2])
	}

	return strings.Join(cleanedNumber, "")
}

func removeSpacesAndDashes(number string) string {
	regularExression := regexp.MustCompile(`\d+`)
	return strings.Join(regularExression.FindAllString(number, -1), "")

}

func swapValue(val1 *string, val2 *string) {
	*val1, *val2 = *val2, *val1
}

func writeToFile(name string, number string, file *os.File) {
	file.WriteString("Name: " + name + "\n")
	file.WriteString("PhoneNumber: " + reformatPhoneNumber(number) + "\n")
	file.WriteString("\n")
}
