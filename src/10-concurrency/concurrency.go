package main

import (
	"sync"
	"time"

	"github.com/pterm/pterm"
)

func main() {
	waitTime := time.Millisecond * 100

	// start the progress bars in go routines
	var wg sync.WaitGroup
	bar1, _ := pterm.DefaultProgressbar.WithTotal(100).WithTitle("Downloading more ram").Start()

	// bar1 := uiprogress.AddBar(20).AppendCompleted().PrependElapsed()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			bar1.Increment() // Increment the progressbar by one. Use Add(x int) to increment by a custom amount.
			time.Sleep(waitTime)
		}
	}()

	bar2, _ := pterm.DefaultProgressbar.WithTotal(100).WithTitle("Contemplating things").Start()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			bar2.Increment() // Increment the progressbar by one. Use Add(x int) to increment by a custom amount.
			time.Sleep(waitTime)
		}
	}()

	time.Sleep(time.Second)
	bar3, _ := pterm.DefaultProgressbar.WithTotal(100).WithTitle("Reticulating Splines").Start()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			bar3.Increment() // Increment the progressbar by one. Use Add(x int) to increment by a custom amount.
			time.Sleep(waitTime)
		}
	}()

	// wait for all the go routines to finish
	wg.Wait()
}
