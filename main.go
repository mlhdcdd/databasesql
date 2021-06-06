package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"log"
)


func main()  {

	var uid = 1
	user, err := QueryUser(uid)
	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("user not found, %v\n",err)
		fmt.Printf("%+v\n", err)
		return
	}
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(user)
}
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func QueryUser(uid int) (*User, error) {
	var err error
	var user = new(User)
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/userdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return user,err
	}
	err = db.QueryRow("select id, username from W_UserBaseInfo where id = ?", uid).Scan(&user.Id, &user.Username)
	switch {
	case err == sql.ErrNoRows :
		return user, errors.Wrap(err,"sql.ErrNoRows")
	case err != nil :
		log.Fatal(err)
	default:
	}
	return user, err
}