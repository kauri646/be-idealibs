package migration

import (
	"fmt"
	"log"

	"github.com/kauri646/be-idealibs/config"
	"github.com/kauri646/be-idealibs/internal/models/users"
)

func RunMigration() {
	err := config.DB.AutoMigrate(&users.User{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}
