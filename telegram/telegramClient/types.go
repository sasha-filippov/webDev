package telegramClient

type Chat struct {
	ChatID int `json:"id"`
}

type InlineButton struct {
	btnTxt string `json:"text"`
}

type ReplyMarkup struct {
	keyboard []InlineButton `json:"inline_keyboard"`
	oneTime  bool           `json:"one_time_keyboard"`
}

type IncomingMessage struct {
	Chat        Chat           `json:"chat"`
	Text        string         `json:"text"`
	ReplyMarkup *[]ReplyMarkup `json:"reply_markup"`
}

type Update struct {
	Message  *IncomingMessage `json:"message"`
	UpdateID int              `json:"update_id"`
}

type UpdatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}
