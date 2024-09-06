package save

import (
	"browser-remote-server/internal/storage"
	resp "browser-remote-server/lib/api/response"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
)

type Request struct {
	URL         string              `json:"url"`
	ElementInfo storage.ElementInfo `json:"element-info"`
}

type Response struct {
	resp.Response
	ElementId int `json:"element-id"`
}

type ElementSaver interface {
	SaveElement(elementInfo storage.ElementInfo) (int, error)
}

func New(log *slog.Logger, elementSaver ElementSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.save.New"

		log = log.With(slog.String("op", op))

		var req Request

		err := render.DecodeJSON(r.Body, &req)

		if err != nil {
			log.Error("failed to decode request body")

			render.JSON(w, r, resp.Error("failed to decode request"))

			return
		}

		id, err := elementSaver.SaveElement(req.ElementInfo)

		if err != nil {
			log.Error("failed to decode request body")

			render.JSON(w, r, resp.Error("failed to decode request"))

			return
		}

		log.Info("new element added", slog.Int("id", id))

		responseOK(w, r, id)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request, id int) {
	render.JSON(w, r, Response{
		Response:  resp.OK(),
		ElementId: id,
	})
}
