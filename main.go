package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

// Variable definition
var naked, wantMore bool = false, true

// If initializer is present, the type can be omitted
var i, j, text = 1, 2, "my text"

// Constant declaration can be character, string, boolean, or numeric and cannot be declared using the := syntax
const pi = 3.14
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

// x int, y int // x, y int // are the same things
func add(x int, y int) int {
	return x + y
}

func minus(x, y int) int {
	return x - y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	// Naked return, return var are defined in the function definition
	x = sum * 4 / 9
	y = sum - x
	return
}

func funWithNumber() {
	// Print random number
	fmt.Println("Rand:", rand.Intn(100))
	fmt.Println("Rand:", rand.Int())
	fmt.Println("Add:", add(10, 15))
	fmt.Println("Minus:", minus(10, 15))
	numberOne, numbertwo := split(15)
	fmt.Println("Split:", numberOne, numbertwo)
}
func getType() {
	// Var declaration can be factored like the imports statement
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	// Print type and value of vars
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func zeroValue() {
	// Var without initializer are given a default zero value
	var (
		i int
		f float64
		b bool
		s string
	)
	fmt.Printf("Zero Value: %v %v %v %q\n", i, f, b, s)
}

func convert() {
	// Type conversion
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	// More simply
	j := 42
	g := float64(j)
	h := uint(g)
	// For useless vars
	_ = u
	_ = h
}

func forLoop() {
	sum := 0
	// Statements: Init: i := 0 Expression: i < 10 Post: i++
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// Init and Post Statements are optional, like this it's a while loop like
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("Sum:", sum)
}

func ifStatement(x float64) string {
	if x < 0 {
		return ifStatement(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func ifStatementWithShort(x, n, lim float64) float64 {
	// If statement can come with a short statement
	// The short statement is a var declaration with the if statement for scope
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// The short statement var can be used in the else statemet too
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	// Printn (with newline) Print without newline
	fmt.Println("Time:", time.Now())
	fmt.Println("TimeStamp in Nanoseconds:", time.Now().UTC().UnixNano())

	// Feed the seed to change the random number
	rand.Seed(time.Now().UTC().UnixNano())

	funWithNumber()

	// Print multiple function output
	// := is the short assignment statement, it works only in function, out we use var statement
	firstString, secondString := swap("culasec", "jean")
	fmt.Println("Swap:", firstString, secondString)

	getType()
	zeroValue()
	forLoop()
	fmt.Println("If St:", ifStatement(15))
	fmt.Println("If St Short:", ifStatementWithShort(3, 3, 20))
}
