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
