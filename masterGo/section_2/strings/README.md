# Time To Practice: strings
## Intro
The world is full of different languages, alphabets, and writing systems, and this is fascinating. The downside is that string handling has become more difficult than it was in the last millennium, where pure and simple ASCII was the predominant text encoding - at least in English speaking countries.

Luckily, Go comes with Unicode support out of the box. So let’s solve a text conversion task with Unicode in mind.

### Your task
Write code that turns a string into an acronym. For example: Turn “Pan Galactic Gargle Blaster” into “PGGB”.

Remember that not all string functionality is unicode-aware. Inspect the unicode package from the standard library at https://golang.org/pkg/unicode/ - it may contain some of the functions you need.

```go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func acronym(s string) (acr string) {

	// TODO: Your code here

}

func main() {
	s := "Pan Galactic Gargle Blaster"
	if len(os.Args) > 1 {
		s = strings.Join(os.Args, " ")
	}
	fmt.Println(acronym(s))
}
```
