package healthcheck

import (
	"diving-log-book-service/internal/pkg/apihelper"
	"net/http"
)

type Healthcheck struct{}

func New() *Healthcheck {
	return &Healthcheck{}
}

func (h Healthcheck) Alive(w http.ResponseWriter, r *http.Request) {
	apihelper.Response(w, "OK")
}
