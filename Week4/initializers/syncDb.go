package initializers

import "jwt-authentication-golang/models"

func SyncDb() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Post{})
	DB.AutoMigrate(&models.Comment{})
	DB.AutoMigrate(&models.Like{})
}
