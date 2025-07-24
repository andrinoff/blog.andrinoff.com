// File: api/index.go
package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/supabase-community/supabase-go"
)

// Post represents the structure of your 'posts' table in Supabase.
// Make sure the json tags match your column names exactly.
type Post struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt string `json:"created_at"`
}

// Handler is the main Vercel serverless function entrypoint.
func Handler(w http.ResponseWriter, r *http.Request) {
	// --- 1. Set up Supabase Client ---
	// Retrieve Supabase URL and Key from environment variables.
	// These will be set in your Vercel project settings.
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_ANON_KEY")

	if supabaseURL == "" || supabaseKey == "" {
		http.Error(w, "Server configuration error: Supabase environment variables not set.", http.StatusInternalServerError)
		fmt.Println("Error: SUPABASE_URL or SUPABASE_ANON_KEY environment variable is not set.")
		return
	}

	// Initialize the Supabase client.
	client, err := supabase.NewClient(supabaseURL, supabaseKey, nil)
	if err != nil {
		http.Error(w, "Failed to initialize Supabase client.", http.StatusInternalServerError)
		fmt.Println("Error initializing Supabase client:", err)
		return
	}

	// --- 2. Fetch Data from Supabase ---
	var results []Post
	// Execute the query to get all posts from the "posts" table.
	// ".order("created_at", &supabase.OrderOpts{Ascending: false})" sorts the results
	// from the newest entry to the oldest.
	err = client.DB.From("posts").Select("*").Order("created_at", &supabase.OrderOpts{Ascending: false}).Execute(&results)

	if err != nil {
		http.Error(w, "Failed to fetch posts from the database.", http.StatusInternalServerError)
		fmt.Println("Error fetching data:", err)
		return
	}

	// --- 3. Send JSON Response ---
	// Set the content type header to application/json.
	w.Header().Set("Content-Type", "application/json")

	// Encode the fetched posts slice into a JSON array and write it to the response.
	if err := json.NewEncoder(w).Encode(results); err != nil {
		http.Error(w, "Failed to encode response.", http.StatusInternalServerError)
		fmt.Println("Error encoding JSON response:", err)
	}
}
