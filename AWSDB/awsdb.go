package AWSDB

import (
	"context"
	"encoding/json"
	"io"
	"os"

	//"database/sql"
	"fmt"
	//"time"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
)



var AWSDB = GetAWSDB()


//var (
//	USER     = os.Getenv("USER")
//	PASSWORD = os.Getenv("PASSWORD")
//	DBNAME     = os.Getenv("DBNAME")
//	HOST     = os.Getenv("HOST")
//	DBPORT     = os.Getenv("DBPORT")
//)


var (
	DBPORT string
	USER string
	PASSWORD string
	DBNAME string
	HOST string
	SECRET_KEY string
)


func AuthDetails(){

	jsonFile, _ := os.Open("db_info.json")
	byteValues,_ := io.ReadAll(jsonFile)

	var data map[string]string
	json.Unmarshal( byteValues, &data)
	USER = data["USER"]
	PASSWORD = data["PASSWORD"]
	HOST = data["HOST"]
	DBPORT = data["DBPORT"]
	DBNAME = data["DBNAME"]
	SECRET_KEY = data["SECRET_KEY"]

}


func GetAWSDB() *pgxpool.Pool{
	AuthDetails()

	qStr := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v",USER,PASSWORD,HOST,DBPORT,DBNAME)

	db, err := pgxpool.Connect(context.Background(), qStr)
	//db, err := sql.Open("pgx",qStr)

	if err!=nil{
		panic(err)
	}

	//// Maximum Idle Connections
	//(*db).SetMaxIdleConns(5)
	//// Maximum Open Connections
	//db.SetMaxOpenConns(20)
	//// Idle Connection Timeout
	//db.SetConnMaxIdleTime(1 * time.Second)
	//// Connection Lifetime
	//db.SetConnMaxLifetime(30 * time.Second)

	return db
}