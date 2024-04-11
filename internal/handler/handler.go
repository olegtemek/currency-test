package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/olegtemek/currency-task/internal/model"
	"github.com/olegtemek/currency-task/internal/response"
)

type service interface {
	Save(date string) error
	Get(date string, code string) ([]*model.Currency, error)
}

type Handler struct {
	log     *slog.Logger
	service service
}

func New(log *slog.Logger, service service) *Handler {
	return &Handler{
		log:     log,
		service: service,
	}
}

// Save godoc
// @Summary Save
// @Description Save currency data by date
// @Tags currency
// @Accept json
// @Produce json
// @Param date path string true "Date in format YYYY.MM.DD"
// @Success 200 {object} response.Response
// @Router /currency/save/{date} [get]
func (h *Handler) Save(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	date, ok := vars["date"]
	if !ok {
		response.New(h.log, &w, r, fmt.Errorf("cannot parse date"), 400)
		return
	}

	err := h.service.Save(date)
	if err != nil {
		response.New(h.log, &w, r, fmt.Errorf("error: %s", err), 400)
		return
	}

	response.New(h.log, &w, r, true, 400)
}

// GetCurrency godoc
// @Summary Get currency data
// @Description Get currency data by date and code
// @Tags currency
// @Accept json
// @Produce json
// @Param date path string true "Date in format YYYY.MM.DD"
// @Param code path string false "Currency code (optional)"
// @Success 200 {object} []model.Currency
// @Router /currency/{date} [get]
// @Router /currency/{date}/{code} [get]
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	date, ok := vars["date"]
	if !ok {
		response.New(h.log, &w, r, fmt.Errorf("cannot parse date"), 400)
		return
	}

	code := vars["code"]

	currencies, err := h.service.Get(date, code)
	if err != nil {
		response.New(h.log, &w, r, fmt.Errorf("error: %s", err), 400)
		return
	}

	response.New(h.log, &w, r, currencies, 200)

}
