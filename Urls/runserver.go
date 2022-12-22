package Urls

import (
	"StudentMerit/BotTest"
	"StudentMerit/auth"
	"encoding/json"

	"github.com/gorilla/mux"
	"net/http"
	//"net/url"
)


// Set mode=1 for local development,
// Set mode=2 for Remote (AWS)
var mode = 1



var APIEP = map[string]string{
	"test"            : "123",
	"manage_classes"  : "/api/manage_classes",
	"get_endpoints"   : "/api/get_endpoints",
	"manage_scores"   : "/api/manage_scores",
	"manage_students" : "/api/manage_students",
	"manage_records"  : "/api/manage_records",
	"student_records" : "/api/student_records",
	"class_students"  : "/api/class_students",
	"top_students"	  : "/api/top_students",
}

func GetAPIEP() {


		APIEP["host"] = "http://localhost:8080"


}


func testingFunc(w http.ResponseWriter, r *http.Request){

	json.NewEncoder(w).Encode(
		map[string]string{
			"Message":"123132312312",
		})

}


func GetEndpoints(w http.ResponseWriter, r *http.Request){

	EnableCORSALL(&w, r)

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
	r.HandleFunc(APIEP["class_students"]+"/{id:[0-9]+}", classStudents)
	r.HandleFunc(APIEP["top_students"], getTopStudents)

	r.HandleFunc(APIEP["manage_records"], recordManager)

	r.HandleFunc(APIEP["manage_records"]+"/{id:[0-9]+}", recordManager)

	r.HandleFunc(APIEP["student_records"]+"/{id:[0-9]+}", getRecordsForStudent)
	r.HandleFunc("/", testingFunc)

	r.HandleFunc("/telegram_bot",BotTest.TelegramBotTest )




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



