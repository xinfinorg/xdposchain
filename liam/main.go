package main

import (
	"log"
	"sync/atomic"
	"time"
)

type Worker struct {
	resetCh     chan time.Duration
	chainHeadCh chan struct{}
	mining      int32
}

func (w *Worker) commitNewWork() {
	log.Println("[worker] committing new work")
}

func getResetTime(chain interface{}, period time.Duration) time.Duration {
	// Simulate dynamic calculation of reset time
	return period
}

func main() {
	// Initialize channels and worker
	resetCh := make(chan time.Duration)
	MinePeriodCh := make(chan time.Duration)
	NewRoundCh := make(chan struct{}, 1)
	chainHeadCh := make(chan struct{})
	c := make(chan struct{}, 1)
	finish := make(chan struct{})

	worker := &Worker{
		resetCh:     resetCh,
		chainHeadCh: chainHeadCh,
		mining:      1, // Simulate active mining
	}

	minePeriod := time.Duration(5) * time.Second // Default mine period
	timeout := time.NewTimer(minePeriod)

	go func() {
		for {
			select {
			case d := <-worker.resetCh:
				log.Println("[worker] receive reset")
				if !timeout.Stop() {
					select {
					case <-timeout.C:
					default:
					}
				}
				timeout.Reset(d)

			case <-timeout.C:
				log.Println("[worker] timeout triggered")
				time.Sleep(10 * time.Second)
				log.Println("[worker] sending timeout event")
				c <- struct{}{}
				log.Println("[worker] sent timeout event")

			case <-finish:
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case v := <-MinePeriodCh:
				log.Println("[worker] process MinePeriodCh event")
				minePeriod = v
				worker.resetCh <- minePeriod

			case <-c:
				log.Println("[worker] process timeout event")
				if atomic.LoadInt32(&worker.mining) == 1 {
					worker.commitNewWork()
				}
				resetTime := getResetTime(nil, minePeriod)
				log.Println("[worker] timeout sending reset channel")
				worker.resetCh <- resetTime
				log.Println("[worker] timeout sent reset channel")

			case <-worker.chainHeadCh:
				log.Println("[worker] process chainHeadCh event")
				worker.commitNewWork()
				resetTime := getResetTime(nil, minePeriod)
				log.Println("[worker] chainHeadCh sending reset channel")
				worker.resetCh <- resetTime
				log.Println("[worker] chainHeadCh sent reset channel")

			case <-NewRoundCh:
				log.Println("[worker] process NewRoundCh event")
				worker.commitNewWork()
				resetTime := getResetTime(nil, minePeriod)
				log.Println("[worker] NewRoundCh sending reset channel")
				worker.resetCh <- resetTime
				log.Println("[worker] NewRoundCh sent reset channel")
			}
		}
	}()

	// Simulate external events
	go func() {
		time.Sleep(2 * time.Second)
		MinePeriodCh <- 2 * time.Second
		time.Sleep(5 * time.Second)
		NewRoundCh <- struct{}{}
		chainHeadCh <- struct{}{}
		time.Sleep(5 * time.Second)
		close(finish)
	}()

	// Wait to finish
	<-finish
	log.Println("Worker finished")
}
