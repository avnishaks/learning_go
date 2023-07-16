package main

import "fmt"

func main() {
	fmt.Println("Hey Geek Welcome to If else in GoLang")

	loginCount:=125
	var loginStatus string
	if loginCount>120{
        loginStatus="You are banned"
    }else if loginCount<120{
       loginStatus="You are allowed"
    }else{
		loginStatus="You are not allowed/banned"
	}

	fmt.Println(loginStatus)

	// Modulo Operator in GoLang
	if 10%2==0{
		fmt.Println("10 is even")
	} else{
		fmt.Println("10 is odd")
	}

	// unique syntax in GoLang for if else

	if value:=10; value%2==0{
		fmt.Println("Value is even")
	}else {
		fmt.Println("Value is odd")
	}

}
