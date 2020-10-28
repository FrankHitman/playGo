package main

import (
	"log"
	"time"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	c.AddFunc("* * * * * *", func() { log.Println("Every hour on the half hour") })
	// c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("* 32/60 * * * *", func() { log.Println("Every hour thirty") })
	// c.AddFunc("* 24 * * * *", func() { fmt.Println("Every hour thirty") })

	c.Start()
	// Funcs are invoked in their own goroutine, asynchronously.
	// Funcs may also be added to a running Cron
	// c.AddFunc("@daily", func() { fmt.Println("Every day") })
	// Inspect the cron job entries' next and previous run times.
	// c.Stop() // Stop the scheduler (does not stop any jobs already running).
	// TestStopCausesJobsToNotRun()
	time.Sleep(10 * time.Second)

}

// const OneSecond = 1*time.Second + 10*time.Millisecond
//
// // Start, stop, then add an entry. Verify entry doesn't run.
// func TestStopCausesJobsToNotRun() {
//	wg := &sync.WaitGroup{}
//	wg.Add(1)
//
//	cron := cron.New()
//	cron.Start()
//	cron.AddFunc("* * * * * ?", func() {
//		log.Println("in cron func")
//		wg.Done()
//	})
//
//	select {
//	case <-time.After(OneSecond):
//		log.Println("in time after")
//		// No job ran!
//	case <-wait(wg):
//		log.Fatal("expected stopped cron does not run any job")
//	}
// }
//
// func wait(wg *sync.WaitGroup) chan bool {
//	ch := make(chan bool)
//	go func() {
//		wg.Wait()
//		ch <- true
//	}()
//	return ch
// }
