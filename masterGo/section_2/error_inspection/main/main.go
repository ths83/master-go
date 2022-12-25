package main

import (
	"errors"
	"fmt"
	"os"
)

type Doc struct {
	ID    int
	Title string
	Text  string
}

func WriteDoc(path string, doc Doc) (error, error) {
	_, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		return fmt.Errorf("cannot write %s (id %d): %w", doc.Title, doc.ID, err), err
	}
	return nil, nil
}

func main() {
	doc := Doc{
		ID:    20,
		Title: "Error Inspection",
		Text:  "In the previous lecture, we learned about wrapping errors...",
	}

	s, err := WriteDoc("/path/to/no_file", doc)
	if err != nil {
		var pathErr *os.PathError
		fmt.Println("pathErr is:", pathErr)
		fmt.Println("err is an os.PathError:", errors.As(err, &pathErr))
		fmt.Println("pathError info - Op:", pathErr.Op, ", Path:",
			pathErr.Path, "Err:", pathErr.Err)
		fmt.Println("Did we hit a timeout: ", pathErr.Timeout())
		return
	}
	fmt.Println(s)
}
