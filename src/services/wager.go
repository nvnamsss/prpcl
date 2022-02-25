package services

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/nvnamsss/prpcl/adapters/database"
	"github.com/nvnamsss/prpcl/dtos"
	"github.com/nvnamsss/prpcl/errors"
	"github.com/nvnamsss/prpcl/logger"
	"github.com/nvnamsss/prpcl/models"
	"github.com/nvnamsss/prpcl/repositories"
)

type WagerService interface {
	Place(ctx context.Context, req *dtos.PlaceWagerRequest) (*dtos.PlaceWagerResponse, error)
	Buy(ctx context.Context, req *dtos.BuyWagerRequest) (*dtos.BuyWagerResponse, error)
	List(ctx context.Context, req *dtos.ListWagersRequest) ([]*dtos.ListWagersResponse, error)
}

type wagerService struct {
	wagerRepository            repositories.WagerRepository
	wagerTransactionRepository repositories.WagerTransactionRepository
	db                         database.DBAdapter
}

func (s *wagerService) Place(ctx context.Context, req *dtos.PlaceWagerRequest) (*dtos.PlaceWagerResponse, error) {
	var (
		wager models.Wager
		res   dtos.PlaceWagerResponse
	)

	if req.SellingPrice <= req.TotalWagerValue*(float64(req.SellingPercentage)/100) {
		logger.Context(ctx).Errorf("selling_price must be greater than total_wager_value * (selling_percentage / 100)")
		return nil, errors.New(errors.ErrInvalidRequest, "selling_price must be greater than total_wager_value * (selling_percentage / 100)")
	}

	_ = copier.Copy(&wager, req)
	wager.CurrentSellingPrice = req.SellingPrice

	if err := s.wagerRepository.Create(ctx, &wager); err != nil {
		logger.Context(ctx).Errorf("create wager got error :%v", err)
		return nil, errors.New(errors.ErrInternalServer, err.Error())
	}

	_ = copier.Copy(&res, wager)
	res.AmountSold = nil
	res.PercentageSold = nil
	res.PlacedAt = wager.CreatedAt.Unix()

	return &res, nil
}

func (s *wagerService) Buy(ctx context.Context, req *dtos.BuyWagerRequest) (*dtos.BuyWagerResponse, error) {
	var (
		wager            *models.Wager
		wagerTransaction models.WagerTransaction
		res              dtos.BuyWagerResponse
		err              error
	)

	tx := s.db.Begin()
	defer tx.RollbackUselessCommitted()

	wager, err = s.wagerRepository.Get(ctx, req.WagerID)
	if err != nil {
		logger.Context(ctx).Errorf("get wager %v got error: %v", req.WagerID, err)
		return nil, errors.New(errors.ErrInternalServer, err.Error())
	}

	if wager == nil {
		logger.Context(ctx).Errorf("wager %v not found", req.WagerID)
		return nil, errors.New(errors.ErrInvalidRequest, fmt.Sprintf("wager %v not found", req.WagerID))
	}

	if req.BuyingPrice > wager.CurrentSellingPrice {
		logger.Context(ctx).Errorf("cannot buy wager %v (%v - %v)", req.WagerID, req.BuyingPrice, wager.CurrentSellingPrice)
		return nil, errors.New(errors.ErrInvalidRequest, "wager is not enough to be bought")
	}

	if err := s.wagerRepository.Buy(ctx, wager, req.BuyingPrice); err != nil {
		logger.Context(ctx).Errorf("buy got error: %v", err.Error())
		return nil, errors.New(errors.ErrInternalServer, err.Error())
	}

	wagerTransaction.BuyingPrice = req.BuyingPrice
	wagerTransaction.WagerID = req.WagerID
	if err := s.wagerTransactionRepository.Create(ctx, &wagerTransaction); err != nil {
		logger.Context(ctx).Errorf("create transaction got error: %v", err)
		return nil, errors.New(errors.ErrInternalServer, err.Error())
	}

	tx.Commit()

	_ = copier.Copy(&res, wagerTransaction)
	res.BoughtAt = wagerTransaction.CreatedAt.Unix()
	return &res, nil
}

func (s *wagerService) List(ctx context.Context, req *dtos.ListWagersRequest) ([]*dtos.ListWagersResponse, error) {
	wagers, err := s.wagerRepository.List(ctx, req)
	if err != nil {
		logger.Context(ctx).Errorf("find wagers got error: %v", err)
		return nil, errors.New(errors.ErrInternalServer, err.Error())
	}

	var data []*dtos.ListWagersResponse = make([]*dtos.ListWagersResponse, 0)
	for _, v := range wagers {
		d := dtos.ListWagersResponse{}
		_ = copier.Copy(&d, v)
		d.PlacedAt = v.CreatedAt.Unix()
		if v.PercentageSold == 0 {
			d.PercentageSold = nil
		}
		if v.AmountSold == 0 {
			d.AmountSold = nil
		}
		data = append(data, &d)
	}

	return data, nil
}

func NewWagerService(wagerRepository repositories.WagerRepository,
	wagerTransactionRepository repositories.WagerTransactionRepository,
	db database.DBAdapter) WagerService {
	return &wagerService{
		wagerRepository:            wagerRepository,
		wagerTransactionRepository: wagerTransactionRepository,
		db:                         db,
	}
}
