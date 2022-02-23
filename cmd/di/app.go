package di

import (
	"log"
	"net/http"

	"github.com/guil95/go-wire-di/internal/usecase"
)

type App struct {
	uc usecase.UseCase
}

func (app *App) Run() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		name := r.URL.Query().Get("name")

		phrase := app.uc.GetName(name)

		_, err := w.Write([]byte(phrase))
		if err != nil {
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
