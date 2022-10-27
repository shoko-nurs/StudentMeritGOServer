package BotTest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func TelegramBotTest(w http.ResponseWriter, r *http.Request){

	var update Update
	body := json.NewDecoder(r.Body).Decode(&update)
	fmt.Println(body)
	json.NewEncoder(w).Encode(update.Message.Text)

}


func SendTextToTelegram(chat_id int, text string) (string, error){
	telegramAPI := "https://api.telegram.org/bot"+os.Getenv("bot_token")
	response,err := http.PostForm()
}