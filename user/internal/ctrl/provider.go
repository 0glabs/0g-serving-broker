package ctrl

import (
	"context"
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
	if err := c.contract.CreateProviderAccount(ctx, providerAddress, *balance); err != nil {
		return errors.Wrap(err, "create provider account in contract")
	}

	err := c.db.CreateProviderAccounts([]model.Provider{account})
	if err != nil {
		rollBackErr := c.SyncProviderAccount(ctx, providerAddress)
		if rollBackErr != nil {
			log.Printf("resync account in db: %s", rollBackErr.Error())
		}
	}
	return errors.Wrap(err, "create provider account in db")
}

func (c Ctrl) GetProviderAccount(ctx context.Context, providerAddress common.Address, mergeDB bool) (model.Provider, error) {
	account, err := c.contract.GetProviderAccount(ctx, providerAddress)
	if err != nil {
		return model.Provider{}, errors.Wrap(err, "get account from contract")
	}
	ret := parse(account)
	if !mergeDB {
		return ret, nil
	}
	rets, err := c.backfillProviderAccount([]contract.Account{account})
	return rets[0], err
}

func (c Ctrl) ListProviderAccount(ctx context.Context, mergeDB bool) ([]model.Provider, error) {
	accounts, err := c.contract.ListProviderAccount(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "list account from contract")
	}
	if mergeDB {
		return c.backfillProviderAccount(accounts)
	}
	list := make([]model.Provider, len(accounts))
	for i, account := range accounts {
		list[i] = parse(account)
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
		list[i] = parse(account)
		if v, ok := accountMap[account.Provider.String()]; ok {
			list[i].LastResponseTokenCount = v.LastResponseTokenCount
		}
	}
	return list, nil
}

func (c Ctrl) SyncProviderAccounts(ctx context.Context) error {
	accounts, err := c.ListProviderAccount(ctx, false)
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
	account, err := c.GetProviderAccount(ctx, providerAddress, false)
	if err != nil {
		return err
	}
	if err := c.db.UpdateProviderAccount(account.Provider, account); err != nil {
		return err
	}

	return c.db.BatchUpdateRefund(account.Refunds)
}

func parse(account contract.Account) model.Provider {
	refunds := make([]model.Refund, len(account.Refunds))
	for i, refund := range account.Refunds {
		refunds[i] = model.Refund{
			Provider:  account.Provider.String(),
			Index:     model.PtrOf(refund.Index.Int64()),
			CreatedAt: model.PtrOf(time.Unix(refund.CreatedAt.Int64(), 0)),
			Amount:    model.PtrOf(refund.Amount.Int64()),
			Processed: refund.Processed,
		}
	}
	return model.Provider{
		Provider:      account.Provider.String(),
		Balance:       model.PtrOf(account.Balance.Int64()),
		PendingRefund: model.PtrOf(account.PendingRefund.Int64()),
		Refunds:       refunds,
		Nonce:         model.PtrOf(account.Nonce.Int64()),
	}
}