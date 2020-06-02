package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"restapiwithgo/models"
)

var (
	file *os.File
	err  error
)

type response struct {
	//ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func create(name string) {
	if _, err := os.Stat(name); err == nil {
		fmt.Printf("File exists\n")
		return

	}

	file, err = os.Create(name)

	if err != nil {
		panic(err)
	}
	log.Println("File Created")
}

func Createfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var f models.File
	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	create(f.Name)
	res := response{
		Message: "File created successfully",
	}
	json.NewEncoder(w).Encode(res)

}

func rename(name, rname string) {
	err = os.Rename(name, rname)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("File Renamed")
}

func Renamefile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var f models.File
	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	rename(f.Name, f.Rname)
	res := response{
		Message: "File renamed successfully",
	}
	json.NewEncoder(w).Encode(res)

}

func delete(name string) {
	err = os.Remove(name)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s Deleted", name)
}

func Deletefile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var f models.File
	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	delete(f.Name)
	res := response{
		Message: "File deleted successfully",
	}
	json.NewEncoder(w).Encode(res)

}

func getall(path string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		return files, err
	}
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}

func Getfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	file, err := getall(".")

	if err != nil {
		log.Fatalf("Unable to get all files. %v", err)
	}
	json.NewEncoder(w).Encode(file)
}

// func get(path string) ([]string, error) {
// 	var files []string
// 	fileInfo, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		return files, err
// 	}
// 	for _, file := range fileInfo {
// 		files =
// 	}
// 	return files, nil
// }

// func Getfile(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	file, err := get("./")

// 	if err != nil {
// 		log.Fatalf("Unable to get file. %v", err)
// 	}
// 	json.NewEncoder(w).Encode(file)
// }
