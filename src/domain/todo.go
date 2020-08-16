package domain

type Todo struct {
	Id int16 `json:"int"`
	UserId int16 `json:"int"`
	Text string `json:"text"`
}