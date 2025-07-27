package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/nedpals/supabase-go"
)

// Post struct defines the shape of our data.
type Post struct {
	ID       int    `json:"id,omitempty"`
	Date     string `json:"date,omitempty"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
}

// Handler is the main entry point for the Vercel Serverless Function.
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	client := supabase.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"))

	var newPost Post
	if err := json.NewDecoder(r.Body).Decode(&newPost); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}


	var results []Post
	err := client.DB.From("posts").Insert(newPost).Execute(&results)
	if err != nil {
		http.Error(w, "Could not add post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)
}