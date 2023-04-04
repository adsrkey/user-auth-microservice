package user

type User struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Hash     []byte `json:"-"`
}
