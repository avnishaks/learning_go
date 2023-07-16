package main

import "fmt"

func main() {
	fmt.Println("Welcome to the world of the Structs in GoLang!")

	user_details:=User{"Avnish Kumar","avns@gmail.com","Abjaja",10}
	fmt.Println(user_details)
	fmt.Printf("User Details: %+v\n",user_details)

}

type User struct {
	Name string
	Email string
	Address string
	Age int
}
