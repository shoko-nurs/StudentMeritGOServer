package Urls

import (
	"StudentMerit/auth"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"os"

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

	json.NewEncoder(w).Encode(
		map[string]string{
			"Message":"Ok",
		})

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


	/// This is used for Heroku deployment
	herokuPort := os.Getenv("PORT")
	http.ListenAndServe(":"+herokuPort, r)

	//// This is used for localserver
	//http.ListenAndServe(":8080", r)

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



