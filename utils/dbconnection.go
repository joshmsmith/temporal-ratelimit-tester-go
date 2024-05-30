package utils

import (
  "fmt"
  "log"
  "os"

  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func GetDBConnection() (db *sql.DB, err error) {

  // Database connection env
  db_host     := os.Getenv("MYSQL_HOST")
  db_database := os.Getenv("MYSQL_DATABASE")
  db_user     := os.Getenv("MYSQL_USER")
  db_password := os.Getenv("MYSQL_PASSWORD")
  connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?multiStatements=true", db_user, db_password, db_host, db_database)

  // Connect to database
  db, err = sql.Open("mysql", connectionString)
  if err != nil {
    log.Println("Failed to connect to MySQL database, check ENVIRONMENT variables!")
    log.Fatal(err)
  }

  return
}
