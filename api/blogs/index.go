package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/nedpals/supabase-go"
)

// Post struct defines the shape of our data.
type Post struct {
	ID       int    `json:"id"`
	Date     string `json:"date"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Slug     string `json:"slug"`
}

// Generates a URL-friendly slug from a title
func generateSlug(title string) string {
	return strings.ToLower(strings.ReplaceAll(title, " ", "-"))
}

// Handler is the main entry point for the Vercel Serverless Function.
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	client := supabase.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"))

	var results []Post
	err := client.DB.From("posts").Select("*").Execute(&results)
	if err != nil {
		log.Printf("Error fetching posts from Supabase: %v", err)
		http.Error(w, "Could not fetch posts", http.StatusInternalServerError)
		return
	}

	// Generate slugs for all posts
	for i := range results {
		results[i].Slug = generateSlug(results[i].Title)
	}

	slug := r.URL.Query().Get("slug")
	w.Header().Set("Content-Type", "application/json") // Set header for all responses

	if slug != "" {
		var foundPost *Post
		for i, post := range results {
			if post.Slug == slug {
				foundPost = &results[i]
				break
			}
		}

		if foundPost != nil {
			if err := json.NewEncoder(w).Encode(foundPost); err != nil {
				log.Printf("Error encoding JSON response: %v", err)
				http.Error(w, "Could not encode response", http.StatusInternalServerError)
			}
		} else {
			http.Error(w, "Could not find post", http.StatusNotFound)
		}

	} else {
		if err := json.NewEncoder(w).Encode(results); err != nil {
			log.Printf("Error encoding JSON response: %v", err)
			http.Error(w, "Could not encode response", http.StatusInternalServerError)
		}
	}
}