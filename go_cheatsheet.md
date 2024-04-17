# Go Cheatsheet

## Variables

**Important**: Go has a garbage collector

### Initialization / Declaration

**var declaration**:

- var hello <type> = "hello"
  - either type or value can be omitted, not both
  - multiple variable declarations per line are possible
  - "Zero values" - **There are no uninitialised variables**:
    - number types: 0
    - string: ""
    - bool: false
    - interfaces / pointers: nil

**shorthand declaration**:

- hello := "hello"
  - multiple variable declarations per line are possible
  - no type, value can not be omitted

### Assignment

- "=" operator
- all variables are mutable
- variable type is immutable after declaration
- multiple assignments per line are possible

### Scope

**package level**:

- outside any function
- accesible in the entire package
- can only be declared with var variant

**block level**:

- inside a block (function, if, loop)
- commonly initialised with the shorthand variant
- implicit blocks: switch case, select

### Shadowing

- variables with the same name can be declared in an inner scope
- variables are not affected by each other
- using the name in the inner block will always refer to the inner variable
- *edge case*:

```go
var a, b int
c, a, b := 3, 1, 2
```

- this works because only the c is being declared, the others are assigned to

### Type Conversion

- T(v), example: i := 42; f := float64(i)

### Consts

- consts can only be declared using var
- you can declare multiple consts in a const () block
- numeric consts are high precision values

### Pointers

```go
x := 10

ptrToX := &x // pointer
*ptrToX++ // dereference

var anotherPtrToX *int
anotherPtrToX = &x
```

### Misc

- only names starting with a capital letter are exported

## Functions

**Example**:

```go
func add(x int, y int) int {
    return x + y
}

// alternative
func add(x, y int) int {
    return x +y
}

// multiple returns
func swap(x, y string) (string, string) {
    return y, x
}

// functions are values!
// functions can be passed around the same way as any other value
func compute(fn func(float64, float64) float64) float64 {
    return fn(4, 5)
}

// Closures
// a function value that references variables from outside its body
// each closure is bound to its outer variables
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}
```


## Flow Control

**IMPORTANT** Go has no while

### For

```go
// basic for loop
for i := 0; i < 10; i++ {
    
}

// without init and post
for ; i < 100 ; {

}

// OR
for i < 100 {

}

// infinite loop
for {

}
```

### If

```go
if x < 10 {

}

// with shorthand statement
// v is only available in the if block (and else block if one exists)
if v := 5; v < lim {

}
```

### Switch

```go
switch os := runtime.GOOS; os {
case "darwin":
    // code here
case "linux":
    // code here
default:
    // code here
}

// switch without condition, equal to switch true
// useful for long else-if blocks
switch {
case condition1:
    //
case condition2:
    //
default:
    //
}
```

### Defer

- you can defer the execution of a function by prepending the defer keyword
- multiple defers get pushed onto a stack, and executed last-in-first-out


## Data Structures

### Structs

```go
type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    v.X = 4
    fmt.Println(v.X)

    // struct pointers
    // no need to write dereference operator
    p := &v
    p.Y = 6
}
```

### Arrays & Slices

```go
var a = [5]int
a[0] = 1 
a[1] = 4

primes := [6]int{2, 3, 5, 7, 11, 13}

// Slices
// a slice doesn't store anything, it just points to the section of the underlying array
// a slice has a length and a capacity
// length: the number of elements contained in the slice
// capacity: total number of elements in the underlying array
// the length cannot exceed the capacity
var s []int = primes[1:4]

// give the slice zero length
s := s[:0]

// extend it's length
s := s[:4]

// drop first two values
s := s[2:]

// omitting bounds uses defaults
// lower bound is 0, upper bound is the length of the slice
s := s[:]

// Slice literal
// Creates the underlying array and points to it
q := []int{2, 3, 5, 7}

// using the make function

// length and capacity are the same
a := make([]int, 5)

// length = 1, capacity = 5
b := make([]int, 1, 5)

// a slice of slices, equal to a two dimensional array in other languages
var x [][]string

// appending to a slice
// if the underlying array is too small, a new array will be allocated
var s := []int

s = append(s, 0)
s = append(s, 1, 2, 3)

// iterating over a slice using the range form
// range returns the index and a copy of the element
// index or value can be omitted using _
// if you just want the index, only use one variable
for i, v := range s {
    fmt.Printf("%d, %d\n", i, v)
}
```

### Maps

```go
var m map[string]float64

m = make(map[string]float64)

m["Nice"] = 6.9

// map literals
var m = map[string]float64 {
    "First": 1.2,
    "Second": 5.1
}


// mutating maps

// inserting / updating an element
m[key] = elem

// get an element
elem = m[key]

// delete an element
delete(m, key)

// test that value is present
elem, ok := m[key]
```
## Methods & Interfaces

### Methods

There are no classes, but methods can be defined on types.

```go
type Vertex struct {
    X, Y float64
}

// function
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    v := Vertex{4, 5}
    fmt.Println(Abs(v))
}

// method equivalent, where v is called a receiver
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
    v := Vertex{4, 5}
    fmt.Println(v.Abs())
}

// methods can be declared on any type that is in the same package as the method
// the type can not be a pointer
type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

func main() {
    f := MyFloat(-5.1)
    fmt.Println(f.Abs())
}

// pointer receivers
// if you want to modify values of the receiver, you need to pass it as a pointer
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

//NOTE: On values and pointers
// functions need to take in a value if it expects a value, same for a pointer
// methods don't care about the receiver they get, they just interpret it to whatever is defined in the method header
// generally just go with a pointer receiver, to be able to modify values and for efficency
// IMPORTANT: all methods on a type should either have a value or a pointer receiver

```

### Interfaces

**A type that is defined as a set of method signatures**

```go
type Abser interface {
    Abs() float64
}

// interfaces are implemented by method, not explictly
type I interface {
    M()
}

type T struct {
    S string
}

func (t T) M() {
    fmt.Println(t.S)
}

// empty interfaces may hold values of any type, used when you deal with an unknown type
var i interface{}

i = 42
i = "hello"

// type assertions
// used to get the underlying value of type T from an interface
var i interface{} = "hello"

s := i.(string) // s = "hello"

s, ok := i.(string) // s = "hello", ok = true

f, ok := i.(float64) // s = 0, ok = false

f := i.(float64) // panics

// type switches
// useful for comparing an interface against multiple types
switch v := i.(type) {
    case int:
        //
    case string:
        //
    default:
        //
}

// the Stringer interface
// similar to implementing a toString method in other languages (I think?)
type Person struct {
    Name string
    Age int
}

func (p Person) String() string {
    return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
```

## Errors

**The error type is a built-in interface**

```go
type error interface {
    Error() string
}

// common error handling
i, err := strconv.Atoi("42")

if err != nil {
    // print error message
}

// implementing the error method on a struct
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	
	z := 1.0
	
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2*z)
	}
	
	return z, nil
}
```
## Readers

**io.Reader is an interface of the io package**

```go
// method definition
// n is the number of bytes populated
// err gets assigned an io.EOF error when the stream ends
func (T) Read(b []byte) (n int, err error)

// example usage
func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
```
