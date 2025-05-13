package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// shadowVariable()
	//initializeIfStatement()
}

func shadowVariable() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
}

// The if statement can include an initialization statement before the conditional expression,
// which allows variables to be declared and assigned within the if statement's scope (only when needed).
// This makes the code more concise and readable, especially when dealing with error handling or temporary variables.
//
// Use this feature only to define new variables that are scoped to the if/else statement; anything else is confusing.
//
//	if initialization; conditional {
//	    // code to execute if the condition is true
//	}
//
//	if response, err := http.Get(url); err != nil {
//	    // handle error
//	    return err
//	} else {
//	    defer response.Body.Close()
//	}
func initializeIfStatement() {
	if n := rand.Intn(10); n < 4 {
		fmt.Println(n, "is too low")
	} else if n > 6 {
		fmt.Println(n, "is too high")
	} else {
		fmt.Println(n, " is just right")
	}
}

// The for statement is the only loop statement in Go.
// for <initialization>; <condition>; <post> { }
// You must use the short declaration operator (:=) to declare a new variable in the initialization statement.
// The condition expression is evaluated before each iteration and must return a boolean value.
// The post statement is executed at the end of each iteration.
// Go allows you to omit one or more of the three statements, but you must include the semicolons, unless it is only using the condition.
func completeForStatement() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

// The for statement can be used without an initialization or post statement (condition only), which is useful for iterating over a collection of data.
// It functions like a while loop in other languages.
func conditionOnlyForStatement() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

// Go's version of a for loop that loops forever.
// It is useful for creating an infinite loop that can be broken out of using a break statement.
func infiniteStatementLoop() {
	for {
		fmt.Println("Looking for a number betwen 5 and 10...")
		n := rand.Intn(15)

		if n <= 5 || n >= 10 {
			continue
		}

		fmt.Println("Found one:", n)
		break
	}
}

// Go encourages short if statement bodies; as lef-aligned as possible.
// Using a continue statement can remove nested code and make the code more readable.
// The statements line up, and the code is easier to read.
func useContinueForClarity() {
	for i := 0; i < 10; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}
}

// The for-range loop returns two values on each iteration: the index and the value of the element at that index.
// The idiomatic names for these values are i and v, when it's a slice or array, and k and v when it's a map.
// If you don't need to access the index or the key, you can use the blank identifier (_) to ignore it.
func rangeForSlice() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
}

// If you want the key, but not the value, you can just leave off the second value.
func rangeKeysOnly() {
	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	for k := range uniqueNames {
		fmt.Println(k)
	}
}

// A for-range loop can be used to iterate over a string.
// It iterates over the runes, not the bytes.  A run can be multiple bytes.
// The first variable holds the number of bytes from the beginning of the string, but the type of the second variable is rune.
func rangeString() {
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
	}
}

// Since Go 1.22, for-range loop provides a copy of the index and value on each iteration.
// If you want to modify the value, you need to use the index to access the value in the slice.
// A for loop doesn't properly handle multi-byte characters.
func forRangeCopy() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals {
		v *= 2
	}
	fmt.Println(evenVals)
}

// If you have a nested loop and want to exit or skip over the iterator of the outer loop, you can use a label.
func labeledForLooop() {
	samples := []string{"hello", "apple_π!"}
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
	}
	fmt.Println()
}

// You can declare a variable that is scoped to all cases in the switch statement.
// By default, cases in switch statements in Go don't fall through; use the fallthrough keyword to allow a case to fall through to the next case.
// You can use the break keyword to exit a switch statement early (but you probably shouldn't).
// You can separate multiple matches with commas.
// Any empty case means nothing happens.
func switchStatement() {
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}
}

// Unlike a regular switch, a blank switch allows you to write logical tests cases.
// If you find that you've written a blank swith where all cases are equality comparisons against the same variable,
// then replace with a regular switch.
func blankSwitch() {
	workds := []string{"hi", "salutations", "hello"}
	for _, word := range workds {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is exactly the right length.")
		}
	}
}

// Use the blank switch over if/else chains when you have multiple related cases.
func FizzBuzz() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
