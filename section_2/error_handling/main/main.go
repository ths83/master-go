package main

import (
	"fmt"
	"log"
)

func propagate(i int) error {
	if err := verify(i); err != nil {
		return fmt.Errorf("propagate(%d): %w", i, err)
	}
	return nil
}

func verify(i int) error {
	if i < 0 || i > 10 {
		return fmt.Errorf("verify: %d is outside the allowed range (0..10)", i)
	}
	return nil
}

func retry(i int) error {
	err := propagate(i)
	if err != nil {
		err = propagate(i / 2)
		if err != nil {
			return fmt.Errorf("retry: %s", err)
		}
	}
	return nil
}

func onlyLog(i int) {
	if err := retry(i); err != nil {
		log.Println("onlyLog:", err)
	}
}

func logAndExit(i int) {
	err := retry(i)
	if err != nil {
		log.Fatalln("exit:", err)
	}
}

func unexpectedError(p *int) {
	if p == nil {
		panic("p must not be nil")
	}
}

func main() {
	defer func() {
		res := recover()
		if res != nil {
			fmt.Println("Recovered from a panic:", res)
		}
	}()

	unexpectedError(nil)
}
