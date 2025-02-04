package providercontract

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/0glabs/0g-serving-broker/common/errors"
	"github.com/0glabs/0g-serving-broker/common/util"
	"github.com/0glabs/0g-serving-broker/fine-tuning/config"
	"github.com/0glabs/0g-serving-broker/fine-tuning/contract"
)

func (c *ProviderContract) AddOrUpdateService(ctx context.Context, service config.Service, occupied bool) error {
	opts, err := c.Contract.CreateTransactOpts()
	if err != nil {
		return err
	}
	cpuCount, err := util.ConvertToBigInt(service.Quota.CpuCount)
	if err != nil {
		return errors.Wrap(err, "convert cpuCount")
	}
	memory, err := util.ConvertToBigInt(service.Quota.Memory)
	if err != nil {
		return errors.Wrap(err, "convert memory")
	}
	storage, err := util.ConvertToBigInt(service.Quota.Storage)
	if err != nil {
		return errors.Wrap(err, "convert storage")
	}
	gpuCount, err := util.ConvertToBigInt(service.Quota.GpuCount)
	if err != nil {
		return errors.Wrap(err, "convert gpuCount")
	}
	pricePerToken, err := util.ConvertToBigInt(service.PricePerToken)
	if err != nil {
		return errors.Wrap(err, "convert PricePerToken")
	}
	quota := contract.Quota{
		CpuCount:    cpuCount,
		NodeMemory:  memory,
		NodeStorage: storage,
		GpuType:     service.Quota.GpuType,
		GpuCount:    gpuCount,
	}
	tx, err := c.Contract.AddOrUpdateService(
		opts,
		service.ServingUrl,
		quota,
		pricePerToken,
		// TODO: replace by real provider signer address
		common.HexToAddress("0x111111"),
		occupied,
	)
	if err != nil {
		return err
	}
	_, err = c.Contract.WaitForReceipt(ctx, tx.Hash())

	return err
}

func (c *ProviderContract) DeleteService(ctx context.Context) error {
	opt, err := c.Contract.CreateTransactOpts()
	if err != nil {
		return err
	}

	tx, err := c.Contract.RemoveService(opt)
	if err != nil {
		return err
	}
	_, err = c.Contract.WaitForReceipt(ctx, tx.Hash())
	return err
}

func (c *ProviderContract) GetService(ctx context.Context) (*contract.Service, error) {
	callOpts := &bind.CallOpts{
		Context: ctx,
	}

	list, err := c.Contract.GetAllServices(callOpts)
	if err != nil {
		return nil, err
	}
	for i := range list {
		if list[i].Provider.String() == c.ProviderAddress {
			return &list[i], nil
		}
	}

	return nil, fmt.Errorf("service not found")
}

func (c *ProviderContract) SyncServices(ctx context.Context, new config.Service) error {
	old, err := c.GetService(ctx)
	if err != nil && err.Error() != "service not found" {
		return err
	}

	if old != nil && identicalService(*old, new) {
		return nil
	}

	if err := c.AddOrUpdateService(ctx, new, false); err != nil {
		return errors.Wrap(err, "add or update service in contract")
	}

	return nil
}

func (c *ProviderContract) AddDeliverable(ctx context.Context, user common.Address, modelRootHash []byte) error {
	opt, err := c.Contract.CreateTransactOpts()
	if err != nil {
		return err
	}

	tx, err := c.Contract.AddDeliverable(opt, user, modelRootHash)
	if err != nil {
		return err
	}
	_, err = c.Contract.WaitForReceipt(ctx, tx.Hash())

	// todo return deliver index?
	return err
}

func identicalService(old contract.Service, new config.Service) bool {
	if old.Url != new.ServingUrl {
		return false
	}
	if old.PricePerToken.Int64() != new.PricePerToken {
		return false
	}
	if old.Occupied {
		return false
	}
	if old.Quota.CpuCount.Int64() != new.Quota.CpuCount {
		return false
	}
	if old.Quota.NodeMemory.Int64() != new.Quota.Memory {
		return false
	}
	if old.Quota.GpuCount.Int64() != new.Quota.GpuCount {
		return false
	}
	if old.Quota.NodeStorage.Int64() != new.Quota.Storage {
		return false
	}
	if old.Quota.GpuType != new.Quota.GpuType {
		return false
	}
	return true
}
