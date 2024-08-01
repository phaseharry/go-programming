package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// importing this way ("_") means we only want the side effects of the package and don't need any of its api's

func main() {
	db, err := sql.Open("postgres", "user=postgres password=Start!123 host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connection to the DB was successfully initialized!")
	}

	TableCreate := `
	CREATE TABLE Number
	(
		Number integer NOT NULL,
		Property text COLLATE pg_catalog."default" NOT NULL
	)
	WITH (
		OIDS = FALSE
	)
	TABLESPACE pg_default;
	ALTER TABLE Number
		OWNER to postgres;
	`
	// executing the above query. We don't care about the response. We only care if there's an error or not
	_, err = db.Exec(TableCreate)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The table called Numbers was successfully created!")
	}

	// inserting nums into the above created table
	insert, insertErr := db.Prepare("INSERT INTO Number VALUES($1, $2)")
	if insertErr != nil {
		panic(insertErr)
	}

	for i := 0; i < 100; i++ {
		var prop string
		if i%2 == 0 {
			prop = "event"
		} else {
			prop = "odd"
		}
		_, err = insert.Exec(i, prop)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("The number:", i, "is:", prop)
		}
	}
	insert.Close()
	fmt.Println("The numbers are ready.")

	db.Close()
}
