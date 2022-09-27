package Urls

import (
	"net/http"
)


func EnableCORSGET(w *http.ResponseWriter){

	(*w).Header().Set("Access-Control-Allow-Methods", "GET")
	(*w).Header().Set("Access-Control-Allow-Origin","http://127.0.0.1:8000")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

}

func EnableCORSALL(w *http.ResponseWriter,mode int){
	if mode==1{
		(*w).Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8000")
	}else{
		(*w).Header().Set("Access-Control-Allow-Origin", "https://shokonurs-random-apps.herokuapp.com")
	}

	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT, PATCH")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

}