package usercontract

import (
	"context"
	"os"
	"time"

	"github.com/0glabs/0g-serving-agent/common/config"
	"github.com/0glabs/0g-serving-agent/common/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type UserContract struct {
	Contract    *contract.ServingContract
	UserAddress string
	LockTime    time.Duration
}

func NewUserContract(conf *config.Config, userAddress string) (*UserContract, error) {
	contract, err := contract.NewServingContract(common.HexToAddress(conf.ContractAddress), conf, os.Getenv("NETWORK"))
	if err != nil {
		return nil, err
	}
	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}
	lockTime, err := contract.LockTime(callOpts)
	if err != nil {
		return nil, err
	}
	return &UserContract{
		Contract:    contract,
		UserAddress: userAddress,
		LockTime:    time.Duration(lockTime.Int64()) * time.Second,
	}, nil
}

func (u *UserContract) Close() {
	u.Contract.Close()
}