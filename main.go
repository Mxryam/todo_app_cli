package main

import (
	"bufio"
	"flag"
	"fmt"
	"strconv"
	"os"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Task struct {
	ID         int
	Title      string
	Duedate    string
	CategoryID int
	IsDone     bool
	UserID     int
}

type Category struct {
	ID     int
	Title  string
	Color  string
	UserID int
}

var categoryStorage []Category
var userStorage []User
var authenticatedUser *User
var taskStorage []Task

func main() {

	//load user storage from file
	loadUserStorageFromFile()

	fmt.Println("Hello to TODO app")
	command := flag.String("command", "no-command", "command to run")
	flag.Parse()

	for {
		runCommand(*command)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("please enter another command")
		scanner.Scan()
		*command = scanner.Text()
	}

}

func runCommand(command string) {

	if command != "register-user" && command != "exit" && authenticatedUser != nil {

		login()

		if authenticatedUser == nil {
			return
		}
	}

	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()

	case "list-task":
		listTask()

	case "login":
		login()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("command is not valid", command)
	}
}

func createTask() {

	scanner := bufio.NewScanner(os.Stdin)
	var title, duedate, category string

	fmt.Println("please enter the task title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the task category id")
	scanner.Scan()
	category = scanner.Text()

	categoryID, err := strconv.Atoi(category)
	if err != nil {
		fmt.Printf("category id is not valid integer, %v\n", err)
		return
	}

	isFound := false
	for _, c := range categoryStorage {
		if c.ID == categoryID && c.UserID == authenticatedUser.ID {
			isFound = true
			break
		}

	}
	if !isFound {
		fmt.Printf("category id is not valid integer, %v\n", err)
		return
	}

	fmt.Println("please enter the task due date")
	scanner.Scan()
	duedate = scanner.Text()

	task := Task{
		ID:       len(taskStorage) + 1,
		Title:    title,
		Duedate:  duedate,
		CategoryID: categoryID,
		IsDone:   false,
		UserID:   authenticatedUser.ID,
	}

	taskStorage = append(taskStorage, task)

}

func createCategory() {

	scanner := bufio.NewScanner(os.Stdin)
	var title, color string

	fmt.Println("please enter the category title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the category color")
	scanner.Scan()
	color = scanner.Text()
	fmt.Println("category", title, color)

	c := Category{
		ID:     len(categoryStorage) + 1,
		Title:  title,
		Color:  color,
		UserID: authenticatedUser.ID,
	}

	categoryStorage = append(categoryStorage, c)
}

func registerUser() {
	scanner := bufio.NewScanner(os.Stdin)
	var id, name, email, password string

	fmt.Println("please enter the name")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter the email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the password")
	scanner.Scan()
	password = scanner.Text()

	id = email

	fmt.Println("user:", id, email, password)

	user := User{
		ID:       len(userStorage) + 1,
		Name:     name,
		Email:    email,
		Password: password,
	}

	userStorage = append(userStorage, user)
	//sve user data in user.txt file
	//create user.txt file
	//write user record in the user.txt file

	path := "user.txt" 
 
	var file *os.File

	file, err :=os.OpenFile(path, os.O_APPEND |os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("path does not exist!", err)

		return

	}
	

	data :=fmt.Sprintf("id: %d, name: %s, email: %s,password: %s\n", user.ID, user.Name,
	 user.Email, user.Password)

	var b = []byte(data)

	numberOfWrittenBytes, wErr := file.Write(b)
	if wErr != nil {
		fmt.Printf("cant write to the file %v\n", wErr)

		return
	}


	fmt.Println("numberOfWrittenBytes", numberOfWrittenBytes)

	file.Write(b)
	file.Close()
}

func login() {
	scanner := bufio.NewScanner(os.Stdin)
	var email, password string

	fmt.Println("please enter email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the password")
	scanner.Scan()
	password = scanner.Text()

	//get the email and password from the client

	for _, user := range userStorage {
		if user.Email == email && user.Password == password {

			authenticatedUser = &user

			break
		}
	}

	if authenticatedUser == nil {
		fmt.Println("the email or password is not correct")

		return
	}

	fmt.Println("category", email, password)

	fmt.Println("user:", email, password)
}

func listTask() {
	for _, task := range taskStorage {
		if task.UserID == authenticatedUser.ID {
			fmt.Println(task)
		}
	}
}