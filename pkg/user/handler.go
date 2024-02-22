package user

import (
	"fmt"

	"gorm.io/gorm"
)

func ConnectAndMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return fmt.Errorf("failed to migrate users database schema: %w", err)
	}
	return nil
}
