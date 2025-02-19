package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)



func main() {

	fmt.Println("End Line")

	command := flag.String("command", "no-command", "create a new task")
	flag.Parse()

	// input: name, 

	scanner := bufio.NewScanner(os.Stdin)
	if *command == "create-task" {
		var name, dueDate, category string

		// name

		fmt.Println("Please Enter The Task Title" )

		scanner.Scan()

		name = scanner.Text()

		// due date

		fmt.Println("Please Enter The Task Due Date" )

		scanner.Scan()

		dueDate = scanner.Text()

		// category

		fmt.Println("Please Enter The Task Category" )

		scanner.Scan()

		category = scanner.Text()
		fmt.Println(name, category, dueDate)
	}

	 
}