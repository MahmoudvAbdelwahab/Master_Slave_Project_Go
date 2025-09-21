# Master_Slave_Project_Go

Simple Master-Slave demo in Go. This project shows a lightweight job-dispatch system where:
- **Master**: HTTP server that accepts job submissions, keeps a job queue, registers slaves, and dispatches jobs to available slaves.
- **Slave**: HTTP worker that registers itself with the master and accepts jobs to execute (simulated work).

**Design goals**
- Pure Go, no external dependencies.
- Simple REST endpoints for clarity.
- Easy to run locally (multiple slaves + one master) — ideal for learning and experimentation.

**Files**
- master/main.go       : Master server source
- slave/main.go        : Slave worker source
- common/types.go      : Shared types used by master and slave
- Makefile             : build helpers
- run_example.sh       : simple script to run master and 2 slaves locally
- README.md            : this file

**How to run locally (Linux/macOS)**
1. Install Go (1.20+ recommended) and set GOPATH properly.
2. Build the master and slave binaries:
   ```bash
   make build
   ```
3. Start the master (default port 8080):
   ```bash
   ./bin/master
   ```
4. Start one or more slaves (each in its own shell). Each slave will auto-register with the master.
   ```bash
   ./bin/slave --master http://127.0.0.1:8080 --port 9001
   ./bin/slave --master http://127.0.0.1:8080 --port 9002
   ```
5. Submit jobs to the master (HTTP POST):
   ```bash
   curl -X POST -H 'Content-Type: application/json' -d '{"payload":"hello world","exec_time":3}' http://127.0.0.1:8080/submit
   ```
6. Check jobs list:
   ```bash
   curl http://127.0.0.1:8080/jobs
   ```

**Notes & extension ideas**
- Replace HTTP dispatch with gRPC or persistent TCP for production use.
- Add authentication and TLS.
- Add persistence (SQLite/postgres) for jobs and results.
- Add retries, backoff, and better scheduling (weighted, capacity-based).
