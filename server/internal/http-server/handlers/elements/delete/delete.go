package delete

import (
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

type ElementDeleter interface {
	DeleteElement(url string, id int) error
}

func New(log *slog.Logger, elementSaver ElementDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.delete.New"

		log = log.With(slog.String("op", op))

		var req Request

		err := render.DecodeJSON(r.Body, &req)

		if err != nil {
			log.Error("failed to decode request body")

			render.JSON(w, r, resp.Error("failed to decode request"))

			return
		}

		err = elementSaver.DeleteElement(req.Url, req.Id)

		if err != nil {
			log.Error("failed to delete element")

			render.JSON(w, r, resp.Error("failed to to delete element"))

			return
		}

		log.Info("element removed", slog.Int("id", req.Id))

		responseOK(w, r)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
	})
}
