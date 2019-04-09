package main

import (
	"fmt"
	"net/http"
	"videoservice/api/defs"
	"videoservice/api/session"
)

var HEAD_FILED_SESSION = "X-Session-Id"
var HEAD_FILED_USERNAME = "X-User-Name"

func validateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEAD_FILED_SESSION)

	username, ok := session.IsSessionExpired(sid)

	if ok {
		return false
	}

	r.Header.Add(HEAD_FILED_USERNAME, username)
	return true
}

func validateUser(w http.ResponseWriter, r *http.Request) bool {
	username := r.Header.Get(HEAD_FILED_USERNAME)
	fmt.Println(username)
	if len(username) == 0 {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}

	return true
}
