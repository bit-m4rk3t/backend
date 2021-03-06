package entities

type User struct {
	ID          int      `json:"id"`
	Email       string   `json:"email"`
	Password    string   `json:"password"`
	Name        string   `json:"name"`
	Surname     string   `json:"surname"`
	PhoneNumber string   `json:"phone_number"`
	Address     string   `json:"address"`
	Role        UserRole `json:"role"`
}

type UserRole int

const (
	NoAdmin UserRole = iota + 1
	Admin
)
