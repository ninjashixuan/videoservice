package dbops

import (
	"log"
	"testing"
)

var tempvid string

func clearTables() {
	connDB.Exec("truncate users")
	connDB.Exec("truncate comments")
	connDB.Exec("truncate sessions")
	connDB.Exec("truncate video")
}

func MainTest(m *testing.M){
	clearTables()
	m.Run()
	//clearTables()
}

func ATestUserWorkFlow(t *testing.T)  {
	t.Run("add",testAddUserCredential)
	t.Run("get",testGerUserCredential)
	//t.Run("del", testDeleteCredential)
}

func testAddUserCredential(t *testing.T) {
	err := AddUserCredential("username", "pwd")
	if err != nil {
		t.Errorf("adduser fail %v", err)
	}
}

func testGerUserCredential(t *testing.T) {
	pwd, err := GerUserCredential("username")
	if pwd != "pwd" || err != nil {
		t.Errorf("getuser fail %v", err)
	}
}

func testDeleteCredential(t *testing.T) {
	err := DeleteCredential("username", "pwd")
	if err != nil {
		t.Errorf("delete user %v", err)
	}
}

func TestVideoWorkFlow(t *testing.T)  {
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
