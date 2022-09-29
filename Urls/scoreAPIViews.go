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



func scoresManager(w http.ResponseWriter, r*http.Request){
	EnableCORSALL(&w, r)
	user, err := auth.Authenticate(r)
	if err!=nil{


		json.NewEncoder(w).Encode(
			map[string]string{
				"message":err.Error(),
			})
		return
	}
	var qStr string
	if r.Method == "GET"{




		qStr = fmt.Sprintf(`SELECT id,action, points,user_added from score ORDER BY date_added ASC`)

		rows, err := HerokuDB.HEROKU_DB.Query(context.Background(), qStr)
		if err!=nil{
			json.NewEncoder(w).Encode(
				map[string]string{ "error":err.Error()})

		}

		var actions []Structures.Score
		for rows.Next(){
			var action Structures.Score
			rows.Scan(&action.Id,&action.Action, &action.Points,&action.UserAdded)
			actions = append(actions,action)
		}

		rp := map[string]interface{}{
			"data":actions,
			"ep":APIEP,
		}


		json.NewEncoder(w).Encode(rp)

		return
	}


	if r.Method == "POST"{
		var ns Structures.Score
		json.NewDecoder(r.Body).Decode(&ns)
		err = ns.Validate(r)


		if err != nil{
			json.NewEncoder(w).Encode(
				map[string]string{
					"message":err.Error(),
				})
			return
		}

		qStr = fmt.Sprintf(`SELECT addScoreType('%s',%v,%v)`, ns.Action, ns.Points, user)

		row := HerokuDB.HEROKU_DB.QueryRow(context.Background(), qStr)


		var result int
		row.Scan(&result)

		// map of statuses
		var message string
		var status int

		if result == 0{
			message = fmt.Sprintf("New Score type is added")
			status = 200
		} else if result==1{
			message = "This score is already added"
			status=400
		} else{
			message="You can't add more than 5 scores"
			status=400
		}

		json.NewEncoder(w).Encode(
			map[string] interface{}{
				"message":message,
				"status":status,
			})
	}

	if r.Method == "PUT"{
		var ns Structures.Score
		err=json.NewDecoder(r.Body).Decode(&ns)

		err = ns.Validate(r)
		if err != nil{
			json.NewEncoder(w).Encode(
				map[string]string{
					"message":err.Error(),
				})
			return
		}



		qStr = fmt.Sprintf(`UPDATE score SET action='%v', points=%v WHERE id=%v`, ns.NewAction, ns.NewPoints, ns.Id)

		_, err = HerokuDB.HEROKU_DB.Exec(context.Background(), qStr)

		if err!=nil{

			json.NewEncoder(w).Encode(map[string]interface{}{

				"message":"This score type is already added",
				"status":400,
			})
			return
		}

		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"message": fmt.Sprintf("Score is updated"),
				"status":200,
				"newName": ns.NewAction,
			})

		return

	}

	if r.Method == "DELETE"{

		var dlt Structures.Score
		err=json.NewDecoder(r.Body).Decode(&dlt)

		qStr = fmt.Sprintf(`SELECT deleteScoreType('%v','%v')`,dlt.Id,user)

		row := HerokuDB.HEROKU_DB.QueryRow(context.Background(), qStr)

		var result int32
		err = row.Scan(&result)

		if err!=nil{

			json.NewEncoder(w).Encode(
				map[string] interface{}{
					"message":err.Error(),
					"status":400,
				})
		}

		var message string
		var status int

		if result == 10{
			message="OK"
			status=200
		} else if result==11 {
			message="Access Denied"
			status=400
		} else{
			message="No such score"
			status=400
		}



		json.NewEncoder(w).Encode(
			map[string] interface{}{
				"message":message,
				"status":status,
			})
	}
}
