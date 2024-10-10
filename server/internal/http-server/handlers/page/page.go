package page

import (
	resp "browser-remote-server/lib/api/response"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
)

type Button struct {
	Name string
}

type PageData struct {
	Buttons []Button
}

func New(log *slog.Logger, htmlpath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		buttons := []Button{
			{Name: "Button 1"},
			{Name: "Button 2"},
			{Name: "Button 3"},
		}

		data := PageData{Buttons: buttons}

		tmpl, err := template.ParseFiles(htmlpath)
		if err != nil {
			log.Error("html parsing error", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("html error"))
			return
		}

		err = tmpl.Execute(w, data)

		if err != nil {
			log.Error("html execution error", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("html error"))
			return
		}
	}
}
