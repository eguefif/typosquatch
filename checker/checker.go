package checker

// TODO:
// - [ ] Check dns for results

import (
	"errors"
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
	domain    string
	state     TaskState
	record    []string
	mxRecords []string
	err       error
}

type Result struct {
	domain    string
	record    []string
	mxRecords []string
}

type ScanErrorStatus int

const (
	Timeout ScanErrorStatus = iota
	NetworkUnreachable
	NormalError
)

type ScanError struct {
	status  ScanErrorStatus
	message string
}

func (se ScanError) Error() string {
	return se.message
}

// Logic *****************************************

func CheckTypoSquatting(domains []string) []Result {
	numWorkers := 50
	if numWorkers > len(domains) {
		numWorkers = len(domains)
	}
	tasks := make([]CheckTask, 0, len(domains))
	for _, domain := range domains {
		tasks = append(tasks, CheckTask{domain: domain, state: StateWaiting})
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
			job.record, job.mxRecords, job.err = checkDns(job.domain)
			job.state = StateFinished
			resultsCh <- job
		}
	}
	return
}

func checkDns(domain string) ([]string, []string, error) {
	record, errLookup := net.LookupHost(domain)
	resultMx, errMx := net.LookupMX(domain)
	if errLookup != nil {
		return []string{}, []string{}, handleDnsError(domain, errLookup)
	}
	if errMx != nil {
		return record, []string{}, handleDnsError(domain, errLookup)
	}

	mxRecords := []string{}
	for _, mxRecord := range resultMx {
		mxRecords = append(mxRecords, mxRecord.Host)
	}
	return record, mxRecords, nil
}

func handleDnsError(domain string, err error) ScanError {
	var dnsError *net.DNSError
	var errorStatus ScanErrorStatus
	if errors.As(err, &dnsError) {
		switch {
		case dnsError.IsTimeout:
			errorStatus = Timeout
		case dnsError.IsTemporary:
			errorStatus = NetworkUnreachable
		default:
			errorStatus = NormalError
		}
	}

	return ScanError{status: errorStatus, message: fmt.Errorf("error lookup %s: %w", domain, err).Error()}
}

func handleResults(domains []string, resultsCh <-chan CheckTask) []Result {
	results := make([]Result, 0, len(domains))

	for taskResult := range resultsCh {
		switch taskResult.state {
		case StateFinished:
			results = append(results, Result{taskResult.domain, taskResult.record, taskResult.mxRecords})
		case StateError:
			fmt.Println("Error in task for domain: ", taskResult.domain)
		}
	}

	return results
}
