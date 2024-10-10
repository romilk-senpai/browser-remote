package save

import (
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
	Host storage.Host `json:"host-info"`
}

type HostProvider interface {
	Read(url string) (storage.Host, error)
}

func New(log *slog.Logger, provider HostProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.host.New"

		log = log.With(slog.String("op", op))

		var req Request

		err := render.DecodeJSON(r.Body, &req)

		if err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))

			render.JSON(w, r, resp.Error("failed to decode request"))

			return
		}

		host, err := provider.Read(req.Url)

		if err != nil {
			log.Error("failed to retreive host info", slog.String("error", err.Error()))

			render.JSON(w, r, resp.Error("failed to retreive host info"))

			return
		}

		responseOK(w, r, host)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request, host storage.Host) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
		Host:     host,
	})
}
