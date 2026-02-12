# Statement 2 – Go String Manager

This is a simple Go program that runs in a loop and lets the user manage a list of strings.

This version is an updated and cleaner implementation. Instead of putting all the logic in one place, the program now uses a switch statement and separate functions for each action (add, remove, check, and list). This makes the code easier to read and maintain.

You can:
Add a string
Remove a string
Check if a string exists
List all strings
Quit the program

The program keeps running until you type `quit`

## LOGIC FLOW
The logic is split into small, focused functions:
-> add
Takes user input and appends a new string to the list.

-> remove
Removes a string from the list if it exists and shows a message if it’s not found.

-> check
Searches the list to see whether a string already exists.

-> list
Prints all stored strings. If the list is empty, it tells the user.

The switch statement in main() is used to route each user command to the correct function, keeping the flow of the program simple and readable.

## How to Run It

Make sure Go is installed on your system.

Check with:

go version

1. Go to the Statement 2 folder
cd golang/statement-2

2. Run the program
go run main.go

## Example
------SELECT YOUR CHOICE OF COMMAND------
Choices: -> add, list, remove, check, quit
Choice: add
Input: Apple
Added!

Choice: check
Input: Apple
Your searched item exists

Choice: list
- Apple

Choice: remove
Input: Apple
Item is Removed

Choice: quit
Bye