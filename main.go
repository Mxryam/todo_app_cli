/*package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"



)


type User struct {
	ID       int
	Name	 string
	Email    string
	Password string
}

var userStorage []User
var authenticatedUser *User

func main(){
	fmt.Println("Hello to TODO app")

	command := flag.String("command", "no command", "command to run")
	flag.Parse()



	// get the password and email from the client
	fmt.Println("you must log in first!")
	scn := bufio.NewScanner(os.Stdin)
	fmt.Println("please enter the email:")
	scn.Scan()
	email := scn.Text()

	fmt.Println("please enter the password:")
	scn.Scan()
	password := scn.Text()


	for _, user := range userStorage {
		if user.Email == email && user.Password == password {

				authenticatedUser = &user

				break
		}
	}

	if authenticatedUser == nil {
		fmt.Println("the email or password is not correct!")
		return
	}

	// if there is a user record with corresponding data allow the user to continue

	for {
		runCommand(*command)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("please enter another command")
		scanner.Scan()
		*command = scanner.Text()
	}

}


	func runCommand(command string) {

		if command != "register_user" && command != "exit" && authenticatedUser == nil {
			login()
			}



		switch command {
		case "create-task":
			createTask()
		case "create-category":
			createCategory()
		case "register-user":
			registerUser()
		case "login":
			login()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("command is not valid", command)
}

}
func createTask() {
	//loggedInUser

	scanner := bufio.NewScanner(os.Stdin)
	var name, duedate, category string

	fmt.Println("please enter the task title")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter the task category")
	scanner.Scan()
	category = scanner.Text()

	fmt.Println("please enter the task due date")
	scanner.Scan()
	duedate = scanner.Text()

	fmt.Println("task:", name, category, duedate)
}

func createCategory() {
	//loggedInUser
	scanner := bufio.NewScanner(os.Stdin)
	var title, color string
	fmt.Println("please enter the category title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the category color")
	scanner.Scan()
	color = scanner.Text()
	fmt.Println("category", title, color)
}

func registerUser() {
	scanner := bufio.NewScanner(os.Stdin)
	var id, email, password string

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
		Email:    email,
		Password: password,
	}

	userStorage = append(userStorage, user)
}

func login() {
	fmt.Println("login process")
	scanner := bufio.NewScanner(os.Stdin)
	var email, password string

	fmt.Println("please enter email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the password")
	scanner.Scan()
	password = scanner.Text()
	fmt.Println("category", email, password)

	fmt.Println("user:", email, password)
}*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	
	"os"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}


type Task struct {
	ID		 int
	Title	 string
	Duedate  string
	Category string
	IsDone	 bool
	UserID   int
}

func (u User) print() {
	fmt.Println("User:", u.ID, u.Email, u.Name)
}

var userStorage []User
var authenticatedUser *User

var taskStorage []Task

func main() {
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

	if command != "register-user" && command!= "exit"&& authenticatedUser== nil {
			
			login()

			if authenticatedUser == nil{
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

	case "list-task" :
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

	fmt.Println("please enter the task category")
	scanner.Scan()
	category = scanner.Text()
	fmt.Println("please enter the task due date")
	scanner.Scan()
	duedate = scanner.Text()

	
		task := Task{
			ID:		 len(taskStorage)+1,
			Title:	 title,
			Duedate:  duedate,
			Category: category,
			IsDone:	 false,
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
		if user.Email == email && user.Password == password{
			
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


func listTask(){
	for _, task := range taskStorage {
		if task.UserID == authenticatedUser.ID {
			fmt.Println(task)
		}
	}
}
