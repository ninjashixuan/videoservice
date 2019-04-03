package defs

type User 

//model struct
type  User struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Pwd string `json:"pwd"`
}

type Video struct {
	 ID string `json:"id"`
	 AuthorId int `json:"author_id"`
	 Name string `json:"name"`
	 DisplayCtime string  `json:"display_ctime"`
}

type  Comment struct {
	ID string `json:"id"`
	Author_id string `json:"author_id"`
	Video_id string `json:"video_id"`
	content string `json:"content"`
}

type Session struct {
	Username string  `json:"username"`
	TTL int64 `json:"ttl"`
}


