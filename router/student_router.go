package router

import (
	"StudentRegistryAPI/controller"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	// api routes
	r.HandleFunc("/student/v1/students", controller.CreateStudent).Methods("POST")
	r.HandleFunc("/student/v1/students", controller.GetAllStudents).Methods("GET")
	r.HandleFunc("/student/v1/students/{studentId}", controller.GetStudent).Methods("GET")
	r.HandleFunc("/student/v1/students/{studentId}", controller.DeleteStudent).Methods("DELETE")
	
	// catch-all route (default route)
	r.PathPrefix("/").HandlerFunc(controller.DefaultHandler)
	
	return r
}
