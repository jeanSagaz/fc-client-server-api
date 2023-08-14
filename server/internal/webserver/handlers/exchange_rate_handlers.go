package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jeanSagaz/server/internal/application/services"
	"github.com/jeanSagaz/server/internal/domain/interfaces"
)

type ExchangeRateHandler struct {
	IUsdBrlRepository interfaces.IUsdBrlRepository
}

func NewExchangeRateHandler(
	IUsdBrlRepository interfaces.IUsdBrlRepository,
) *ExchangeRateHandler {
	return &ExchangeRateHandler{
		IUsdBrlRepository: IUsdBrlRepository,
	}
}

func (h *ExchangeRateHandler) GetQuotationTax(w http.ResponseWriter, r *http.Request) {
	exchangeRateService := services.NewExchangeRateService(h.IUsdBrlRepository)
	resp, err := exchangeRateService.GetQuotation()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
