package persons

import (
	Persons "golang-cell5/src/controllers/v1/models/persons"

	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var offset int
	limit := int(25)

	if keys, ok := r.URL.Query()["limit"]; ok {
		limit, _ = strconv.Atoi(keys[0])
	}
	if keys, ok := r.URL.Query()["page"]; ok {
		page, _ := strconv.Atoi(keys[0])
		offset = (page - 1) * limit
	}

	persons, err := Persons.FindBy(db, limit, offset)

	packet, err := json.Marshal(persons)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse persons", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}