package session

import "sync"

var sessionMap  *sync.Map

func init()  {
	sessionMap =  &sync.Map{}
}

