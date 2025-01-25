package resend

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/resend/resend-go/v2"
)

type verificationCode struct {
	code      string
	expiresAt time.Time
}

var (
	codeStore = make(map[string]verificationCode)
	mutex     sync.Mutex
)

func generateRandomCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	code := ""
	for i := 0; i < length; i++ {
		num := rand.Intn(10)
		code += fmt.Sprintf("%d", num)
	}
	return code
}

func sendVerificationEmailService(to string) error {
	verificationCode := generateRandomCode(6)

	storeVerificationCode(to, verificationCode, 1*time.Minute)

	apikey := os.Getenv("RESEND_APIKEY")
	fromEmail := os.Getenv("RESEND_DOMAIN")

	ctx := context.TODO()
	client := resend.NewClient(apikey)

	htmlBody := fmt.Sprintf("<p>Tu código de verificación es:</p><h2>%s</h2>", verificationCode)

	params := &resend.SendEmailRequest{
		From:    fromEmail,
		To:      []string{to},
		Subject: "Verifica tu correo electrónico",
		Html:    htmlBody,
	}

	sent, err := client.Emails.SendWithContext(ctx, params)
	if err != nil {
		return fmt.Errorf("error al enviar el correo: %v", err)
	}
	fmt.Println("Correo enviado con ID:", sent.Id)
	return nil
}

func storeVerificationCode(email, code string, duration time.Duration) {
	mutex.Lock()
	defer mutex.Unlock()

	expirationTime := time.Now().Add(duration)
	codeStore[email] = verificationCode{
		code:      code,
		expiresAt: expirationTime,
	}

	go func() {
		time.Sleep(duration)
		mutex.Lock()
		defer mutex.Unlock()
		delete(codeStore, email)
	}()
}

func validateverificationCode(email, code string) bool {
	mutex.Lock()
	defer mutex.Unlock()

	if storedCode, exists := codeStore[email]; exists {
		if time.Now().Before(storedCode.expiresAt) && storedCode.code == code {
			return true
		}
	}
	return false
}
