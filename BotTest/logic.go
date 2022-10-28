package BotTest

import (
	"StudentMerit/HerokuDB"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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
	qStr := fmt.Sprintf(`INSERT INTO telegram(chat_id, text) VALUES(%v, %s)`,chatId, text)
	HerokuDB.HEROKU_DB.Exec(context.Background(), qStr)
	SendTextToTelegram(chatId, toSend)


}


func SendTextToTelegram(chat_id int64, text string) (string){
	telegramAPI := "https://api.telegram.org/bot"+os.Getenv("bot_token")+"/"+"sendMessage"

	bodyOBj := map[string]string{
		"chat_id": strconv.FormatInt(chat_id,10),
		"text":text,
	}

	marsh,_ := json.Marshal(bodyOBj)
	bodyReader := bytes.NewBuffer(marsh)

	req, _ := http.NewRequest("POST", telegramAPI, bodyReader)

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