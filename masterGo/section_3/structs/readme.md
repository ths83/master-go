# Time To Practice: structs

## Intro

Custom, configurable types are a central aspect of flexibility in programming languages. The struct type is THE classic
customizable type. In its pure form, it is nothing but a collection of other types, but in conjunction with embedding
and methods–the topics of the upcoming lessions–, structs turn into a very versatile building block of apps and
libraries.

## Part 1: are these structs comparable?

Two instances of a struct type can be compared to each other as long as each of their fields is a comparable type.

For example, two instances of this struct…

```go
package main

type s struct {
	n int
}
```

…are comparable because an integer is a comparable type.

### Task: verify if instances of the following structs are comparable

Without running the code, can you tell for each of the following structs if two instances of that struct are comparable
to each other?

```go
package main

type s1 struct {
	n int
	b bool
}
s11 := s1{n: 4, b: true}
s12 := s1{n: 4, b: true}
fmt.Println(s11 == s12) // Does this line compile?
type s2 struct {
	r []rune
}
s21 := s2{r: []rune{'a', 'b', '🎵'}}
s22 := s2{r: []rune{'a', 'b', '🎶'}}
fmt.Println(s21 == s22) // Does this line compile?
type s3 struct {
	r [3]rune
}
s31 := s3{r: [3]rune{'a', 'b', '🎵'}}
s32 := s3{r: [3]rune{'a', 'b', '🎶'}}
fmt.Println(s31 == s32) // Does this line compile?
```

## Part 2: fun with empty structs

Can a struct contain no elements at all? It can, and then it is called an empty struct.

`type emptyStruct struct{}`

An empty struct takes zero bytes of memory, and so does a struct of empty structs. We can test this with a function from
the unsafe package, unsafe.Sizeof(<variable>).

```go
package main

type metaEmptyStruct struct {
	A struct{}
	B struct{}
}
fmt.Println("Size of emptyStruct:", unsafe.Sizeof(emptyStruct{}))
fmt.Println("Size of metaEmptyStruct:", unsafe.Sizeof(metaEmptyStruct{}))
Even a slice of empty structs consumes no memory except for the slice header.

type sliceOfNothings []struct{}
sOfN := make(sliceOfNothings, math.MaxInt64)
```

Even when instantiating the slice with a length and a capacity, it still only consumes memory for the header.

`voidSlice := make([]struct{}, 1000, 5000)`

The literal value of an empty struct looks a bit funny though:

`empty := struct{}{}`

The first pair of parens are the empty struct definition body, and the second pair belong to the empty literal value.

Now empty structs seem a bit superfluous–or can we maybe do something useful with them?

Well, there are indeed some use cases for empty structs. In the following two tasks, you will explore two of those use
cases.

### Task 1: define a Set type

Go has no native Set type. A set is a container for values where each value can only be stored once. An array of ints,
for example, can have multiple elements with value 37, but in a set, there can be only one such element.

Your task is to create a data structure that can hold elements of a given type, but each value exists only once within a
given instance of this data structure. Inserting the same value a second time does not change the contents of the set.

> Hint: Which of the available “container” types–arrays, slices, or maps–could be a suitable data structure?

> Another hint: Do not think too complicated! The solution is really nothing but a single type declaration:

`type ...`

and the declared type somehow includes a struct{} type.

### Task 2: create an integer iterator

Problem: The standard for loop for i := 0; i < 10; i++ looks soo old school. You want a sleeker way of iterating over
ints, like the Java folks do it. You want an iterator that returns the numbers from 0 to n, one number at a time.

Your task is to write a function iter(n int) <return value> that takes a number n and returns something that can be
passed to a range loop, like so:

```go
package main
for i := range iter(7) {
...
}
```

So the result of iter(7) must be something that a range loop can loop over n times. Of course, empty structs play a role
here. Remember the zero-sized data types discussed in the intro.

## Bonus task: a duplicate finder

As a bonus task, use a Set type for finding duplicate lines in a text file.

Write a small program that opens a file, reads the file line by line, and verifies if the line already exists in the
set. If so, it shall print the line; otherwise, it shall add the line to the set.

> Tip for scanning text files: Use a bufio.Scanner. The following functions and methods help with easy scanning:

os.Open(string) opens a file at path path and returns an os.File and an error value.
bufio.NewScanner(io.Reader) takes an io.Reader and creates a scanner. An os.File can act as an io.Reader!
The scanner has these methods (among others):
Scan() advances to the next token and returns a boolean true on success. Handy for use in a for loop!
Text() returns the current token as a string.
Try starting from a blank file! There is no template file included for this bonus task. See the extra file
“structbonussolution.go” for a possible solution.

Happy coding!