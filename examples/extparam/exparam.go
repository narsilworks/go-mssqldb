package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

var (
	debug         = flag.Bool("debug", false, "enable debugging")
	password      = flag.String("password", "password101", "the database password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", "localhost", "the database server")
	user          = flag.String("user", "sa", "the database user")
	database      = flag.String("database", "APPSDB", "the database")
)

func main() {
	flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", *server, *user, *password, *port, *database)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	// var val string
	// err = conn.QueryRow(`SELECT ?`, mssql.NewExtParam(mssql.VarChar("TestValue"), 50)).Scan(&val)
	// if err != nil {
	// 	log.Fatal("Error:", err.Error())
	// }
	// fmt.Printf("%s\n", val)

	//rows, err := conn.Query(`SELECT FirstName, LastName FROM tcoEmployee WHERE Whse = ?`, mssql.NewExtParam(mssql.VarChar("QUEZON CITY"), 50))
	rows, err := conn.Query(`SELECT FirstName, LastName FROM tcoEmployee WHERE Whse = ?`, "QUEZON CITY")
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	defer rows.Close()

	var lname, fname string
	for rows.Next() {
		err = rows.Scan(&fname, &lname)
		if err != nil {
			log.Fatal("Error:", err.Error())
		}
		fmt.Printf("Name: %s %s\n", lname, fname)
	}

}
