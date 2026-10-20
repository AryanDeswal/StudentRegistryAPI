# StudentRegistryAPI

A small RESTful API in Go for managing student records, built with an MVC-style layout and a thread-safe in-memory store.

![Go](https://img.shields.io/badge/Go-1.22-00ADD8?logo=go&logoColor=white)
![gorilla/mux](https://img.shields.io/badge/router-gorilla%2Fmux-2E7D32)

## About

StudentRegistryAPI exposes JSON endpoints to create, list, fetch and soft-delete student records. Records are kept in a mutex-guarded Go map behind a singleton (`db/in_memory_db.go`), so there is no external database to set up. All data is lost when the process stops.

## Features

- Create, list, get-by-ID and soft-delete student records
- UUID assigned to each student on creation, plus a `created_at` timestamp
- Soft delete: deleted records are flagged and hidden from reads, not removed from the map
- Concurrency-safe in-memory storage (`sync.Mutex` + `sync.Once` singleton)
- JSON error responses (`{"error": "..."}`) and a catch-all `404 not found` route
- Info and error logging to `app.log`
- Port configured through a `.env` file

## Tech Stack

- **Language:** Go (module targets Go 1.22.2)
- **Router:** [gorilla/mux](https://github.com/gorilla/mux)
- **IDs:** [google/uuid](https://github.com/google/uuid)
- **Config:** [joho/godotenv](https://github.com/joho/godotenv)
- **Storage:** in-memory Go map (no database)

## Project Structure

```
StudentRegistryAPI/
├── main.go                     # Loads .env, starts the HTTP server
├── router/student_router.go    # Route definitions
├── controller/
│   ├── student_controller.go   # Student request handlers
│   └── default_controller.go   # Catch-all 404 handler
├── model/student.go            # Student struct and UUID helper
├── db/in_memory_db.go          # Thread-safe in-memory store (singleton)
├── view/student_view.go        # JSON response helpers
├── utils/logger/logger.go      # Info/Error loggers writing to app.log
├── sample.env                  # Example environment file
├── go.mod / go.sum
└── app.log                     # Log output
```

## Getting Started

### Prerequisites

- Go 1.22 or newer

### Install and run

```sh
git clone https://github.com/AryanDeswal/StudentRegistryAPI.git
cd StudentRegistryAPI
go mod download
cp sample.env .env
go run main.go
```

The server prints `server running on port <PORT>` once it starts. A `.env` file is required: the app exits if it cannot load one or if `PORT` is unset.

### Environment variables

| Name   | Purpose                               |
| ------ | ------------------------------------- |
| `PORT` | Port the HTTP server listens on (`sample.env` uses `8080`) |

## API Reference

Base path: `/student/v1`

| Method   | Route                              | Description                                   |
| -------- | ---------------------------------- | --------------------------------------------- |
| `POST`   | `/student/v1/students`             | Create a student, returns `201` with `{"id": "<uuid>"}` |
| `GET`    | `/student/v1/students`             | List all students that are not deleted        |
| `GET`    | `/student/v1/students/{studentId}` | Get one student, or `404` if missing/deleted  |
| `DELETE` | `/student/v1/students/{studentId}` | Soft-delete a student, returns `204 No Content` (or `404`) |

Any other route returns `404 not found`.

### Student object

| Field        | Type   | Notes                          |
| ------------ | ------ | ------------------------------ |
| `id`         | string | UUID, set by the server        |
| `name`       | string |                                |
| `age`        | int    |                                |
| `class`      | string |                                |
| `subject`    | string |                                |
| `created_at` | string | RFC 3339 timestamp, set by the server |

Missing fields in the request body default to `""` or `0`. Unknown keys are ignored.

## Usage

```sh
# Create
curl -X POST http://localhost:8080/student/v1/students \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "age": 20, "class": "Physics", "subject": "Quantum Mechanics"}'
# {"id":"<uuid>"}

# List
curl http://localhost:8080/student/v1/students

# Get one
curl http://localhost:8080/student/v1/students/<uuid>

# Soft delete
curl -X DELETE http://localhost:8080/student/v1/students/<uuid>
```

Example student response:

```json
{
  "id": "student-uuid",
  "name": "John Doe",
  "age": 20,
  "class": "Physics",
  "subject": "Quantum Mechanics",
  "created_at": "2024-05-26T12:08:08Z"
}
```

## Logging

`utils/logger` opens `app.log` in the working directory (append mode) and provides two loggers:

- `INFO:` successful operations (create, list, get, delete)
- `ERROR:` failures such as invalid request bodies or unknown IDs
