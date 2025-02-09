package resend

import (
	"context"
	"fmt"
	"math/rand"
	"os"

	"github.com/NetKBs/backend-reviewapp/src/social/user"
	"github.com/resend/resend-go/v2"
)

func generateRandomCode(length int) string {
	const letterBytes = "0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func sendVerificationEmailService(userId uint, to string) error {
	apikey := os.Getenv("RESEND_APIKEY")
	fromEmail := os.Getenv("RESEND_DOMAIN")

	if apikey == "" || fromEmail == "" {
		return fmt.Errorf("RESEND_APIKEY or RESEND_DOMAIN environment variable not set")
	}

	verificationCode := generateRandomCode(6)
	if err := sendVerificationEmailRepository(userId, verificationCode); err != nil {
		return err
	}

	ctx := context.TODO()
	client := resend.NewClient(apikey)
	htmlBody := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
			<head>
				<style>
					p, h1, h2 {
						font-family: Arial, sans-serif;
					}
					.container {
						max-width: 600px;
						margin: 40px auto;
						padding: 20px;
						border-radius: 10px;
						box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
						background-color: #1e1e1e;
					}
					.header {
						background-color: hsl(151, 55.0%%, 41.5%%);
						padding: 20px;
						border-radius: 10px 10px 0 0;
						text-align: center;
						color: white;
					}
					.code {
						font-size: 32px;
						font-weight: bold;
						color: hsl(151, 55.0%%, 41.5%%);
						margin-bottom: 20px;
						text-align: center;
					}
					p {
						font-size: 16px;
						line-height: 1.5;
						color: #e0e0e0;
					}
					.footer {
						margin-top: 20px;
						font-size: 14px;
						color: #bbbbbb;
						text-align: center;
					}
				</style>
			</head>
			<body>
				<div class="container">
					<div class="header">
						<h1>Verificación de cuenta para Leif</h1>
					</div>
					<div class="main">
						<p><b>Estimado usuario de Leif,</b></p>
						<p>Gracias por crear una cuenta con nosotros. Para completar el proceso de registro, necesitamos verificar tu dirección de correo electrónico.</p>
						<p><b>Tu código de verificación es:</b></p>
						<h2 class="code">%s</h2>
						<p>Por favor, ingresa este código en la app para activar tu cuenta.</p>
						<p>Si no creaste esta cuenta, ignora este correo y asegúrate de que tu dirección de correo electrónico no haya sido comprometida.</p>
						<p class="footer">Atentamente,</p>
						<p class="footer">El equipo de Leif</p>
					</div>
				</div>
			</body>
		</html>

	`, verificationCode)

	params := &resend.SendEmailRequest{
		From:    fromEmail,
		To:      []string{to},
		Subject: "Código de verificación de Leif",
		Html:    htmlBody,
	}

	sent, err := client.Emails.SendWithContext(ctx, params)
	if err != nil {
		return fmt.Errorf("error sending verification email: %v", err)
	}
	fmt.Println("Correo enviado con ID:", sent.Id)
	return nil
}

func verifyVerificationCodeService(userId uint, code string) error {
	if err := verifyVerificationCodeRepository(userId, code); err != nil {
		return err
	}

	if err := user.VerifyUserService(userId); err != nil {
		return err
	}

	return nil
}
