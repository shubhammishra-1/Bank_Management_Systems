package setup

import (

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"log"
)



func Open_DB()(db *sql.DB, err error){

    log.Println("MySQL connection Established")

    // Open up our database connection.

    // I've set up a database on my local machine using Go_DB.

    db, err = sql.Open("mysql", "root:55120@tcp(127.0.0.1:3306)/Bank_Systems")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

   return db,err

}


func Close_DB(db *sql.DB){

    // defer the close till after the main function has finished
    log.Println("Database has been Closed!")
    db.Close()

}
