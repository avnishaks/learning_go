package main

import "fmt"

func main() {
	fmt.Println("Hey Geek Welcome to Learning sections of Maps in Golang")

	languages := make(map[string]string)
	languages["GO"] = "Golang"
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"]="Ruby"

	fmt.Println("The list of all languages ",languages)
	fmt.Println("The full form of JS language ",languages["JS"])

	for key,value := range languages{
		fmt.Printf("For the Key %v mapped value is %v \n",key,value) 
	}

	delete(languages,"JS")
	fmt.Println("The list of all languages after deletion ",languages)


}
