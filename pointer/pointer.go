package pointer

import "fmt"

func square(answer [5]float64) [5]float64 {
	//for i := range answer {
	//	answer[i] = answer[i] * answer[i]
	//}

	for i := 0; i < len(answer); i++ {
		answer[i] = answer[i] * answer[i]
	}
	return answer
}

func Pointer() {
	fmt.Println("Hello , Welcome to the world of Pointer ! Enjoy the journey of Pointer !!")

	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value of the p point to is : %v\n", *p)
	fmt.Printf("The value of the i point to is : %v\n", i)
	*p = 20
	fmt.Printf("The value of the p point to is : %v\n", *p)
	fmt.Printf("The value of the p point address is : %v\n", p)

	// under the hood the reference of the both the variable slice and sliceCopy address location is same
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	fmt.Println(slice)
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)

	// lets take an example and print the square of the number here
	var number1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("The memory location of the number1 array is : %p\n", &number1)
	var finalAnswer [5]float64 = square(number1)
	fmt.Printf("\nThe result is: %v\n", finalAnswer)
}
