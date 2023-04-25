package telegram

type Message struct {
	UpdateId int    `json:"update_id"`
	Message  string `json:"message"`
}

type response struct {
	Ok     bool      `json:"ok"`
	Result []Message `json:"result"`
}
