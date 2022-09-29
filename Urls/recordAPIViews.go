package Urls

import (
	"StudentMerit/HerokuDB"
	"StudentMerit/Structures"
	"StudentMerit/auth"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func recordManager(w http.ResponseWriter, r *http.Request){
	EnableCORSALL(&w, r)
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
		defer table.Close()

		var myRecords []Structures.MyRecords
		for table.Next(){
			var rc Structures.MyRecords
			table.Scan(
				&rc.Id,
				&rc.StdId,
				&rc.StdName,
				&rc.StdSurname,
				&rc.StdClass,
				&rc.Action,
				&rc.Points,
				&rc.DateScoreAdded)
			rc.Formatted = rc.DateScoreAdded.Format(" 02-Jan-2006, 15:04")
			myRecords=append(myRecords,rc)

		}

		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"data":myRecords,
				"ep":APIEP,
			})

	}

	if r.Method == "DELETE"{

		record_id:=mux.Vars(r)["id"]

		qStr := fmt.Sprintf(`SELECT deleteRecord(%v,%v)`, user,record_id)

		row := HerokuDB.HEROKU_DB.QueryRow(context.Background(), qStr)

		var result byte
		err:=row.Scan(&result)

		var message string
		var status int

		if result == 10{
			message="OK"
			status=200
		}else{
			message=err.Error()
			status=400
		}

		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"message":message,
				"status":status,
			})

	}
}


func getRecordsForStudent(w http.ResponseWriter, r*http.Request){
	EnableCORSALL(&w, r)
	_, err := auth.Authenticate(r)
	if err != nil {
		json.NewEncoder(w).Encode(
			map[string]string{
				"message": err.Error(),
			})
		return
	}

	if r.Method == "GET"{
		id := mux.Vars(r)["id"]

		qStr := fmt.Sprintf(`SELECT * FROM getrecordsforstudent(%v) `,id)

		rows, err := HerokuDB.HEROKU_DB.Query(context.Background(),qStr)
		if err!=nil{
			json.NewEncoder(w).Encode(
				map[string]string{
					"message":err.Error(),
				})
		}
		defer rows.Close()

		var Records []Structures.RecordsForStudent

		for rows.Next(){
			var sr Structures.RecordsForStudent
			rows.Scan(&sr.FromUser, &sr.Action, &sr.Points, &sr.Date)
			sr.DateFormatted = sr.Date.Format("02-Jan-2006")
			sr.TimeFormatted = sr.Date.Format("15:04")

			Records=append(Records,sr)
		}

		json.NewEncoder(w).Encode(
			map[string]interface{}{
			"message":"OK",
			"data":Records,
		})
		return
	}


}