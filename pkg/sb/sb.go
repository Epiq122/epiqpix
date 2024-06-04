package sb

import (
	"os"

	"github.com/nedpals/supabase-go"
)

var Client *supabase.Client

func Init() error {
	sbHost := os.Getenv("SUPABASE_URL")
	sbKey := os.Getenv("SUPABASE_KEY")
	Client = supabase.CreateClient(sbHost, sbKey)
	return nil
}
