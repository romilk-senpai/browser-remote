package trigger

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
	Id  int    `json:"id"`
}

type Response struct {
	resp.Response
}

type ElementProvider interface {
	GetElementById(url string, id int) (storage.Element, error)
}

func New(log *slog.Logger, eventController *events.EventController, provider ElementProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.trigger.New"

		log := log.With(slog.String("op", op))

		var req Request

		err := render.DecodeJSON(r.Body, &req)

		if err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))

			render.JSON(w, r, resp.Error("failed to decode request"))

			return
		}

		element, err := provider.GetElementById(req.Url, req.Id)

		if err != nil {
			log.Error("invalid event", slog.String("error", err.Error()), slog.String("host", req.Url), slog.Int("id", req.Id))

			render.JSON(w, r, resp.Error("invalid event"))

			return
		}

		eventController.PushEvent(events.EventData{
			Host:    req.Url,
			Element: element,
		})

		log.Info("pushed new event", slog.String("host", req.Url), slog.Int("id", req.Id))

		responseOK(w, r)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
	})
}
