package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

var run = true

func main() {
	fmt.Println("Booting")
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	a := make(chan string, 100)
	b := make(chan string, 100)
	c := make(chan string, 100)
	d := make(chan string, 100)

	// handle SIGINT, SIGTERM
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		mutex.Lock()
		run = false
		mutex.Unlock()
	}()

	go A(a, mutex)
	go B(a, b)
	for index := 0; index < 5; index++ {
		wg.Add(1)
		go C(b, c, d, &wg)
	}
	go SUM(c)
	go D(d)
	fmt.Println("Started")

	wg.Wait()
	close(c)
	close(d)
	fmt.Println("Shutting down")
}

func Interrupt() {
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
}

// create work
func A(a chan<- string, mutex *sync.Mutex) {
	for i := 0; ; i++ {
		mutex.Lock()
		if run {
			mutex.Unlock()
			time.Sleep(time.Second * 1)
			a <- strconv.Itoa(i)
		} else {
			mutex.Unlock()
			break
		}
	}
	close(a)
}

// multiply work
func B(a <-chan string, b chan<- string) {
	for as := range a {
		time.Sleep(time.Second * 1)
		b <- as
		b <- as
		b <- as
		b <- as
		b <- as
	}
	close(b)
}

// split work
func C(b <-chan string, c chan<- string, d chan<- string, wg *sync.WaitGroup) {
	for bs := range b {
		c <- bs
		d <- bs
	}
	wg.Done()
}

// document work
func D(d <-chan string) {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		Interrupt()
		return
	}
	var sum int
	for ds := range d {
		l, err := f.WriteString(ds + "\n")
		if err != nil {
			fmt.Println(err)
			f.Close()
			Interrupt()
			return
		}
		sum += l
	}
	fmt.Println(sum, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		Interrupt()
		return
	}
}

// calculate sum
func SUM(c <-chan string) {
	var sum int
	var num int
	var err error
	for n := range c {
		num, err = strconv.Atoi(n)
		if err != nil {
			fmt.Println("There was an error")
			Interrupt()
			return
		}
		sum += num
	}
	fmt.Println("The Sum is: ", sum)
}
