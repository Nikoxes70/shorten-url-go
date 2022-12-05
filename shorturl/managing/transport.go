package managing

import (
	"encoding/json"
	"fmt"
	"net/http"

	"shorten-url-go/shorturl/shorturl"
)

type parser interface {
	RawToRedirections(raw map[string]interface{}) ([]shorturl.Redirection, error)
}

type transporter struct {
	service Servicer
	parser  parser
}

type CreateShortURLRequest struct {
	Redirections map[string]interface{}
	Type         int64
}

type CreateShortURLResponse struct {
	ID string `json:"id"`
}

type Servicer interface {
	CreateShortURL(redirections []shorturl.Redirection, idType IdType) (id string, err error)
}

func NewTransport(s Servicer, parser parser) *transporter {
	return &transporter{
		s,
		parser,
	}
}

func (t transporter) HandleManage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/shorturl" || r.Method != http.MethodPost {
		http.Error(w, "404 not found.", http.StatusMethodNotAllowed)
		return
	}

	var reqBody map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "server error - failed to parse request body", http.StatusInternalServerError)
		return
	}

	idType := alphanumeric
	idTypeRaw := r.URL.Query().Get("type")
	if idTypeRaw != "" {
		idType = IdType(idTypeRaw)
	}

	if !idType.IsValid() {
		http.Error(w, "unsupported id type", http.StatusBadRequest)
		return
	}

	redirections, err := t.parser.RawToRedirections(reqBody)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse request body - %v", err), http.StatusBadRequest)
		return
	}

	id, err := t.service.CreateShortURL(redirections, idType)
	if err != nil {
		http.Error(w, "server error - failed to create short url", http.StatusInternalServerError)
		return
	}

	responseBody := CreateShortURLResponse{
		ID: id,
	}

	b, err := json.Marshal(responseBody)
	if err != nil {
		http.Error(w, "server error - failed to create short url", 500)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Printf("failed to martshal response - %v\n", err)
	}

	return
}
