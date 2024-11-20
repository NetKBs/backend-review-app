package user

import (
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func GetUserByIdRepository(id uint) (user schema.User, err error) {
	db := config.DB

	if err = db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func UpdateUserRepository(user *schema.User) error {
	db := config.DB
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Save(user).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return err
}
func CreateUserRepository(user *schema.User) error {
	db := config.DB
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Create(user).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return err
}

// Preguntar si hay que crear la funcion de recuperar cuenta
func DeleteUserbyIDRepository(id uint) error {
	db := config.DB

	// Iniciar transacción
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Buscar usuario por ID
	var user schema.User
	if err := tx.Where("id = ?", id).First(&user).Error; err != nil {
		tx.Rollback()
		return err // Retornar error si no se encuentra el usuario
	}

	// Actualizar campo DeletedAt con soft delete
	if err := tx.Delete(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Confirmar transacción
	return tx.Commit().Error
}
