package services

import (
	"context"
	"errors"
	"testing"

	"github.com/nvnamsss/prpcl/adapters/database"
	"github.com/nvnamsss/prpcl/dtos"
	mocksDB "github.com/nvnamsss/prpcl/mocks/adapters/database"
	mocksRepo "github.com/nvnamsss/prpcl/mocks/repositories"
	"github.com/nvnamsss/prpcl/models"
	"github.com/nvnamsss/prpcl/repositories"
	"github.com/stretchr/testify/mock"
)

func Test_wagerService_Place(t *testing.T) {
	type fields struct {
		wagerRepository repositories.WagerRepository
	}
	type args struct {
		ctx context.Context
		req *dtos.PlaceWagerRequest
	}
	var (
		wagerRepository                                       = &mocksRepo.WagerRepository{}
		errWagerRepository                                    = &mocksRepo.WagerRepository{}
		reqs               map[string]*dtos.PlaceWagerRequest = map[string]*dtos.PlaceWagerRequest{
			"good": {
				TotalWagerValue:   10000.1,
				Odds:              30,
				SellingPercentage: 5,
				SellingPrice:      1000.64,
			},
			"sell below percentage": {
				TotalWagerValue:   10000.1,
				Odds:              30,
				SellingPercentage: 100,
				SellingPrice:      1000.64,
			},
			"create error": {
				TotalWagerValue:   10000,
				Odds:              1,
				SellingPercentage: 10,
				SellingPrice:      10000,
			},
		}
	)
	wagerRepository.On("Create", mock.Anything, mock.Anything).Return(nil)
	errWagerRepository.On("Create", mock.Anything, mock.Anything).Return(errors.New("just an error"))

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "good",
			fields: fields{wagerRepository: wagerRepository},
			args: args{
				ctx: context.Background(),
				req: reqs["good"],
			},
			wantErr: false,
		},
		{
			name:   "sell below percentage",
			fields: fields{wagerRepository: wagerRepository},
			args: args{
				ctx: context.Background(),
				req: reqs["sell below percentage"],
			},
			wantErr: true,
		},
		{
			name:   "create error",
			fields: fields{wagerRepository: errWagerRepository},
			args: args{
				ctx: context.Background(),
				req: reqs["create error"],
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &wagerService{
				wagerRepository: tt.fields.wagerRepository,
			}
			got, err := s.Place(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Place() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if got.CurrentSellingPrice != got.SellingPrice {
				t.Errorf("Place() current_selling_price != selling_price")
			}

			if got.AmountSold != nil {
				t.Error("Place() amount_sold != nil")
			}

			if got.PercentageSold != nil {
				t.Error("Place() percentage_sold != nil")
			}
		})
	}
}

