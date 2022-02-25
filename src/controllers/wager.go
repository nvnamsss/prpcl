package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nvnamsss/prpcl/dtos"
	"github.com/nvnamsss/prpcl/errors"
	"github.com/nvnamsss/prpcl/logger"
	"github.com/nvnamsss/prpcl/services"
	"github.com/nvnamsss/prpcl/utils"
)

type WagerController struct {
	Base
	wagerService services.WagerService
}

// @Summary Place wager
// @Description Place wager
// @Tags Wagers
// @Accept json
// @Produce json
// @Param values	body dtos.PlaceWagerRequest	true "body"
// @Success 200 {object} dtos.ListWagersResponse
// @Failure 401 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router /wagers [post]
func (h *WagerController) Place(c *gin.Context) {
	var (
		req dtos.PlaceWagerRequest
		res *dtos.PlaceWagerResponse
		err error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		logger.Context(c.Request.Context()).Errorf("validation error: %v", err)
		h.HandleError(c, errors.New(errors.ErrInvalidRequest, err.Error()))
		return
	}

	req.SellingPrice = utils.Round2(req.SellingPrice)

	if res, err = h.wagerService.Place(c.Request.Context(), &req); err != nil {
		h.HandleError(c, err)
		return
	}

	h.JSON201(c, res)
}

// @Summary Buy wager
// @Description Buy wager
// @Tags Wagers
// @Accept json
// @Produce json
// @Param id	path int64	true "wager id"
// @Param values	body dtos.BuyWagerRequest	true "body"
// @Success 200 {object} dtos.BuyWagerResponse
// @Failure 401 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router /buy/:id [post]
func (h *WagerController) Buy(c *gin.Context) {
	var (
		req dtos.BuyWagerRequest
		res *dtos.BuyWagerResponse
		err error
	)

	if req.WagerID, err = strconv.ParseInt(c.Param("id"), 10, 64); err != nil {
		logger.Context(c.Request.Context()).Errorf("validation error: %v", err)
		h.HandleError(c, errors.New(errors.ErrInvalidRequest, err.Error()))
		return
	}

	if err = c.ShouldBindJSON(&req); err != nil {
		logger.Context(c.Request.Context()).Errorf("validation error: %v", err)
		h.HandleError(c, errors.New(errors.ErrInvalidRequest, err.Error()))
		return
	}
	req.BuyingPrice = utils.Round2(req.BuyingPrice)

	if res, err = h.wagerService.Buy(c.Request.Context(), &req); err != nil {
		h.HandleError(c, err)
		return
	}

	h.JSON201(c, res)
}

// @Summary List wagers
// @Description List wagers
// @Tags Wagers
// @Accept json
// @Produce json
// @Param values	query dtos.ListWagersRequest	true "query"
// @Success 200 {object} dtos.ListWagersResponse
// @Failure 401 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router /wagers [get]
func (h *WagerController) List(c *gin.Context) {
	var (
		req dtos.ListWagersRequest
		res []*dtos.ListWagersResponse
		err error
	)

	if err = c.ShouldBindQuery(&req); err != nil {
		logger.Context(c.Request.Context()).Errorf("validation error: %v", err)
		h.HandleError(c, errors.New(errors.ErrInvalidRequest, err.Error()))
		return
	}

	if res, err = h.wagerService.List(c.Request.Context(), &req); err != nil {
		h.HandleError(c, err)
		return
	}

	h.JSON(c, res)
}

func NewWagerController(restaurantService services.WagerService) *WagerController {
	return &WagerController{
		wagerService: restaurantService,
	}
}
