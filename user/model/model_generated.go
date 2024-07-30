// Code generated by gen; DO NOT EDIT.

package model

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ================================= Model =================================
func (d *Model) BeforeCreate(tx *gorm.DB) error {
	if d.ID == nil {
		d.ID = PtrOf(uuid.New())
	}
	return nil
}

func (d *Model) Bind(ctx *gin.Context) error {
	var r Model
	if err := ctx.ShouldBindJSON(&r); err != nil {
		return err
	}

	return nil
}

func (d *Model) BindWithReadonly(ctx *gin.Context, old Model) error {
	if err := d.Bind(ctx); err != nil {
		return err
	}
	if d.ID == nil {
		d.ID = old.ID
	}

	return nil
}

// ================================= Provider =================================
func (d *Provider) BeforeCreate(tx *gorm.DB) error {
	if d.ID == nil {
		d.ID = PtrOf(uuid.New())
	}
	return nil
}

func (d *Provider) Bind(ctx *gin.Context) error {
	var r Provider
	if err := ctx.ShouldBindJSON(&r); err != nil {
		return err
	}
	d.Provider = r.Provider
	d.Balance = r.Balance
	d.PendingRefund = r.PendingRefund
	d.LastResponseTokenCount = r.LastResponseTokenCount
	d.Nonce = r.Nonce

	return nil
}

func (d *Provider) BindWithReadonly(ctx *gin.Context, old Provider) error {
	if err := d.Bind(ctx); err != nil {
		return err
	}
	if d.ID == nil {
		d.ID = old.ID
	}

	return nil
}

// ================================= Refund =================================
func (d *Refund) BeforeCreate(tx *gorm.DB) error {
	if d.ID == nil {
		d.ID = PtrOf(uuid.New())
	}
	return nil
}

func (d *Refund) Bind(ctx *gin.Context) error {
	var r Refund
	if err := ctx.ShouldBindJSON(&r); err != nil {
		return err
	}
	d.Provider = r.Provider
	d.Amount = r.Amount
	d.Processed = r.Processed

	return nil
}

func (d *Refund) BindWithReadonly(ctx *gin.Context, old Refund) error {
	if err := d.Bind(ctx); err != nil {
		return err
	}
	if d.ID == nil {
		d.ID = old.ID
	}
	if d.Index == nil {
		d.Index = old.Index
	}

	return nil
}

// ================================= Service =================================
func (d *Service) BeforeCreate(tx *gorm.DB) error {
	if d.ID == nil {
		d.ID = PtrOf(uuid.New())
	}
	return nil
}

func (d *Service) Bind(ctx *gin.Context) error {
	var r Service
	if err := ctx.ShouldBindJSON(&r); err != nil {
		return err
	}
	d.Name = r.Name
	d.Type = r.Type
	d.URL = r.URL
	d.InputPrice = r.InputPrice
	d.OutputPrice = r.OutputPrice

	return nil
}

func (d *Service) BindWithReadonly(ctx *gin.Context, old Service) error {
	if err := d.Bind(ctx); err != nil {
		return err
	}
	if d.ID == nil {
		d.ID = old.ID
	}

	return nil
}
