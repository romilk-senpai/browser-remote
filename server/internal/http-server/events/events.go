package events

import (
	"browser-remote-server/internal/storage"
	"log/slog"
)

type EventData struct {
	Host    string
	Element storage.Element
}

type Event struct {
	EventData EventData
	Processed bool
}

type EventController struct {
	currentEvent Event
	logger       *slog.Logger
}

func New(log *slog.Logger) *EventController {
	return &EventController{
		currentEvent: Event{
			EventData{},
			true,
		},
		logger: log,
	}
}

func (c *EventController) PushEvent(eventData EventData) {
	if !c.currentEvent.Processed {
		c.logger.Warn("pushing new event despite previous event haven't beed processed")
	}

	c.currentEvent = Event{
		eventData,
		false,
	}
}

func (c *EventController) Current() *Event {
	if c.currentEvent.Processed {
		return nil
	}

	return &c.currentEvent
}
