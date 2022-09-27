package Urls

import (
	"StudentMerit/HerokuDB"
	"StudentMerit/auth"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	//"net/url"
)



var APIEP = map[string]string{

	"host"            : "https://shokonurs-student-merit.herokuapp.com",
	"manage_classes"  : "/api/manage_classes",
	"get_endpoints"   : "/api/get_endpoints",
	"manage_scores"   : "/api/manage_scores",
	"manage_students" : "/api/manage_students",
	"manage_records"  : "/api/manage_records",
	"student_records" : "/api/student_records",

}


func testingFunc(w http.ResponseWriter, r *http.Request){
	qStr := fmt.Sprintf(`SELECT addClass('10b',1)`)
	row := HerokuDB.HEROKU_DB.QueryRow(context.Background(), qStr)

	var x []interface{}
	row.Scan(&x)
	fmt.Println(x)

}


func GetEndpoints(w http.ResponseWriter, r *http.Request){
	EnableCORSALL(&w)
	_, err := auth.Authenticate(r)
	if err!=nil{

		json.NewEncoder(w).Encode(
			map[string]string{
				"message":err.Error(),
			})
		return
	}

	rp := map[string] interface{}{
		"ep":APIEP,
	}

	json.NewEncoder(w).Encode(rp)

}


func RunServerFunc(){

	r := mux.NewRouter()
	r.HandleFunc(APIEP["get_endpoints"], GetEndpoints)
	r.HandleFunc(APIEP["manage_classes"], classMainPageAPIView)
	r.HandleFunc(APIEP["manage_scores"], scoresManager)
	r.HandleFunc(APIEP["manage_students"], studentsManager)
	r.HandleFunc(APIEP["manage_students"]+"/{id:[0-9]+}", studentsManager)
	r.HandleFunc(APIEP["manage_records"], recordManager)

	r.HandleFunc(APIEP["manage_records"]+"/{id:[0-9]+}", recordManager)

	r.HandleFunc(APIEP["student_records"]+"/{id:[0-9]+}", getRecordsForStudent)
	r.HandleFunc("/", testingFunc)
	http.ListenAndServe(":8080", r)

}



// This will make subrouters. Like include urlpatterns in django
//func APIListener(r *mux.Router){
//	api := r.PathPrefix("/api").Subrouter()
//	api.HandleFunc("/test", testingAPIVIew)
//
//
//	// Add class subrouter
//	class := api.PathPrefix(APIEP["manage_classes"]).Subrouter()
//	class.HandleFunc("",classMainPageAPIView)
//
//}



