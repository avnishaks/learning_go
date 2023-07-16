package main

import "fmt"

func main() {
	fmt.Println("Loops , Break and Continue in GoLang")

	// Loop in GoLang

	days:=[]string{ "Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	for i:=0; i<len(days); i++ {
		fmt.Println(days[i])
	}

	fmt.Println(" ")

	for idx:=range days {
		fmt.Println(days[idx])
	}

	fmt.Println(" ")

	for idx, day := range days {
        fmt.Printf("Value is %v Index is %v\n",idx, day)
    }

	// Break in GoLang

	fmt.Println(" ")
	fmt.Println("Break in GoLang")

	BreakValue:=1

	for BreakValue<10{
		if(BreakValue==6){
			break
		}
		fmt.Println("The value is ",BreakValue)
        BreakValue++
	}


	fmt.Println(" ")
	fmt.Println("Continue in GoLang")

	BreakValues:=1
	for BreakValues<10{
		if(BreakValues==6){
			BreakValues++
			continue;
		}
		fmt.Println("The value is ",BreakValues)
        BreakValues++
	}





}
