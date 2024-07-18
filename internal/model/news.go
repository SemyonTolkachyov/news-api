package model

type NewsWithCategories struct {
	Id         int    `json:"Id" binding:"required"`
	Title      string `json:"Title"`
	Content    string `json:"Content"`
	Categories []int  `json:"Categories"`
}
