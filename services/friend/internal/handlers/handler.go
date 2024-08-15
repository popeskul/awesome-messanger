package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/popeskul/awesome-messanger/services/friend/internal/models"
)

type Service interface {
	GetFriends(ctx context.Context) ([]*models.Friend, error)
	AddFriend(ctx context.Context, inout *models.Friend) (*models.Friend, error)
	RespondToFriendRequest(ctx context.Context, inout *models.Friend) (*models.Friend, error)
}

type Validator interface {
	Struct(interface{}) error
}

type Handler struct {
	service   Service
	validator Validator
}

func NewHandler(service Service, validator Validator) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}

// PostAddFriend godoc
// @Summary Add a friend
// @Description Add a friend
// @Tags friends
// @Accept json
// @Produce json
// @Param body body FriendRequest true "Friend request"
// @Success 201 {string} string "Friend added"
// @Failure 400 {string} string "Invalid request payload"
// @Failure 400 {string} string "Validation failed"
// @Failure 500 {string} string "Failed to add friend"
// @Router /add-friend [post]
func (h *Handler) PostAddFriend(w http.ResponseWriter, r *http.Request) {
	var req FriendRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err := h.service.AddFriend(r.Context(), req.ConvertToModel())
	if err != nil {
		http.Error(w, "Failed to add friend", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetFriends godoc
// @Summary Get friends
// @Description Get friends
// @Tags friends
// @Accept json
// @Produce json
// @Param userId query string true "User ID"
// @Success 200 {array} Friend "Friends"
// @Failure 500 {string} string "Failed to retrieve friends"
// @Router /friends [get]
func (h *Handler) GetFriends(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")

	friends, err := h.service.GetFriends(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve friends", http.StatusInternalServerError)
		return
	}

	var result []models.Friend
	for _, friend := range friends {
		if friend.UserId == userId {
			result = append(result, models.Friend{
				UserId:   friend.UserId,
				FriendId: friend.FriendId,
			})
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// PostRespondFriendRequest godoc
// @Summary Respond to a friend request
// @Description Respond to a friend request
// @Tags friends
// @Accept json
// @Produce json
// @Param body body PostRespondFriendRequestJSONRequestBody true "Friend request response"
// @Success 200 {string} string "Friend request responded"
// @Failure 400 {string} string "Invalid request payload"
// @Failure 400 {string} string "Validation failed"
// @Failure 500 {string} string "Failed to respond to friend request"
// @Router /respond-friend-request [post]
func (h *Handler) PostRespondFriendRequest(w http.ResponseWriter, r *http.Request) {
	var req PostRespondFriendRequestJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err := h.service.RespondToFriendRequest(r.Context(), req.ConvertToModel())
	if err != nil {
		http.Error(w, "Failed to respond to friend request", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetLive godoc
// @Summary Get live
// @Description Get live
// @Tags health
// @Success 200 {string} string "OK"
// @Router /live [get]
func (h *Handler) GetLive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// GetReady godoc
// @Summary Get ready
// @Description Get ready
// @Tags health
// @Success 200 {string} string "OK"
// @Router /ready [get]
func (h *Handler) GetReady(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
