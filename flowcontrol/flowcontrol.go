package flowcontrol

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func Main() {
	var wg sync.WaitGroup
	wg.Add(1)
	jobs := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- i
		}
		time.Sleep(time.Second * 30)
		wg.Done()
	}()
	go func() {
		for i := 0; i < 2; i++ {
			go worker(jobs)
		}
	}()
	wg.Wait()
}

func worker(jobs <-chan int) {
	for job := range jobs {
		j := job
		if rest := math.Mod(float64(job), 4); rest == 0 {
			for j > 0 {
				fmt.Println(fmt.Sprintf("#%d job sleeps in #%d", job, j))
				time.Sleep(1 * time.Second)
				j--
			}
			fmt.Println(fmt.Sprintf("#%d job done", job))
		}
	}
}
