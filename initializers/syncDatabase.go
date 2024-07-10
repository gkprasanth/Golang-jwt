package initializers

import "gin-backend/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}