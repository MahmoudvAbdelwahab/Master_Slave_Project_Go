package main

import (
    "flag"
    "fmt"
    "math/rand"
    "os"
    "os/signal"
    "sync"
    "syscall"
    "time"
)

// Job represents a unit of work
type Job struct {
    ID  int
    Num int
}

// Result represents a processed job result
type Result struct {
    JobID  int
    Input  int
    Output int
}

// worker simulates processing jobs and sending results
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        // simulate variable processing time
        time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
        out := job.Num * job.Num
        fmt.Printf("Worker %d processed job %d: %d^2 = %d\n", id, job.ID, job.Num, out)
        results <- Result{JobID: job.ID, Input: job.Num, Output: out}
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())

    workerCount := flag.Int("workers", 4, "Number of workers")
    jobCount := flag.Int("jobs", 20, "Number of jobs")
    flag.Parse()

    jobs := make(chan Job, *jobCount)
    results := make(chan Result, *jobCount)

    // Wait group to track workers
    var wg sync.WaitGroup

    // Start workers
    for w := 1; w <= *workerCount; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }

    // Graceful shutdown
    sig := make(chan os.Signal, 1)
    signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-sig
        fmt.Println("\nReceived interrupt, shutting down...")
        close(jobs)
    }()

    // Master: enqueue jobs
    go func() {
        for i := 1; i <= *jobCount; i++ {
            jobs <- Job{ID: i, Num: rand.Intn(100)}
        }
        close(jobs)
    }()

    // Close results channel when all workers finish
    go func() {
        wg.Wait()
        close(results)
    }()

    // Collect results
    var collected []Result
    for r := range results {
        collected = append(collected, r)
    }

    fmt.Println("\nAll jobs processed. Summary:")
    for _, r := range collected {
        fmt.Printf("Job %d: %d^2 = %d\n", r.JobID, r.Input, r.Output)
    }
}
