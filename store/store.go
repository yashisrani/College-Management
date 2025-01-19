package store

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
  
type Postgress struct {
	DB *gorm.DB
}
 func (store *Postgress) NewStore() error  {
	dsn := "host=localhost user=yash1 password=password1 dbname=project1 port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("connection failed")
		return err
	}else{
		fmt.Println("Database connected successfully")
		store.DB = db
	}
	return nil
 }

//  store (store.go) -> controller(server.go) -> cmd(main.go)

type StoreOperation interface {
	NewStore() error
}