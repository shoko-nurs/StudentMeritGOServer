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



func DeleteClassAPIView(w http.ResponseWriter, r *http.Request){
	EnableCORSALL(&w)
	user, err := auth.Authenticate(r)
	if err!=nil{


		json.NewEncoder(w).Encode(
			map[string]string{
				"message":err.Error(),
			})
		return
	}

	Id := mux.Vars(r)["id"]

	qStr := fmt.Sprintf(`DELETE FROM class WHERE id=%v`,Id)

	rows, err := HerokuDB.HEROKU_DB.Query(context.Background(),qStr)
	defer rows.Close()

	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{
			"message":"No such class",
		})
		return
	}

	// MOdify added_classes number in DB
	var classesAdded int64

	qStr = fmt.Sprintf(`SELECT added_classes FROM user_class_number_added WHERE user_added=%v`, user)
	row := HerokuDB.HEROKU_DB.QueryRow(context.Background(), qStr)

	row.Scan(&classesAdded)
	classesAdded -= 1
	qStr = fmt.Sprintf(`UPDATE user_class_number_added SET added_classes=%v WHERE user_added=%v`,classesAdded, user)
	row = HerokuDB.HEROKU_DB.QueryRow(context.Background(), qStr)


	json.NewEncoder(w).Encode(map[string]string{
		"message":"OK",
	})
}


func classMainPageAPIView(w http.ResponseWriter, r *http.Request){
	// CORS Headers assignment must be executed before Method checking


	EnableCORSALL(&w)

	if r.Method == "GET"{

		_, err := auth.Authenticate(r)

		if err!=nil{


			json.NewEncoder(w).Encode(
				map[string]string{
					"message":err.Error(),
				})
			return
		}

		qStr := "SELECT id, class, user_added FROM class ORDER BY class"

		rows, err := HerokuDB.HEROKU_DB.Query(context.Background(), qStr)
		if err!=nil{

			json.NewEncoder(w).Encode(err)
			return
		}
		defer rows.Close()

		var Classes []Structures.Class

		for rows.Next(){
			var c Structures.Class
			rows.Scan(&c.Id, &c.Class, &c.UserAdded)
			Classes = append(Classes, c)
		}


		rp := map[string]interface{}{
			"data":Classes,
			"ep":APIEP,
		}

		json.NewEncoder(w).Encode(rp)

		return
	}

	if r.Method == "POST"{



		user, err := auth.Authenticate(r)
		if err!=nil{


			json.NewEncoder(w).Encode(
				map[string]string{
					"message":err.Error(),
				})
			return
		}

		var newClass Structures.Class
		json.NewDecoder(r.Body).Decode(&newClass)
		err = newClass.Validate(r)

		if err != nil{
			json.NewEncoder(w).Encode(
				map[string]string{
				"message":err.Error(),
				})
			return
		}

		qStr := fmt.Sprintf(`SELECT addClass('%s',%v)`, newClass.Class,user)


		row := HerokuDB.HEROKU_DB.QueryRow(context.Background(), qStr)


		if err!= nil{
			json.NewEncoder(w).Encode(
				map[string]string{
					"error":err.Error(),
				})

		}
		var result int
		row.Scan(&result)


		// map of statuses
		var message string
		var status int

		if result == 0{
			message = fmt.Sprintf("Class %v is added", newClass.Class)
			status = 200
		} else if result==1{
			message = "This class is already added"
			status=400
		} else{
			message="You can't add more than 3 classes"
			status=400
		}


		json.NewEncoder(w).Encode(
			map[string] interface{}{
				"message":message,
				"status":status,
			})



	}


	if r.Method == "PUT"{

		_, err := auth.Authenticate(r)

		if err!=nil{


			json.NewEncoder(w).Encode(
				map[string]string{
					"message":err.Error(),
				})
			return
		}


		var editedClass Structures.Class
		json.NewDecoder(r.Body).Decode(&editedClass)

		err = editedClass.Validate(r)

		if err!=nil{

			json.NewEncoder(w).Encode(map[string]interface{}{

				"message":err.Error(),
				"status":400,
			})
			return
		}


		qStr := fmt.Sprintf(`UPDATE class SET class='%s' WHERE id='%v'`,editedClass.NewName,editedClass.Id)

		_, err = HerokuDB.HEROKU_DB.Exec(context.Background(), qStr)

		if err!=nil{

			json.NewEncoder(w).Encode(map[string]interface{}{

				"message":"This class is already added",
				"status":400,
			})
			return
		}

		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"message": fmt.Sprintf("Class is updated"),
				"status":200,
				"newName": editedClass.NewName,
			})

		return



	}

	if r.Method == "DELETE"{
		user, err := auth.Authenticate(r)

		if err!=nil{


			json.NewEncoder(w).Encode(
				map[string]string{
					"message":err.Error(),
				})
			return
		}


		var classId struct{
			Id uint64 `json:"id"`
		}

		err=json.NewDecoder(r.Body).Decode(&classId)


		qStr := fmt.Sprintf(`CALL deleteClass(%v,%v,1)`,user, classId.Id)
		_, err = HerokuDB.HEROKU_DB.Exec(context.Background(), qStr)

		json.NewEncoder(w).Encode(
			map[string]string{
				"message":"OK",
			})



	}
}

func EditClass(w http.ResponseWriter, r *http.Request){


}