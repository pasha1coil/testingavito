package segment

import "time"

type History struct {
	User_id int       `json:"id"`
	Name    string    `json:"Name"`
	Mode    string    `json:"Mode"`
	Time    time.Time `json:"time"`
}

type GetHistory struct {
	User_id int    `json:"id"`
	Start   string `json:"start"`
	End     string `json:"end"`
}
