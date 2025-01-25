package resend

type verifyinputDTO struct {
	
	Email string `json:"email" binding:"required"`
	Code  string `json:"code,omitempty"`
}
