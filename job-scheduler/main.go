package main

import (
	"fmt"
	"time"
)

func main() {
	jobScheduler := NewJobScheduler(1 * time.Second)
	jobScheduler.Start()
	fmt.Print("Job started")

	job := PrintJob{Message: "Hello, World!"}
	jobScheduler.JobQueue <- job

	time.Sleep(10 * time.Second)
}

func (s *JobScheduler) ScheduleOnce(duration time.Duration, job Job) {
	go func() {
		time.Sleep(duration)
		s.JobQueue <- job
	}()
}

func (s *JobScheduler) Start() {
	go func() {
		ticker := time.NewTicker(s.Interval)

		for {
			select {
			case job := <-s.JobQueue:
				job.Execute()
			case <-ticker.C:
				for job := range s.JobQueue {
					job.Execute()
				}
			}
		}
	}()
}

type JobScheduler struct {
	JobQueue chan Job
	Interval time.Duration
}

func NewJobScheduler(interval time.Duration) *JobScheduler {
	return &JobScheduler{
		JobQueue: make(chan Job),
		Interval: interval,
	}
}

type PrintJob struct {
	Message string
}

func (p PrintJob) Execute() {
	fmt.Println(p.Message)
}

type Job interface {
	Execute()
}
