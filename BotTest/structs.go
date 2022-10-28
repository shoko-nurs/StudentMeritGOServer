package BotTest

type User struct{
	UserId int64	`json:"id"`
	Username string `json:"username"`

}

type Update struct{
	UpdateId int64     `json:"update_id"`
	Message Message  `json:"message"`

}


type Message struct{
	Text string   `json:"text"`
	Chat Chat     `json:"chat"`
	From User	  `json:"from"`

}


type Chat struct{
	Id int64 `json:"id"`
}
