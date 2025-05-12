# Go Control Structures

Answers to all questions are in the `main.go` file.

## Initialization with if statements

Write a Go function named `initializeIfStatement` that uses an if statement with an initialization clause to:
  - Generate a random integer n in the range 0–9.
  - Prints "n is too low" if n < 4.
  - Prints "n is too high" if n > 6.
  - Otherwise, prints "n is just right".

## The Four Ways of Looping

## The Complete `for` Statement
Write a function named `completeForStatement` that uses a for statement with an initialization, condition, and post clause to print the numbers 0 through 9.

Ensure that you:
- Declare the loop variable `i` in the initialization clause.
- Use i < 10 as the loop's condition.
- Increment `i` in the post statement.
- Print each value of `i` inside the loop.

## Condition-Only `for` Statement
Write a function named `conditionOnlyForStatement` that demonstrates a condition-only for loop (omitting the initialization and post statements) to print the numbers 0 through 9. Inside the function, you should:

- Initialize a counter variable `i` before the loop.
- Use a conditional-only loop that exits when i is less than 10.
- Inside the loop, print `i` and then increment it.

## The Infinite `for` Statement
Write a function named `infiniteStatementLoop` that uses an infinite for loop to repeatedly execute until a certain condition is met.

Inside the loop:
- Print the message "Looking for a number between 5 and 10...".
- Generate a random integer `n` in the range 0–14.
- Skip to the next iteration if `n` is not between 5 and 10 inclusive.
- When `n` is between 5 and 10 inclusive, print "Found one: n" and break out of the loop.

The solution should demonstrate the infinite loop construct, as well as the use of continue and break to control loop execution.

## Using `continue` for Clarity
Refactor the following nested-if FizzBuzz loop into a new function named `useContinueForClarity` that uses `continue` statements to eliminate nesting and improve readability. Your refactored version should check for "FizzBuzz" first, then "Fizz", then "Buzz", and finally print the number if none of those cases match.

```go
for i := 1; i <= 100; i++ {
  if i%3 == 0 {
    if i%5 == 0 {
      fmt.Println("FizzBuzz")
    } else {
      fmt.Println("Fizz")
    }
  } else if i%5 == 0 {
    fmt.Println("Buzz")
  } else {
    fmt.Println(i)
  }
}
```

## The `for` Range Statement
The `for-range` statement simplifies the process of iterating over slices, arrays, maps, and channels (or types based on these).

The orders of keys and values in a map are not guaranteed to be the same, which is intended to prevent Hash DOS attacks.  Exception - To help debug, formatting functions like `fmt.Printf` and `fmt.Sprintf` are guaranteed to use the same order for each invocation.

Write a Go function named `rangeForSlice` that:
- Declares a slice of integers `evenVals` with the values `{2, 4, 6, 8, 10, 12}`.
- Uses a for-range to iterate over the slice.
- Prints each element's index `i` and value `v` inside the loop.
- Bonus: Explain how you would modify the loop if you only needed the values (or only the indices) and wanted to ignore the other.


## The `for` Range Statement with Strings
Write a Go function named `rangeString` that:
- Defines a slice of strings containing "hello" and "apple_π!".
- Iterates over the slice with for-range.
- For each item, print the byte offset `i`, the rune value `r`, and its string conversion.
- Bonus: Explain why for-range on a string iterates runes (not bytes) and how it handles multi-byte Unicode characters.

## The `for` Range Value is a Copy
Write a Go function named `forRangeCopy` that:

- Declares a slice of integers: `evenVals := []int{2, 4, 6, 8, 10, 12}`
- Uses a for-range loop that trys to double each element by writing `v *= 2`.
- Prints the slice after the loop.
- Bonus: Describe how you would modify the loop to actually update each element in `evenVals`.

## Labeling `for` Loops
Write a Go function named `labeledForLooop` that demonstrates using a labeled `continue` to skip the rest of an inner loop and move on to the next iteration of an outer loop. The function should:
- Declare a slice of strings: `samples := []string{"hello", "apple_π!"}`
- Label the outer loop as `outer`:
  - Use a nested for-range inside the outer loop to iterate over each rune of the current string.
  - Print the byte offset `i`, the rune value `r`, and its string form.
  When the rune `r` equals 'l', use continue outer to skip processing the rest of the current string and jump to the next one.
  - After both loops complete, print a blank line.

## The `switch` Statement
Write a Go function named `switchStatement` that demonstrates a switch with an initialization statement and multiple case features. Your function should:

- Declare a slice of words: `words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}`
- Loop over each word in words.
- Inside the loop, use a switch to branch on the word's length:
  - For lengths 1, 2, 3, or 4, print: "*word* is a short word!"
  - For length 5, declare a new variable (e.g., wordLen) inside the case and print: "*word* is exactly the right length: *length*".
  - For lengths 6, 7, 8, or 9, leave the case body empty (do nothing).
  - In a default case, print: "*word*, is a long word!"
  - Use comma-separated values in cases.

## A case for using `break` in a `switch` Statement
There is one case where you might want to use `break` in a `switch` statement. In the following code, what would you need to change to make it work?

```go
for i := 0; i < 10; i++ {
  switch i {
  case 0, 2, 4,6:
    fmt.Println("Even")
  case 3:
    fmt.Println(i, "is divisible by 3 but not 2")
  case 7:
    fmt.Println("exit the loop!")
    break
  default:
    fmt.Println(i,"is boring")
  }
}
```

Add the following:
```go
loop:
  for i := 0; i < 10; i++ {

break loop
```

## Using a Blank Switch
A regular switch only allows you to compare for equality.

A blank switch doesn't specify the value you're comparing against, and allows you to use any boolean comparison for each case.

Write a Go function named `blankSwitch` that:
- Declares a slice of strings: `words := []string{"hi", "salutations", "hello"}`
- Loops over each word in the slice.
- Uses a blank switch with an initialization clause to perform logical tests on the word length:
  - if wordLen < 5, print: "*word*, is a short word!"
  - if wordLen > 10, print: "*word*, is a long word!"
  - Default, print: "*word*, is exactly the right length."
- Hint: A blank switch (no tag expression after switch) lets you write arbitrary boolean expressions in each case.

## Write Fizzbuzz using a Blank Switch
- Loop through 100 numbers.  For each number:
  - If divisible by 3 and 5, print "FizzBuzz".
  - If divisible by 3, print "Fizz".
  - If divisible by 5, print "Buzz".
  - Otherwise, print the number.