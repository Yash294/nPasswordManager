package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"yash294/nPasswordManager/server/models"
)

const (
	StoragePath = "./passwords/"
	NumParts    = 3
)

func distributeAndSavePassword(userID, passwordID, password string) error {
	passwordParts := splitPassword(password)

	for i, part := range passwordParts {
		partFilename := fmt.Sprintf("%s_%s_%d.json", userID, passwordID, i+1)
		partData, err := json.Marshal(models.Password{UserID: userID, PasswordID: passwordID, Part: i + 1, Password: part})
		if err != nil {
			return fmt.Errorf("failed to marshal password part to JSON: %v", err)
		}
		file, err := os.Create(StoragePath + partFilename)
		if err != nil {
			return fmt.Errorf("failed to create password part file: %v", err)
		}
		defer file.Close()
		if _, err := file.Write(partData); err != nil {
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
	var passwordPart models.Password
	if err := json.NewDecoder(r.Body).Decode(&passwordPart); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to decode request body: %v", err)
		return
	}

	if err := distributeAndSavePassword(passwordPart.UserID, passwordPart.PasswordID, passwordPart.Password); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to distribute and save password: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func verifyPassword(userID, passwordID string) (bool, error) {
	files, err := os.ReadDir(StoragePath)
	if err != nil {
		return false, fmt.Errorf("failed to read passwords directory: %v", err)
	}
	for _, file := range files {
		filename := file.Name()
		if strings.HasPrefix(filename, userID+"_"+passwordID+"_") && strings.HasSuffix(filename, ".json") {
			return true, nil
		}
	}
	return false, nil
}

func removePassword(userID, passwordID string) error {
	files, err := os.ReadDir(StoragePath)
	if err != nil {
		return fmt.Errorf("failed to read passwords directory: %v", err)
	}
	for _, file := range files {
		filename := file.Name()
		if strings.HasPrefix(filename, userID+"_"+passwordID+"_") && strings.HasSuffix(filename, ".json") {
			err := os.Remove(StoragePath + filename)
			if err != nil {
				return fmt.Errorf("failed to remove password part file %s: %v", filename, err)
			}
		}
	}
	return nil
}

func readHashedPasswords(userID, passwordID string) ([]models.Password, error) {
	var passwordParts []models.Password

	files, err := os.ReadDir(StoragePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read passwords directory: %v", err)
	}
	for _, file := range files {
		filename := file.Name()
		if strings.HasPrefix(filename, userID+"_"+passwordID+"_") && strings.HasSuffix(filename, ".json") {
			fileContents, err := os.ReadFile(StoragePath + filename)
			if err != nil {
				return nil, fmt.Errorf("failed to read file %s: %v", filename, err)
			}
			var passwordPart models.Password
			if err := json.Unmarshal(fileContents, &passwordPart); err != nil {
				return nil, fmt.Errorf("failed to unmarshal JSON from file %s: %v", filename, err)
			}
			passwordParts = append(passwordParts, passwordPart)
		}
	}
	return passwordParts, nil
}

func main() {
	os.Mkdir(StoragePath, os.ModePerm)

	http.HandleFunc("/passwords", handlePasswords)
	http.ListenAndServe(":8080", nil)
}