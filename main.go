package main

import (
	"context"
	"log"
	"time"
)

func main() {
}

func subtest(ctx context.Context, name string, doneCallback func()) {
	index := 0
	for {
		time.Sleep(100 * time.Millisecond)
		index++
		select {
		case <-ctx.Done():
			log.Printf("%s: %s\n", name, context.Cause(ctx))
			if value, ok := ctx.Value("parent").(string); ok {
				log.Printf("%s: parent:%s\n", name, value)
			}
			if doneCallback != nil {
				doneCallback()
			}
			return
		default:
			log.Printf("%s: %d working...\n", name, index)
		}
	}
}

// clientChan must be initialized before use
var clientChan chan string

func producer(content string) {
	clientChan <- content
}

func consumer() *string {
	if content, ok := <-clientChan; ok {
		return &content
	}
	return nil
}
