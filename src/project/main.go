package project

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}

type Worker struct {
	Id         int
	JobQueue   chan Job
	WorkerPool chan chan Job
	QuitChan   chan bool
}

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
	JobQueue   chan Job
}

func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		Id:         id,
		JobQueue:   make(chan Job),
		WorkerPool: workerPool,
		QuitChan:   make(chan bool),
	}
}

func (worker *Worker) Start() {
	go func() {
		for {
			worker.WorkerPool <- worker.JobQueue

			select {
			case job := <-worker.JobQueue:
				fmt.Printf("Worker with id %d Started\n", worker.Id)
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("Worker with id %d finished has result %d\n", worker.Id, fib)
			case <-worker.QuitChan:
				fmt.Printf("Worker with id %d has quit\n", worker.Id)
			}
		}
	}()
}

func (worker *Worker) Stop() {
	go func() {
		worker.QuitChan <- true
	}()
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	return &Dispatcher{
		MaxWorkers: maxWorkers,
		JobQueue:   jobQueue,
		WorkerPool: make(chan chan Job, maxWorkers),
	}
}

func (dispatcher *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-dispatcher.JobQueue:
			go func() {
				workerJobQueue := <-dispatcher.WorkerPool
				workerJobQueue <- job
			}()
		}
	}
}

func (dispatcher *Dispatcher) Run() {
	for i := 0; i < dispatcher.MaxWorkers; i++ {
		worker := NewWorker(i, dispatcher.WorkerPool)
		worker.Start()
	}

	go dispatcher.Dispatch()
}

func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	if r.Method != "POST" {
		http.Error(w, "POST Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	dalay, err := time.ParseDuration(r.FormValue("delay"))

	if err != nil {
		http.Error(w, "Invalid delay", http.StatusBadRequest)
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Invalid value", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Invalid name", http.StatusBadRequest)
		return
	}

	job := Job{
		Name:   name,
		Delay:  dalay,
		Number: value,
	}

	jobQueue <- job

	w.WriteHeader(http.StatusCreated)
}

func Main() {
	const (
		maxWorkers   = 4
		maxQueueSize = 20
		port         = ":8080"
	)

	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(jobQueue, maxWorkers)

	dispatcher.Run()

	// http://localhost:8080/fib
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jobQueue)
	})

	log.Fatal(http.ListenAndServe(port, nil))

}
