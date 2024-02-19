package config

import (
	"fmt"
	"path/filepath"
	"runtime"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	_, b, _, _ = runtime.Caller(0)

	ProjectRootPath = filepath.Join(filepath.Dir(b), "../")
)

const (
	host     = "aws-0-ap-southeast-1.pooler.supabase.com"
	port     = 5432
	user     = "postgres.admkecqjbucsholomvpm"
	password = "#Kaka060406!"
	dbname   = "postgres"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database")
	}

	fmt.Println("Connected to database")
}
