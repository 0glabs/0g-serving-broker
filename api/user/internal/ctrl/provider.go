package ctrl

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/0glabs/0g-serving-agent/common/contract"
	"github.com/0glabs/0g-serving-agent/common/errors"
	"github.com/0glabs/0g-serving-agent/user/model"
	"github.com/ethereum/go-ethereum/common"
)

func (c Ctrl) CreateProviderAccount(ctx context.Context, providerAddress common.Address, account model.Provider) error {
	balance := big.NewInt(0)
	balance.SetInt64(*account.Balance)

	signer, err := c.getSigner(ctx)
	if err != nil {
		return errors.Wrap(err, "get signer from db")
	}
	if err := c.contract.CreateProviderAccount(ctx, providerAddress, *balance, signer); err != nil {
		return errors.Wrap(err, "create provider account in contract")
	}
	if err := c.db.CreateProviderAccounts([]model.Provider{account}); err != nil {
		rollBackErr := c.SyncProviderAccount(ctx, providerAddress)
		if rollBackErr != nil {
			log.Printf("resync account in db: %s", rollBackErr.Error())
		}
	}
	return errors.Wrap(err, "create provider account in db")
}

// GetProviderAccount get account information from contract and database, and ignore processed items
func (c Ctrl) GetProviderAccount(ctx context.Context, providerAddress common.Address) (model.Provider, error) {
	account, err := c.contract.GetProviderAccount(ctx, providerAddress)
	if err != nil {
		return model.Provider{}, errors.Wrap(err, "get account from contract")
	}
	rets, err := c.backfillProviderAccount([]contract.Account{account})
	return rets[0], err
}

// ListProviderAccount get account information from contract and database, and ignore processed items
func (c Ctrl) ListProviderAccount(ctx context.Context) ([]model.Provider, error) {
	accounts, err := c.contract.ListProviderAccount(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "list account from contract")
	}
	return c.backfillProviderAccount(accounts)
}

func (c Ctrl) ChargeProviderAccount(ctx context.Context, providerAddress common.Address, account model.Provider) error {
	if account.Balance == nil {
		return fmt.Errorf("nil account.Balance")
	}
	_, err := c.getProviderAccountFromContract(ctx, providerAddress)
	if err != nil {
		return err
	}

	signer, err := c.getSigner(ctx)
	if err != nil {
		return errors.Wrap(err, "get signer from db")
	}
	amount := big.NewInt(0)
	amount.SetInt64(*account.Balance)
	if err := c.contract.DepositFund(ctx, providerAddress, *amount, signer); err != nil {
		return errors.Wrap(err, "deposit fund in contract")
	}

	return errors.Wrapf(c.SyncProviderAccount(ctx, providerAddress), "update charged account in db")
}

func (c *Ctrl) updateSigner(ctx context.Context, providerAddress common.Address, pubKey [2]string) error {
	signer, err := c.parseBigIntStringKey(pubKey)
	if err != nil {
		return errors.Wrap(err, "parse signer")
	}
	amount := big.NewInt(0)
	if err := c.contract.DepositFund(ctx, providerAddress, *amount, signer); err != nil {
		return errors.Wrap(err, "update signer by calling depositFund in contract")
	}
	err = c.db.UpdateProviderAccount(providerAddress.String(), model.Provider{
		Provider: providerAddress.String(),
		Signer:   model.StringSlice{pubKey[0], pubKey[1]},
	})
	return errors.Wrap(err, "update signer in db")
}

func (c Ctrl) SyncProviderAccounts(ctx context.Context) error {
	accounts, err := c.listProviderAccountFromContract(ctx)
	if err != nil {
		return err
	}
	refunds := []model.Refund{}
	for i := range accounts {
		refunds = append(refunds, accounts[i].Refunds...)
	}

	if err := c.db.BatchUpdateProviderAccount(accounts); err != nil {
		return err
	}

	return c.db.BatchUpdateRefund(refunds)
}

func (c Ctrl) SyncProviderAccount(ctx context.Context, providerAddress common.Address) error {
	account, err := c.getProviderAccountFromContract(ctx, providerAddress)
	if err != nil {
		return err
	}
	if err := c.db.UpdateProviderAccount(account.Provider, account); err != nil {
		return err
	}

	return c.db.BatchUpdateRefund(account.Refunds)
}

// getProviderAccountFromContract get account information from contract
func (c Ctrl) getProviderAccountFromContract(ctx context.Context, providerAddress common.Address) (model.Provider, error) {
	account, err := c.contract.GetProviderAccount(ctx, providerAddress)
	if err != nil {
		return model.Provider{}, errors.Wrap(err, "get account from contract")
	}
	ret := parseAccount(account, false)
	return ret, nil
}

// listProviderAccountFromContract get account information from contract
func (c Ctrl) listProviderAccountFromContract(ctx context.Context) ([]model.Provider, error) {
	accounts, err := c.contract.ListProviderAccount(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "list account from contract")
	}
	list := make([]model.Provider, len(accounts))
	for i, account := range accounts {
		list[i] = parseAccount(account, false)
	}
	return list, nil
}

func (c Ctrl) backfillProviderAccount(accounts []contract.Account) ([]model.Provider, error) {
	list := make([]model.Provider, len(accounts))
	dbAccounts, err := c.db.ListProviderAccount()
	if err != nil {
		return nil, errors.Wrap(err, "list account from db")
	}
	accountMap := make(map[string]model.Provider, len(dbAccounts))
	for i, account := range dbAccounts {
		accountMap[account.Provider] = dbAccounts[i]
	}
	for i, account := range accounts {
		list[i] = parseAccount(account, true)
		if v, ok := accountMap[account.Provider.String()]; ok {
			list[i].LastResponseTokenCount = v.LastResponseTokenCount
		}
	}
	return list, nil
}

func parseAccount(account contract.Account, ignoreProcessedRefund bool) model.Provider {
	refunds := []model.Refund{}
	for _, refund := range account.Refunds {
		if ignoreProcessedRefund && refund.Processed {
			continue
		}
		refunds = append(refunds, model.Refund{
			Provider:  account.Provider.String(),
			Index:     model.PtrOf(refund.Index.Int64()),
			CreatedAt: model.PtrOf(time.Unix(refund.CreatedAt.Int64(), 0)),
			Amount:    model.PtrOf(refund.Amount.Int64()),
			Processed: &refund.Processed,
		})
	}
	return model.Provider{
		Provider:      account.Provider.String(),
		Balance:       model.PtrOf(account.Balance.Int64()),
		PendingRefund: model.PtrOf(account.PendingRefund.Int64()),
		Refunds:       refunds,
		Signer:        []string{account.Signer[0].String(), account.Signer[1].String()},
		Nonce:         model.PtrOf(account.Nonce.Int64()),
	}
}
