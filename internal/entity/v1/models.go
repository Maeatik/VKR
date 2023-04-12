package entity

type User struct{
	Id			int		`json:"-" db:"id"`
	Name 		string 	`json:"name"  db:"login"`
	Password 	string 	`json:"password" db:"password"`
}
type User2 struct{
	Id			int		`json:"-" db:"id"`
	Name 		string 	`json:"name" binding:"required"`
	Password 	string 	`json:"password" binding:"required"`
}