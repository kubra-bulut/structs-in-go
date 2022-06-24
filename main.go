package main

import "fmt"

//Defining the struct
type Engineer struct {
	Name     string
	Age      int
	Project1 Project
}

type Project struct {
	Name         string
	Priority     string
	Technologies []string
}

func main() {
	engineer := Engineer{"Kübra", 21,
		Project{"Beginner's Guide", "High", []string{"Go"}}}
	fmt.Println(engineer)
	//Prints with the field names
	fmt.Printf("%+v\n", engineer)
}
