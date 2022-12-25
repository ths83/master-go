package main

import "fmt"

//func fp(a *int) *int {
//	b := 2
//	a = &b
//	return a
//}
//func main() {
//	x := 1
//	p := &x // to make clear that we deal with two pointers
//	a := fp(p)
//	fmt.Println(*a)
//	fmt.Println(x)
//}

//func f(a *int) {
//	*a = *a + 1
//}
//
//func main() {
//	x := 1
//	p := &x
//	f(p)
//	fmt.Println(x)
//}

func f() *int {
	a := 7
	return &a
}
func main() {
	p := f()
	fmt.Println(*p)
}
