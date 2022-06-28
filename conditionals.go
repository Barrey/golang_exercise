package main

import "fmt"

func main() {
	var number int
	number = 10

	if number > 5 {
		fmt.Println("number is greater than 5")
	} else {
		fmt.Println("number is less than 5")
	}

	number2 := 8

	if number2 > 5 {
		fmt.Println("number2 is greater than 5")
	} else if number2 < 5 {
		fmt.Println("number2 is less than 5")
	}

	number3 := 2

	switch number3 {
	case 1:
		fmt.Println("number3 is 1")
	case 2:
		fmt.Println("number3 is 2")
	case 3:
		fmt.Println("number3 is 3")
	case 4:
		fmt.Println("number3 is 4")
	}

	size := "XL"

	switch size {
	case "S":
		fmt.Println("size is S")
	case "M":
		fmt.Println("size is M")
	case "L":
		fmt.Println("size is L")
	case "XL":
		fmt.Println("size is XL")
	default:
		fmt.Println("size is unknown")
	}

	x := 8
	y := 9
	if product := x * y; product > 60 {
		fmt.Println(product, "  is greater than 60")
	}

	//shorthand conditionals with If
	if success := true; success {
		fmt.Println("We're rich!")
	} else {
		fmt.Println("Where did we go wrong?")
	}

	//shorthand conditionals with Switch
	amountStolen := 50000

	switch numOfThieves := 5; numOfThieves {
	case 1:
		fmt.Println("I get $", amountStolen)
	case 5:
		fmt.Println("Everyone gets $", amountStolen/5)
	}
}
