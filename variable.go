package main

import (
	f "fmt"
	"math"
	"time"
)

func main() {
	/*
		Variable usage in Golang
		* 	var x uint
			x = 100
		* 	var x uint = 100
		* 	var y int8 = 100
		* 	var z = 100 	=> Golang automatically infers the data type

		or you just can declare with (without var)
		x := 100			=> Golang infers the variable data type as an int

		Multiple variable declaration
		var x, y, z int = -1, 5, 9
		a, b := 7, 2
		fmt.Println(x, y, z, a, b)

		Multiple variable declaration with different type
		found, answer := true, "yes"
		var name, age = "Steve", 35
	*/

	var nomer int

	x := 100
	f.Println(x + 12)
	f.Println("I use f.Println instead of fmt.Println")
	f.Println(time.Now())
	f.Println(nomer)

	// int
	f.Printf("int min - max: %d - %d\n", math.MinInt, math.MaxInt)
	// int8
	f.Printf("int8 min - max: %d - %d\n", math.MinInt8, math.MaxInt8)
	// int16
	f.Printf("int16 min - max: %d - %d\n", math.MinInt16, math.MaxInt16)
	// int32
	f.Printf("int32 min - max: %d - %d\n", math.MinInt32, math.MaxInt32)
	// int64
	f.Printf("int64 min - max: %d - %d\n", math.MinInt64, math.MaxInt64)

	// unsigned
	// uint
	f.Printf("uint min - max: %d - %d\n", uint(0), uint(math.MaxUint))
	// uint8
	f.Printf("uint8 min - max: %d - %d\n", 0, math.MaxUint8)
	// uint16
	f.Printf("uint16 min - max: %d - %d\n", 0, math.MaxUint16)
	// uint32
	f.Printf("uint32 min - max: %d - %d\n", 0, math.MaxUint32)
	// uint64
	f.Printf("uint64 min - max: %d - %d\n", 0, uint64(math.MaxUint64))
}
