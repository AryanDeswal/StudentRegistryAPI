package controller

import (
	"StudentRegistryAPI/db"
	"StudentRegistryAPI/model"
	"StudentRegistryAPI/utils/logger"
	"StudentRegistryAPI/view"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student model.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		view.ReaderError(w, http.StatusBadRequest, err.Error())
		logger.ErrorLogger.Println("Failed to decode request body:", err)
	}

	id := db.GetDBInstance().CreateStudent(student)
	logger.InfoLogger.Println("Student created with ID:", id)
	view.RenderJSON(w, http.StatusCreated, map[string]string{"id": id})
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	students := db.GetDBInstance().GetAllStudents()
	logger.InfoLogger.Println("Retrieved all students")
	view.RenderJSON(w, http.StatusOK, students)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentId"]

	student, exists := db.GetDBInstance().GetStudent(studentID)
	if !exists {
		view.ReaderError(w, http.StatusNotFound, "Student not found")
		logger.ErrorLogger.Println("Student not found with ID:", studentID)
		return
	}

	logger.InfoLogger.Println("Retrived student with ID:", studentID)
	view.RenderJSON(w, http.StatusOK, student)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentId"]

	if !db.GetDBInstance().DeleteStudent(studentID) {
		view.ReaderError(w, http.StatusNotFound, "Student not found")
		logger.ErrorLogger.Println("Student not found with ID:", studentID)
		return
	}

	logger.InfoLogger.Println("Soft deleted student with ID:", studentID)
	w.WriteHeader(http.StatusNoContent)
}
