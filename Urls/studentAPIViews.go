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



func studentsManager(w http.ResponseWriter, r*http.Request) {
	EnableCORSALL(&w, r)
	user, err := auth.Authenticate(r)
	if err != nil {
		json.NewEncoder(w).Encode(
			map[string]string{
				"message": err.Error(),
			})
		return
	}

	if r.Method == "GET" {

		vars := mux.Vars(r)
		var qStr string
		if vars["id"]==""{
			qStr = fmt.Sprintf(`SELECT * FROM student ORDER BY class_name`)
		}else{
			id:=vars["id"]
			qStr = fmt.Sprintf(`SELECT * FROM student WHERE id=%v`,id)
		}


		rows, err := HerokuDB.HEROKU_DB.Query(context.Background(), qStr)
		if err != nil {
			//fmt.Println(err)
			json.NewEncoder(w).Encode(
				map[string]string{
					"message": err.Error(),
				})
			return
		}

		var allStudents []Structures.Student
		for rows.Next() {
			var s Structures.Student
			rows.Scan(&s.Id, &s.Name, &s.Surname, &s.ClassId, &s.CurrentScore, &s.ClassName, &s.UserAdded)

			allStudents = append(allStudents, s)
		}

		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"students": allStudents,
				"ep":       APIEP,
			})

	}

	if r.Method == "POST" {
		var ns Structures.Student
		json.NewDecoder(r.Body).Decode(&ns)

		ns.UserAdded = user
		err= ns.Validate(r)
		if err!=nil{
			json.NewEncoder(w).Encode(
				map[string]string{
					"message": err.Error(),
				})
			return
		}

		qStr := fmt.Sprintf(`select addNewStudent('%s', '%s',%v, '%s',%v)`,
			ns.Name, ns.Surname, ns.ClassId, ns.ClassName, ns.UserAdded)

		row := HerokuDB.HEROKU_DB.QueryRow(context.Background(), qStr)
		var result int
		err = row.Scan(&result)

		if err!=nil{

			json.NewEncoder(w).Encode(
				map[string]string{
					"message": err.Error(),
				})
			return
		}

		var message string
		var status int
		if result == 10{
			message="Student is added"
			status = 200
		}else if result == 11 {
			message="This student is already added"
			status=400
		}else{
			message="You can't add more than 10 students"
			status=400
		}

		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"message": message,
				"status": status,
			})

		return
	}

	if r.Method == "DELETE"{
		var s Structures.Student
		json.NewDecoder(r.Body).Decode(&s)

		qStr := fmt.Sprintf(`SELECT deleteStudent(%v,%v,%v)`,s.Id,s.UserAdded,user)

		row := HerokuDB.HEROKU_DB.QueryRow(context.Background(), qStr)

		var result int
		err = row.Scan(&result)


		if err!=nil{
			json.NewEncoder(w).Encode(

				map[string]interface{}{
					"message":err.Error(),
					"status":400,
				})
				return
		}
		var message string
		var status int

		if result == 10{
			message="OK"
			status=200
		}else if result == 12{
			message="You can't add more than 10 students"
			status=400
		}else{
			message="Access Denied!"
			status=400
		}

		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"message":message,
				"status":status,
			})



	}

	if r.Method == "PUT"{
		var uS Structures.Student
		json.NewDecoder(r.Body).Decode(&uS)

		err = uS.Validate(r)
		if err!=nil{
			json.NewEncoder(w).Encode(
				map[string]string{
					"message":err.Error(),
				})
				return
		}

		qStr := fmt.Sprintf(`UPDATE student SET name='%s', surname='%s',class_id=%v, class_name='%s' where id=%v`,
			uS.Name, uS.Surname, uS.ClassId,uS.ClassName, uS.Id )
		
		_, err = HerokuDB.HEROKU_DB.Exec(context.Background(), qStr)
		if err!=nil{
			json.NewEncoder(w).Encode(
				map[string]string{
					"message":err.Error(),
				})
				return
		}

		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"message":"OK",
				"status":200,
			})




	}
}



func classStudents(w http.ResponseWriter, r *http.Request){
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
		class_id := mux.Vars(r)["id"]

		qStr := fmt.Sprintf(`SELECT * FROM STUDENT WHERE class_id=%v ORDER BY current_score DESC`,class_id)

		rows, err:= HerokuDB.HEROKU_DB.Query(context.Background(), qStr)
		if err!=nil{
			json.NewEncoder(w).Encode(
				map[string]string{
					"message":err.Error(),
				})
			return
		}

		var classStudents []Structures.Student
		for rows.Next(){
			var s Structures.Student
			rows.Scan(&s.Id, &s.Name, &s.Surname, &s.ClassId, &s.CurrentScore, &s.ClassName, &s.UserAdded)
			classStudents = append(classStudents,s)
		}

		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"students": classStudents,
				"ep":       APIEP,
			})

	}


}