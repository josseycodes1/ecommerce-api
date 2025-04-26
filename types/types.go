package types

type RegisterUserPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lirstName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
