package requisicao

import (
	"encoding/json"
	"net/http"
)

func LerRequisicao(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(data); err != nil {
		return err
	}

	return nil
}
