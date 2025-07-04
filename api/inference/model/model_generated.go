// Code generated by gen; DO NOT EDIT.

package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

// ================================= Request =================================
func (d *Request) Bind(ctx *gin.Context) error {
	var r Request
	if err := ctx.ShouldBindJSON(&r); err != nil {
		return err
	}
	d.UserAddress = r.UserAddress
	d.Nonce = r.Nonce
	d.ServiceName = r.ServiceName
	d.InputFee = r.InputFee
	d.OutputFee = r.OutputFee
	d.Fee = r.Fee
	d.Signature = r.Signature
	d.TeeSignature = r.TeeSignature
	d.RequestHash = r.RequestHash
	d.Processed = r.Processed
	d.VLLMProxy = r.VLLMProxy

	return nil
}

func (d *Request) BindWithReadonly(ctx *gin.Context, old Request) error {
	if err := d.Bind(ctx); err != nil {
		return err
	}

	return nil
}

// ================================= Service =================================
func (d *Service) Bind(ctx *gin.Context) error {
	var r Service
	if err := ctx.ShouldBindJSON(&r); err != nil {
		return err
	}
	d.Name = r.Name
	d.Type = r.Type
	d.URL = r.URL
	d.ModelType = r.ModelType
	d.Verifiability = r.Verifiability
	d.InputPrice = r.InputPrice
	d.OutputPrice = r.OutputPrice

	return nil
}

func (d *Service) BindWithReadonly(ctx *gin.Context, old Service) error {
	if err := d.Bind(ctx); err != nil {
		return err
	}

	return nil
}

// ================================= User =================================
func (d *User) Bind(ctx *gin.Context) error {
	var r User
	if err := ctx.ShouldBindJSON(&r); err != nil {
		return err
	}
	d.User = r.User
	d.LastRequestNonce = r.LastRequestNonce
	d.LockBalance = r.LockBalance
	d.LastBalanceCheckTime = r.LastBalanceCheckTime
	d.Signer = r.Signer
	d.UnsettledFee = r.UnsettledFee

	return nil
}

func (d *User) BindWithReadonly(ctx *gin.Context, old User) error {
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
