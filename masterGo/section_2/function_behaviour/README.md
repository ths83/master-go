# Time To Practice: Function Behavior

## Task 1: Variadic functions

Write a function longest() that takes an arbitrary number of strings and prints the length of the longest one.

```go
func longest( /* TODO */) int {
// TODO
}

func main() {
fmt.Println(longest("Six", "sleek", "swans", "swam", "swiftly", "southwards"))
}
```

## Task 2: Scope

Are you sure you remember all the scope levels? Especially in and around loops…

This task is an analytical one: The code below apparently uses the same variable s all over the place! This cannot be,
or can it?

Look closely, and try to identify all scope levels in this code.
Then rename the variables so that each scope has its unique variable name instead of s and no shadowing occurs anymore.

```go
func main() {
s := "abcde"
for _, s := range s {
s := unicode.ToUpper(s)
fmt.Print(string(s))
}
fmt.Println("\n" + s)
}
```
