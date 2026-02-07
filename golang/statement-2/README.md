This is a small command-line program written in Go.
It lets you store and manage a list of strings using simple commands.

You can add items, list them, check if something exists, remove items, and quit the program.

##  Workflow

When you run the program, it keeps asking you for a command:
• add → add a new string
• list → show everything you’ve added
• remove → delete a string
• check → see if a string exists
• quit → exit the program

For commands like add, remove, and check, the program will ask you to type a string.

## How to Run

Make sure Go is installed:

go version


Then run:

cd golang/statement-2
go run main.go

## Example
Command: add
Input: hello
Added!

Command: check
Input: hello
Exists!

Command: list
hello

Command: remove
Input: hello
Removed!

Command: quit
Bye!
