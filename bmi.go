package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Please input your height")
		height_text, _ := reader.ReadString('\n')
		height_text = strings.TrimRight(height_text, "\n")
		if height_text == "end" || height_text == "quit" {
			break
		}
		height, err := strconv.ParseFloat(height_text, 64)
		if err != nil {
			continue
		}

		if height >= 10 {
			height /= 100
		}
		fmt.Println("Please input your weight")
		weight_text, _ := reader.ReadString('\n')
		weight_text = strings.TrimRight(weight_text, "\n")
		if weight_text == "end" || weight_text == "quit" {
			break
		}
		weight, err := strconv.ParseFloat(weight_text, 64)
		if err != nil {
			continue
		}

		bmi := weight / math.Pow(height, 2)
		fmt.Println("\n" + "Your BMI is " + fmt.Sprintln(bmi))
	}

}
