# Time To Practice: Channels

## Task 1: What do these two main() functions print?

Here are two quite similar main() functions. Can you predict the output of each of them?

```go
package main

import "fmt"

func main() {
	c := make(chan int)
	c <- 1024
	fmt.Println(<-c)
}
```

```go
package main

import "fmt"

func main() {
	c := make(chan int, 1)
	c <- 1024
	fmt.Println(<-c)
}
```

(There is no solution file for this task. Simply run the two code snippets and compare the result. If you want, you can
use the Go playground for this test.)

## Task 2: Waiting for a goroutine

Remember channel axiom #4:

> A receive from a closed channel returns the zero value immediately

In other words, once the sender closes the channel, the channel starts delivering zero values to the reader.

Can you think of a way to use this behavior for making the main goroutine wait for a goroutine to finish work?

Let’s start from this code template:

```go
package main

import (
	"fmt"
	"time"
)

func worker( /* TODO: Maybe receive a channel here? */) {
	// Do something for some time
	for i := 0; i < 1000; i++ {
		fmt.Print(".")
	}
	fmt.Println()
	// TODO: How to tell main() that work is done?
}

func main() {

	// start the goroutine.
	go worker( /* TODO: Maybe pass a channel here? */)

	fmt.Println("Waiting for the goroutine")
	// TODO: Add code to wait for the goroutine.
	// time.Sleep() doesn't count.
	fmt.Println("Done")

}
```

This technique is also discussed in the lecture about managing goroutines.

## Task 3: Sharing one channel between more than two goroutines

Channels are variables, and therefore can be passed around like variables. What if we pass a channel to multiple
functions? Can two or more goroutines send to or receive from the same channel at the same time?

Write a program that starts three goroutines and passes a channel (the same channel!) to each one.

Let one goroutine send a number of consecutive integers (1,2,3,…), and let each of the other two continuously read
elements from the channel.

Will both goroutines receive the same data? Or will each one receive only a subset of the data?

(This task comes with no code template.)

## Task 4: Select

In this select statement, both cases send into the same channel. Can you predict what the receiver gets when
continuously reading from that channel?

Does the behavior depend on whether the channel is buffered?

```
for {
    select {
        case c <- 0:
        case c <- 1:
    }
}
```
