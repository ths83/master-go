package main

import (
	"fmt"
	"github.com/AppliedGoCourses/ConcurrencyDeepDive/mockdb"
	"golang.org/x/sync/errgroup"
	"os"
	"strconv"
	"sync"
	"time"
)

// If all goroutines start at once, none of them would ever get the chance
// to get a previously used connection from the pool.
// Therefore, we use a slight delay between the creation of goroutines,
// to simulate new work coming in gradually.
//
// Play with this value to see how the goroutine timing affects the
// time savings that can be achieved by pooling. Try e.g., 10, 20, and 30
const msec = 10

// Used for muting output during benchmarking.
var stdout = os.Stdout

func batchQueryWithoutPool() {
	var g errgroup.Group

	for i := 0; i < 100; i++ {
		time.Sleep(msec * time.Millisecond)
		g.Go(func() error {
			fmt.Println("Creating a new connection")
			db, _ := mockdb.Open("Server1")
			defer db.Close()
			_, _ = db.Query("select ingredients from recipes where name = 'rice bowl'")
			return nil
		})
	}
	g.Wait()
}

func batchQueryWithPool() {
	pool := &sync.Pool{}
	var g errgroup.Group
	for i := 0; i < 100; i++ {
		time.Sleep(msec * time.Millisecond)
		g.Go(func() (err error) {
			var db *mockdb.MockDB
			item := pool.Get()
			if item == nil {
				fmt.Println("Creating a new connection")
				db, _ = mockdb.Open("Server1")
			} else {
				db = item.(*mockdb.MockDB)
				fmt.Println("Reusing a connection from the pool")
			}
			_, _ = db.Query("select ingredients from recipes where name = 'rice bowl'")
			pool.Put(db)
			return nil
		})
	}
	g.Wait()
}

func batchQueryWithAutoPool() {
	pool := &sync.Pool{
		New: func() any {
			fmt.Println("Pool: Creating a new connection")
			db, _ := mockdb.Open("Server1")
			return db
		},
	}

	var g errgroup.Group

	for i := 0; i < 100; i++ {
		time.Sleep(msec * time.Millisecond)
		g.Go(func() (err error) {
			var db *mockdb.MockDB
			db = pool.Get().(*mockdb.MockDB)
			_, _ = db.Query("select ingredients from recipes where name = 'rice bowl'")
			pool.Put(db)
			return nil
		})
	}
	g.Wait()
}

// batchQueryWithLimitedAutoPool uses a buffered channel to limit the
// number of connections created by the sync.Pool.
//
// This poses a new situation. When connections are in use for a long time,
// other goroutines may be blocked waiting for a connection to become available.
// To address this, we use a select statement with a timeout case.
func batchQueryWithLimitedAutoPool() {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Fprintln(stdout, "Pool: Creating a new connection")
			db, _ := mockdb.Open("Server1")
			return db
		},
	}

	// A buffered channel limits the number of items in the pool.
	// Why a channel? Because a simple integer would not be
	// safe for concurrent use.
	limit := make(chan struct{}, 10)

	var g errgroup.Group

	for i := 0; i < 100; i++ {
		time.Sleep(msec * time.Millisecond)
		g.Go(func() (err error) {
			var db *mockdb.MockDB

			// Only allow 10 concurrent connections,
			select {
			case limit <- struct{}{}:
				// The limit chan is not yet full, so we can get a DB connection.
				db = pool.Get().(*mockdb.MockDB)
				fmt.Fprintln(stdout, "Got a connection from the pool -", len(limit))
			case <-time.After(msec * 20 * time.Millisecond):
				// The limit chan is still full, let's stop waiting.
				fmt.Fprintln(stdout, "Timeout waiting for a connection")
				return nil // returning an error would make the errgroup stop the other goroutines
			default:
				// This default case is ONLY necessary for printing the wait status.
				// Usually, you do not want a busy loop here. Remove this default case,
				// and the select statement then blocks until either a connection
				// becomes available or the timeout occurs.
				fmt.Fprintln(stdout, "    Waiting for an available connection")
				<-time.After(msec * 5 * time.Millisecond)
			}
			_, _ = db.Query("select ingredients from recipes where name = 'rice bowl'")

			// Do not add more connections back to the pool
			// than the limit would allow.
			select {
			case <-limit:
				fmt.Fprintln(stdout, "Put a connection back to the pool -", len(limit))
				pool.Put(db)
			default:
				fmt.Fprintln(stdout, "Pool is full, closing the connection")
				db.Close()
			}
			return nil
		})
	}
	g.Wait()
}

// limitedBatchQuery limits the number of goroutines, which indirectly limits
// the number of connections in the pool. The code is much simpler than the
// batchQueryWithLimitedAutoPool example, where we had to deploy two select
// blocks for getting and returning a connection, respectively.
func limitedBatchQuery() {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Fprintln(stdout, "Pool: Creating a new connection")
			db, _ := mockdb.Open("Server1")
			return db
		},
	}

	// A buffered channel limits the number of goroutines that
	// run at an given timne.
	limit := make(chan struct{}, 10)

	var g errgroup.Group

	for i := 0; i < 100; i++ {
		time.Sleep(msec * time.Millisecond)

		// Only allow 10 concurrent goroutines,
		select {
		case limit <- struct{}{}:
			g.Go(func() (err error) {
				var db *mockdb.MockDB
				fmt.Fprintln(stdout, len(limit), "goroutines are running")

				db = pool.Get().(*mockdb.MockDB)
				_, _ = db.Query("select ingredients from recipes where name = 'rice bowl'")
				pool.Put(db)
				<-limit
				return nil
			})
		}
	}
	g.Wait()
}

type funcs []struct {
	name string
	fn   func()
}

func usage(fns funcs) {
	fmt.Fprintln(stdout, "Usage: syncpool n")
	for i, fn := range fns {
		fmt.Printf("\tn = %d: %s\n", i, fn.name)
	}
}

func main() {
	fns := funcs{
		{"batch query without pool", batchQueryWithoutPool},
		{"batch query with pool", batchQueryWithPool},
		{"batch query with auto pool", batchQueryWithAutoPool},
		{"batch query with limited auto pool", batchQueryWithLimitedAutoPool},
		{"limited batch query", limitedBatchQuery},
	}
	if len(os.Args) <= 1 {
		usage(fns)
		return
	}
	selection := os.Args[1]
	n, err := strconv.Atoi(selection)
	if err != nil {
		usage(fns)
		return
	}
	fmt.Fprintln(stdout, fns[n].name)
	fns[n].fn()
}
