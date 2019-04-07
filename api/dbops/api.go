package dbops

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"videoservice/api/defs"
	"videoservice/api/util"
)

func AddUserCredential(username, pwd string) error {
	stmt, err := connDB.Prepare("INSERT INTO users (username, pwd) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(username, pwd)
	if err != nil {
		return err
	}

	defer stmt.Close()

	return nil
}

func GerUserCredential(username string) (string, error) {
	stmt, err := connDB.Prepare("SELECT pwd FROM users WHERE username = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	err = stmt.QueryRow(username).Scan(&pwd)

	if err != nil && err != sql.ErrNoRows {
		log.Printf("select fail ")
		return "", err
	}

	defer stmt.Close()
	return pwd, nil
}

func DeleteUser(username, pwd string) error {
	stmt, err := connDB.Prepare("DELETE FROM users WHERE username = ? AND pwd = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(username, pwd)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return nil
}

func GetUser(username string) (*defs.User, error) {
	stmt, err := connDB.Prepare("SELECT id, pwd FROM user WHERE username = ?")
	if err != nil {
		return nil, err
	}

	var id int
	var pwd string

	err = stmt.QueryRow(username).Scan(&id, &pwd)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("query fail %v", err)
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	res := &defs.User{id, username, pwd}
	defer stmt.Close()

	return res, nil
}

func AddNewVideo(author_id int, name string) (*defs.Video, error) {
	vid := util.Newuuid()

	stmt, err := connDB.Prepare("INSERT INTO video (id, author_id, info, display_ctime) VALUE (?, ?, ?, ? )")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	t := time.Now()
	ctime := t.Format("Jan 2 2006, 15:04:05")

	_, err = stmt.Exec(vid, author_id, name, ctime)
	if err != nil {
		log.Printf("insert into fail %v", err)
		return nil, err
	}

	res := &defs.Video{
		ID:           vid,
		AuthorId:     author_id,
		Name:         name,
		DisplayCtime: ctime,
	}

	return res, nil
}

func GetVideo(vid string) (*defs.Video, error) {
	stmt, err := connDB.Prepare("SELECT author_id, info, display FROM video WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var (
		author_id    int
		name         string
		diaplayctime string
	)

	err = stmt.QueryRow(vid).Scan(&author_id, &name, &diaplayctime)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("query fail %v", err)
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	res := &defs.Video{
		ID:           vid,
		AuthorId:     author_id,
		Name:         name,
		DisplayCtime: diaplayctime,
	}

	return res, nil

}

func DeleteVideo(vid string) error {
	stmt, err := connDB.Prepare("DELETE FROM video WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(vid)

	if err != nil && err != sql.ErrNoRows {
		log.Printf("delete rows fail %v", err)
		return err
	}

	if err == sql.ErrNoRows {
		return nil
	}
	return nil
}

//
//func ListVideo(uname string) ([]*defs.Video, error) {
//	stmt, err := connDB.Prepare("")
//}

func AddComment(vid string, aid int, content string) error {
	id := util.Newuuid()
	stmt, err := connDB.Prepare("INSERT INTO comments (id, author_id, video_id, content) VALUE (?, ?, ?, ?)")
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}

	_, err = stmt.Exec(id, aid, vid, content)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}

	defer stmt.Close()

	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmt, err := connDB.Prepare(`SELECT comments.id, users.username, comments.content FROM
        comments INNER JOIN users ON users.id = comments.author_id WHERE comments.video_id = ? AND
        comments.create_time > FROM_UNIXTIME(?) AND comments.create_time <= FROM_UNIXTIME(?)
        ORDER BY comments.create_time DESC
    `)

	defer stmt.Close()

	var res []*defs.Comment

	if err != nil {
		return res, err
	}

	rows, err := stmt.Query(vid, from, to)
	if err != nil {
		log.Printf("select fail %v", err)
		return res, err
	}

	for rows.Next() {
		var comment_id, username, context string
		if err := rows.Scan(&comment_id, &username, &context); err != nil {
			return res, err
		}
		cmt := &defs.Comment{ID:comment_id, Author_name:username, Video_id:vid, Content:context}
		res = append(res, cmt)
	}
	return  res, nil
}
