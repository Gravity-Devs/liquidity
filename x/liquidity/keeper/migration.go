package keeper

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
	return SafeForceWithdrawal(ctx, m.keeper, m.bankKeeper, m.accountKeeper)
}

// SafeForceWithdrawal call ForceWithdrawal safely by recover, cached ctx
func SafeForceWithdrawal(ctx sdk.Context, keeper Keeper, bankKeeper liquiditytypes.BankKeeper, accountKeeper liquiditytypes.AccountKeeper) error {
	broken := false
	defer func() {
		if r := recover(); r != nil {
			broken = true
		}
	}()

	cachedCtx, writeCache := ctx.CacheContext()
	err := ForceWithdrawal(cachedCtx, keeper, bankKeeper, accountKeeper)
	if err == nil && !broken {
		writeCache()
	}
	return nil
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
	bankKeeper.IterateAllBalances(ctx, func(address sdk.AccAddress, coin sdk.Coin) (stop bool) {
		if strings.HasPrefix(coin.Denom, "pool") {
			pool := poolByPoolCoinDenom[coin.Denom]
			if _, ok := accMap[address.String()]; !ok {
				accMap[address.String()] = accountKeeper.GetAccount(ctx, address)
				accList = append(accList, address.String())
			}

			// Skip pool reserve accounts
			acc := accMap[address.String()]
			if acc.GetSequence() != 0 || acc.GetPubKey() != nil {
				if _, err := keeper.WithdrawWithinBatch(ctx, &liquiditytypes.MsgWithdrawWithinBatch{
					WithdrawerAddress: address.String(),
					PoolId:            pool.GetId(),
					PoolCoin:          coin,
				}); err != nil {
					logger.Debug(
						"failed force withdrawal",
						"withdrawer", address.String(),
						"poolcoin", coin,
						"error", err,
					)
				}
			}
		}
		return false
	})

	// Execute batch manually
	keeper.ExecutePoolBatches(ctx)

	// Delete all batches
	keeper.DeleteAndInitPoolBatches(ctx)

	// iterating and withdraw again if there is any pool coin left due to decimal error
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

	// Delete all batches
	keeper.DeleteAndInitPoolBatches(ctx)

	// TODO: Need to decide whether fund remaining reserve balance to community pool or not
	return nil
}
