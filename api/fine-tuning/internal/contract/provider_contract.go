package providercontract

import (
	"context"
	"os"
	"time"

	"github.com/0glabs/0g-serving-broker/common/log"
	"github.com/0glabs/0g-serving-broker/fine-tuning/config"
	"github.com/0glabs/0g-serving-broker/fine-tuning/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type ProviderContract struct {
	Contract        *contract.ServingContract
	ProviderAddress string
	LockTime        time.Duration
	logger          log.Logger
}

func NewProviderContract(conf *config.Config, logger log.Logger) (*ProviderContract, error) {
	contract, err := contract.NewServingContract(common.HexToAddress(conf.ContractAddress), &conf.Networks, os.Getenv("NETWORK"), conf.GasPrice, conf.MaxGasPrice, logger)
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
	wallets, err := contract.Client.Network.Wallets()
	if err != nil {
		return nil, err
	}
	return &ProviderContract{
		Contract:        contract,
		ProviderAddress: wallets.Default().Address(),
		LockTime:        time.Duration(lockTime.Int64()) * time.Second,
		logger:          logger,
	}, nil
}

func (u *ProviderContract) Close() {
	u.Contract.Close()
}
