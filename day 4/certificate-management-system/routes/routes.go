package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"

	"certificate-management-system/controllers"
)

func SetupRoutes(router *mux.Router, db *mongo.Database) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Certificate Management System!"))
	}).Methods("GET")

	    router.HandleFunc("/create-certificate", func(w http.ResponseWriter, r *http.Request) {
        controllers.CreateCertificate(db, w, r) // Pass the http.ResponseWriter and *http.Request arguments
    }).Methods("POST")


  router.HandleFunc("/get-all-certificates", func(w http.ResponseWriter, r *http.Request) {
        controllers.GetAllCertificates(db, w) 
    }).Methods("GET")

	router.HandleFunc("/update-certificate", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateCertificate(db, w, r)
	}).Methods("PUT")

	router.HandleFunc("/delete-certificate", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteCertificate(db, w, r)
	}).Methods("DELETE")

	router.HandleFunc("/getCertificateById", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetCertificateByID(db, w, r)
	}).Methods("GET")

	
}
