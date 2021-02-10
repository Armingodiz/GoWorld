package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	// adding table to db
	//_, err = conn.Exec(context.Background(), "CREATE TABLE product (\nid SERIAL PRIMARY KEY,\nname VARCHAR NOT NULL,\nprice NUMERIC(10,2) DEFAULT 0\n);")
	if err != nil {
		fmt.Fprintf(os.Stderr, "table failed: %v\n", err)
		os.Exit(1)
	}
	// inserting new row to table :
	_, err = conn.Exec(context.Background(), "INSERT INTO product(name, price)\nVALUES ('armin', 12);")
	if err != nil {
		fmt.Fprintf(os.Stderr, "insert failed: %v\n", err)
		os.Exit(1)
	}
	// selecting row with condition :
	var num int
	err = conn.QueryRow(context.Background(), "SELECT price  FROM product\nWHERE name='armin';").Scan(&num)
	if err != nil {
		fmt.Fprintf(os.Stderr, "selest failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(num)
}
