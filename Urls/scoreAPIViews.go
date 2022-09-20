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



func scoreManager(w http.ResponseWriter, r*http.Request){
	EnableCORSALL(&w)
	_, err := auth.Authenticate(r)
	if err!=nil{


		json.NewEncoder(w).Encode(
			map[string]string{
				"message":err.Error(),
			})
		return
	}

	if r.Method == "GET"{
		qStr := fmt.Sprintf(`SELECT id,action, points from score`)

		rows, err := HerokuDB.HEROKU_DB.Query(context.Background(), qStr)
		if err!=nil{
			json.NewEncoder(w).Encode(
				map[string]string{ "error":err.Error()})

		}

		var actions []Structures.Score
		for rows.Next(){
			var action Structures.Score
			rows.Scan(&action.Id,&action.Action, &action.Points)
			actions = append(actions,action)
		}

		rp := map[string]interface{}{
			"data":actions,
			"ep":APIEP,
		}

		fmt.Println(actions)
		json.NewEncoder(w).Encode(rp)

		return
	}


}
