package HerokuDB

import (
	"context"
	"os"

	//"database/sql"
	"fmt"
	//"time"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
)



var HEROKU_DB = GetHerokuDB()


var (
	USER     = os.Getenv("USER")
	PASSWORD = os.Getenv("PASSWORD")
	DBNAME     = os.Getenv("DBNAME")
	HOST     = os.Getenv("HOST")
	DBPORT     = os.Getenv("DBPORT")
)


func GetHerokuDB() *pgxpool.Pool{


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