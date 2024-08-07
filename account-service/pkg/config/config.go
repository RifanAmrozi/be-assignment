package config

import (
	"os"

	"github.com/nedpals/supabase-go"
)

var SupabaseClient *supabase.Client

func InitSupabase() {
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	SupabaseClient = supabase.CreateClient(supabaseURL, supabaseKey)
}
