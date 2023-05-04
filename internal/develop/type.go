package develop

// собиратель сообщений
type Collector interface {
	Collect(limit int) ([]Event, error)
}

type Process interface {
	Process(e Event) error
}

const Unknown = 0
const Message = 1

type Type int

type Event struct {
	Type Type
	Text string
	Data interface{}
}