func Test_wagerService_Buy(t *testing.T) {
	type fields struct {
		wagerRepository            repositories.WagerRepository
		wagerTransactionRepository repositories.WagerTransactionRepository
		db                         database.DBAdapter
	}
	type args struct {
		ctx context.Context
		req *dtos.BuyWagerRequest
	}
	var (
		wagerRepository                                                = &mocksRepo.WagerRepository{}
		errWagerRepository                                             = &mocksRepo.WagerRepository{}
		wagerTransactionRepository                                     = &mocksRepo.WagerTransactionRepository{}
		errWagerTransactionRepository                                  = &mocksRepo.WagerTransactionRepository{}
		db                                                             = &mocksDB.DBAdapter{}
		reqs                          map[string]*dtos.BuyWagerRequest = map[string]*dtos.BuyWagerRequest{
			"good": {
				WagerID:     1,
				BuyingPrice: 10,
			},
			"wager not found": {
				WagerID:     0,
				BuyingPrice: 10,
			},
			"get wager error": {},
			"buy error": {
				WagerID:     2,
				BuyingPrice: 10,
			},
			"cannot buy more": {
				WagerID:     1,
				BuyingPrice: 1200,
			},
		}
		wagers map[string]*models.Wager = map[string]*models.Wager{
			"good": {
				ID:                  1,
				CurrentSellingPrice: 1000,
			},
			"buy error": {
				ID:                  2,
				CurrentSellingPrice: 1000,
			},
		}
	)

	wagerRepository.On("Get", mock.Anything, int64(1)).Return(wagers["good"], nil)
	wagerRepository.On("Get", mock.Anything, int64(2)).Return(wagers["buy error"], nil)
	wagerRepository.On("Get", mock.Anything, int64(0)).Return(nil, nil)
	wagerRepository.On("Buy", mock.Anything, wagers["good"], mock.Anything, mock.Anything).Return(nil)
	wagerRepository.On("Buy", mock.Anything, wagers["buy error"], mock.Anything, mock.Anything).Return(errors.New("just an error"))
	errWagerRepository.On("Get", mock.Anything, mock.Anything).Return(nil, errors.New("just an error"))

	wagerTransactionRepository.On("Create", mock.Anything, mock.Anything).Return(nil)
	errWagerTransactionRepository.On("Create", mock.Anything, mock.Anything).Return(errors.New("just an error"))

	db.On("Begin", mock.Anything).Return(db)
	db.On("RollbackUselessCommitted").Return()
	db.On("Commit").Return()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "good",
			fields: fields{
				wagerRepository:            wagerRepository,
				wagerTransactionRepository: wagerTransactionRepository,
				db:                         db,
			},
			args: args{
				ctx: context.Background(),
				req: reqs["good"],
			},
			wantErr: false,
		},
		{
			name: "wager not found",
			fields: fields{
				wagerRepository:            wagerRepository,
				wagerTransactionRepository: wagerTransactionRepository,
				db:                         db,
			},
			args: args{
				ctx: context.Background(),
				req: reqs["wager not found"],
			},
			wantErr: true,
		},
		{
			name: "get wager error",
			fields: fields{
				wagerRepository:            errWagerRepository,
				wagerTransactionRepository: errWagerTransactionRepository,
				db:                         db,
			},
			args: args{
				ctx: context.Background(),
				req: reqs["get wager error"],
			},
			wantErr: true,
		},
		{
			name: "buy error",
			fields: fields{
				wagerRepository:            wagerRepository,
				wagerTransactionRepository: errWagerTransactionRepository,
				db:                         db,
			},
			args: args{
				ctx: context.Background(),
				req: reqs["buy error"],
			},
			wantErr: true,
		},
		{
			name: "cannot buy more",
			fields: fields{
				wagerRepository:            wagerRepository,
				wagerTransactionRepository: errWagerTransactionRepository,
				db:                         db,
			},
			args: args{
				ctx: context.Background(),
				req: reqs["cannot buy more"],
			},
			wantErr: true,
		},
		{
			name: "create transaction error",
			fields: fields{
				wagerRepository:            wagerRepository,
				wagerTransactionRepository: errWagerTransactionRepository,
				db:                         db,
			},
			args: args{
				ctx: context.Background(),
				req: reqs["good"],
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &wagerService{
				wagerRepository:            tt.fields.wagerRepository,
				wagerTransactionRepository: tt.fields.wagerTransactionRepository,
				db:                         tt.fields.db,
			}
			_, err := s.Buy(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Buy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

		})
	}
}

func Test_wagerService_List(t *testing.T) {
	type fields struct {
		wagerRepository repositories.WagerRepository
	}
	type args struct {
		ctx context.Context
		req *dtos.ListWagersRequest
	}

	var (
		wagerRepository    = &mocksRepo.WagerRepository{}
		errWagerRepository = &mocksRepo.WagerRepository{}
		wagers             = []*models.Wager{
			{
				ID: 1,
			},
			{
				ID: 2,
			},
		}
	)
	wagerRepository.On("List", mock.Anything, mock.Anything).Return(wagers, nil)
	errWagerRepository.On("List", mock.Anything, mock.Anything).Return(nil, errors.New("just an error"))

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "good",
			fields: fields{
				wagerRepository: wagerRepository,
			},
			args: args{
				ctx: context.Background(),
				req: &dtos.ListWagersRequest{},
			},
			wantErr: false,
		},
		{
			name: "list error",
			fields: fields{
				wagerRepository: errWagerRepository,
			},
			args: args{
				ctx: context.Background(),
				req: &dtos.ListWagersRequest{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &wagerService{
				wagerRepository: tt.fields.wagerRepository,
			}
			_, err := s.List(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

		})
	}
}
