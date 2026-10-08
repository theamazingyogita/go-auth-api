package config

import (
	"fmt" // to print strings
	"log"
	"os"
	"time" // to get durations related tings

	"github.com/joho/godotenv"
	// to task to the os  - read the env file
	// print th log
)

// a struct class in dart needs no {} if no data inside and work with only param

type Config struct {
	Port      string        // whicj port to listem to
	DB        DBConfig      // all our db setting. - one more class holding all this
	JWTSecret string        // secret key to sign all login tokens
	JWTExpiry time.Duration //  token expiry time
}
type DBConfig struct {
	Host     string // where Postgres runs, e.g. "localhost"
	Port     string // Postgres port, usually "5432"
	User     string // DB username
	Password string // DB password
	Name     string // database name
	SSLMode  string // "disable" locally; "require" on most cloud databases
}

// a method using dbconfig
// here we use this fucniton feed it out env and pss it to  use to to connect it to out db
func (d DBConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Name, d.SSLMode)
}

func Load() *Config {

	if err := godotenv.Load(); err != nil {
		log.Println("no env found")
	}
	return &Config{
		Port: getEnv("PORT", "8080"),
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: os.Getenv("DB_PASSWORD"), // read as-is; a password can be empty locally
			Name:     mustEnv("DB_NAME"),       // required: the app stops if missing
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWTSecret: mustEnv("JWT_SECRET"), // required: never run without a secret
		JWTExpiry: 24 * time.Hour,        // tokens last one day
	}
}
func getEnv(key, fallback string) string {

	if v := os.Getenv(key); v != "" {

		return v
	}
	return fallback
}
func mustEnv(key string) string {
	// "must" is a Go naming convention: this function stops the app if it fails.

	v := os.Getenv(key) // := declares a new variable and infers its type

	if v == "" {
		log.Fatalf("%s is not set", key)
		// Prints the message and exits the program immediately.
		// It's better to fail at startup than to fail later on a real request.
	}
	return v
}
