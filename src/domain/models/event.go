package models

type Event struct {
	Service string
	Event   string
	Date    string
	Data    string
}

func NewEvent(service string, event string, date string, data string) *Event {
	return &Event{
		Service: service,
		Event:   event,
		Date:    date,
		Data:    data,
	}
}
