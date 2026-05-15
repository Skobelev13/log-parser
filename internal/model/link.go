package model

type Link struct {
	Switch string `json:"switch"`
	Port   string `json:"port"`
	Peer   string `json:"peer"`
}