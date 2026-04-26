package handlers

import (
	"sync"
	"typosquatch/checker"
)

// TODO: Add mutex

type Handler struct {
	lock      sync.RWMutex
	jobs      map[int64][]checker.Result
	nextJobId int64
}

func NewHandler() *Handler {
	jobsStore := make(map[int64][]checker.Result)
	return &Handler{
		jobs:      jobsStore,
		nextJobId: 0,
	}
}

func (h *Handler) AddJob() int64 {
	jobId := h.nextJobId
	h.nextJobId = jobId + 1

	return jobId
}

func (h *Handler) AddResult(jobId int64, result []checker.Result) {
	h.jobs[jobId] = result
}

func (h *Handler) PopResult(jobId int64) []checker.Result {
	result := h.jobs[jobId]

	delete(h.jobs, jobId)

	return result
}
