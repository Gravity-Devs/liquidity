package keeper

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	v043 "github.com/gravity-devs/liquidity/x/liquidity/legacy/v043"
	liquiditytypes "github.com/gravity-devs/liquidity/x/liquidity/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper        Keeper
	bankKeeper    liquiditytypes.BankKeeper
	accountKeeper liquiditytypes.AccountKeeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper, bankKeeper liquiditytypes.BankKeeper, accountKeeper liquiditytypes.AccountKeeper) Migrator {
	return Migrator{keeper: keeper, bankKeeper: bankKeeper, accountKeeper: accountKeeper}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v043.MigrateStore(ctx, m.keeper.storeKey)
}

// Migrate2to3 migrates from version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	SafeForceWithdrawal(ctx, m.keeper, m.bankKeeper, m.accountKeeper)
	// Even if it fails, it is reverted, so return nil to prevent panic due to the force withdrawal failure
	return nil
}

// SafeForceWithdrawal call ForceWithdrawal safely by recover, cached ctx
func SafeForceWithdrawal(ctx sdk.Context, keeper Keeper, bankKeeper liquiditytypes.BankKeeper, accountKeeper liquiditytypes.AccountKeeper) (err error) {
	logger := keeper.Logger(ctx)
	defer func() {
		if r := recover(); r != nil {
			err = sdkerrors.Wrapf(sdkerrors.ErrPanic, "panic recovered on force withdrawal, %v", r)
			logger.Error(err.Error())
		}
	}()

	cachedCtx, writeCache := ctx.CacheContext()
	err = ForceWithdrawal(cachedCtx, keeper, bankKeeper, accountKeeper)
	if err == nil {
		writeCache()
	} else {
		logger.Error("error occurred on force withdrawal", "error", err)
	}
	return
}

// ForceWithdrawal Forcefully withdraw pool token holders once migration
func ForceWithdrawal(ctx sdk.Context, keeper Keeper, bankKeeper liquiditytypes.BankKeeper, accountKeeper liquiditytypes.AccountKeeper) error {
	logger := keeper.Logger(ctx)
	poolByPoolCoinDenom := map[string]liquiditytypes.Pool{} // PoolCoinDenom => Pool
	for _, pool := range keeper.GetAllPools(ctx) {
		poolByPoolCoinDenom[pool.PoolCoinDenom] = pool
	}

	// Iterate over all the balances of all accounts and find the accounts that hold poolxxx... coin.
	// Unless it is pool reserve account, forcefully withdraw their pool coin and transfer
	// the corresponding amount of respective reserve coins back to their accounts.
	// Lastly, burn pool coins that are withdrawn.
	accMap := map[string]authtypes.AccountI{}
	accList := []string{}
	errorCount := 0
	bankKeeper.IterateAllBalances(ctx, func(address sdk.AccAddress, coin sdk.Coin) (stop bool) {
		if strings.HasPrefix(coin.Denom, "pool") {
			addrStr := address.String()
			pool := poolByPoolCoinDenom[coin.Denom]
			if _, ok := accMap[addrStr]; !ok {
				accMap[addrStr] = accountKeeper.GetAccount(ctx, address)
				accList = append(accList, addrStr)
			}

			// Skip pool reserve accounts
			acc := accMap[addrStr]
			if acc.GetSequence() != 0 || acc.GetPubKey() != nil {
				if res, err := keeper.WithdrawWithinBatch(ctx, &liquiditytypes.MsgWithdrawWithinBatch{
					WithdrawerAddress: addrStr,
					PoolId:            pool.GetId(),
					PoolCoin:          coin,
				}); !res.Succeeded || err != nil {
					errorCount++
				}
			}
		}
		return false
	})

	// Execute batch manually
	keeper.ExecutePoolBatches(ctx)
	keeper.DeleteAndInitPoolBatches(ctx)

	// iterating and withdraw again if there is any pool coin left due to decimal error by too small pool coin value
	if errorCount > 0 {
		for _, address := range accList {
			acc := accMap[address]
			balances := keeper.bankKeeper.GetAllBalances(ctx, acc.GetAddress())
			for _, coin := range balances {
				if strings.HasPrefix(coin.Denom, "pool") {
					// Skip pool reserve accounts
					if acc.GetSequence() != 0 || acc.GetPubKey() != nil {
						if _, err := keeper.WithdrawWithinBatch(ctx, &liquiditytypes.MsgWithdrawWithinBatch{
							WithdrawerAddress: address,
							PoolId:            poolByPoolCoinDenom[coin.Denom].GetId(),
							PoolCoin:          coin,
						}); err != nil {
							logger.Debug(
								"failed force withdrawal",
								"withdrawer", address,
								"poolcoin", coin,
								"error", err,
							)
						}
					}
				}
			}
		}

		// Execute batch manually
		keeper.ExecutePoolBatches(ctx)
		keeper.DeleteAndInitPoolBatches(ctx)
	}

	if len(keeper.GetAllDepositMsgStates(ctx)) > 0 {
		return fmt.Errorf("deposit msg states must be empty for migration")
	}
	if len(keeper.GetAllSwapMsgStates(ctx)) > 0 {
		return fmt.Errorf("swap msg states must be empty for migration")
	}
	if len(keeper.GetAllWithdrawMsgStates(ctx)) > 0 {
		return fmt.Errorf("withdraw msg states must be empty for migration")
	}

	// Fund remaining reserve balance to community pool
	for _, pool := range keeper.GetAllPools(ctx) {
		reserveAcc := pool.GetReserveAccount()
		balances := keeper.bankKeeper.GetAllBalances(ctx, reserveAcc)
		if balances.IsZero() {
			continue
		}
		err := keeper.distrKeeper.FundCommunityPool(ctx, balances, reserveAcc)
		if err != nil {
			logger.Debug(
				"failed fund community pool",
				"pool id", pool.Id,
				"error", err,
			)
		}
	}

	// Delete pools and pool batches to remove this module
	for _, poolBatch := range keeper.GetAllPoolBatches(ctx) {
		keeper.DeletePoolBatch(ctx, poolBatch)
	}

	for _, pool := range keeper.GetAllPools(ctx) {
		keeper.DeletePool(ctx, pool)
	}
	return nil
}
