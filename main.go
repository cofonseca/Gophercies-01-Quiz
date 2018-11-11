package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

var filePath = flag.String("Filepath", "", "The path to a quiz file in .csv format.")
var interval = flag.Int("Time", 30, "The amount of time allotted to answer each question.")

func main() {

	flag.Parse()

	numCorrect := 0
	total := 0

	csvFile, _ := os.Open(*filePath)
	reader := csv.NewReader(csvFile)
	csvLines, _ := reader.ReadAll()

	fmt.Println("Welcome! Press ENTER to begin.")
	fmt.Scanf("\r")

	timer := time.NewTimer(time.Duration(*interval) * time.Second)
	for _, line := range csvLines {

		fmt.Println(line[0])
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == line[1] {
			timer.Reset(time.Duration(*interval) * time.Second)
			numCorrect++
		}

		go func() {
			<-timer.C
			fmt.Println("Time expired.")
			os.Exit(1)
		}()

		total++

	}

	resultPercent := ((float32(numCorrect) / float32(total)) * 100)

	fmt.Println("Results:", numCorrect, "of", total)
	fmt.Println(int32(resultPercent), "%")

}
