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

var numCorrect = 0
var total = 0

func main() {

	flag.Parse()

	csvFile, _ := os.Open(*filePath)
	csvLines, _ := csv.NewReader(csvFile).ReadAll()

	fmt.Println("Welcome! Press ENTER to begin.")
	fmt.Scanf("\r")

	timerDuration := time.Duration(*interval) * time.Second

	timer := time.NewTimer(timerDuration)
	for _, line := range csvLines {

		fmt.Println(line[0])
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == line[1] {
			timer.Reset(timerDuration)
			numCorrect++
		}

		go func() {
			<-timer.C
			fmt.Println("Time expired.")
			returnResults(numCorrect, len(csvLines))
		}()

	}

	returnResults(numCorrect, len(csvLines))

}

func returnResults(numCorrect int, total int) {
	resultPercent := ((float32(numCorrect) / float32(total)) * 100)
	fmt.Println("Result:", numCorrect, "of", total)
	fmt.Println(int32(resultPercent), "%")
	os.Exit(0)
}
