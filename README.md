# Master_Slave_Project_Go
Quick start:  Install Go 1.20+.  Run make build to build bin/master and bin/slave.  Start the master: ./bin/master  Start one or more slaves: ./bin/slave --master http://127.0.0.1:8080 --port 9001  Submit a job: curl -X POST -H 'Content-Type: application/json' -d '{"payload":"hello","exec_time":3}' http://127.0.0.1:8080/submit  Want me to: 
