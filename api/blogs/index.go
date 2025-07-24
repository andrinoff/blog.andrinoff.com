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

func getPosts(w http.ResponseWriter, r *http.Request) {
	var results []Post

	err := client.DB.From("posts").
		Select("title,created_at", "", false).
		Order("created_at", "desc", "").
		Execute(&results)

	if err != nil {
		log.Printf("Error executing Supabase query: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "apikey, X-Client-Info, Authorization")

	if err := json.NewEncoder(w).Encode(results); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "apikey, X-Client-Info, Authorization, Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}
	
	if r.Method == http.MethodGet {
		getPosts(w, r)
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
