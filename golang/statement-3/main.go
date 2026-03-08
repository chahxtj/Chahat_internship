package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	next map[rune]*Node
	end  bool
}
type Trie struct {
	head *Node
}

func NewTrie() *Trie {
	return &Trie{head: &Node{next: make(map[rune]*Node)}}
}
func (t *Trie) Add(word string) {
	if word == "" {
		return
	}

	curr := t.head
	for _, ch := range word {
		if curr.next[ch] == nil {
			curr.next[ch] = &Node{next: make(map[rune]*Node)}
		}
		curr = curr.next[ch]
	}
	curr.end = true
}
func (t *Trie) Exists(word string) bool {
	if word == "" {
		return false
	}

	curr := t.head
	for _, ch := range word {
		if curr.next[ch] == nil {
			return false
		}
		curr = curr.next[ch]
	}
	return curr.end
}

func (t *Trie) Delete(word string) {
	if word == "" {
		return
	}

	curr := t.head
	for _, ch := range word {
		if curr.next[ch] == nil {
			return
		}
		curr = curr.next[ch]
	}
	curr.end = false
}

func (t *Trie) Print() {
	printFromNode(t.head, "")
}

func printFromNode(n *Node, prefix string) {
	if n.end {
		fmt.Println(prefix)
	}
	for ch, child := range n.next {
		printFromNode(child, prefix+string(ch))
	}
}

func main() {
	trie := NewTrie()
	sc := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("------STRING MANAGER------")
		fmt.Print("CHOOSE A COMMAND : add , remove , check , list , quit: ")

		if !sc.Scan() {
			fmt.Println("Input error, try ahain")
			return
		}

		choice := strings.ToLower(strings.TrimSpace(sc.Text()))

		switch choice {
		case "add":
			fmt.Print("Enter the word: ")
			if !sc.Scan() {
				fmt.Println("Input error, try again")
				return
			}

			word := strings.TrimSpace(sc.Text())
			if word == "" {
				fmt.Println("Empty input isn't allowed")
				continue
			}

			trie.Add(word)
			fmt.Println("Word added, Successfully!")

		case "check":
			fmt.Print("Enter the word: ")
			if !sc.Scan() {
				fmt.Println("Input error, try again")
				return
			}

			word := strings.TrimSpace(sc.Text())
			if trie.Exists(word) {
				fmt.Println("The word exists")
			} else {
				fmt.Println("The word is Not found")
			}

		case "remove":
			fmt.Print("Enter the word: ")
			if !sc.Scan() {
				fmt.Println("Input error, try again")
				return
			}
			word := strings.TrimSpace(sc.Text())

			trie.Delete(word)
			fmt.Println("The word is Removed")

		case "list":
			trie.Print()

		case "quit":
			fmt.Println("---QUITING---")
			return

		default:
			fmt.Println("Invalid command, TRY AGAIN")
		}
	}
}
