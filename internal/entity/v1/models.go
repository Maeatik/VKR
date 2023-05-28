package entity

import "time"

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name"  db:"login"`
	Password string `json:"password" db:"password"`
}

type UserChange struct {
	Id          int    `json:"-" db:"id"`
	Password    string `json:"password" db:"password"`
	NewPassword string `json:"new_password" db:"password"`
}

type Site struct {
	Id     int    `json:"id" db:"id"`
	UserID int    `json:"-"  db:"user_id(FK)"`
	Url    string `json:"url"  db:"url"`
	Tag    string `json:"tag" db:"tag"`
}

type Sites struct {
	IdSite int       `json:"id_site" db:"id"`
	UserID int       `json:"-"  db:"user_id(FK)"`
	Url    string    `json:"url"  db:"url"`
	Tag    string    `json:"tag" db:"tag"`
	Date   time.Time `json:"date"  db:"date"`
	IdText string    `json:"id_text" db:"id"`
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

type DownloadData struct {
	Id int
}
