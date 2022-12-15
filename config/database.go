package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gobuffalo/packr/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

var (
	DbConnection *sql.DB
)

func Database() *sql.DB {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("failed load file environment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	DbConnection, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = DbConnection.Ping()
	if err != nil {
		panic(err)
	}

	return DbConnection
}

func DBMigrate(dbParam *sql.DB, status migrate.MigrationDirection) {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "../migrations"),
	}

	n, errs := migrate.Exec(dbParam, "postgres", migrations, status)
	if errs != nil {
		panic(errs.Error())
	}

	DbConnection = dbParam

	fmt.Println("Applied", n, " migrations!")
}
