package dbops

import (
	"fmt"
	"log"
	"testing"
	"time"
)

var tempvid string

func clearTables() {
	connDB.Exec("truncate users")
	connDB.Exec("truncate comments")
	connDB.Exec("truncate sessions")
	connDB.Exec("truncate video")
}

//func MainTest(m *testing.M){
//	clearTables()
//	m.Run()
//	//clearTables()
//}

func aTestUserWorkFlow(t *testing.T)  {
	t.Run("add",testAddUserCredential)
	t.Run("get",testGerUserCredential)
	//t.Run("del", testDeleteCredential)
}

func testAddUserCredential(t *testing.T) {
	err := AddUserCredential("username4", "pwd")
	if err != nil {
		t.Errorf("adduser fail %v", err)
	}
}

func testGerUserCredential(t *testing.T) {
	pwd, err := GerUserCredential("username1")
	if pwd != "pwd" || err != nil {
		t.Errorf("getuser fail %v", err)
	}
}

func testDeleteCredential(t *testing.T) {
	err := DeleteUser("username", "pwd")
	if err != nil {
		t.Errorf("delete user %v", err)
	}
}

func testVideoWorkFlow(t *testing.T)  {
	//t.Run("addvideo", testAddNewVideo)
	t.Run("delete", testDeleteVideo)
}

func testAddNewVideo(t *testing.T) {
	info, err := AddNewVideo(1, "myvidoe")
	if err != nil {
		log.Printf("add fail %v", err)
	}
	tempvid = info.ID
}

func testDeleteVideo(t *testing.T) {
	err := DeleteVideo("9133b895-bfe4-46f9-9771-e8555055dc51")
	if err != nil {
		log.Printf("delele fail %v", err)
	}
}

func aTestAddComment(t *testing.T) {
	vid := "123456789"
	err := AddComment(vid, 7, "this is new comment tt")
	if err != nil {
		log.Printf("newed comment fail %v", err)
	}
}

func TestListComments(t *testing.T) {
	vid := "12345678"
	from := int(time.Now().Unix() - 360000)
	to := int(time.Now().Unix())
	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("get commentlist fail %v", err)
	}

	for i, comment := range res {
		fmt.Printf("commet: %d %v \n", i, comment)
	}

}