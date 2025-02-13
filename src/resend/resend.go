package resend

type verifyinputDTO struct {
	UserId uint   `json:"user_id" binding:"required"`
	Code   string `json:"code" binding:"required"`
}

type generateInputDTO struct {
	UserId uint   `json:"user_id" binding:"required"`
	Email  string `json:"email" binding:"required"`
}
