package defs

type UserCredential struct {
	Username string `json:"username"`
	Pwd string `json:"pwd"`
}

type NewComment struct {
	AuthorId int  `json:"author_id"`
	Context string `json:"context"`
}

type NewVideo struct {
	AuthorId int `json:"author_id"`
	Name string `json:"name"`
}


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
	Author_name string `json:"author_name"`
	Video_id string `json:"video_id"`
	Content string `json:"content"`
}

type SimpleSession struct {
	Username string  `json:"username"`
	TTL int64 `json:"ttl"`
}

//type SimpleSession struct {
//    Username string  `json:"username"`
//    TTL int64 `json:"ttl"`
//}

