package session

import (
	"sync"
	"time"
	"videoservice/api/dbops"
	"videoservice/api/defs"
	"videoservice/api/util"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

func DelExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DelSession(sid)
}

func LoadAllSession() {
	r, err := dbops.RetrieveAllSession()
	if err != nil {
		return
	}

	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

func GenerateNewSessionId(username string) string {
	id := util.Newuuid()
	ct := nowInMilli() + 30 * 60 * 1000

	ss := &defs.SimpleSession{Username: username, TTL: ct}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ct, username)

	return id
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	ct := nowInMilli()
	if ok {
		if ss.(*defs.SimpleSession).TTL < ct {
			DelExpiredSession(sid)
			return "", true
		}

		return ss.(*defs.SimpleSession).Username, false
	} else {
		ss, err := dbops.RetrieveSession(sid)
		if err != nil || ss == nil {
			return "", true
		}

		if ss.TTL < ct {
			return "", true
		}

		sessionMap.Store(sid, ss)
		return  ss.Username, false
	}

	return "", true
}


