package main

import "fmt"

// package level constants
const pi float32 = 3.14
const minutesInHour uint8 = 60
const hoursInDay uint = 24
const comp = 2 + 3i

func main() {
	fmt.Println("Package constants:")
	fmt.Println("---------------------------")
	fmt.Println(pi)
	fmt.Println(minutesInHour)
	fmt.Println(hoursInDay)
	fmt.Println(comp)

	x := explicitTypedConstants()
	y := inferredTypeConstants()
	fmt.Println(x)
	fmt.Println(y)
}

func explicitTypedConstants() bool {
	// this is how "variables" are declared and used
	var x int32 = 2147483647
	var y float32 = 4.0
	fmt.Printf("int values are %d", x)
	fmt.Printf("float values are %f", y)
	x--
	y += 4
	fmt.Printf("int values are %d", x)
	fmt.Printf("float values are %f", y)
	// this is how "constants" are declared and used
	const xx int32 = 2222
	const yy float32 = 2222.0
	fmt.Printf("int values are %d", xx)
	fmt.Printf("float values are %f", yy)

	const printingEnabled bool = true
	const myPi float32 = 3.14
	const myMinutesInHour uint8 = 60
	const myHoursInDay uint = 8
	const number int = 4
	const aBigNumber uint16 = 64353
	const aBigNegNumber int32 = -67353
	const aBiggerNumber float32 = 909938892178012238923475684 // too big for int64 or even uint64
	const aBiggerNegNumber float32 = -909938892178012238923475684
	const favoriteLanguage string = "Go Language"

	fmt.Println(printingEnabled)
	fmt.Println(myPi)
	fmt.Println(myMinutesInHour)
	fmt.Println(myHoursInDay)
	fmt.Println(number)
	fmt.Println(aBigNumber)
	fmt.Println(aBigNegNumber)
	fmt.Println(aBiggerNumber)
	fmt.Println(aBiggerNegNumber)
	fmt.Println(favoriteLanguage)
	return true
}

func inferredTypeConstants() bool {

	// this is how "variables" are declared and used without declaring type
	x := 2147483647
	y := 4.0
	fmt.Printf("int values are %d", x)
	fmt.Printf("float values are %f", y)
	x--
	y += 4
	fmt.Printf("int values are %d", x)
	fmt.Printf("float values are %f", y)
	// this is how "variables" are declared and used without declaring type
	var xx = 44
	var yy = -4.0
	fmt.Printf("int values are %d", xx)
	fmt.Printf("float values are %f", yy)
	xx--
	yy += 4
	// this is how "constants" are declared and used with declaring type
	const xxx int = 44
	const yyy float32 = -55.0
	fmt.Printf("int values are %d", xxx)
	fmt.Printf("float values are %f", yyy)
	return false
}
