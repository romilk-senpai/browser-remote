package save

import (
	"browser-remote-server/internal/storage"
	resp "browser-remote-server/lib/api/response"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
)

type Request struct {
	Url   string `json:"url"`
	Name  string `json:"name"`
	Query string `json:"query"`
}

type Response struct {
	resp.Response
	ElementId int `json:"element-id"`
}

type ElementSaver interface {
	SaveElement(url string, name string, query string) (storage.Element, error)
}

func New(log *slog.Logger, elementSaver ElementSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.save.New"

		log := log.With(slog.String("op", op))

		var req Request

		err := render.DecodeJSON(r.Body, &req)

		if err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))

			render.JSON(w, r, resp.Error("failed to decode request"))

			return
		}

		el, err := elementSaver.SaveElement(req.Url, req.Name, req.Query)

		if err != nil {
			log.Error("failed to save element", slog.String("error", err.Error()))

			render.JSON(w, r, resp.Error("failed to save element"))

			return
		}

		log.Info("new element added", slog.Int("id", el.Id))

		responseOK(w, r, el.Id)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request, id int) {
	render.JSON(w, r, Response{
		Response:  resp.OK(),
		ElementId: id,
	})
}
