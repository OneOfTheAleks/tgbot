package storage

type Storage interface {
	Save(m *Msg) error
	Show() (*Msg, error)
	Remove(tag string) error
}

const ErrorSave = "Error saving"

type Msg struct {
	Txt string
	Tag string
}
