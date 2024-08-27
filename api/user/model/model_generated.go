// Code generated by gen; DO NOT EDIT.

package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

// ================================= Provider =================================
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
	d.Signer = r.Signer

	return nil
}

func (d *Provider) BindWithReadonly(ctx *gin.Context, old Provider) error {
	if err := d.Bind(ctx); err != nil {
		return err
	}

	return nil
}

// ================================= Refund =================================
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
	if d.Index == nil {
		d.Index = old.Index
	}

	return nil
}

// ================================= SystemInfo =================================
func (d *SystemInfo) Bind(ctx *gin.Context) error {
	var r SystemInfo
	if err := ctx.ShouldBindJSON(&r); err != nil {
		return err
	}
	d.K = r.K
	d.V = r.V

	return nil
}

func (d *SystemInfo) BindWithReadonly(ctx *gin.Context, old SystemInfo) error {
	if err := d.Bind(ctx); err != nil {
		return err
	}

	return nil
}

//=============== implementation of sql.scanner and sql.valuer  ===============
func (m StringSlice) Value() (driver.Value, error) {
	return json.Marshal(m)
}
func (m *StringSlice) Scan(value interface{}) error {
	if v, ok := value.([]byte); ok {
		return json.Unmarshal(v, m)
	}
	return fmt.Errorf("can't convert %T to StringSlice", value)
}	
