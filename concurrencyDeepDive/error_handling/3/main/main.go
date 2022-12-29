package main

import (
	"fmt"
	"github.com/AppliedGoCourses/ConcurrencyDeepDive/mockdb"
	"golang.org/x/sync/errgroup"
	"log"
)

func main() {

	var g errgroup.Group

	conns := []string{"db1", "db2", "db3", "db4", "db5", "db6"}

	res := make(chan string)

	for _, conn := range conns {
		c := conn
		g.Go(func() error {
			return checkDBstatus(c, res)
		})
	}

	done := make(chan struct{})
	go func() {
		for {
			r, ok := <-res
			if !ok {
				close(done)
				return
			}
			fmt.Println(r)
		}
	}()

	err := g.Wait()
	if err != nil {
		log.Printf("Monitor error: %s\n", err)
	}
	close(res)
	<-done
	fmt.Println("\nDone.")
}

func checkDBstatus(conn string, res chan<- string) error {
	db, err := mockdb.Open(conn)
	if err != nil {
		return fmt.Errorf("checkDBstatus: cannot open DB: %w", err)
	}
	defer db.Close()

	status, err := db.Status()
	if err != nil {
		return fmt.Errorf("checkDBstatus: cannot check status: %w", err)
	}
	res <- status
	return nil
}
