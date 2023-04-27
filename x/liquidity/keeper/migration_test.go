package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/gravity-devs/liquidity/app"
	"github.com/gravity-devs/liquidity/x/liquidity"
	"github.com/gravity-devs/liquidity/x/liquidity/keeper"
	"github.com/gravity-devs/liquidity/x/liquidity/types"
)

func TestMigration(t *testing.T) {
	simapp, ctx := createTestInput()
	simapp.LiquidityKeeper.SetParams(ctx, types.DefaultParams())
	params := simapp.LiquidityKeeper.GetParams(ctx)

	keeper.BatchLogicInvariantCheckFlag = false

	// define test denom X, Y for Liquidity Pool
	denomX := "denomX"
	denomY := "denomY"
	denomX, denomY = types.AlphabeticalDenomPair(denomX, denomY)

	X, Y := sdk.NewInt(1000000), sdk.NewInt(1000000)
	deposit := sdk.NewCoins(sdk.NewCoin(denomX, X), sdk.NewCoin(denomY, Y))
	depositAdd := sdk.NewCoins(sdk.NewCoin(denomX, X.MulRaw(1000000000000000000)), sdk.NewCoin(denomY, Y.MulRaw(1000000000000000000)))

	poolCreator, _ := app.GenAccount(simapp, ctx, 1, true, deposit.Add(params.PoolCreationFee...))
	depositer, _ := app.GenAccount(simapp, ctx, 2, true, depositAdd)
	smallHolder, _ := app.GenAccount(simapp, ctx, 2, true, nil)
	smallHolder2, _ := app.GenAccount(simapp, ctx, 3, true, nil)
	smallHolder3, _ := app.GenAccount(simapp, ctx, 4, true, nil)
	derivedAcc, _ := app.GenAccount(simapp, ctx, 0, false, nil)

	depositA := simapp.BankKeeper.GetBalance(ctx, poolCreator, denomX)
	depositB := simapp.BankKeeper.GetBalance(ctx, poolCreator, denomY)
	depositBalance := sdk.NewCoins(depositA, depositB)
	require.Equal(t, deposit, depositBalance)

	// create Liquidity pool
	poolTypeID := types.DefaultPoolTypeID
	msg := types.NewMsgCreatePool(poolCreator, poolTypeID, depositBalance)
	_, err := simapp.LiquidityKeeper.CreatePool(ctx, msg)
	require.NoError(t, err)

	// verify created liquidity pool
	pools := simapp.LiquidityKeeper.GetAllPools(ctx)
	pool := pools[0]
	poolID := pool.Id
	require.Equal(t, 1, len(pools))
	require.Equal(t, uint64(1), poolID)
	require.Equal(t, denomX, pool.ReserveCoinDenoms[0])
	require.Equal(t, denomY, pool.ReserveCoinDenoms[1])

	// verify minted pool coin
	poolCoin := simapp.LiquidityKeeper.GetPoolCoinTotalSupply(ctx, pool)
	creatorBalance := simapp.BankKeeper.GetBalance(ctx, poolCreator, pool.PoolCoinDenom)
	require.Equal(t, poolCoin, creatorBalance.Amount)

	depositMsg := types.NewMsgDepositWithinBatch(depositer, 1, depositAdd)
	_, err = simapp.LiquidityKeeper.DepositWithinBatch(ctx, depositMsg)
	require.NoError(t, err)

	poolBatch, found := simapp.LiquidityKeeper.GetPoolBatch(ctx, depositMsg.PoolId)
	require.True(t, found)
	msgs := simapp.LiquidityKeeper.GetAllPoolBatchDepositMsgs(ctx, poolBatch)
	require.Equal(t, 1, len(msgs))

	err = simapp.LiquidityKeeper.ExecuteDeposit(ctx, msgs[0], poolBatch)
	require.NoError(t, err)

	liquidity.EndBlocker(ctx, simapp.LiquidityKeeper)
	liquidity.BeginBlocker(ctx, simapp.LiquidityKeeper)

	sendCoin := creatorBalance
	sendCoin.Amount = sendCoin.Amount.QuoRaw(2)
	simapp.BankKeeper.SendCoins(ctx, poolCreator, derivedAcc, sdk.Coins{sendCoin})
	simapp.BankKeeper.SendCoins(ctx, poolCreator, smallHolder, sdk.Coins{sdk.NewCoin(pool.PoolCoinDenom, sdk.NewInt(100))})
	simapp.BankKeeper.SendCoins(ctx, poolCreator, smallHolder2, sdk.Coins{sdk.NewCoin(pool.PoolCoinDenom, sdk.NewInt(99))})
	simapp.BankKeeper.SendCoins(ctx, poolCreator, smallHolder3, sdk.Coins{sdk.NewCoin(pool.PoolCoinDenom, sdk.NewInt(1))})

	err = keeper.SafeForceWithdrawal(ctx, simapp.LiquidityKeeper, simapp.BankKeeper, simapp.AccountKeeper)
	require.NoError(t, err)

	require.Equal(t, "100denomX,100denomY", simapp.BankKeeper.GetAllBalances(ctx, smallHolder).String())
	require.Equal(t, "99denomX,99denomY", simapp.BankKeeper.GetAllBalances(ctx, smallHolder2).String())
	require.Equal(t, "1denomX,1denomY", simapp.BankKeeper.GetAllBalances(ctx, smallHolder3).String())
	require.Equal(t, "499800denomX,499800denomY", simapp.BankKeeper.GetAllBalances(ctx, poolCreator).String())
	require.Equal(t, "1000000000000000000000000denomX,1000000000000000000000000denomY", simapp.BankKeeper.GetAllBalances(ctx, depositer).String())
	require.Equal(t, "500000poolD35A0CC16EE598F90B044CE296A405BA9C381E38837599D96F2F70C2F02A23A4", simapp.BankKeeper.GetAllBalances(ctx, derivedAcc).String())

	// fund remaining reserve
	reserveAcc, err := sdk.AccAddressFromBech32(pool.ReserveAccountAddress)
	require.NoError(t, err)
	err = simapp.DistrKeeper.FundCommunityPool(ctx, simapp.BankKeeper.GetAllBalances(ctx, reserveAcc), reserveAcc)
	require.NoError(t, err)

	withdrawMsg := types.NewMsgWithdrawWithinBatch(derivedAcc, 1, sdk.NewCoin(pool.PoolCoinDenom, sdk.NewInt(500000)))
	_, err = simapp.LiquidityKeeper.WithdrawWithinBatch(ctx, withdrawMsg)
	require.ErrorIs(t, err, types.ErrDepletedPool)

	require.Equal(t, "500000poolD35A0CC16EE598F90B044CE296A405BA9C381E38837599D96F2F70C2F02A23A4", simapp.BankKeeper.GetAllBalances(ctx, derivedAcc).String())

	err = keeper.SafeForceWithdrawal(ctx, simapp.LiquidityKeeper, simapp.BankKeeper, simapp.AccountKeeper)
	require.NoError(t, err)

	// TODO: add set invariant, and panic case
	// TODO: test coverage case
}

