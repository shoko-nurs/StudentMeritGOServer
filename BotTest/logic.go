package BotTest

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var Actions = map[string]interface{}{

	"string":1,
}




func TelegramBotTest(w http.ResponseWriter, r *http.Request){


	var update Update
	json.NewDecoder(r.Body).Decode(&update)
	text := update.Message.Text
	userName := update.Message.From.Username
	chatId := update.Message.Chat.Id

	toSend := userName+" "+text
	SendTextToTelegram(chatId, toSend)


}


func SendTextToTelegram(chat_id int64, text string) (string){
	telegramAPI := "https://api.telegram.org/bot"+os.Getenv("bot_token")+"/"+"sendMessage"

	sendBody := url.Values{
		"chat_id": {strconv.FormatInt(chat_id,10)},
		"text": {text},

	}
	req, _ := http.NewRequest("POST", telegramAPI, strings.NewReader(sendBody.Encode()))

	hc := http.Client{}
	req.Header.Add("Content-Type", "application/json")

	hc.Do(req)
	return "Success"

	//response,err := http.PostForm(
	//	telegramAPI,
	//	url.Values{
	//		"chat_id": {strconv.FormatInt(chat_id,10)},
	//		"text": {text},
	//
	//	})
	//
	//if err!=nil{
	//	log.Printf("error when posting")
	//	return "", err
	//}
	//
	//defer response.Body.Close()
	//
	//var bodyBytes,_ = ioutil.ReadAll(response.Body)
	//bodyString := string(bodyBytes)
	//log.Printf("%s", bodyString)
	//return bodyString, nil
}