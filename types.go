package common

// Shared types between master and slaves
type Job struct {
    ID       string `json:"id"`
    Payload  string `json:"payload"`
    ExecTime int    `json:"exec_time"` // seconds to simulate work
    Status   string `json:"status"`    // queued|running|done|failed
    Result   string `json:"result"`    
}

type RegisterRequest struct {
    Address string `json:"address"` // http://ip:port of slave
}
