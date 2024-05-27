# StudentRegistryAPI

## Overview
The **StudentRegistryAPI** is a Go-based RESTful API designed for managing student records, following the MVC (Model-View-Controller) architecture pattern. It supports basic CRUD (Create, Read, Update, Delete) operations with an in-memory database.

## Table of Contents
- [Project Structure](#project-structure)
- [Setting Up](#setting-up)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
  - [Create a Student](#create-a-student)
  - [Get All Students](#get-all-students)
  - [Get a Student by ID](#get-a-student-by-id)
  - [Soft Delete a Student](#soft-delete-a-student)
- [Logging](#logging)
- [Additional Notes](#additional-notes)

## Project Structure
```
StudentRegistryAPI/
├── controller/
│   └── default_controller.go
│   └── student_controller.go
├── db/
│   └── in_memory_db.go
├── model/
│   └── student.go
├── router/
│   └── student_router.go
├── utils/
│   └── logger
│       └── logger.go
├── view/
│   └── student_view.go
├── .gitignore
├── app.log
├── go.mod
├── go.sum
├── main.go
└── sample.env
```

## Setting Up
Clone the repository:
```sh
git clone <repository_url>
cd StudentRegistryAPI
```

Initialize the Go module:
```sh
go mod init StudentRegistryAPI
go mod tidy
```

## Running the Application
Run the application:
```sh
go run main.go
```

## API Endpoints

### Create a Student
- **URL:** `/student/v1/students`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "name": "John Doe",
    "age": 20,
    "class": "Physics",
    "subject": "Quantum Mechanics"
  }
  ```
- If you don't provide the required fields in the request body, those fields will be initialized as follows:
  1. Empty strings for string fields.
  2. Zero for integer fields.
- Extra key-value pairs will be ignored
- **Response:**
  ```json
  {
    "id": "generated-uuid"
  }
  ```

### Get All Students
- **URL:** `/student/v1/students`
- **Method:** `GET`
- **Response:**
  ```json
  [
    {
      "id": "student-uuid-1",
      "name": "John Doe",
      "age": 20,
      "class": "Physics",
      "subject": "Quantum Mechanics",
      "created_at": "2023-05-26T13:47:00Z"
    },
    {
      "id": "student-uuid-2",
      "name": "Jane Doe",
      "age": 22,
      "class": "Chemistry",
      "subject": "Organic Chemistry",
      "created_at": "2023-05-26T13:47:00Z"
    }
  ]
  ```

### Get a Student by ID
- **URL:** `/student/v1/students/{studentId}`
- **Method:** `GET`
- **Response:**
  ```json
  {
    "id": "student-uuid",
    "name": "John Doe",
    "age": 20,
    "class": "Physics",
    "subject": "Quantum Mechanics",
    "created_at": "2023-05-26T13:47:00Z"
  }
  ```

### Soft Delete a Student
- **URL:** `/student/v1/students/{studentId}`
- **Method:** `DELETE`
- **Response:** `204 No Content`

## Logging
The application logs events to `app.log` with two log levels:
- **Info Logger:** Logs informational messages, such as successful operations.
- **Error Logger:** Logs error messages, such as failed operations.

## Additional Notes
- **Soft Delete:** Deleting a student marks the record as deleted but does not remove it from the database.
- **UUID Generation:** Each student record is assigned a unique UUID upon creation.
- **In-Memory Database:** The application uses an in-memory database (a map) for simplicity. Data will be lost when the application stops.