func TestMigrationFailCase(t *testing.T) {
	simapp, ctx := createTestInput()
	simapp.LiquidityKeeper.SetParams(ctx, types.DefaultParams())
	params := simapp.LiquidityKeeper.GetParams(ctx)

	// on invariant checking flag for panic case
	keeper.BatchLogicInvariantCheckFlag = true

	// define test denom X, Y for Liquidity Pool
	denomX := "denomX"
	denomY := "denomY"
	denomX, denomY = types.AlphabeticalDenomPair(denomX, denomY)

	X, Y := sdk.NewInt(1000000), sdk.NewInt(1000000)
	deposit := sdk.NewCoins(sdk.NewCoin(denomX, X), sdk.NewCoin(denomY, Y))
	// attempt decimal error case
	depositAdd := sdk.NewCoins(sdk.NewCoin(denomX, X.MulRaw(1000000000000000000).MulRaw(100000000000)), sdk.NewCoin(denomY, Y.MulRaw(1000000000000000000).MulRaw(100000000000)))

	poolCreator, _ := app.GenAccount(simapp, ctx, 1, true, deposit.Add(params.PoolCreationFee...))
	depositer, _ := app.GenAccount(simapp, ctx, 2, true, depositAdd)
	smallHolder, _ := app.GenAccount(simapp, ctx, 2, true, nil)
	smallHolder2, _ := app.GenAccount(simapp, ctx, 3, true, nil)
	smallHolder3, _ := app.GenAccount(simapp, ctx, 4, true, nil)
	derivedAcc, _ := app.GenAccount(simapp, ctx, 0, false, nil)

	depositA := simapp.BankKeeper.GetBalance(ctx, poolCreator, denomX)
	depositB := simapp.BankKeeper.GetBalance(ctx, poolCreator, denomY)
	depositBalance := sdk.NewCoins(depositA, depositB)
	require.Equal(t, deposit, depositBalance)

	// create Liquidity pool
	poolTypeID := types.DefaultPoolTypeID
	msg := types.NewMsgCreatePool(poolCreator, poolTypeID, depositBalance)
	_, err := simapp.LiquidityKeeper.CreatePool(ctx, msg)
	require.NoError(t, err)

	// verify created liquidity pool
	pools := simapp.LiquidityKeeper.GetAllPools(ctx)
	pool := pools[0]
	poolID := pool.Id
	require.Equal(t, 1, len(pools))
	require.Equal(t, uint64(1), poolID)
	require.Equal(t, denomX, pool.ReserveCoinDenoms[0])
	require.Equal(t, denomY, pool.ReserveCoinDenoms[1])

	// verify minted pool coin
	poolCoin := simapp.LiquidityKeeper.GetPoolCoinTotalSupply(ctx, pool)
	creatorBalance := simapp.BankKeeper.GetBalance(ctx, poolCreator, pool.PoolCoinDenom)
	require.Equal(t, poolCoin, creatorBalance.Amount)

	depositMsg := types.NewMsgDepositWithinBatch(depositer, 1, depositAdd)
	_, err = simapp.LiquidityKeeper.DepositWithinBatch(ctx, depositMsg)
	require.NoError(t, err)

	poolBatch, found := simapp.LiquidityKeeper.GetPoolBatch(ctx, depositMsg.PoolId)
	require.True(t, found)
	msgs := simapp.LiquidityKeeper.GetAllPoolBatchDepositMsgs(ctx, poolBatch)
	require.Equal(t, 1, len(msgs))

	err = simapp.LiquidityKeeper.ExecuteDeposit(ctx, msgs[0], poolBatch)
	require.NoError(t, err)

	liquidity.EndBlocker(ctx, simapp.LiquidityKeeper)
	liquidity.BeginBlocker(ctx, simapp.LiquidityKeeper)

	sendCoin := creatorBalance
	sendCoin.Amount = sendCoin.Amount.QuoRaw(2)
	simapp.BankKeeper.SendCoins(ctx, poolCreator, derivedAcc, sdk.Coins{sendCoin})
	simapp.BankKeeper.SendCoins(ctx, poolCreator, smallHolder, sdk.Coins{sdk.NewCoin(pool.PoolCoinDenom, sdk.NewInt(100))})
	simapp.BankKeeper.SendCoins(ctx, poolCreator, smallHolder2, sdk.Coins{sdk.NewCoin(pool.PoolCoinDenom, sdk.NewInt(99))})
	simapp.BankKeeper.SendCoins(ctx, poolCreator, smallHolder3, sdk.Coins{sdk.NewCoin(pool.PoolCoinDenom, sdk.NewInt(1))})

	// panic recovered, reverted by cached ctx
	err = keeper.SafeForceWithdrawal(ctx, simapp.LiquidityKeeper, simapp.BankKeeper, simapp.AccountKeeper)
	require.NoError(t, err)

	require.Equal(t, "100poolD35A0CC16EE598F90B044CE296A405BA9C381E38837599D96F2F70C2F02A23A4", simapp.BankKeeper.GetAllBalances(ctx, smallHolder).String())
	require.Equal(t, "99poolD35A0CC16EE598F90B044CE296A405BA9C381E38837599D96F2F70C2F02A23A4", simapp.BankKeeper.GetAllBalances(ctx, smallHolder2).String())
	require.Equal(t, "1poolD35A0CC16EE598F90B044CE296A405BA9C381E38837599D96F2F70C2F02A23A4", simapp.BankKeeper.GetAllBalances(ctx, smallHolder3).String())
	require.Equal(t, "499800poolD35A0CC16EE598F90B044CE296A405BA9C381E38837599D96F2F70C2F02A23A4", simapp.BankKeeper.GetAllBalances(ctx, poolCreator).String())
	require.Equal(t, "100000000000000000000000000000000000poolD35A0CC16EE598F90B044CE296A405BA9C381E38837599D96F2F70C2F02A23A4", simapp.BankKeeper.GetAllBalances(ctx, depositer).String())
	require.Equal(t, "500000poolD35A0CC16EE598F90B044CE296A405BA9C381E38837599D96F2F70C2F02A23A4", simapp.BankKeeper.GetAllBalances(ctx, derivedAcc).String())
}
