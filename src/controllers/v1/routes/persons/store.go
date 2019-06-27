package persons

import (
	Persons "golang-cell5/src/controllers/v1/models/persons"
	DB "golang-cell5/src/system/db"

	"encoding/json"
	"log"
	"net/http"
)

func Store(w http.ResponseWriter, r *http.Request) {
	var person Persons.Person

	person.Name = r.PostFormValue("name")
	person.Bio = r.PostFormValue("bio")
	person.Dob = r.PostFormValue("dob")

	DB.Store(db, &person)

	packet, err := json.Marshal(person)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse persons", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
