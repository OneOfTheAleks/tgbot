package telegramm

import (
	"strings"
	"tgbot/internal/clients"
	"tgbot/internal/storage"
)

const (
	Show  = "/show"
	Save  = "/save"
	Help  = "/help"
	Start = "/start"
	//----------------------------------------------------------------
	msgHelp           = `help help`
	msgStart          = `start`
	msgUnknown        = `unknown`
	msgNoSave         = `no-save`
	msgSave           = ` save`
	msgUnknownCommand = "Unknown command"
)

func (p *Proc) Cmd(text string, ChatId int, userName string) error {
	text = strings.TrimSpace(text)

	switch text {
	case Help:
		return p.sendHelp(ChatId)
	case Show:
		return p.showMessage(ChatId)
	case Save:
		return p.saveMessage(ChatId, text, userName)
	case Start:
		return p.sendHello(ChatId)
	default:
		p.tg.Send(ChatId, msgUnknownCommand)
	}

	return nil
}

func (p *Proc) saveMessage(chatId int, text, userName string) (err error) {
	defer func() { err = clients.Wrap("error save message", err) }()
	// to do добавить парсинг тега
	msg := &storage.Msg{
		Txt: text,
		Tag: "тест",
	}
	err = p.Storage.Save(msg)
	if err != nil {
		return err
	}
	err = p.tg.Send(chatId, msgSave)
	if err != nil {
		return err
	}
	return nil
}

func (p *Proc) showMessage(chatId int) (err error) {
	defer func() { err = clients.Wrap("error show message", err) }()
	show, err := p.Storage.Show()
	if err != nil {
		return err
	}
	err = p.tg.Send(chatId, show.Txt)
	if err != nil {
		return err
	}
	return nil
}
func (p *Proc) sendHelp(chatId int) error {
	return p.tg.Send(chatId, msgHelp)
}

func (p *Proc) sendHello(chatId int) error {
	return p.tg.Send(chatId, msgStart)
}
