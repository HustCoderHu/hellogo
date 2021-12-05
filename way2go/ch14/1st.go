package ch14

import (
	"fmt"
	"log"
	"time"

	"github.com/sirupsen/logrus"
)

func pump1(out chan int) chan int {
	for i := 0; i < 10; i++ {
		out <- i * 2
	}
	costTime := make(chan int)
	costTime <- 2222
	return costTime
}

func pump2(out chan int) chan int {
	for i := 0; i < 10; i++ {
		out <- i * 3
	}
	costTime := make(chan int)
	costTime <- 3333
	return costTime
}

func F4_select() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go pump1(ch1)
	go pump2(ch2)

	for {
		select {
		case v := <-ch1:
			logrus.Info("ch1:", v)
		case v := <-ch2:
			logrus.Info("ch2:", v)
		}
	}
}

func F5_ticker() {
	ticker := time.NewTicker(1e8)
	defer ticker.Stop()

	done := false
	timeOutNs := time.Duration(5e8)
	boomCh := time.After(timeOutNs)
	for {
		select {
		case <-ticker.C:
			logrus.Info("tick.")
		case <-boomCh:
			logrus.Warn("boom!")
			done = true
			break
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
		if done {
			break
		}
	}
}

type Work struct {
	In  int
	Out int
}

func (work *Work) String() string {
	return fmt.Sprintf("in:%v", work.In)
}

func server(workChan <-chan *Work) {
	for work := range workChan {
		go safelyDo(work) // start the goroutine for that work
	}
}
func do(work *Work) {
	logrus.Info("work:", work.In)
	if 3 == work.In {
		panic("work.In==3")
	}
	work.Out = work.In * 10
}
func safelyDo(work *Work) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Work failed with %s in %v", err, work)
		}
	}()
	do(work)
}

func F6_recover() {
	workch := make(chan *Work)
	go func() {
		for i := 0; i < 4; i++ {
			workch <- &Work{i, 0}
		}
	}()
	server(workch)
}
