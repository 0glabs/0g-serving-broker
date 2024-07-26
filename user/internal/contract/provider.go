package usercontract

import (
	"context"
	"math/big"

	"github.com/0glabs/0g-serving-agent/common/contract"
	"github.com/0glabs/0g-serving-agent/common/errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (c *UserContract) CreateProviderAccount(ctx context.Context, provider common.Address, balance big.Int) error {
	account, _ := c.GetProviderAccount(ctx, provider)
	zeroAddress := common.Address{}
	if account.User != zeroAddress {
		return errors.New("account already exists")
	}
	return c.DepositFund(ctx, provider, balance)
}

func (c *UserContract) GetProviderAccount(ctx context.Context, provider common.Address) (contract.Account, error) {
	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}
	return c.contract.GetAccount(callOpts, common.HexToAddress(c.userAddress), provider)
}

func (c *UserContract) ListProviderAccount(ctx context.Context) ([]contract.Account, error) {
	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}
	accounts, err := c.contract.GetAllAccounts(callOpts)
	if err != nil {
		return nil, err
	}
	ret := []contract.Account{}
	for i := range accounts {
		if accounts[i].User.String() != c.userAddress {
			continue
		}
		ret = append(ret, accounts[i])
	}
	return ret, nil
}

func (c *UserContract) DepositFund(ctx context.Context, provider common.Address, balance big.Int) error {
	opts, err := c.contract.CreateTransactOpts()
	if err != nil {
		return err
	}

	opts.Value = &balance
	tx, err := c.contract.DepositFund(opts, provider)
	if err != nil {
		return err
	}
	_, err = c.contract.WaitForReceipt(ctx, tx.Hash())
	return err
}

func (c *UserContract) RequestRefund(ctx context.Context, provider common.Address, refund *big.Int) (*contract.ServingRefundRequested, error) {
	opts, err := c.contract.CreateTransactOpts()
	if err != nil {
		return nil, err
	}
	tx, err := c.contract.RequestRefund(opts, provider, refund)
	if err != nil {
		return nil, err
	}
	receipt, err := c.contract.WaitForReceipt(ctx, tx.Hash())
	if err != nil {
		return nil, err
	}

	return c.contract.Serving.ParseRefundRequested(*receipt.Logs[0])
}