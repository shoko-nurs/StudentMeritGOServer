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

	//text := update.Message.Text
	//userName := update.Message.From.Username
	//chatId := update.Message.Chat.Id

	//toSend := userName+" "+text
	//usr := "user"
	//qStr := fmt.Sprintf(`INSERT INTO telegram(chat_id) VALUES(%v)`,chatId)
	qStr1 := `INSERT INTO telegram (chat_id) VALUES (123456)`
	HerokuDB.HEROKU_DB.Exec(context.Background(), qStr1)
	//SendTextToTelegram(chatId, toSend)
	//testingTelegram(chatId, toSend)

}

func testingTelegram(chat_id int64, text string){
	telegramAPI := "https://api.telegram.org/bot"+os.Getenv("bot_token") + "/sendMessage"

	bodyOBj := map[string]string{
		"chat_id": strconv.FormatInt(chat_id,10),
		"text":text,
	}

	marsh,_ := json.Marshal(bodyOBj)
	bodyReader := bytes.NewBuffer(marsh)

	req, _ := http.NewRequest("POST", telegramAPI, bodyReader)

	//hc := http.Client{}
	req.Header.Add("Content-Type", "application/json")

	//response, _ := hc.Do(req)

	//qStr := fmt.Sprintf(`INSERT INTO telegram(chat_id, username) values(%v, %v)`,response.StatusCode, response.Status)
	//HerokuDB.HEROKU_DB.Exec(context.Background(), qStr)

	qStr1 := fmt.Sprintf(`INSERT INTO telegram (chat_id) VALUES (%v)`,chat_id)
	HerokuDB.HEROKU_DB.Exec(context.Background(), qStr1)
}

func SendTextToTelegram(chat_id int64, text string) {
	telegramAPI := "https://api.telegram.org/bot"+os.Getenv("bot_token") + "/sendMessage"

	bodyOBj := map[string]string{
		"chat_id": strconv.FormatInt(chat_id,10),
		"text":text,
	}

	marsh,_ := json.Marshal(bodyOBj)
	bodyReader := bytes.NewBuffer(marsh)

	req, _ := http.NewRequest("POST", telegramAPI, bodyReader)

	hc := http.Client{}
	req.Header.Add("Content-Type", "application/json")

	response, _ := hc.Do(req)


	qStr:= fmt.Sprintf(`INSERT INTO telegram(text) values(%s)`,telegramAPI)
	qStr2:= fmt.Sprintf(`INSERT INTO telegram(response) values(%s)`,response.StatusCode)
	qStr3 := fmt.Sprintf(`insert into telegram(chat_id) values(%v)`,response.StatusCode)
	qStr4 := fmt.Sprintf(`INSERT INTO telegram(chat_id) values(1)`)
	HerokuDB.HEROKU_DB.Exec(context.Background(), qStr)
	HerokuDB.HEROKU_DB.Exec(context.Background(), qStr2)
	HerokuDB.HEROKU_DB.Exec(context.Background(), qStr3)
	HerokuDB.HEROKU_DB.Exec(context.Background(), qStr4)

	//qStr := fmt.Sprintf(`INSERT INTO telegram(chat_id) values(%v)`,response.StatusCode)
	//HerokuDB.HEROKU_DB.Exec(context.Background(), qStr)

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