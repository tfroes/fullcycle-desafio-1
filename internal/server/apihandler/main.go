package apihandler

import (
	"client-server/internal/server/cotacaoservice"
	"client-server/internal/server/dbservice"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type CotacaoApiHandler struct {
	Repo *dbservice.CotacaoRepo
}

func NewCotacaoApiHandler(repo dbservice.CotacaoRepo) *CotacaoApiHandler {
	return &CotacaoApiHandler{
		Repo: &repo,
	}
}

func (api CotacaoApiHandler) GetCotacaoHandler(w http.ResponseWriter, r *http.Request) {
	var ctx, cancel = context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	cotacao, err := cotacaoservice.BuscaCotacao(ctx)

	if errors.Is(err, context.DeadlineExceeded) {
		w.WriteHeader(http.StatusRequestTimeout)
		json.NewEncoder(w).Encode(&ErrorModel{ctx.Err().Error()})
		return
	}

	if err != nil {
		fmt.Printf("[Error] Busca cotação com timetout!\n")

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&ErrorModel{err.Error()})
		return
	}

	cotacaoDto := CotacaoMapInserecotacao(cotacao)
	cotacaoDB := dbservice.NewCotacao(*cotacaoDto)

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	err = api.Repo.Create(ctx, cotacaoDB)

	if errors.Is(err, context.DeadlineExceeded) {
		fmt.Printf("[Error] Salvar cotação DB timetout!\n")
		w.WriteHeader(http.StatusRequestTimeout)
		json.NewEncoder(w).Encode(&ErrorModel{ctx.Err().Error()})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&ErrorModel{err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotacao)
}

func CotacaoMapInserecotacao(cotacao *cotacaoservice.Cotacao) *dbservice.InsereCotacaoDTO {
	return &dbservice.InsereCotacaoDTO{
		Code:        cotacao.USDBRL.Code,
		Codein:      cotacao.USDBRL.Codein,
		Name:        cotacao.USDBRL.Name,
		High:        cotacao.USDBRL.High,
		Low:         cotacao.USDBRL.Low,
		VarBid:      cotacao.USDBRL.VarBid,
		PctChange:   cotacao.USDBRL.PctChange,
		Bid:         cotacao.USDBRL.Bid,
		Ask:         cotacao.USDBRL.Ask,
		Timestamp:   cotacao.USDBRL.Timestamp,
		Create_date: cotacao.USDBRL.Create_date,
	}
}
