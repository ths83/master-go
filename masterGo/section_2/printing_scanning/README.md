# Time To Practice: Printing And Scanning

The broad range of formatting options available for printing and scanning can be overwhelming. Plus, the programming
languages you know might do it differently.

Practicing is the best strategy to get familiar with Go’s format options.

For the following practice, open the page https://golang.org/pkg/fmt/ in your browser. Here you find all available formatting placeholders (called “verbs”) that you need for the tasks.

## Part 1: Printing

### What to do

Inspect the following Printf() statements. Each has a comment that shows the desired output format.

```go
package main

import "fmt"

func main() {
	// Print RGB values...
	r, g, b := 124, 87, 3

	// ...as #7c5703  (specifying hex format, fixed width, and leading zeroes)
	// Hint: don't forget to add a newline at the end of the format string.
	fmt.Printf("", r, g, b)

	// ...as rgb(124, 87, 3)
	fmt.Printf("", r, g, b)

	// ...as rgb(124, 087, 003) (specifying fixed width and leading zeroes)
	fmt.Printf("", r, g, b)

	// ...as rgb(48%, 34%, 1%) (specifying a literal percent sign)
	fmt.Printf("", r, g, b)

	// Print the type of r.
	fmt.Printf("", r)
}
```

Copy the code into your editor (or use the template that you can find in the downloaded source code). Then fill the
format strings so that the output matches the format in the comments.

## Part 2: Scanning
### What to do
In the code below, str1 and str2 contain strings with some numbers to read. Complete each Sscanf statement with format
string and variables to read into.

```go
package main

import "fmt"

func main() {

    var n1, n2, n3, n4 int
    var f1 float64

    // Scan the card number.
    str1 := "Card number: 1234 5678 0123 4567"
    _, err := fmt.Sscanf(str1, "", ... )
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("%04d %04d %04d %04d\n", n1, n2, n3, n4)

    // Scan the numeric values into a floating-point variable, and an integer.
    str2 := "Brightness is 50.0% (hex #7ffff)"
    _, err = fmt.Sscanf(str2, "", ... )
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(f1, n1)

}
```
## Bonus task: Advanced printing
### What to do
Print the same variable twice, once as a decimal value and once as a hex value, without repeating the variable in
Printf’s argument list.

Hint: Search for “Explicit argument indexes” in the fmt documentation!

```go
package main

import "fmt"

func main() {
	n := 49374
	fmt.Printf("", n)
}
```
