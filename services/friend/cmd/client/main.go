package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type FriendRequest struct {
	UserId   string `json:"user_id" validate:"required"`
	FriendId string `json:"friend_id" validate:"required"`
}

type PostRespondFriendRequestJSONRequestBody struct {
	UserId    string `json:"user_id"`
	FriendId  string `json:"friend_id"`
	Responded bool   `json:"responded"`
}

const (
	baseURL = "http://localhost:8080"
)

func main() {
	client := &http.Client{Timeout: 10 * time.Second}

	addFriendReq := FriendRequest{
		UserId:   "user1",
		FriendId: "user2",
	}
	addFriend(client, addFriendReq)

	getFriends(client, "user1")

	respondFriendRequest(client, PostRespondFriendRequestJSONRequestBody{
		UserId:    "user1",
		FriendId:  "user2",
		Responded: true,
	})
}

func addFriend(client *http.Client, req FriendRequest) {
	url := baseURL + "/add-friend"
	body, err := json.Marshal(req)
	if err != nil {
		log.Fatalf("Error marshalling request: %v", err)
	}

	resp, err := client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Printf("Unexpected status code: %d", resp.StatusCode)
		responseBody, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Response body: %s", responseBody)
		log.Fatalf("Failed to add friend")
	}

	log.Println("Friend added successfully")
}

func getFriends(client *http.Client, userId string) {
	url := baseURL + "/friends?user_id=" + userId

	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code: %d", resp.StatusCode)
		responseBody, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Response body: %s", responseBody)
		log.Fatalf("Failed to get friends")
	}

	var friends []FriendRequest
	if err := json.NewDecoder(resp.Body).Decode(&friends); err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}

	log.Printf("Friends: %v", friends)
}

func respondFriendRequest(client *http.Client, req PostRespondFriendRequestJSONRequestBody) {
	url := baseURL + "/respond-friend-request"
	body, err := json.Marshal(req)
	if err != nil {
		log.Fatalf("Error marshalling request: %v", err)
	}

	resp, err := client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code: %d", resp.StatusCode)
		responseBody, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Response body: %s", responseBody)
		log.Fatalf("Failed to respond to friend request")
	}

	log.Println("Response to friend request sent successfully")
}
