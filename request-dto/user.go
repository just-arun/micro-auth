package requestdto

type AddUser struct {
	Email    string `json:"email"`
	UserName string `json:"userName"`
}

type UserList struct {
	Data []AddUser `json:"data"`
}
