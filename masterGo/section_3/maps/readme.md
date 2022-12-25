# Time To Practice: Maps

Maps are a pretty versatile data structure, and make a good ad-hoc storage in many situations. In this Time To Practice,
we want to create a word counter.

## Task 1: count words

Implement a function that receives a string and a map[string]int. The function shall split the string into words, turn
the words into lowercase, and adds the word to the map and/or increases the counter for this word.

### Tips

What could be the easiest way of increasing a counter for a map element? Would it matter whether or not the element is
already in the map or whether it has to be added?
Remember moons["Jupiter"]++ and the fact that a map element springs into existence if it does not exist. (An advantage
of the “zero value” concept in Go.)
After splitting the string, the resulting words may still have punctuation attached - quotation marks, colons, question
marks, etc. Trim all these away using strings.Trim(<wordvariable>, " \t\n\"'.,:;?!()-".
Turn each word into lowercase before counting it - have a look into the strings package for a suitable function.
This is the function to fill:

```go
package main

func count(s string, m map[string]int) {
	// your code here
}
```

## Task 2: Print the word counts

Write a function that receives a map[string]int and prints out each key (the word) together with its value (the count),
but only if the count is greater than 1.

```go
package main

func printCount(m map[string]int) {
	// your code here
}
```

Find these two Go files attached to this lecture:

maps.go as a starting point. It contains a main() that feeds your functions from either a file (if you pass a file name
on the command line) or from a test text at the end of the file.
mapssolution.go that contains one possible way of implementing the two functions. If your solution differs from
this–don’t worry, as long as it counts like it should.
Happy coding!