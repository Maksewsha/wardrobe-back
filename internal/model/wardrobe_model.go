package model

import "time"

type WardrobeDTO struct {
	Id          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Poster      []byte `json:"poster"`
}

type WardrobeDAO struct {
	Id          uint64    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Poster      []byte    `json:"poster"`
	CreatedTime time.Time `json:"created_time"`
}
