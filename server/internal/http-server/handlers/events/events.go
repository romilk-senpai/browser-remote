package events

import (
	"browser-remote-server/internal/http-server/events"
	"browser-remote-server/internal/storage"
	resp "browser-remote-server/lib/api/response"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
)

type Request struct {
	Url string `json:"url"`
}

type Response struct {
	resp.Response
	Element *storage.Element
}

func New(log *slog.Logger, eventController *events.EventController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.event.New"
		log = log.With(slog.String("op", op))

		var req Request
		err := render.DecodeJSON(r.Body, &req)

		if err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))

			render.JSON(w, r, resp.Error("failed to decode request"))

			return
		}

		currentEvent := eventController.Current()

		if currentEvent == nil {
			log.Info("event already processed at", slog.String("host", req.Url))

			responseOK(w, r, nil)

			return
		}

		currentEvent.Processed = true

		log.Info("processed new event", slog.String("host", req.Url), slog.Int("id", currentEvent.EventData.Element.Id))

		responseOK(w, r, &currentEvent.EventData.Element)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request, element *storage.Element) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
		Element:  element,
	})
}
