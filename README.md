# typosquatch

A typosquatting domain detection tool with a REST API, built in Go.

---

## Design

The program generates a list of domain permutations, then uses a worker pool to resolve them concurrently.

* One goroutine distributes tasks via a buffered channel (size: `numWorkers * 2`)
* `numWorkers` goroutines resolve DNS and send results to a results channel (same size)
* One goroutine collects results

The API layer submits scan jobs in the background and exposes endpoints to poll their status.

---

## Roadmap

### Phase 1 — Project setup and CLI basics
**Go concepts:** `go mod init`, package structure, `flag` package, `os.Args`, error handling idioms, table-driven tests with `testing`

- [x] Initialize module and project structure
- [x] Parse `--domain` flag from CLI
- [x] Validate domain input (non-empty, strip `www.`, basic format check)
- [x] Unit tests for the validation function

---

### Phase 2 — Permutation engine
**Go concepts:** slices, maps, `rune` vs `byte`, range loops, `iter.Seq`

- [x] Character deletion (`exaple.com`)
- [ ] Character transposition (`exmaple.com`)
- [ ] Adjacent key substitution (QWERTY map)

---

### Phase 3 — Concurrent DNS resolution
**Go concepts:** goroutines, channels, `sync.WaitGroup`, worker pool pattern, `net` package

- [x] Worker pool reading candidates from an input channel
- [x] A record DNS lookup per candidate
- [x] MX record lookup

---

### Phase 4 — Enrichment
**Go concepts:** `net/http` client, struct design, JSON marshal/unmarshal, interfaces, error wrapping with `fmt.Errorf`

- [x] MX record lookup
- [ ] HTTP/HTTPS probe (grab page title and `Server` header)
- [ ] TLS certificate inspection (issuer, expiry, CN match)
- [ ] Define a `Checker` interface — each enrichment is a pluggable implementation

---

### Phase 5 — Error handling
**Go concepts:** custom error types, `fmt.Errorf` + `%w`, `errors.Is` / `errors.As`, sentinel errors

- [x] Define a `ScanError` type with a `Code` field (e.g. `ErrNotFound`, `ErrInvalidDomain`)
- [x] Wrap DNS errors with context (`fmt.Errorf("lookup %s: %w", domain, err)`)
- [x] Use `errors.Is` / `errors.As` in handlers to map errors to HTTP status codes

---

### Phase 6 — REST API
**Go concepts:** `net/http` handlers, `http.ServeMux` (Go 1.22 routing), `json.Decoder` / `json.Encoder`, `sync.RWMutex`, proper HTTP status codes

- [ ] `POST /scan` — validate body `{"domain": "..."}`, start background scan, return `{"id": "..."}` with `202 Accepted`
- [ ] `GET /scan/{id}` — return job status (`pending | running | done | error`) and results when done
- [ ] In-memory job store protected by `sync.RWMutex`
- [ ] Map `ScanError` codes to correct HTTP status codes in a single `writeError` helper
- [ ] Table-driven tests for each handler using `httptest.NewRecorder`
