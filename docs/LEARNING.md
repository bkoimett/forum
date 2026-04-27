# 🚀 Learning Notes: Forum

## 🎯 Project Goals
* build server & template rendering
* add sqllit database and sql
* add authentication and sessions (cookies)
* add containerisation (Docker)

## 🛠️ Key Go Concepts Applie d


## 🏗️ Project Structure & Architecture
* `cmd/`: CLI entry points.
* `pkg/`: Library code used by external projects.
* `internal/`: Private application code (ensures encapsulated code).


* **Key Learning:** Learned to use `internal` for strict boundary enforcement.

## 💡 "Aha!" Moments & Solutions
* **Problem:** [E.g., Handling JSON efficiently]
* **Solution:** [E.g., Used streaming decoder `json.NewDecoder`]
* **Resource:** [Link to tutorial/documentation]

## 🚩 Pitfalls & Common Mistakes (Go-Specific)
* **Variable Shadowing:** Be careful with `:=` inside `if` or `for` blocks.
* **Defer Ordering:** `defer` runs in LIFO order (last in, first out).
* **Pointer vs Value:** Passed structs by pointer `*T` to avoid copying large data.

## 🚀 Future Learning
* [ ] Advanced concurrency patterns (`context`, `sync.WaitGroup`).
* [ ] Testing best practices (table-driven tests).
* [ ] Profiling and optimization (`pprof`).

## 📚 References
* [Go by Example](https://gobyexample.com/)
* [Effective Go](https://golang.org/doc/effective_go)
