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

	res := make(chan DBStatus)

	for _, conn := range conns {
		wg.Add(1)
		go checkDBstatus(conn, res, &wg)
	}

	done := make(chan struct{})

	go func() {
		for {
			r, ok := <-res
			if !ok {
				close(done)
				return
			}
			if r.Err != nil {
				log.Printf("Monitor error: %s\n", r.Err)
			} else {
				fmt.Println(r.Status)
			}
		}
	}()

	wg.Wait()
	close(res)
	<-done
	fmt.Println("\nDone.")
}

type DBStatus struct {
	Status string
	Err    error
}

func checkDBstatus(conn string, res chan<- DBStatus, wg *sync.WaitGroup) {

	defer wg.Done()

	result := DBStatus{}

	db, err := mockdb.Open(conn)
	if err != nil {
		result.Err = fmt.Errorf("checkDBstatus: cannot open DB: %s", err)
		res <- result
		return
	}
	defer db.Close()

	status, err := db.Status()
	if err != nil {
		result.Err = fmt.Errorf("checkDBstatus: cannot check status: %s", err)
		res <- result
		return
	}

	result.Status = status
	res <- result
}
