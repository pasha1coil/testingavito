package segment

type User struct {
	User_number int `json:"user_number"`
}

type DataBase struct {
	Name string `json:"name_table"`
}

type DbOutput struct {
	ID      int    `json:"ID,omitempty"`
	User_id int    `json:"User_id,omitempty"`
	Name    string `json:"Name,omitempty"`
	Mode    string `json:"Mode,omitempty"`
	Created string `json:"created,omitempty"`
}
