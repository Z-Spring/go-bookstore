package model

type User struct {
	Id       int    `json:"id,omitempty"`
	Uid      string `json:"uid,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}
