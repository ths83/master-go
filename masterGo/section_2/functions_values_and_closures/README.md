# Time To Practice: Function Values and Closures

## Task 1: Two closures

### Intro

As you learned in the lecture about function values and closures, a closure can reference the outer function’s variables
even after the outer function has terminated.

But what happens if the outer function generates and returns two closures?

Do they access the same outer variables, or does each of them get its own copy?

### Your task

Copy and paste the code from below into your editor and name the file “closures.go”.

Add code to newClosure so that it returns two closures. The first one is of type func(), the second one of type func()
int.

Both closures shall modify an integer variable defined in the outer function as follows:

- The first closure shall just set the outer variable to 5. It returns nothing.

- The second closure shall multiply the outer variable by 7 and return the value.

main() calls newClosure to create the new closures, and then calls both closures and prints out the result.

Run the code via

```shell
go run closures.go
```

and inspect the outcome - is this what you have expected?

```go
package main

import "fmt"

func newClosures() (func(), func() int) {
	a := 0
	// Your code here!
}
func main() {
	f1, f2 := newClosures()
	f1()      // sets "a" to 5
	n := f2() // multiplies "a" by 7 - is f2's internal value of "a" 0 or 5 before the call?
	fmt.Println(n)
}
```

## Task 2: Clever tracing with “defer”

### Intro

The defer keyword allows to specify a function that is called whenever the current function ends. What if we could call
one function at the beginning of the current function, and one at the end, with only one function call?

Like so:

```go
package main

import "fmt"

func f() {
	trace("f")
	fmt.Println("Doing something")
}
```

And when calling function f() it would print:

```
Entering f
Doing something
Leaving f
```

With a tricky use of the defer statement and a closure, we can do that!

### Your task

Write a function trace() that receives a string - the name of the current function - and does the following:

Print “Entering <name>” where <name> is the string parameter that trace receives
Create and return a function that prints “Leaving <name>”
Then call trace() via the defer keyword in such a way that trace() runs immediately, and returns its result to defer.

```go
package main

import "fmt"

func trace(name string) func() {
	// TODO:
	// 1. Print "Entering <name>"
	// 2. return a func() that prints "Leaving <name>"
}

func f() {
	defer // TODO: add trace() here so the defer receives the returned function
		fmt.Println("Doing something")
}

func main() {
	fmt.Println("Before f")
	f()
	fmt.Println("After f")
}
```
