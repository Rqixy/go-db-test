package main

import (
	"fmt"
	"log"
	"testDB/models"
)

func main() {
	// データベースに接続する
	db, err := models.DBConnect()
	if err != nil {
		log.Fatal(err)
	}
	// 最後に接続を閉じる
	defer db.Close()

	// Serverの構造体を初期化する
	var s = models.Server{DB: db}

	// Select
	users, err := s.SelectUsers()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)
	//
	name := ""
	users, err = s.SelectUsersByName(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)
	//
	user, err := s.SelectUserByID(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)

	// Insert
	name = ""
	age := 1
	err = s.InsertUser(name, age)
	if err != nil {
		log.Fatal(err)
	}
	//
	users = []models.User{
		{Name: "", Age: 1},
		{Name: "", Age: 2},
	}
	err = s.InsertUsers(users)
	if err != nil {
		log.Fatal(err)
	}

	// Update
	name = ""
	err = s.UpdateUser(1, name)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	err = s.DeleteUser(1)
	if err != nil {
		log.Fatal(err)
	}
}