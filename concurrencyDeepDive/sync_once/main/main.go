package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {
	b := FileBuffer{}

	var wg sync.WaitGroup
	for i := 1; i < 1000; i++ {
		wg.Add(1)
		go func() {
			data := b.GetFile()
			if data == nil {
				log.Println("error: file not loaded")
			}
		}()
		wg.Done()
	}
	wg.Wait()
}

type FileBuffer struct {
	data []byte
	once sync.Once
}

func (f *FileBuffer) GetFile() []byte {
	if f == nil {
		log.Fatalln("receiver must not be nil")
	}
	f.once.Do(func() {
		fmt.Println("loading 'data'")

		var err error
		f.data, err = os.ReadFile("data")
		if err != nil {
			log.Fatalln(err)
		}
	})
	return f.data
}
