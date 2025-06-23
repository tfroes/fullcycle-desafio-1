package dbservice

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CotacaoRepo struct {
	DB *gorm.DB
}

func NewCotacaoRepo(db *gorm.DB) *CotacaoRepo {
	return &CotacaoRepo{DB: db}
}

func NewCotacao(dto InsereCotacaoDTO) *Cotacao {
	entidade := &Cotacao{
		Id:          uuid.New(),
		Code:        dto.Code,
		Codein:      dto.Codein,
		Name:        dto.Name,
		High:        dto.High,
		Low:         dto.Low,
		VarBid:      dto.VarBid,
		PctChange:   dto.PctChange,
		Bid:         dto.Bid,
		Ask:         dto.Ask,
		Timestamp:   dto.Timestamp,
		Create_date: dto.Create_date,
	}

	return entidade
}

func (c *CotacaoRepo) Create(ctx context.Context, cotacao *Cotacao) error {
	return c.DB.WithContext(ctx).Create(cotacao).Error
}
