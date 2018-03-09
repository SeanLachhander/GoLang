package main

import(
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)

/* Create a, 'Person' struct object that will contain data.
   The omitempty param simply means if the property is null,
   it will be excluded from the JSON data, and not show up as an empty string.
*/

type Person struct {
	ID string `json:"id, omitempty:`
	Firstname string `json:"Firstname,omitempty"`
	Lastname string `json:"Lastname,omitempty"`
	Address *Address `json:"Address,omitempty"`

}

/* Nested struct.
   Inside the Person struct, there's an Address property which is of type pointer,
   since omitempty will not work if not.
 */

type Address struct {
	City string `json:"city,omitempty"`
	State string `json:"state,omitempty"`

}

var people []Person // Global public slice of Person

// Get single record, use Mux library and get parameters passed in with the request.
// Loop over the global slice, and look for any IDs that match the ID found in the request parameters.
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {      // if match is found, use the JSON encoder to display it
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// Returns the person variable
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// Receiving JSON data
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(people)
}


func main(){
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Sean", Lastname: "Lachhander", Address: &Address{City: "Clifton Park", State: "New York"}})
	people = append(people, Person{ID: "2", Firstname: "Ian", Lastname: "Lachhander",})
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))
}


























