package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/wcharczuk/go-chart/v2"
)

func main() {
	// reading from file
	content, err := ioutil.ReadFile("./input/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	data := string(content)

	// converting data to ascii number
	encodedText := stringToASCII(data)
	allString := strings.Join(encodedText, "")

	// counting occurance
	digitFrequency := make(map[rune]int)

	for _, char := range allString {
		if unicode.IsDigit(char) {
			digitFrequency[char]++
		}
	}

	fmt.Printf("%d", digitFrequency['2'])

	// Create a bar chart
	sbc := chart.BarChart{
		Title: "Benford's Law",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 30,
			},
		},
		Height:   512,
		BarWidth: 100,
		Bars: []chart.Value{
			{Value: float64(digitFrequency['1']), Label: "1"},
			{Value: float64(digitFrequency['2']), Label: "2"},
			{Value: float64(digitFrequency['3']), Label: "3"},
			{Value: float64(digitFrequency['4']), Label: "4"},
			{Value: float64(digitFrequency['5']), Label: "5"},
			{Value: float64(digitFrequency['6']), Label: "6"},
			{Value: float64(digitFrequency['7']), Label: "7"},
			{Value: float64(digitFrequency['8']), Label: "8"},
			{Value: float64(digitFrequency['9']), Label: "9"},
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	sbc.Render(chart.PNG, f)
}

func stringToASCII(input string) []string {
	var asciiValues []string

	for _, char := range input {
		asciiValues = append(asciiValues, strconv.Itoa(int(char)))
	}
	return asciiValues
}

type pair struct {
	digit     int
	frequency int
}
