package generics

import "fmt"

type gasEngine struct {
	gallons float32
	mpg     float32
}

type electricEngine struct {
	kwph   float32
	mpkwph float32
}

type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func Generics() {
	fmt.Printf("\nHello welcome to world of generics how the world of generics works in Go Programming Languages")

	var gasCar = car[gasEngine]{
		carMake:  "Honda",
		carModel: "Civic",
		engine: gasEngine{
			gallons: 10.8,
			mpg:     12.2,
		},
	}

	var electricCar = car[electricEngine]{
		carMake:  "Tesla",
		carModel: "Model 3",
		engine: electricEngine{
			kwph:   2.9,
			mpkwph: 12.9,
		},
	}

	fmt.Printf("The value of the gasCar is : %v\n", gasCar)

	fmt.Printf("The value of the electricCar is : %v\n", electricCar)
}
