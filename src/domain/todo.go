package domain

type Todo struct {
	Id int `json:"int"`
	UserId int16 `json:"int"`
	Text string `json:"text"`
}
