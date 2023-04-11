/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Tue Apr 11 2023 11:20:04 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// data source
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := "postgres"
	dbName := "testdb"

	// postgres://postgres:postgres@localhost:5432/testdb?sslmode=disable
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sqlx.Connect("postgres", dataSourceName)

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected")
	}

	query := `INSERT INTO customers (customer_id, first_name, last_name) VALUES (:id, :first_name, :last_name)`

	_, err = db.NamedExec(query, map[string]interface{}{
		"id": 102,
		"first_name": "Ismail",
		"last_name": "Ash Shidiq",
	})

	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Insert Successful")
	}
}