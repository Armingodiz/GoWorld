package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql",
		"root:armin3011@tcp(127.0.0.1:3306)/loginSystem")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	rows , err2 :=db.Query("SELECT * FROM loginSystem.users")
	if err2 != nil {
		log.Fatalln(err)
	}
	var x int
	var a , b,c string
	for rows.Next() {
		fmt.Println("###############################################################")
		err := rows.Scan(&x,&b,&a, &c)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(a,b,c)
	}

	stmt, err := db.Prepare(`INSERT INTO users VALUES (?,?,?,?);`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec(4,"fdsfd","fdsf","rew")
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Println("INSERTED RECORD", n)
	defer db.Close()
}

func check(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}