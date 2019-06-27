package persons

import (
	Persons "golang-cell5/src/controllers/v1/models/persons"
	DB "golang-cell5/src/system/db"

	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	newId, _ := strconv.Atoi(id)
	person := Persons.Person{Id: int64(newId)}
	person.Name = r.PostFormValue("name")
	person.Bio = r.PostFormValue("bio")
	person.Dob = r.PostFormValue("dob")

	if err := DB.Update(db, person.Id, &person); err != nil {
		log.Println(err)
		http.Error(w, "Unable to get person", http.StatusInternalServerError)
		return
	}

	packet, err := json.Marshal(person)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse person", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
