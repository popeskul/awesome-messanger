package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "http://localhost:8081/graphql" // Убедитесь, что это правильный путь
	query := `
    {
        "query": "{ search(nickname: \"example_nickname\") { id nickname description } }"
    }`

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(query)))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Print status code
	log.Printf("Response status code: %d", resp.StatusCode)

	// Print response body for debugging
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	log.Printf("Response body: %s", body)

	// Attempt to decode JSON response
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}

	log.Printf("Decoded response: %v", result)
}
