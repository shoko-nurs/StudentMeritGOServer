package Urls

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)




func testingAPIVIew (w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{

		var testMap = map[string]interface{}{
			"id":123,
			"name":"Nurs",
			"surname":"Shoko",
			"date":time.Now(),
		}

		json.NewEncoder(w).Encode(testMap)
	}



	if r.Method == "POST"{

		data := map[string]string{}

		err := json.NewDecoder(r.Body).Decode(&data)

		if err!=nil{
			fmt.Println(err.Error()+"\n")
		}

		fmt.Println(data)
		for key, value := range data{
			fmt.Printf("%T ,%T\n", key, value)
		}
		json.NewEncoder(w).Encode(data)
	}
}


type registrationStr struct {
	Email string		`json:"email"`
	Password1 string	`json:"password1"`
	Password2 string	`json:"password2"`

}

func registrationAPIView(w http.ResponseWriter, r *http.Request){

	if r.Method == "POST"{
		var newUser registrationStr
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err!=nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		

	}
}