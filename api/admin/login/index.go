package handler

import (
	"encoding/json"
	"net/http"
	"os"
)

// LoginRequest defines the structure for the login request body.
type LoginRequest struct {
	Password string `json:"password"`
}

// Handler processes the login request.
func Handler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	// Decode the request body into the LoginRequest struct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Get the master password from environment variables.
	masterPassword := os.Getenv("MASTER_PASSWORD")
	
	// Compare the provided password with the master password.
	if req.Password != masterPassword {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// If the password is correct, respond with a success message.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Authentication successful"})
}
