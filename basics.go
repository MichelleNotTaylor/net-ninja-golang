package main

import (
	"fmt"
)

func main() {
	// Loops
	// x := 0
	// for x < 5 {
	// 	fmt.Println("The value of x is:", x)
	// 	x++
	// }
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Value of i is:", i)
	// }

	names := []string{"Ilana", "Michelle"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println("Value of i is:", i)
	// }

	// How to cycle through a slice
	// for index, value := range names {
	// 	fmt.Printf("The value at index %v is %v\n", index, value)
	// }

	for _, value := range names {
		fmt.Printf("The value is %v\n", value)
	}

	fmt.Println(names)
	// The Standard Library
	// greeting := "Hello friends"
	// // fmt.Println(strings.Contains(greeting, "Hello"))
	// // fmt.Println(strings.ReplaceAll(greeting, "Hello", "What's up"))

	// fmt.Println(strings.Index(greeting, "ll"))
	// // Strings.Split returns a slice
	// fmt.Println(strings.Split(greeting, " "))

	// fmt.Println("Original string value = ", greeting)

	// ages := []int{27, 2, 30, 4, 11, 56}
	// sort.Ints(ages)
	// fmt.Println(ages)

	// index := sort.SearchInts(ages, 30)
	// fmt.Println(index)

	// names := []string{"Michelle", "Derrick", "Ilana"}
	// sort.Strings(names)
	// fmt.Println(names)
	// fmt.Println(sort.SearchStrings(names, "Derrick"))
	// String
	// var nameOne string = "Michelle"
	// var nameTwo = "Derrick"
	// var nameThree string

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameThree = "Ilana"
	// nameFour := "Layla"
	// fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	// // Integer
	// var ageOne int = 27
	// var ageTwo = 30
	// ageThree := 2

	// fmt.Println(ageOne, ageTwo, ageThree)

	// // Bits and Memory
	// var numOne int8 = 25
	// var numTwo int8 = -128
	// var numThree uint8 = 25

	// fmt.Println(numOne, numTwo, numThree)

	// var scoreOne float32 = 25.98
	// var scoreTwo float64 = 219387120847019201.2
	// scoreThree := 1.5

	// fmt.Println(scoreOne, scoreTwo, scoreThree)

	//Print
	// age := "27"
	// name := "Michelle"
	// fmt.Print("Hello, ")
	// fmt.Print("world!\n")
	// fmt.Print("New line\n")

	// //Println
	// fmt.Println("Hello ninjas!")
	// fmt.Println("Goodbye ninjas!")
	// fmt.Println("My age is,", age, "and my name is", name)

	// //Printf (formatted strings)
	// fmt.Printf("My age is %v and my name is %v \n", age, name)
	// fmt.Printf("My age is %q and my name is %q \n", age, name)
	// fmt.Printf("My age is of type %T \n", age)
	// fmt.Printf("You scored  %f points \n", 225.55)
	// fmt.Printf("You scored %0.1f points! \n", 225.55)

	// Sprintf (save formatted strings)
	// var str = fmt.Sprintf("You scored %0.1f points! \n", 225.55)
	// fmt.Println(str)

	// Arrays and Slices
	// var ages [3]int = [3]int{20, 25, 27}
	// var ages = [3]int{20, 25, 27}
	// names := [4]string{"Yoshi", "Mario", "Peach", "Bowser"}
	// names[1] = "Luigi"

	// fmt.Println(ages, len(ages))
	// fmt.Println(names, len(names))

	// // Slices (use arrays under the hood)
	// var scores = []int{100, 60, 40}
	// scores[2] = 25
	// scores = append(scores, 85)

	// fmt.Println(scores, len(scores))

	// Slice Ranges
	// rangeOne := names[1:3]
	// rangeTwo := names[2:]
	// rangeThree := names[:3]

	// fmt.Println(rangeOne)
	// fmt.Println(rangeTwo)
	// fmt.Println(rangeThree)

	// rangeOne = append(rangeOne, "Koopa")
	// fmt.Println(rangeOne)
}
