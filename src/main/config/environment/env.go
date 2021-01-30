package environment

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	// Port returns the server listening port
	Port = getEnv("PORT", "7979")
	// JWT_SECRET returns the database hostname
	JWT_SECRET = getEnv("JWT_SECRET", "ichbindeinvater")
)

func getEnv(name string, fallback string) string {
	godotenv.Load()
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}
