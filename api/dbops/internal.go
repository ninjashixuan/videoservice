package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"sync"
	"videoservice/api/defs"
)

func InsertSession(sid string, ttl int64, username string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmt, err := connDB.Prepare("INSERT INTO sessions (session_id, ttl, username) VALUE (?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(sid, ttlstr, username)
	if err != nil {
		return err
	}

	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmt, err := connDB.Prepare("SELECT ttl, username FROM sessions WHERE session_id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var ttl string
	var username string
	err = stmt.QueryRow(sid).Scan(&ttl, &username)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.Username = username
	} else {
		return nil, err
	}

	return ss, nil
}

func RetrieveAllSession() (*sync.Map, error) {
	ms := &sync.Map{}
	stmt, err := connDB.Prepare("SELECT * FROM sessions")
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil{
		log.Printf("%v", err)
		return nil, err
	}

	for rows.Next() {
		var id, ttl, username string
		if err := rows.Scan(&id, &ttl, &username); err != nil {
			log.Printf("%v", err)
			return nil, err
		}

		if TTL, err := strconv.ParseInt(ttl, 10, 64); err != nil {
			return nil, err
		} else {
			ss := &defs.SimpleSession{TTL:TTL, Username:username}
			ms.Store(id, ss)
			log.Printf("session_id:  %s TTL: %d", id, ss.TTL)
		}
	}

	return ms, nil
}

func DelSession(sid string) error {
	stmt, err := connDB.Prepare("DELETE FROM sessions WHERE session_id = ?")
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(sid)
	if err != nil {
		return err
	}

	return nil
}
