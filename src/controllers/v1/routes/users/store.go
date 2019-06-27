package users

import (
	Users "golang-cell5/src/controllers/v1/models/users"
	DB "golang-cell5/src/system/db"
	Passwords "golang-cell5/src/system/passwords"

	"encoding/json"
	"log"
	"net/http"
)

func Store(w http.ResponseWriter, r *http.Request) {
	var user Users.User

	user.First = r.PostFormValue("first")
	user.Last = r.PostFormValue("last")
	user.Email = r.PostFormValue("email")
	password := r.PostFormValue("password")

	encryptedPassword, err := Passwords.Encrypt(password)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to encrypt password", http.StatusInternalServerError)
		return
	}

	user.Password = encryptedPassword

	if err = DB.Store(db, &user); err != nil {
		log.Println(err)
		http.Error(w, "Unable to store users", http.StatusInternalServerError)
		return
	}

	packet, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse users", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
