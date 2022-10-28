package main

import (
	"StudentMerit/BotTest"
	"StudentMerit/Urls"
)


func main() {

	BotTest.SetBotEnv()
	Urls.GetAPIEP()
	Urls.RunServerFunc()


}
