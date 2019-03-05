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
	var billAmount float64 = 0
	var tip float64 = 0
	var tipRate float64 = 0
	var total float64 = 0
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is the bill? ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	billAmount, _ = strconv.ParseFloat(text, 64)
	fmt.Print("What is the tip percentage? ")
	text, _ = reader.ReadString('\n')
	text = strings.TrimSpace(text)
	tipRate, _ = strconv.ParseFloat(text, 64)
	tip = math.Ceil(billAmount*tipRate) / 100
	total = billAmount + tip
	fmt.Printf("The tip is $%.2f\n", tip)
	fmt.Printf("The total is $%.2f\n", total)
}
