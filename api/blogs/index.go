package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/nedpals/supabase-go"
)

// Post struct defines the shape of our data.
type Post struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}

// Handler is the main entry point for the Vercel Serverless Function.
// It handles incoming HTTP requests.
func Handler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 1. Initialize the Supabase client from environment variables
	client := supabase.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"))

	// 2. Fetch all records from the "posts" table
	var results []Post 
	err := client.DB.From("posts").Select("*").Execute(&results)

	

	if err != nil {
		log.Printf("Error fetching data from Supabase: %v", err)
		http.Error(w, "Could not fetch posts", http.StatusInternalServerError)
		return
	} else {
		log.Printf("Fetched %d posts from Supabase", len(results))
		log.Printf ("Posts: %+v", results)
	}

	// 3. Set the response header and encode the results as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(results); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Could not encode response", http.StatusInternalServerError)
	}
}