#!/usr/bin/env bash
set -euo pipefail
echo "Build..."
make build
echo "Start master in background..."
./bin/master &
MASTER_PID=$!
echo "Starting two slaves in background..."
./bin/slave --master http://127.0.0.1:8080 --port 9001 &
./bin/slave --master http://127.0.0.1:8080 --port 9002 &
echo "Master PID: $MASTER_PID"
echo "Submit a job:"
echo "curl -X POST -H 'Content-Type: application/json' -d '{"payload":"hello","exec_time":3}' http://127.0.0.1:8080/submit"
echo "To stop everything: kill $MASTER_PID and killall bin/slave || true"
