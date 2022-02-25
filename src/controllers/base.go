package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nvnamsss/prpcl/errors"
)

type Base struct {
}

// JSON responds a HTTP request with JSON data.
func (h *Base) JSON(c *gin.Context, data interface{}) {
	if data != nil {
		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(errors.ErrNoResponse.Status(), errors.New(errors.ErrNoResponse))
	}
}

// JSON responds a HTTP request with JSON data.
func (h *Base) JSON201(c *gin.Context, data interface{}) {
	if data != nil {
		c.JSON(http.StatusCreated, data)
	} else {
		c.JSON(errors.ErrNoResponse.Status(), errors.New(errors.ErrNoResponse))
	}
}

// HandleError handles error of HTTP request.
func (h *Base) HandleError(c *gin.Context, err error) {
	if err != nil {
		appErr, ok := err.(errors.AppError)
		if ok {
			c.JSON(appErr.ErrorCode.Status(), appErr)
		} else {
			c.JSON(errors.ErrInternalServer.Status(), errors.New(errors.ErrInternalServer))
		}
	} else {
		c.JSON(errors.ErrNoResponse.Status(), errors.New(errors.ErrNoResponse))
	}
}
