package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // driver for postgresSQL
	"github.com/subosito/gotenv"
)

func fatal(err error) {
	log.Fatal(err)
}

func init() {
	gotenv.Load()
}

//ConnectDB initialises the postgres database
func ConnectDB() *sql.DB {
	config := initConfig()

	plsqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s"+" dbname=%s",
		config["host"], config["port"], config["user"], config["pass"],
		config["dbname"])

	db, err := sql.Open("postgres", plsqlInfo)

	if err != nil {
		fatal(err)
		//panic(err)
	}

	err2 := db.Ping()

	if err2 != nil {
		fatal(err)
	}

	return db
}

func initConfig() map[string]string {

	config := make(map[string]string)

	config["host"] = os.Getenv("DBHOST")
	config["port"] = os.Getenv("DBPORT")
	config["user"] = os.Getenv("DBUSER")
	config["pass"] = os.Getenv("DBPASS")
	config["dbname"] = os.Getenv("DBNAME")
	return config
}
