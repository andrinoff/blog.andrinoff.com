package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	
	supabase "github.com/nedpals/supabase-go"
)

type Post struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}


var client *supabase.Client



func init() {
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_ANON_KEY")

	if supabaseURL == "" || supabaseKey == "" {
		log.Fatal("FATAL: SUPABASE_URL and SUPABASE_ANON_KEY environment variables must be set.")
	}

	client = supabase.CreateClient(supabaseURL, supabaseKey)

	fmt.Println("Successfully initialized Supabase client!")
}

// getPosts is the core logic for fetching posts using the Supabase client.
func getPosts(w http.ResponseWriter, r *http.Request) {
	// Create a slice to hold the results from the database.
	var results []Post

	// Build and execute the query using the supabase-go library's DB client.
	// This is equivalent to: SELECT title, created_at FROM posts ORDER BY created_at DESC
	err := client.DB.From("posts").
		Select("title,created_at", "", false).
		Order("created_at", &supabase.OrderOpts{Ascending: false}).
		Execute(&results)

	if err != nil {
		log.Printf("Error executing Supabase query: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the response headers.
	w.Header().Set("Content-Type", "application/json")
	// Add a CORS header to allow requests from any origin.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "apikey, X-Client-Info, Authorization")


	// Encode the results slice to JSON and send it as the response.
	if err := json.NewEncoder(w).Encode(results); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Handler is the public function that Vercel will invoke for each request.
// It acts as a simple router.
func Handler(w http.ResponseWriter, r *http.Request) {
	// For a serverless function, we should handle preflight OPTIONS requests for CORS.
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-control-Allow-Headers", "apikey, X-Client-Info, Authorization, Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}
	
	// Route GET requests to our main logic.
	if r.Method == http.MethodGet {
		getPosts(w, r)
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
