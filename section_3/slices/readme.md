# Time To Practice: Slices

## Task 1: Appending through a pointer

appendOne is a simple function that takes a pointer to a slice of ints and appends a 1 to this slice.

```go
package main

func appendOne(s *[]int) {
	*s = append(*s, 1)
}

func main() {
	s1 := []int{0, 0, 0, 0}
	// s1 := make([]int, 4, 8) // capacity is twice the initial size
	s2 := s1
	fmt.Printf("Before appendOne:\ns1: %v\ns2: %v\n", s1, s2)
	appendOne(&s1)
	fmt.Printf("After appendOne:\ns1: %v\ns2: %v\n", s1, s2)
	s1[0] = 2
	fmt.Printf("After changing s1:\ns1: %v\ns2: %v\n", s1, s2)

}
```

Does appendOne(&s2) also change s1?

Create a new .go file, add the above code, and run the file via go run.

Then change the first two lines of main() to:

```go
// s1 := []int{0, 0, 0, 0}
s1 := make([]int, 4, 8) // capacity is twice the initial size
```

Repeat the test.

If one of the results (or both) is different from what you expected, inspect the properties of s1 and s2 (address,
length and capacity) - they might give you a hint about what’s going on.

## Task 2: The slice is copied! Or is it?

Go has pass-by-value semantics, as we learned in the lecture about functions and pointers. The following code passes a
slice to the function changeSlice(). The slice is not passed as a pointer like in the first task, so apparently it is
passed by value, and the local variable s is therefore just a copy of s1 that is passed to changeSlice() in main().

```go
package main

func changeSlice1(s []int) {
	s[0] = 7
}

func main() {
	s1 := []int{1}
	fmt.Println("s1 before changeSlice1:", s1)
	changeSlice1(s1)
	fmt.Println("s1 after changeSlice1:", s1)
}
```

Still, assigning a new value to s[0] changes the value of s1[0] as well.

Why?

## Task 3: A little change

A slight modification of task #2: Instead of assigning a new value to s[0], we set s to a new slice of ints. Now s1 is
not affected by the change.

```go
package main

func changeSlice2(s []int) {
	s = []int{7}
}

func main() {
	s1 := []int{1}
	fmt.Println("s1 before changeSlice2:", s1)
	changeSlice2(s1)
	fmt.Println("s1 after changeSlice2:", s1)
}
```

Why does this behave differently than the code in task #2?

## Task 4: An append() gotcha

The built-in function append() is a convenient way of appending data to a slice without worrying about the capacity. But
in the code below it does not seem to work as intended:

```go
package main

func main() {
	src := []int{}
	src = append(src, 0)
	src = append(src, 1)
	src = append(src, 2)
	dest1 := append(src, 3)
	dest2 := append(src, 4)
	fmt.Println(src, dest1, dest2)
}
```

Run this code and inspect the output. What would you have expected? Why do you see this result instead?

Find the solutions and the answers in slicesolutions.go, attached to this lecture.
