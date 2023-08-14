package services

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/jeanSagaz/server/internal/domain/entity"
	"github.com/jeanSagaz/server/internal/domain/interfaces"
	"github.com/jeanSagaz/server/internal/dto"
)

type ExchangeRateService struct {
	IUsdBrlRepository interfaces.IUsdBrlRepository
}

func NewExchangeRateService(
	IUsdBrlRepository interfaces.IUsdBrlRepository,
) *ExchangeRateService {
	return &ExchangeRateService{
		IUsdBrlRepository: IUsdBrlRepository,
	}
}

func (h *ExchangeRateService) GetQuotation() (*dto.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	// resp, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var dollarRealExchangeRate dto.DollarRealExchangeRate
	err = json.Unmarshal(res, &dollarRealExchangeRate)
	if err != nil {
		return nil, err
	}

	usdBrl := entity.NewUsdBrl(dollarRealExchangeRate.UsdBrl.Code,
		dollarRealExchangeRate.UsdBrl.Codein,
		dollarRealExchangeRate.UsdBrl.Name,
		dollarRealExchangeRate.UsdBrl.High,
		dollarRealExchangeRate.UsdBrl.Low,
		dollarRealExchangeRate.UsdBrl.VarBid,
		dollarRealExchangeRate.UsdBrl.PctChange,
		dollarRealExchangeRate.UsdBrl.Bid,
		dollarRealExchangeRate.UsdBrl.Ask,
		dollarRealExchangeRate.UsdBrl.Timestamp,
		dollarRealExchangeRate.UsdBrl.CreateDate)
	_, err = h.IUsdBrlRepository.Insert(usdBrl)
	if err != nil {
		return nil, err
	}

	//return &dollarRealExchangeRate, nil
	r := &dto.Response{
		Bid: dollarRealExchangeRate.UsdBrl.Bid,
	}
	return r, nil
}
