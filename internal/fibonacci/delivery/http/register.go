package http

import (
	"net/http"

	"github.com/RinatNamazov/fbs-golang-test-task/internal/fibonacci"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	usecase fibonacci.UseCase
}

func NewHandler(r *httprouter.Router, uc fibonacci.UseCase) {
	h := &Handler{uc}

	r.HandlerFunc(http.MethodGet, "/getFibonacciSequence", h.getFibonacciSequence)
}
