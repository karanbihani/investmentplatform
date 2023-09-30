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

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Database connection is successful!")
	return &PostgresStore{
		db:db,
	},nil
}

func (s *PostgresStore) Init() error{
	return s.CreateAccountTable()
}

func (s *PostgresStore) CreateAccountTable() error{
	query := `create table if not exists account(
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		number serial,
		balance serial,
		created_at timestamp
	)`
	_,err := s.db.Exec(query)
	return err
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