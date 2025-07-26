package main

import (
	"fmt"
	"os"

	"github.com/nedpals/supabase-go"
)

type Post struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}

func main() {
	// 1. Initialize the Supabase client
	// Replace with your actual Supabase project URL and anon key.
	// It's best practice to load these from environment variables.
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	if supabaseURL == "" || supabaseKey == "" {
		fmt.Println("Error: SUPABASE_URL and SUPABASE_KEY environment variables must be set.")
		fmt.Println("Example: export SUPABASE_URL=\"https://your-project-ref.supabase.co\"")
		fmt.Println("Example: export SUPABASE_KEY=\"your-anon-key\"")
		return
	}

	client := supabase.CreateClient(supabaseURL, supabaseKey)

	// 2. Fetch and sort the data
	// We will store the results in a slice of Post structs.
	 var results map[string]interface{}

	 err := supabase.DB.From("posts").Select("*").Single().Execute(&results)
 	 if err != nil {
   		 panic(err)
  	}

	
	for _, post := range results {
		fmt.Printf("ID: %d, CreatedAt: %s, Title: %s\n", post.ID, post.CreatedAt, post.Title)
	}
}
