package http

import (
	"encoding/json"
	"net/http"

	"github.com/popeskul/awesome-messanger/services/friend/internal/core/models"
	"github.com/popeskul/awesome-messanger/services/friend/internal/core/ports"
)

type Validator interface {
	Struct(interface{}) error
}

type handler struct {
	useCases  ports.UseCase
	validator Validator
}

func NewHandler(useCases ports.UseCase, validator Validator) ports.HandlerFriends {
	return &handler{
		useCases:  useCases,
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
// @Success 201 {object} Friend "Friend"
// @Failure 400 {string} string "Invalid request payload"
// @Failure 400 {string} string "Validation failed"
// @Failure 500 {string} string "Failed to add friend"
// @Router /add-friend [post]
func (h *handler) PostAddFriend(w http.ResponseWriter, r *http.Request) {
	var req FriendRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.useCases.FriendUseCase().AddFriend(r.Context(), req.ConvertToModel())
	if err != nil {
		http.Error(w, "Failed to add friend", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	if err = json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
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
func (h *handler) GetFriends(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")

	friends, err := h.useCases.FriendUseCase().GetFriends(r.Context())
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
func (h *handler) PostRespondFriendRequest(w http.ResponseWriter, r *http.Request) {
	var req PostRespondFriendRequestJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err := h.useCases.FriendUseCase().RespondToFriendRequest(r.Context(), req.ConvertToModel())
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
func (h *handler) GetLive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// GetReady godoc
// @Summary Get ready
// @Description Get ready
// @Tags health
// @Success 200 {string} string "OK"
// @Router /ready [get]
func (h *handler) GetReady(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
