package users

import "github.com/wgarcia4190/bookstore_users_api/internal/utils/crypto"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (request LoginRequest) GetEncryptedPassword() string {
	pass, _ := crypto.GetMd5(request.Password)
	return pass
}
