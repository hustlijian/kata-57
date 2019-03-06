package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getFloat64(reader *bufio.Reader, prompt string) (float64, error) {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	number, err := strconv.ParseFloat(text, 64)
	if err != nil {
		number, err = getFloat64(reader, prompt)
	}
	if number < 0 {
		fmt.Println("please input a positive number")
		number, err = getFloat64(reader, prompt)
	}
	return number, err
}

func calculate(bill, rate float64) (tip, total float64) {
	tip = math.Ceil(bill*rate) / 100
	total = bill + tip
	return tip, total
}

func main() {
	var billAmount float64 = 0
	var tip float64 = 0
	var tipRate float64 = 0
	var total float64 = 0
	reader := bufio.NewReader(os.Stdin)
	billAmount, _ = getFloat64(reader, "What is the bill? ")
	tipRate, _ = getFloat64(reader, "What is the tip percentage? ")
	tip, total = calculate(billAmount, tipRate)
	fmt.Printf("The tip is $%.2f\n", tip)
	fmt.Printf("The total is $%.2f\n", total)
}
