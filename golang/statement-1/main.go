package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Simple Go Calculator")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter first number: ")
	scanner.Scan()
	n1, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("Enter second number: ")
	scanner.Scan()
	n2, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("Enter the choice of operation (+, -, *, /): ")
	scanner.Scan()
	choice := scanner.Text()

	switch choice {
	case "+":
		fmt.Printf("Result: %.2f", n1+n2)
	case "-":
		fmt.Printf("Result: %.2f", n1-n2)
	case "*":
		fmt.Printf("Result: %.2f", n1*n2)
	case "/":
		if n2 == 0 {
			fmt.Println("Error: division by zero is not allowed")
		} else {
			fmt.Printf("Result: %.2f", n1/n2)
		}
	default:
		fmt.Println("Invalid choice. Please choose correct option ")
	}
}
