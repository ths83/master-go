package main

import (
	"fmt"
	"github.com/AppliedGoCourses/ConcurrencyDeepDive/mockdb"
	"log"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	conns := []string{"db1", "db2", "db3", "db4", "db5", "db6"}

	res := make(chan string)
	err := make(chan error)

	for _, conn := range conns {
		wg.Add(1)
		go checkDBstatus(conn, res, err, &wg)
	}

	done := make(chan struct{})

	go func() {
		for {
			select {
			case r, ok := <-res:
				if !ok {
					close(done)
					return
				}
				fmt.Println(r)
			case e := <-err:
				log.Printf("Monitor error: %s\n", e)
			}
		}
	}()

	wg.Wait()

	// Tell the result-reading goroutine to stop
	close(res)
	// Wait for the result-reading goroutine to finish
	<-done
	fmt.Println("\nDone.")
}

func checkDBstatus(conn string, res chan<- string, errc chan<- error, wg *sync.WaitGroup) {

	defer wg.Done()

	db, err := mockdb.Open(conn)
	if err != nil {
		errc <- fmt.Errorf("checkDBstatus: cannot open DB: %s", err)
		return
	}
	defer db.Close()

	status, err := db.Status()
	if err != nil {
		errc <- fmt.Errorf("checkDBstatus: cannot check status: %s", err)
		return
	}
	res <- status
}
