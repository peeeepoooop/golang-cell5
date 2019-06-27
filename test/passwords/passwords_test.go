package password

import (
	"golang-cell5/src/system/passwords"

	"log"
	"testing"
)

func init() {
	log.Println("Testing password")
}

func TestBasicLog(t *testing.T) {
	pass := "test"
	hash, err := passwords.Encrypt(pass)
	if err != nil {
		t.Error(err.Error())
	}
	log.Println(hash)
	ok := passwords.IsValid(hash, pass)
	if !ok {
		t.Error("Password not matching... hashing it not working")
	}
	log.Println("Successfully tested hashing!")
}
