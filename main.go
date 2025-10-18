package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// CatFactResponse represents the response from the Cat Facts API
type CatFactResponse struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

// User represents the user information
type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Stack string `json:"stack"`
}

// ProfileResponse represents the API response structure
type ProfileResponse struct {
	Status    string `json:"status"`
	User      User   `json:"user"`
	Timestamp string `json:"timestamp"`
	Fact      string `json:"fact"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// fetchCatFact fetches a random cat fact from the Cat Facts API
func fetchCatFact() (string, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get("https://catfact.ninja/fact")
	if err != nil {
		return "", fmt.Errorf("failed to fetch cat fact: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("cat facts API returned status: %d", resp.StatusCode)
	}

	var catFact CatFactResponse
	if err := json.NewDecoder(resp.Body).Decode(&catFact); err != nil {
		return "", fmt.Errorf("failed to decode cat fact: %w", err)
	}

	return catFact.Fact, nil
}

// profileHandler handles the /me endpoint
func profileHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight requests
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Only allow GET requests
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{
			Status:  "error",
			Message: "Method not allowed. Use GET.",
		})
		return
	}

	log.Printf("Received request from %s at %s", r.RemoteAddr, time.Now().UTC().Format(time.RFC3339))

	// Fetch cat fact
	catFact, err := fetchCatFact()
	if err != nil {
		log.Printf("Error fetching cat fact: %v", err)
		catFact = "Unable to fetch cat fact at this moment: '( Please try again later."
	}

	// Get user information from environment variables with defaults
	email := getEnv("USER_EMAIL", "your.email@example.com")
	name := getEnv("USER_NAME", "Your Full Name")
	stack := getEnv("USER_STACK", "Go/Native HTTP")

	// Create response
	response := ProfileResponse{
		Status: "success",
		User: User{
			Email: email,
			Name:  name,
			Stack: stack,
		},
		Timestamp: time.Now().UTC().Format(time.RFC3339Nano),
		Fact:      catFact,
	}

	// Set content type and return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

// healthHandler handles health check endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
		"time":   time.Now().UTC().Format(time.RFC3339),
	})
}

// getEnv gets an environment variable with a fallback default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	// Get port from environment variable or use default
	port := getEnv("PORT", "8080")

	// Register handlers
	http.HandleFunc("/me", profileHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Profile API - Use GET /me to retrieve profile information",
		})
	})

	// Start server
	addr := "0.0.0.0:" + port
	log.Printf("Server starting on %s", addr)
	log.Printf("Environment: USER_EMAIL=%s, USER_NAME=%s, USER_STACK=%s",
		getEnv("USER_EMAIL", "not set"),
		getEnv("USER_NAME", "not set"),
		getEnv("USER_STACK", "not set"))

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
