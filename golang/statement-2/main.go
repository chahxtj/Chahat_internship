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

		fmt.Println("-----Select your choice of command-----")
		fmt.Println("Choices: add, list, remove, check, quit")
		fmt.Print("Choice: ")
		scanner.Scan()

		ch := strings.ToLower(scanner.Text())

		if ch == "add" {
			fmt.Print("Input: ")
			scanner.Scan()
			items = append(items, scanner.Text())
			fmt.Println("Added!")

		} else if ch == "check" {
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

		} else if ch == "list" {
			if len(items) == 0 {
				fmt.Println("List is empty")
			}
			for _, v := range items {
				fmt.Println("-", v)
			}

		} else if ch == "remove" {
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
				fmt.Println("Not found.")
			}

		} else if ch == "quit" {
			fmt.Println("Bye")

			break

		} else {
			fmt.Println("Invalid choice : Please try again")
		}
	}
}
