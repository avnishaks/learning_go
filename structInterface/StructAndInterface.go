package structInterface

import "fmt"

type gasEngine struct {
	mpg      uint8
	galllons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.galllons
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

func canMakeIt(e engine, miles uint8) {
	if miles >= e.milesLeft() {
		fmt.Println("You can able to complete your journey")
	} else {
		fmt.Println("You can able to cover your complete journey need to get some fuel")
	}
}

type engine interface {
	milesLeft() uint8
}

func Struct() {
	fmt.Println("Hello welcome to the world of Struct")
	var hpEngine gasEngine = gasEngine{10, 20}
	fmt.Printf("Below is quantity of the hpEngine : { mpg : %v } and { gallons: %v }\n", hpEngine.mpg, hpEngine.galllons)
	fmt.Printf("Miles amount left in the gasEngine is %v\n", hpEngine.milesLeft())

	// above point is all about the struct in go language

}

func Interface() {
	fmt.Println("Hello welcome to the world of Interface")
	var typeOfEngine electricEngine = electricEngine{10, 25}
	canMakeIt(typeOfEngine, 50)
}
