package structs

type Users struct {
	ID         int64  `json:"id"`
	First_name string `json:"firts_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
}