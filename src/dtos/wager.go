package dtos

type PlaceWagerRequest struct {
	TotalWagerValue   float64 `json:"total_wager_value" binding:"min=1"`
	Odds              int     `json:"odds" binding:"min=1"`
	SellingPercentage int     `json:"selling_percentage" binding:"min=1,max=100"`
	SellingPrice      float64 `json:"selling_price" binding:"gt=0"`
}

type PlaceWagerResponse struct {
	ID                  int64    `gorm:"primary_key" json:"id"`
	TotalWagerValue     float64  `json:"total_wager_value" `
	Odds                int      `json:"odds" `
	SellingPercentage   float64  `json:"selling_percentage"`
	SellingPrice        float64  `json:"selling_price"`
	CurrentSellingPrice float64  `json:"current_selling_price"`
	PercentageSold      *float64 `json:"percentage_sold"`
	AmountSold          *int64   `json:"amount_sold"`
	PlacedAt            int64    `json:"placed_at"`
}

type BuyWagerRequest struct {
	WagerID     int64   `json:"-"`
	BuyingPrice float64 `json:"buying_price" binding:"gt=0"`
}

type BuyWagerResponse struct {
	ID          int64   `json:"id"`
	WagerID     int64   `json:"wager_id"`
	BuyingPrice float64 `json:"buying_price"`
	BoughtAt    int64   `json:"bought_at"`
}

type ListWagersRequest struct {
	Page     int `form:"page" json:"place"`
	PageSize int `form:"page_size" json:"cursor"`
}

type ListWagersResponse struct {
	ID                  int64    `gorm:"primary_key" json:"id"`
	TotalWagerValue     float64  `json:"total_wager_value" binding:"min=1"`
	Odds                int      `json:"odds" binding:"min=1"`
	SellingPercentage   float64  `json:"selling_percentage"`
	SellingPrice        float64  `json:"selling_price"`
	CurrentSellingPrice *float64 `json:"current_selling_price"`
	PercentageSold      *float64 `json:"percentage_sold"`
	AmountSold          *int64   `json:"amount_sold"`
	PlacedAt            int64    `json:"placed_at"`
}
