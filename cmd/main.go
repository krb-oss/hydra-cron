package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	sh := gocron.NewScheduler(time.UTC)
	sh.Cron("0 5 * * *").Do(task)
	sh.StartBlocking()
}

func task() {
	fmt.Println("Worker running...")
}
