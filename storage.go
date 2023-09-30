package main

import (
	"database/sql"
	"fmt"
	"log"

	//"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)
// func goDotEnvVariable(key string) string {

// 	// load .env file
// 	err := godotenv.Load(".env")
  
// 	if err != nil {
// 	  log.Fatalf("Error loading .env file")
// 	}
  
// 	return os.Getenv(key)
// }

type Storage interface{
	CreateAccount(*Account) error
	UpdateAccount(*Account) error
	DeleteAccount(int) error
	GetAccountByID(int)(*Account, error)
}

type PostgresStore struct{
	db *sql.DB
}

func NewPostgresStore()(*PostgresStore, error){
	connStr := "user=postgres dbname=postgres password=bihani123 sslmode=disable"
    
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Database connection is successful!")
	return &PostgresStore{
		db:db,
	},nil
}

func (s *PostgresStore) CreateAccount(*Account) error{
	return nil
}

func (s *PostgresStore) UpdateAccount(*Account) error{
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error{
	return nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error){
	return nil,nil
}