package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
	"videoservice/api/dbops"
	"videoservice/api/defs"
	"videoservice/api/session"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success: true, SessionId: id}

	if rsp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(rsp), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	if err := json.Unmarshal(res, &ubody); err != nil {
		log.Printf("%s", err)
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	uname := p.ByName("username")
	log.Printf("login url name %s", uname)
	log.Printf("login body name %s", ubody.Username)
	if uname != ubody.Username {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}

	pwd, err := dbops.GetUserCredential(ubody.Username)
	pass := bcrypt.CompareHashAndPassword([]byte(pwd), []byte(ubody.Pwd))
	if err != nil || pass != nil {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}

	id := session.GenerateNewSessionId(ubody.Username)
	ss := defs.SignedIn{Success: true, SessionId: id}
	if rsp, err := json.Marshal(ss); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		log.Printf("logined successed")
		sendNormalResponse(w, string(rsp), 200)
	}
}

func UserInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !validateUser(w, r) {
		log.Printf("not auth user")
		return
	}

	uname := p.ByName("username")
	uid, err := dbops.GetUser(uname)
	if err != nil {
		log.Printf("get user fail")
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	user := &defs.UserInfo{Id: uid.ID}
	if rsp, err := json.Marshal(user); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(rsp), 200)
	}
}

func AddVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
