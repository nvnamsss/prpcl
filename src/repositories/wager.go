package repositories

import (
	"context"
	"errors"

	"github.com/nvnamsss/prpcl/adapters/database"
	"github.com/nvnamsss/prpcl/dtos"
	"github.com/nvnamsss/prpcl/models"
	"gorm.io/gorm"
)

type WagerRepository interface {
	Get(ctx context.Context, id int64) (*models.Wager, error)
	Create(ctx context.Context, wager *models.Wager) error
	Buy(ctx context.Context, wager *models.Wager, price float64, txn ...database.DBAdapter) error
	List(ctx context.Context, req *dtos.ListWagersRequest) ([]*models.Wager, error)
}

type wagerRepository struct {
	db database.DBAdapter
}

func (r *wagerRepository) Get(ctx context.Context, id int64) (*models.Wager, error) {
	var (
		rs *models.Wager = &models.Wager{}
	)
	if err := r.db.Gormer().First(rs, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *wagerRepository) Create(ctx context.Context, wager *models.Wager) error {
	return r.db.Gormer().WithContext(ctx).Create(wager).Error
}

func (r *wagerRepository) Update(ctx context.Context, wager *models.Wager) error {
	return r.db.Gormer().WithContext(ctx).Model(wager).UpdateColumns(map[string]interface{}{
		"current_selling_price": wager.CurrentSellingPrice,
		"percentage_sold":       wager.PercentageSold,
		"amount_sold":           wager.AmountSold,
	}).Error
}

func (r *wagerRepository) Buy(ctx context.Context, wager *models.Wager, price float64, txs ...database.DBAdapter) error {
	var (
		tx database.DBAdapter = r.db
	)

	if len(txs) > 0 {
		tx = txs[0]
	}

	db := tx.Gormer().WithContext(ctx).Model(wager).
		Where("current_selling_price >= ?", price).
		Updates(map[string]interface{}{
			"current_selling_price": gorm.Expr("current_selling_price - ?", price),
			"percentage_sold":       gorm.Expr("percentage_sold + ?", price/wager.SellingPrice),
			"amount_sold":           gorm.Expr("amount_sold + ?", 1),
		})
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("cannot buy anymore")
	}
	return nil
}

func (r *wagerRepository) List(ctx context.Context, req *dtos.ListWagersRequest) ([]*models.Wager, error) {
	var (
		rs []*models.Wager
	)
	if err := r.db.Gormer().WithContext(ctx).
		Limit(req.PageSize).
		Offset((req.PageSize - 1) * req.PageSize).
		Find(&rs).Error; err != nil {
		return nil, err
	}
	return rs, nil
}

func NewWagerRepository(dbAdapter database.DBAdapter) WagerRepository {
	return &wagerRepository{
		db: dbAdapter,
	}
}
