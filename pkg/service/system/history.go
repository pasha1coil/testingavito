package segment

import "time"

type History struct {
	User_id int       `json:"id"`
	Name    string    `json:"Name"`
	Mode    string    `json:"Mode"`
	Time    time.Time `json:"time"`
}

type GetHistory struct {
	User_id int    `json:"id" binding:"required"`
	Start   string `json:"start" binding:"required"`
	End     string `json:"end" binding:"required"`
}
