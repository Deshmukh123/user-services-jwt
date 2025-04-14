package db

import (
	"net/http"
	"os"
)

var (
	SupabaseURL = os.Getenv("SUPABASE_URL")
	SupabaseKey = os.Getenv("SUPABASE_KEY")
	HttpClient  = &http.Client{}
)
