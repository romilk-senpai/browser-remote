package page

import (
	resp "browser-remote-server/lib/api/response"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
)

func New(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("/page/index.html")
		if err != nil {
			log.Error("html parsing error", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("html error"))
			return
		}

		err = tmpl.Execute(w, r)

		if err != nil {
			log.Error("html execution error", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("html error"))
			return
		}
	}
}
