package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	items := []string{}

	for {

		fmt.Println("------SELECT YOUR CHOICE OF COMMAND------")
		fmt.Println("Choices: add, list, remove, check, quit")
		fmt.Print("Choice: ")
		scanner.Scan()

		choice := strings.ToLower(scanner.Text())

		if choice == "add" {
			fmt.Print("Input: ")
			scanner.Scan()
			items = append(items, scanner.Text())
			fmt.Println("Added!")

		} else if choice == "check" {
			fmt.Print("Input: ")
			scanner.Scan()
			input := scanner.Text()
			found := false

			for _, v := range items {
				if v == input {
					found = true
				}
			}

			if found {
				fmt.Println("Exists!")
			} else {

				fmt.Println("Does not exist")
			}

		} else if choice == "list" {
			if len(items) == 0 {
				fmt.Println("List empty")
			}
			for _, v := range items {
				fmt.Println("-", v)
			}

		} else if choice == "remove" {
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

			items = newItems

			if removed {
				fmt.Println("Removed!")
			} else {
				fmt.Println("Not found")
			}

		} else if choice == "quit" {
			fmt.Println("Bye")
			break

		} else {
			fmt.Println("Invalid choice: TRY AGAIN")
		}
	}
}
