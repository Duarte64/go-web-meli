package users

type User struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Lastname  string  `json:"lastname"`
	Email     string  `json:"email"`
	Age       int     `json:"age"`
	Height    float64 `json:"height"`
	Active    bool    `json:"active"`
	CreatedAt string  `json:"created_at"`
}
