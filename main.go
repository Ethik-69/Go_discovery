package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"strings"
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

func switchcase() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s", os)
	}
}

func anotherswitch() {
	// Switch case true is a good way to do a long if then else statement
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning !")
	case t.Hour() < 1:
		fmt.Println("Good afternoon !")
	default:
		fmt.Println("Good evening !")
	}
}

func pointer() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

func array() {
	// Array hav fixed size
	fmt.Println("Array:")
	var a [2]string
	a[0] = "hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	fmt.Println("Slice:")
	// Slice are dynamically-sized
	// Slice does not store any data, it just describes section of an underlying array
	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	b := names[0:2]
	c := names[1:3]
	fmt.Println(b, c)

	c[0] = "XXX"
	fmt.Println(b, c)
	fmt.Println(names)
}

func moreSlice() {
	fmt.Println("More Slice:")
	// Zero(default) value of a slice
	var def []int
	fmt.Println(def, len(def), cap(def))
	if def == nil {
		fmt.Println("nil !")
	}

	// Next
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	fmt.Println(s[0])

	// default value for slice
	// var d [10]int
	// d[0:10]
	// d[:10]
	// d[0:]
	// d[:]
	// All of this are the same things

	t := []int{2, 3, 5, 7, 11, 13}

	t = t[1:4]
	fmt.Println(t)

	t = t[:2]
	fmt.Println(t)

	t = t[1:]
	fmt.Println(t)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func sliceLenghtAndCapacity() {
	fmt.Println("Slice Lenght and Capacity:")
	// Slice lenght is the number of items in it
	// Slice capacity is the number og items in the underlying (original) array
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero lenght
	s = s[:0]
	printSlice(s)

	// Extend its lenght
	s = s[:4]
	printSlice(s)

	// Drop its first two values
	s = s[2:]
	printSlice(s)
}

func sliceMake() {
	fmt.Println("Slice Make:")
	// Create a slice with make is the way to create dynamically-sized array
	// The make function allocates a zeroed array and returns a slice that refers to it
	a := make([]int, 5)
	printSlice(a)

	// make with a specified cap
	b := make([]int, 0, 5)
	printSlice(b)

	c := b[:2]
	printSlice(c)

	d := c[2:5]
	printSlice(d)
}

func sliceContainedDataType() {
	fmt.Println("Tic Tac Toe:")
	// Slice can contain any datatype
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func sliceAppend() {
	// If the backing array of s is too small to fit all the given values
	// a bigger array will be allocated.
	var s []int
	printSlice(s)

	// sppend works on nil slices
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed
	s = append(s, 1)
	printSlice(s)

	// We cen add more than one element at a time
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func hash() {
	fmt.Println("Struct:")
	type Vertex struct {
		x int
		y int
	}
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	fmt.Println(v.x)
	fmt.Println(v.y)

	// Pointer
	p := &v
	p.x = 1e9
	fmt.Println(v)

	v1 := Vertex{1, 2} // Hes type Vertex
	v2 := Vertex{x: 1} // y:0 is implicit
	v3 := Vertex{}     // x:0 and y:0
	p = &Vertex{1, 2}  // Has type *Vertex
	fmt.Println(v1, p, v2, v3)
}

func rangeuh() {
	fmt.Println("Range:")
	// Yeah, i can use the word range...
	// With range on a slice, two values are returned for each iteration
	// The first is the index
	// The second is a copy of the element
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// You can skip the index or the value
	// by assigning it to _
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func main() {
	// Defer wait for the end of the current function
	// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
	defer fmt.Println("Job Done =D")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	defer fmt.Println("Job Almost Done =D")
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
	switchcase()
	anotherswitch()
	pointer()
	array()
	moreSlice()
	sliceLenghtAndCapacity()
	sliceMake()
	sliceContainedDataType()
	sliceAppend()
	hash()
	rangeuh()
}
