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
