package main

import (
	"database/sql"
	"log"

	"github.com/code-farms/go-backend/cmd/api"
	"github.com/code-farms/go-backend/configs"
	"github.com/code-farms/go-backend/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User: configs.Envs.DBUser,
		Passwd: configs.Envs.DBPassword,
		Addr: configs.Envs.DBAddress,
		DBName: configs.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	intiStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func intiStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	} 
	log.Println("Connected to MySQL database")
}