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

	// fetching all data
	// rows, err := db.Query("SELECT customer_id, first_name, last_name FROM customers WHERE customer_id < 50")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer rows.Close()

	// for rows.Next() {
	// 	var customerId, firstName, lastName string
	// 	if err := rows.Scan(&customerId, &firstName, &lastName); err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Printf("customer_id: %s\n first_name: %s\n last_name: %s\n", customerId, firstName, lastName)
	// }

	// fetching one data
	row := db.QueryRow("SELECT customer_id, first_name, last_name FROM customers WHERE customer_id = $1", 1)

	var customerId, firstName, lastName string
	err = row.Scan(&customerId, &firstName, &lastName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("customer_id: %s\n first_name: %s\n last_name: %s\n", customerId, firstName, lastName)
}

// nama repo: app-mahasiswa-db