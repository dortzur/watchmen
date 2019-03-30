package model

type UserData struct {
	User     string `json:"user" binding:"required" example:"1234"`
	Company  string `json:"company" binding:"required" example:"5687"`
	Password string `json:"password" binding:"required" example:"Baba"`
}
