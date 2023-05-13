package entity

import "time"

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name"  db:"login"`
	Password string `json:"password" db:"password"`
}

type Site struct {
	Id     int    `json:"-" db:"id"`
	UserID int    `json:"-"  db:"user_id(FK)"`
	Url    string `json:"url"  db:"url"`
	Tag    string `json:"tag" db:"tag"`
}

type SiteID struct {
	Id int `json:"id" db:"id"`
}

type MainText struct {
	Id     int       `json:"id" db:"id"`
	UserID int       `json:"-"  db:"user_id(FK)"`
	SiteID int       `json:"site_id"  db:"site_id(FK)"`
	Date   time.Time `json:"date"  db:"date"`
	Text   string    `json:"text" db:"text"`
}

type TextID struct {
	Id     int `json:"id" db:"id"`
	SiteId int `json:"site_id" db:"site_id(FK)"`
}
