package storage

import (
	"github.com/JasurbekUz/order-service/storage/postgres"
	"github.com/JasurbekUz/order-service/storage/repo"
	"github.com/jmoiron/sqlx"
)

//
type Istorage interface {
	Order() repo.OrderStorageI
}

type storagePg struct {
	db        *sqlx.DB
	orderRepo repo.OrderStorageI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:        db,
		orderRepo: postgres.NewOrderRepo(db),
	}
}
func (s storagePg) Order() repo.OrderStorageI {
	return s.orderRepo
}
