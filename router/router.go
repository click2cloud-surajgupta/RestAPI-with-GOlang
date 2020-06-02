package router

import (
	"restapiwithgo/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/create_file", middleware.Createfile).Methods("POST")
	router.HandleFunc("/delete_file_name", middleware.Deletefile).Methods("DELETE")
	router.HandleFunc("/change_file_name", middleware.Renamefile).Methods("PUT")
	router.HandleFunc("/get_all_files", middleware.Getfiles).Methods("GET")
	// router.HandleFunc("/get_file", middleware.Getfile).Methods("GET")

	return router
}
