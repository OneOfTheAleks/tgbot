package telegram

type Message struct {
	UpdateId int        `json:"update_id"`
	Message  *InMessage `json:"message"`
}

type Response struct {
	Ok     bool      `json:"ok"`
	Result []Message `json:"result"`
}

type InMessage struct {
	Text string `json:"text"`
	From From   `json:"from"`
	Chat Chat   `json:"chat"`
}

type From struct {
	Uname string `json:"username"`
}

type Chat struct {
	Id int `json:"id"`
}
