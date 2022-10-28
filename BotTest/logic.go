package BotTest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
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


func SendTextToTelegram(chat_id int64, text string) (string, error){
	telegramAPI := "https://api.telegram.org/bot"+os.Getenv("bot_token")
	response,_ := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id": {strconv.FormatInt(chat_id,10)},
			"text": {text},

		})

	defer response.Body.Close()

	var bodyBytes,_ = ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	return bodyString, nil
}