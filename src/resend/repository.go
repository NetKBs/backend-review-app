package resend

import (
	"errors"
	"time"

	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func sendVerificationEmailRepository(userId uint, code string) error {
	db := config.DB
	model := schema.ValidationCode{
		UserId: userId,
		Code:   code,
	}

	if err := db.Where("id = ?", userId).First(&schema.User{}).Error; err != nil {
		return errors.New("user not found")
	}

	if err := db.Create(&model).Error; err != nil {
		return err
	}

	return nil
}

func verifyVerificationCodeRepository(userId uint, code string) error {
	bd := config.DB
	var lastCode schema.ValidationCode

	if err := bd.Where("id = ?", userId).First(&schema.User{}).Error; err != nil {
		return errors.New("user not found")
	}

	if err := bd.Order("created_at DESC").Where("user_id = ?", userId).First(&lastCode).Error; err != nil {
		return err
	}

	if time.Since(lastCode.CreatedAt).Minutes() > 5 {
		return errors.New("the verification code has expired")
	}

	if lastCode.Code != code {
		return errors.New("the verification code is incorrect")
	}

	return nil
}
