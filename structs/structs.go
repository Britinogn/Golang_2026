package main

import (
	"fmt"
	"time"
)


type user struct{
	firstName string
	lastName string
	birthDate string
	age int
	createdAt time.Time
}

// (u user) receiver argument
func (u user) outputUserDetails(){
	

	fmt.Println(u.firstName, u.lastName, u.birthDate)
	//fmt.Println((*u).firstName, (*u).lastName, (*u).birthDate)
}

func (u *user) clearUserName(){
	u.firstName = ""
	u.lastName = ""
}


func  newUser(firstName, lastName, birthDate string ) *user {
	return &user{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")


	var appUser *user 

	appUser = newUser(userFirstName, userLastName, userBirthDate)
	//...

	//fmt.Println(firstName, lastName ,birthDate)
	//outputUserDetails(&appUser)
	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}