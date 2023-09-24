package requestdto

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type VerifyUser struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

type GetNewVerifyUserOTP struct {
	Email string `json:"email"`
}

type ResetPassword struct {
	Password string `json:"password"`
}
