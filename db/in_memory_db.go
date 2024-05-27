package db

import (
	"StudentRegistryAPI/model"
	"sync"
	"time"
)

// InMemoryDB simulates a database
type InMemoryDB struct {
	mu       sync.Mutex
	students map[string]model.Student
}

var dbInstance *InMemoryDB
var once sync.Once

// GetDBInstance returns the singleton instance of the in-memory database
func GetDBInstance() *InMemoryDB {
	once.Do(func() {
		dbInstance = &InMemoryDB{
			students: make(map[string]model.Student),
		}
	})
	return dbInstance
}

func (db *InMemoryDB) CreateStudent(student model.Student) string {
	db.mu.Lock()
	defer db.mu.Unlock()
	student.ID = model.GenerateUUID()
	student.CreatedAt = time.Now()
	db.students[student.ID] = student
	return student.ID
}

func (db *InMemoryDB) GetAllStudents() []model.Student {
	db.mu.Lock()
	defer db.mu.Unlock()
	var students []model.Student
	for _, student := range db.students {
		if !student.Deleted {
			students = append(students, student)
		}
	}
	return students
}

func (db *InMemoryDB) GetStudent(id string) (model.Student, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()
	student, exists := db.students[id]
	if exists && !student.Deleted {
		return student, true
	}
	return model.Student{}, false
}

func (db *InMemoryDB) DeleteStudent(id string) bool {
	db.mu.Lock()
	defer db.mu.Unlock()
	student, exists := db.students[id]
	if exists && !student.Deleted {
		student.Deleted = true
		db.students[id] = student
		return true
	}
	return false
}
