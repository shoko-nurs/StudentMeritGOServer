package main

import (
	"StudentMerit/Urls"
)

func main() {
	//fmt.Println(Urls.GetAllowedOrigins())
	SetENV()
	//fmt.Println(HerokuDB.GetHerokuDB())
	Urls.RunServerFunc()


}
