package checker

// TODO:
// - [ ] Check dns for results

import (
	"fmt"
	"net"
	"sync"
)

// Types ***********************************

type TaskState int

const (
	StateWaiting TaskState = iota
	StateProcessing
	StateFinished
	StateError
)

var stateName = map[TaskState]string{
	StateWaiting:    "waiting",
	StateProcessing: "processing",
	StateFinished:   "finished",
	StateError:      "error",
}

type CheckResult int

type CheckTask struct {
	domain   string
	state    TaskState
	squatted bool
}

type Result struct {
	domain   string
	squatted bool
}

// Logic *****************************************

func CheckTypoSquatting(domains []string) []Result {
	numWorkers := 50
	if numWorkers > len(domains) {
		numWorkers = len(domains)
	}
	tasks := make([]CheckTask, 0, len(domains))
	for _, domain := range domains {
		tasks = append(tasks, CheckTask{domain, StateWaiting, false})
	}

	jobsCh := make(chan CheckTask, numWorkers*2)
	resultsCh := make(chan CheckTask, numWorkers*2)

	scheduleWorkers(numWorkers, jobsCh, resultsCh, tasks)

	results := handleResults(domains, resultsCh)
	return results
}

func scheduleWorkers(numWorkers int, jobsCh chan CheckTask, resultsCh chan CheckTask, tasks []CheckTask) {
	var wg sync.WaitGroup
	wg.Add(numWorkers + 1)

	go func() {
		defer wg.Done()
		tasksDistributor(tasks, jobsCh)
	}()

	for _ = range numWorkers {
		go func() {
			defer wg.Done()
			tasksWorker(jobsCh, resultsCh)
		}()
	}

	go func() {
		wg.Wait()
		close(resultsCh)
	}()
}

func tasksDistributor(jobsList []CheckTask, jobsCh chan<- CheckTask) {
	for _, job := range jobsList {
		jobsCh <- job
	}
	close(jobsCh)
	return
}

func tasksWorker(jobsCh <-chan CheckTask, resultsCh chan<- CheckTask) {
	for job := range jobsCh {
		if job.state == StateWaiting {
			job.state = StateProcessing
			job.squatted = checkDns(job.domain)
			job.state = StateFinished
			resultsCh <- job
		}
	}
	return
}

func checkDns(domain string) bool {
	_, ok := net.LookupHost(domain)
	if ok != nil {
		return false
	}
	return true
}

func handleResults(domains []string, resultsCh <-chan CheckTask) []Result {
	results := make([]Result, 0, len(domains))

	for taskResult := range resultsCh {
		switch taskResult.state {
		case StateFinished:
			results = append(results, Result{taskResult.domain, taskResult.squatted})
		case StateError:
			fmt.Println("Error in task for domain: ", taskResult.domain)
		}
	}

	return results
}
