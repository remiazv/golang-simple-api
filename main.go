package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var people []Person

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Person{})
}
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person

	_ = json.NewDecoder(r.Body).Decode(&person)

	person.ID = params["id"]
	people = append(people, person)

	json.NewEncoder(w).Encode(people)
}
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var deletedPerson Person
	for index, item := range people {
		if item.ID == params["id"] {
			deletedPerson = item
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(deletedPerson)
	}
}

func PopulatePeople() {
	people = append(people, Person{
		ID:        "1",
		Firstname: "Joh",
		Lastname:  "Doe",
		Address: &Address{
			City:  "City X",
			State: "State X",
		},
	})

	people = append(people, Person{
		ID:        "2",
		Firstname: "Mar",
		Lastname:  "Help",
		Address: &Address{
			City:  "City A",
			State: "State A",
		},
	})

	people = append(people, Person{
		ID:        "3",
		Firstname: "Clark",
		Lastname:  "Kent",
		Address: &Address{
			City:  "Metropolis",
			State: "State X",
		},
	})
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/person", GetPeople).Methods("GET")
	router.HandleFunc("/person/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/person/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/person/{id}", DeletePerson).Methods("DELETE")

	log.Println("Server listening at: ", "8000")

	PopulatePeople()

	log.Fatal(http.ListenAndServe(":8000", router))
}
