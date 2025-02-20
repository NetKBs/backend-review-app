package resend

type verifyinputDTO struct {
	Code string `json:"code" binding:"required"`
}

type generateInputDTO struct {
	Email string `json:"email" binding:"required"`
}
