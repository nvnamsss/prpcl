package models

import (
	"time"

	"gorm.io/gorm"
)

type Wager struct {
	ID                  int64          `gorm:"primary_key" json:"id"`
	TotalWagerValue     float64        `json:"total_wager_value"`
	Odds                int            `json:"odds"`
	SellingPercentage   int            `json:"selling_percentage"`
	SellingPrice        float64        `json:"selling_price"`
	CurrentSellingPrice float64        `json:"current_selling_price"`
	PercentageSold      float64        `json:"percentage_sold"`
	AmountSold          int64          `json:"amount_sold"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `sql:"index"`
}

type WagerTransaction struct {
	ID          int64          `gorm:"primary_key" json:"id"`
	WagerID     int64          `json:"wager_id"`
	BuyingPrice float64        `json:"buying_price"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `sql:"index"`
}
