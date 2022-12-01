package response

import "time"

type UserRegisterResponse struct {
	ID        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserPatchTopUpResponse struct {
	Message string `json:"message"`
}
