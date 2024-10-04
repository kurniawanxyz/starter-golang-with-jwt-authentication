package requests

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Name     string `json:"name" validate:"required"`
	Telp     string `json:"telp" validate:"required,e164"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type VerifyUserRequest struct {
	Token  string `json:"token" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type CreateTokenRequest struct {
	Email string `json:"email" validate:"required,email"`
	Type string `json:"type" validate:"required,oneof=email_verification forgot_password"`
}

type ResetPasswordRequest struct {
	Token string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
	KonfirmasiPassword string `json:"konfirmasi_password" validate:"required,min=8,eqfield=Password"`
}