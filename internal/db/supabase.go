// package db

// import (
// 	"github.com/nedpals/supabase-go"
// 	"os"
// )

// var Supabase *supabase.Client

// func ConnectSupabase() {
// 	Supabase = supabase.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"))
// }
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
