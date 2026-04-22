# typosquatch

A typosquatting domain detection tool with a CLI and REST API, built in Go.

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
**Go concepts:** slices, maps, `rune` vs `byte` (essential for homoglyphs), range loops, functions as values, package design

- [x] Character deletion (`exaple.com`)
- [ ] Character transposition (`exmaple.com`)
- [ ] Adjacent key substitution (QWERTY map)
- [ ] Homoglyph substitution (`examp1e.com`)
- [ ] TLD variants (`.net`, `.org`, `.co`, `.io` ...)
- [ ] Subdomain insertion (`mail-example.com`)
- [ ] Bitsquatting (XOR each byte with powers of 2)
- [ ] Unit tests per strategy

---

### Phase 3 — Concurrent DNS resolution
**Go concepts:** goroutines, channels, `sync.WaitGroup`, worker pool pattern, `context` with timeout and cancellation, `net` package

- [ ] Worker pool reading candidates from an input channel
- [ ] A record DNS lookup per candidate
- [ ] Configurable concurrency (`--concurrency` flag)
- [ ] Context-based timeout and cancellation
- [ ] Benchmark at 50 / 100 / 200 workers

---

### Phase 4 — Enrichment
**Go concepts:** `net/http` client, struct design, JSON marshal/unmarshal, interfaces, error wrapping with `fmt.Errorf`

- [ ] MX record lookup
- [ ] HTTP/HTTPS probe (grab page title and `Server` header)
- [ ] TLS certificate inspection (issuer, expiry, CN match)
- [ ] WHOIS registration date via public API
- [ ] Define a `Checker` interface — each enrichment is a pluggable implementation

---

### Phase 5 — CLI output
**Go concepts:** `io.Writer` abstraction, `text/tabwriter`, `encoding/json`, `sort` package, struct tags

- [ ] Sort results by score descending
- [ ] Default pretty table output (`text/tabwriter`)
- [ ] `--output json` flag
- [ ] `--min-score` filter flag

---

### Phase 6 — REST API
**Go concepts:** `net/http` handlers, `json.Decoder` / `json.Encoder`, `sync.RWMutex`, goroutines for background jobs, proper HTTP status codes

- [ ] `POST /scan` — submit domain, return job ID immediately (non-blocking)
- [ ] `GET /scan/:id` — poll status (`pending | running | done | error`) and results
- [ ] `GET /scan/:id/export` — JSON or CSV (`?format=csv`)
- [ ] In-memory job store protected by `sync.RWMutex`
- [ ] CORS middleware (understand the headers, don't just use a library)
- [ ] Request logging middleware
- [ ] Panic recovery middleware

---

### Phase 7 — Production quality
**Go concepts:** `slog`, `cobra`, `errgroup`, benchmarks (`testing.B`), build pipeline

- [ ] Structured logging with `slog` throughout
- [ ] Migrate CLI to `cobra`
- [ ] Replace manual `WaitGroup` patterns with `errgroup`
- [ ] `golangci-lint` clean
- [ ] Makefile: `build`, `test`, `lint`, `bench`
- [ ] Multi-stage `Dockerfile`
- [ ] GitHub Actions: test + lint on push

---

### Phase 8 — Stretch goals

- [ ] Persist scan results to SQLite (`database/sql` + `mattn/go-sqlite3`)
- [ ] OpenAPI spec (`openapi.yaml`)
- [ ] Rate limiting middleware
- [ ] `--watch` mode: re-scan on interval, diff against previous results
