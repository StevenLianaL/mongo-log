package test

import (
	"context"
	"fmt"
	"mongo-log"
	"sync"
	"testing"
	"time"
)

// should use like this case
func TestManagerInWg(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("tag")
	collection := mongo_log.GetCollection("test", "test", "test", "test", "localhost", &ctx)
	var manager = mongo_log.Manager{
		Project:    "test",
		App:        "test",
		Collection: collection,
		Ctx:        &ctx,
	}
	t.Log(manager)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			manager.Log("a new log", mongo_log.INFO, i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func TestManagerInRoutine(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := mongo_log.GetCollection("test", "test", "test", "test", "localhost", &ctx)
	var manager = mongo_log.Manager{
		Project:    "test",
		App:        "test",
		Collection: collection,
		Ctx:        &ctx,
	}
	t.Log(manager)
	for i := 0; i < 10; i++ {
		go manager.Log("a new log", mongo_log.INFO, i)
	}
	time.Sleep(time.Second * 1)
}