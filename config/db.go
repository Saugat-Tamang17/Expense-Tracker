package config

import(
	"database/sql"
	"fmt"
	"log"
	"os"
	_"github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB(){
	var err error
	dbURL:= os.Getenv("DB_URL")

	DB,err:=sql.Open("postgres",dbURL)
	if err!=nil{
		log.Fatal("Error opening the database",err)
		return
	}

	err =DB.Ping()
	if err!=nil{
		log.Fatal("Cannot connect to Database: ",err)
	}
	
	fmt.Println("\nConnected To Database Sucessfully")
}