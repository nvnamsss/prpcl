package repositories

import (
	"context"

	"github.com/nvnamsss/prpcl/adapters/database"
	"github.com/nvnamsss/prpcl/models"
)

type WagerTransactionRepository interface {
	Create(ctx context.Context, transaction *models.WagerTransaction, txs ...database.DBAdapter) error
}

type wagerTransactionRepository struct {
	db database.DBAdapter
}

func (w *wagerTransactionRepository) Create(ctx context.Context, transaction *models.WagerTransaction, txs ...database.DBAdapter) error {
	var (
		tx = w.db
	)
	if len(txs) > 0 {
		tx = txs[0]
	}
	return tx.Gormer().WithContext(ctx).Create(transaction).Error
}

func NewWagerTransactionRepository(db database.DBAdapter) WagerTransactionRepository {
	return &wagerTransactionRepository{
		db: db,
	}
}
