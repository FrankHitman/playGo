package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)

	// 起一个新的线程，10秒钟之后通过 channel 通知另外一个线程结束。
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}
}

// output
// Current time:  2024-06-13 12:21:31.538028 +0800 CST m=+1.000670773
// Current time:  2024-06-13 12:21:32.537635 +0800 CST m=+2.000269659
// Current time:  2024-06-13 12:21:33.537567 +0800 CST m=+3.000193273
// Current time:  2024-06-13 12:21:34.538167 +0800 CST m=+4.000785183
// Current time:  2024-06-13 12:21:35.537658 +0800 CST m=+5.000268400
// Current time:  2024-06-13 12:21:36.53801 +0800 CST m=+6.000612387
// Current time:  2024-06-13 12:21:37.538571 +0800 CST m=+7.001165347
// Current time:  2024-06-13 12:21:38.538052 +0800 CST m=+8.000638433
// Current time:  2024-06-13 12:21:39.537892 +0800 CST m=+9.000470375
// Current time:  2024-06-13 12:21:40.53779 +0800 CST m=+10.000359836
// Done!
