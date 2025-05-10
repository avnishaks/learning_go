package main

import (
	"fmt"
	datastructurealgo "main/Data_Structure_Algo"
	"main/channels"
	functions "main/functions"
	"main/generics"
	"main/goroutines"
	"main/operations"
	"main/pointer"
	"main/structInterface"
	"unicode/utf8"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	// this file contains information around the const and variable
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Println("Hello and welcome, %s!", s)

	var instNumber int8 = 10
	instNumber = instNumber + 10
	fmt.Println("instNumber=", instNumber)

	// float number
	var floatNum float32 = 122882.29
	fmt.Println("floatNum=", floatNum)

	// typecasting and performing operations

	var floatNumber32Bit float32 = 10.7
	var intNumber232bit int32 = 2
	var result float32 = floatNumber32Bit + float32(intNumber232bit)
	fmt.Println("Post addition operation number is ", result)

	// finding the length of the string

	fmt.Println(utf8.RuneCountInString("abc334433"))

	fmt.Println('a')

	// boolean type value

	var myBoolean bool = false
	fmt.Println(!myBoolean)

	// Default value for all the int , unSignInt , float and ruin is set to 0
	// Default value for the String is empty string
	// Default value for the boolean is false

	// In go we define the variable directly no need to specify it explicitly

	myVarNext := "text in which variable is not define explicitly"
	fmt.Println(myVarNext)

	var1, var2 := 10, 12
	fmt.Println("The first variable is :", var1, "\nThe second variable is :", var2)

	const myConst string = "const value"
	//myConst ="new const value"
	fmt.Println(myConst)

	functions.Function()

	var num int = 10
	var den int = 2
	var divisionResult, remainder, err = functions.IntDivFunc(num, den)
	if err != nil {
		fmt.Println("Encountered error while processing division : ", err)
	} else {
		fmt.Printf("Division result is :%v \nRemainder is :%v\n", divisionResult, remainder)
	}

	operations.Array()

	operations.Map()

	operations.Loop()

	operations.StringRunesByte()

	structInterface.Struct()

	structInterface.Interface()

	pointer.Pointer()

	goroutines.Goroutines()

	channels.Channels()

	generics.Generics()

	datastructurealgo.DriverFunc()

	datastructurealgo.DriverFuncRabbit()

	datastructurealgo.RunCountSubArray()

	datastructurealgo.RunFindNumber()

	datastructurealgo.RunCodeForMinimumSumAfterReplacingZero()

}
