package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
)

type User struct {
	ID int
	Email string
	Passkey string
}

var userStorage []User
  
func main() {

	fmt.Println("End Line")

	command := flag.String("command", "no-command", "create a new task")
	flag.Parse()

	// input: name,

	runCommand(*command)
	fmt.Println("User List", userStorage)
}

func runCommand(command string) {
	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "login":
		login()
	case "register":
		register()
	}
}

func createTask() {
	scanner := bufio.NewScanner(os.Stdin)
	var name, dueDate, category string
	
	// name
	
	fmt.Println("Please Enter The Task Title")
	
	scanner.Scan()
	
	name = scanner.Text()
	
	// due date
	
	fmt.Println("Please Enter The Task Due Date")
	
	scanner.Scan()
	
	dueDate = scanner.Text()
	
	// category
	
	fmt.Println("Please Enter The Task Category")
	
	scanner.Scan()
	
	category = scanner.Text()
	
	fmt.Println(name, category, dueDate)
}

func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)
	var title, color string
	
	fmt.Println("please enter category title")
	
	scanner.Scan()
	
	title = scanner.Text()
	
	fmt.Println("please enter category color")
	
	scanner.Scan()
	
	color = scanner.Text()
	
	fmt.Println(title, color)
}

func register() {
	scanner := bufio.NewScanner(os.Stdin)
	var id, email, password string

	fmt.Println("please enter user password")

	scanner.Scan()

	password = scanner.Text()

	fmt.Println("please enter user email")

	scanner.Scan()

	email = scanner.Text()

	id = email

	fmt.Println(id, email, password)

	user := User{
		ID: rand.Int(),
		Email: email,
		Passkey: password,
	}

	userStorage = append(userStorage, user)
}

func login() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter Your Username")
	scanner.Scan()

	var username = scanner.Text()

	fmt.Println("Enter your password")

	scanner.Scan()

	password := scanner.Text()

	fmt.Println(username, password)
}
