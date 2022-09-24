package Urls

import (
	"StudentMerit/HerokuDB"
	"StudentMerit/Structures"
	"StudentMerit/auth"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func recordManager(w http.ResponseWriter, r *http.Request){
	EnableCORSALL(&w)
	user, err := auth.Authenticate(r)
	if err != nil {
		json.NewEncoder(w).Encode(
			map[string]string{
				"message": err.Error(),
			})
		return
	}

	if r.Method=="POST"{

	var record Structures.Record
	json.NewDecoder(r.Body).Decode(&record)
	record.UserAdded = user

	err = record.Validate(r)
	if err!=nil{
		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"message":err.Error(),
				"status":400,
			})
		return
	}

	qStr := fmt.Sprintf(`select addRecord(%v,%v,%v)`,
		record.StudentId, record.ScoreId,record.UserAdded,
		)

	fmt.Println(qStr)

	_, err = HerokuDB.HEROKU_DB.Exec(context.Background(), qStr)
	if err!=nil{
		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"message":err.Error(),
				"status":400,
			})
		return
	}

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"message":"New record added",
			"status":200,
		})

}

	if r.Method=="GET"{
		qStr := fmt.Sprintf(`SELECT * FROM getMyRecords(%v)`, user)

		table, err:= HerokuDB.HEROKU_DB.Query(context.Background(), qStr)

		if err!=nil{
			json.NewEncoder(w).Encode(
				map[string]string{
					"message": err.Error(),
				})
			return
		}

		var myRecords []Structures.MyRecords
		for table.Next(){
			var rc Structures.MyRecords
			table.Scan(&rc.StdName,&rc.StdSurname,&rc.StdClass,&rc.Action,&rc.Points,&rc.DateScoreAdded)
			rc.Formatted = rc.DateScoreAdded.Format(" 02-Jan-2006, 15:04")
			myRecords=append(myRecords,rc)

		}

		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"data":myRecords,
				"ep":APIEP,
			})

	}

}
