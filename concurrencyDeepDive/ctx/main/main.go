package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"math/rand"
	"time"
)

func main() {
	var g errgroup.Group
	bgctx := context.Background()
	ctx, cancel := context.WithTimeout(bgctx, 3*time.Second)

	const numShards = 5
	queries := []string{"pid=5543", "HTTP_418", "CON_RST"}
	logs := make(chan string, 1000)

	start := time.Now()

	for _, query := range queries {
		q := query
		for shard := 0; shard < numShards; shard++ {
			sh := shard
			g.Go(func() error {
				queryShard(ctx, q, sh, logs)
				return nil
			})
		}
	}

	limit := 100 // only the first result counts.

	g.Go(func() error {
		for i := 1; i <= limit; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("Receiving goroutine: %s after %s\n", ctx.Err(), time.Since(start))
				return nil
			case log := <-logs:
				fmt.Printf("Result %d: %s\n", i, log)
			}
		}
		fmt.Printf("Receiving goroutine: all results received. Query completed after %s\n", time.Since(start))
		cancel() // <- cancel the context.
		return nil
	})

	g.Wait()
	fmt.Printf("All goroutines finished after %s\n", time.Since(start))
}

func queryShard(ctx context.Context, query string, shard int, shardRes chan<- string) {
	start := time.Now()

	queryTime := time.After(time.Duration(5+rand.Intn(2)) * time.Second)

	for {
		select {
		case <-time.After(time.Duration(rand.Intn(1500)+500) * time.Millisecond):
			shardRes <- fmt.Sprintf("queryShard: found an occurrence of %s in shard %d at index %d", query, shard, 100000+rand.Intn(899999))
		case <-queryTime:
			fmt.Printf("queryShard: finished query '%s' on shard %d after %s\n", query, shard, time.Since(start))
			return
		case <-ctx.Done():
			fmt.Printf("queryShard: %s on shard %d after %s\n", ctx.Err(), shard, time.Since(start))
			return
		}
	}
}
