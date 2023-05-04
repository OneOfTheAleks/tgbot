package telegramm

import (
	"errors"
	"tgbot/internal/clients"
	"tgbot/internal/clients/telegram"
	"tgbot/internal/develop"
	"tgbot/internal/storage"
)

type Proc struct {
	tg      *telegram.Client
	Offset  int
	Storage storage.Storage
}

type Data struct {
	ChatId   int    `json:"chatId"`
	UserName string `json:"username"`
}

func New(client *telegram.Client, storage storage.Storage) *Proc {
	return &Proc{
		tg:      client,
		Storage: storage,
	}
}

// получаем месенджи
// преобразуем месенджы телеграмма в наши абстрактные эвенты
func (p *Proc) Collect(limit int) ([]develop.Event, error) {
	update, err := p.tg.Update(p.Offset, limit)
	if err != nil {
		return nil, clients.Wrap("collect save", err)
	}

	if len(update) == 0 {
		return nil, nil
	}

	res := make([]develop.Event, 0, len(update))

	for _, e := range update {
		res = append(res, proc(e))
	}
	p.Offset = update[len(update)-1].UpdateId + 1
	return res, nil
}

func (p *Proc) Process(ev develop.Event) error {
	switch ev.Type {
	case develop.Message:
		return p.procMessage(ev)

	default:
		return clients.Wrap("process error", errors.New("unknown message type"))
	}
}

func (p *Proc) procMessage(ev develop.Event) error {
	data, ok := getData(ev)
	if !ok {
		return clients.Wrap("process error", errors.New("error get data"))
	}
	err := p.Cmd(ev.Text, data.ChatId, data.UserName)
	if err != nil {
		return clients.Wrap("process error", err)
	}
	return nil
}

func getData(ev develop.Event) (data Data, ok bool) {
	data, ok = ev.Data.(Data)
	return

}

// ----------------------------------------------------------------
func proc(u telegram.Message) develop.Event {
	uType := getType(u)
	res := develop.Event{
		Type: uType,
		Text: getTex(u),
	}
	if uType == develop.Message {
		res.Data = Data{
			ChatId:   u.Message.Chat.Id,
			UserName: u.Message.From.Uname,
		}
	}
	return res
}

func getTex(u telegram.Message) string {
	if u.Message == nil {
		return ""
	}

	return u.Message.Text
}

func getType(u telegram.Message) develop.Type {
	if u.Message == nil {
		return develop.Unknown
	}
	return develop.Message
}
