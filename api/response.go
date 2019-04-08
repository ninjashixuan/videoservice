package main

import (
	"encoding/json"
	"io"
	"net/http"
	"videoservice/api/defs"
)

func sendErrorResponse(w http.ResponseWriter, rps defs.ErrResponse) {
	w.WriteHeader(rps.HttpSC)

	reStr, _ := json.Marshal(&rps.Error)
	io.WriteString(w, string(reStr))
}

func sendNormalResponse(w http.ResponseWriter, rsp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, rsp)
}
