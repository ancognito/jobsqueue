# jobsqueue

Example implementation of a basic in-memory job queue.

Not implemented:

* initial refactoring
* context propagation
* test coverage / benchmarks
* concurrency
* containerization / deployment
* CI
* API specification
* ser/de / message schema
* authentication
* CORS (if front-end consumers)
* tracing, logger, metrics
* better documentation


# Run it

In one terminal:

    go run cmd/server/main.go


In another terminal:

# Using `curl`

## Enqueue:

    curl -H 'Content-Type: application/json' -X POST \
      --data '{"Type": "TIME_CRITICAL"}' \
      http://localhost:8080/jobs/enqueue

> {"ID":100}

    curl http://localhost:8080/jobs/100

> {"ID":100,"Type":"TIME_CRITICAL","Status":"QUEUED"}


## Dequeue:

    curl http://localhost:8080/jobs/dequeue

> {"ID":100,"Type":"TIME_CRITICAL","Status":"IN_PROGRESS"}

    curl http://localhost:8080/jobs/100

> {"ID":100,"Type":"TIME_CRITICAL","Status":"IN_PROGRESS"}


## Conclude:

    curl http://localhost:8080/jobs/100/conclude

> (no content)

    curl http://localhost:8080/jobs/100

> {"ID":100,"Type":"TIME_CRITICAL","Status":"CONCLUDED"}
