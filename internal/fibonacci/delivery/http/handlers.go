package http

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (m *Handler) getFibonacciSequence(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	x, y := r.FormValue("from"), r.FormValue("to")

	from, err := strconv.ParseUint(x, 10, 32)
	if err != nil {
		_ = enc.Encode("Invalid parameter 'from'")
		return
	}

	to, err := strconv.ParseUint(y, 10, 32)
	if err != nil {
		_ = enc.Encode("Invalid parameter 'to'")
		return
	}

	sequence, err := m.usecase.GetFibonacciSequence(r.Context(), uint32(from), uint32(to))
	if err != nil {
		_ = enc.Encode(err.Error())
		return
	}

	enc.Encode(sequence)
}
