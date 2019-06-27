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

func Destroy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	newId, _ := strconv.Atoi(id)
	person := Persons.Person{Id: int64(newId)}

	if err := DB.Destroy(db, person.Id, &person); err != nil {
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
