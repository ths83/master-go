# Time To Practice: If, Switch, And For

The “Time To Practice” lectures contain mini-exercises for recapitulating the previous lecture(s) and strengthen the
knowledge you just learned.

The programming tasks are intentionally not very complex. If they feel too easy for you, keep in mind that the tasks
first and foremost shall help you getting a feeling for the language - after all, practicing is the most efficient way
of internalizing what you have learned.

So here we go:

## Part 1: if and for

### What to do

Read the task description, create a new file named collatz.go in your editor, then copy & paste the code below into, and
complete the function so that it accomplishes the task.

Execute the file in a shell or command prompt by cd’ing to the directory where your collatz.go file is and running

```shell
go run collatz.go
```

_Optional: To test if your function returns correct results, download collatz_test.go, place it into the same directory
as collatz.go, and run_

```shell
go test
```

Your collatz() function is then tested against a list of different inputs and results. This is Go’s built-in unit
testing feature; you will learn more about it in section 4 of this course.

#### Task: Collatz Conjecture

Pick a number n > 1. Apply the following process repeatedly until n becomes 1.

If n is even, divide it by 2.
If n is odd, multiply it by 3, then add 1.
Return the number of steps needed.

In case you need the remainder of a division, use “a % b”.

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func collatz(n int) int {
	count := 0
	// Your code here!
	return count
}
func main() {
	var n int
	var err error
	if len(os.Args) > 1 { // Read the number from the command line
		n, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
	} else { // Read the number interactively
		fmt.Println("Input a number > 1: ")
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if n <= 1 {
		fmt.Println("n must be larger than 1.")
		return
	}
	c := collatz(n)
	fmt.Println(n, "requires", c, "steps to reach 1.")
}
```

### Part 2: switch

#### What to do

Similar to the Collatz Conjecture part, create a new file in your editor called fizzbuzz.go, copy and paste the code
from below into that file, and implement the task. Then run

```shell
go run fizzbuzz.go
```

to see if your code works.

#### The Task: FizzBuzz

Write a program that plays the game of FizzBuzz. The rules are simple:

Count from 1 to a given number n.
Print out each number, with the following exceptions:
If the number is divisible by 3, print “Fizz” instead of the number.
If the number is divisible by 5, print “Buzz”.
If the number is divisible by 15, print “FizzBuzz”.
The code below reads a number from the command line (default is 50 if no number is given) and calls function fizzbuzz
with that number.

Fill the body of fizzbuzz, using a for loop and a switch statement.

````go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func fizzbuzz(n int) {
	// Your code here!
}
func main() {
	n := 50
	if len(os.Args) > 1 {
		max, err := strconv.Atoi(os.Args[1])
		if err == nil {
			n = max
		}
	}
	fizzbuzz(n)
}
````
