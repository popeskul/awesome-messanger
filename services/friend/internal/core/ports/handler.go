package ports

import "net/http"

type HandlerFriends interface {
	PostAddFriend(w http.ResponseWriter, r *http.Request)
	GetFriends(w http.ResponseWriter, r *http.Request)
	PostRespondFriendRequest(w http.ResponseWriter, r *http.Request)
	GetLive(w http.ResponseWriter, r *http.Request)
	GetReady(w http.ResponseWriter, r *http.Request)
}

type Handler interface {
	HandlerFriends
}
