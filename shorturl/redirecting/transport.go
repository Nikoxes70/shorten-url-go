package redirecting

import (
	"net/http"
)

type transporter struct {
	service Servicer
}

type CreateShortURLRequest struct {
	Redirections map[string]interface{}
	Type         int64
}

type CreateShortURLResponse struct {
	ID string `json:"id"`
}

type Servicer interface {
	Redirect(id string) (string, error)
}

func NewTransport(s Servicer) *transporter {
	return &transporter{
		s,
	}
}

func (t transporter) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "404 not found.", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Path[1:len(r.URL.Path)]

	url, err := t.service.Redirect(id)
	if err != nil {
		http.Error(w, "failed to redirect", http.StatusNotFound)
	}

	http.Redirect(w, r, url, http.StatusFound)

	return
}
