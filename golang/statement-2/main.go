package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func add(scanner *bufio.Scanner, items []string) []string {
	fmt.Print("Input: ")
	scanner.Scan()
	input := scanner.Text()
	items = append(items, input)
	fmt.Println("Added!")
	return items
}

func check(scanner *bufio.Scanner, items []string) {
	fmt.Print("Input: ")
	scanner.Scan()
	input := scanner.Text()

	for _, v := range items {
		if v == input {
			fmt.Println("Your searched item exists")
			return
		}
	}
	fmt.Println("Searched item doesn't exist")
}

func list(items []string) {
	if len(items) == 0 {
		fmt.Println("List is empty")
		return
	}
	for _, v := range items {
		fmt.Println("-", v)
	}
}

func remove(scanner *bufio.Scanner, items []string) []string {
	fmt.Print("Input: ")
	scanner.Scan()
	input := scanner.Text()
	newItems := []string{}
	removed := false

	for _, v := range items {
		if v != input {
			newItems = append(newItems, v)
		} else {
			removed = true
		}
	}
	if removed {
		fmt.Println("Item is Removed")
	} else {
		fmt.Println("Not found")
	}
	return newItems
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	items := []string{}

	for {
		fmt.Println("------SELECT YOUR CHOICE OF COMMAND------")
		fmt.Println("Choices: -> add, list, remove, check, quit")
		fmt.Print("Choice: ")
		scanner.Scan()

		choice := strings.ToLower(scanner.Text())

		switch choice {
		case "add":
			items = add(scanner, items)

		case "check":
			check(scanner, items)

		case "list":
			list(items)

		case "remove":
			items = remove(scanner, items)

		case "quit":
			fmt.Println("Bye")
			return

		default:
			fmt.Println("Invalid choice: TRY AGAIN")
		}
	}
}
