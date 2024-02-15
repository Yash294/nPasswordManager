package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"yash294/nPasswordManager/server/models"
)

const (
	StoragePath = "./passwords/"
	NumParts    = 3
)

func distributeAndSavePassword(userID, password string) error {
	passwordParts := splitPassword(password)

	for i, part := range passwordParts {
		partFilename := fmt.Sprintf("%s_%d.json", userID, i+1)
		partData, err := json.Marshal(models.Password{ID: userID, Part: i + 1, Password: []byte(part)})
		if err != nil {
			return fmt.Errorf("failed to marshal password part to JSON: %v", err)
		}
		if err := ioutil.WriteFile(StoragePath+partFilename, partData, 0644); err != nil {
			return fmt.Errorf("failed to write password part to file: %v", err)
		}
	}

	return nil
}

func splitPassword(password string) []string {
	partLength := (len(password) + NumParts - 1) / NumParts

	var passwordParts []string
	for i := 0; i < len(password); i += partLength {
		end := i + partLength
		if end > len(password) {
			end = len(password)
		}
		passwordParts = append(passwordParts, password[i:end])
	}

	return passwordParts
}

func main() {
	os.Mkdir(StoragePath, os.ModePerm)

	http.HandleFunc("/passwords", handlePasswords)
	http.ListenAndServe(":8080", nil)
}

func handlePasswords(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createPassword(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
	}
}

func createPassword(w http.ResponseWriter, r *http.Request) {
	var password models.Password
	if err := json.NewDecoder(r.Body).Decode(&password); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to decode request body: %v", err)
		return
	}

	if err := distributeAndSavePassword(password.ID, string(password.Password)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to distribute and save password: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}



// user could send password and I need to check if password exists
func verifyPassword() {}

// remove password http handler
func removePassword() {}

// list all of user passwords
func readHashedPasswords() {}