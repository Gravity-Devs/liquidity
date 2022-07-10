package keeper

import (
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/gravity-devs/liquidity/v2/x/liquidity/types"
)

// RegisterInvariants registers all liquidity invariants.
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "escrow-amount",
		LiquidityPoolsEscrowAmountInvariant(k))
}

// AllInvariants runs all invariants of the liquidity module.
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		res, stop := LiquidityPoolsEscrowAmountInvariant(k)(ctx)
		return res, stop
	}
}

// LiquidityPoolsEscrowAmountInvariant checks that outstanding unwithdrawn fees are never negative.
func LiquidityPoolsEscrowAmountInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		remainingCoins := sdk.NewCoins()
		batches := k.GetAllPoolBatches(ctx)
		for _, batch := range batches {
			swapMsgs := k.GetAllPoolBatchSwapMsgStatesNotToBeDeleted(ctx, batch)
			for _, msg := range swapMsgs {
				remainingCoins = remainingCoins.Add(msg.RemainingOfferCoin)
			}
			depositMsgs := k.GetAllPoolBatchDepositMsgStatesNotToBeDeleted(ctx, batch)
			for _, msg := range depositMsgs {
				remainingCoins = remainingCoins.Add(msg.Msg.DepositCoins...)
			}
			withdrawMsgs := k.GetAllPoolBatchWithdrawMsgStatesNotToBeDeleted(ctx, batch)
			for _, msg := range withdrawMsgs {
				remainingCoins = remainingCoins.Add(msg.Msg.PoolCoin)
			}
		}

		batchEscrowAcc := k.accountKeeper.GetModuleAddress(types.ModuleName)
		escrowAmt := k.bankKeeper.GetAllBalances(ctx, batchEscrowAcc)

		broken := !escrowAmt.IsAllGTE(remainingCoins)

		return sdk.FormatInvariant(types.ModuleName, "batch escrow amount invariant broken",
			"batch escrow amount LT batch remaining amount"), broken
	}
}

// These invariants cannot be registered via RegisterInvariants since the module uses per-block batch execution.
// We should approach adding these invariant checks inside actual logics of deposit / withdraw / swap.

var (
	BatchLogicInvariantCheckFlag = false // It is only used at the development stage, and is disabled at the product level.
	// For coin amounts less than coinAmountThreshold, a high errorRate does not mean
	// that the calculation logic has errors.
	// For example, if there were two X coins and three Y coins in the pool, and someone deposits
	// one X coin and one Y coin, it's an acceptable input.
	// But pool price would change from 2/3 to 3/4 so errorRate will report 1/8(=0.125),
	// meaning that the price has changed by 12.5%.
	// This happens with small coin amounts, so there should be a threshold for coin amounts
	// before we calculate the errorRate.
	errorRateThreshold, _ = math.NewIntFromString(".05") // 5%
	coinAmountThreshold   = math.NewInt(20)              // If a decimal error occurs at a value less than 20, the error rate is over 5%.
)

func errorRate(expected, actual math.Int) math.Int {
	// To prevent divide-by-zero panics, return 1.0(=100%) as the error rate
	// when the expected value is 0.
	if expected.IsZero() {
		return math.OneInt()
	}
	return actual.Sub(expected).Quo(expected).Abs()
}

// MintingPoolCoinsInvariant checks the correct ratio of minting amount of pool coins.
func MintingPoolCoinsInvariant(poolCoinTotalSupply, mintPoolCoin, depositCoinA, depositCoinB, lastReserveCoinA, lastReserveCoinB, refundedCoinA, refundedCoinB math.Int) {
	if !refundedCoinA.IsZero() {
		depositCoinA = depositCoinA.Sub(refundedCoinA)
	}

	if !refundedCoinB.IsZero() {
		depositCoinB = depositCoinB.Sub(refundedCoinB)
	}

	poolCoinRatio := mintPoolCoin.Quo(poolCoinTotalSupply)
	depositCoinARatio := depositCoinA.Quo(lastReserveCoinA)
	depositCoinBRatio := depositCoinB.Quo(lastReserveCoinB)
	expectedMintPoolCoinAmtBasedA := depositCoinARatio.Mul(poolCoinTotalSupply)
	expectedMintPoolCoinAmtBasedB := depositCoinBRatio.Mul(poolCoinTotalSupply)

	// NewPoolCoinAmount / LastPoolCoinSupply == AfterRefundedDepositCoinA / LastReserveCoinA
	// NewPoolCoinAmount / LastPoolCoinSupply == AfterRefundedDepositCoinA / LastReserveCoinB
	if depositCoinA.GTE(coinAmountThreshold) && depositCoinB.GTE(coinAmountThreshold) &&
		lastReserveCoinA.GTE(coinAmountThreshold) && lastReserveCoinB.GTE(coinAmountThreshold) &&
		mintPoolCoin.GTE(coinAmountThreshold) && poolCoinTotalSupply.GTE(coinAmountThreshold) {
		if errorRate(depositCoinARatio, poolCoinRatio).GT(errorRateThreshold) ||
			errorRate(depositCoinBRatio, poolCoinRatio).GT(errorRateThreshold) {
			panic("invariant check fails due to incorrect ratio of pool coins")
		}
	}

	if mintPoolCoin.GTE(coinAmountThreshold) &&
		(sdk.MaxInt(mintPoolCoin, expectedMintPoolCoinAmtBasedA).Sub(sdk.MinInt(mintPoolCoin, expectedMintPoolCoinAmtBasedA))).Quo(mintPoolCoin).GT(errorRateThreshold) ||
		sdk.MaxInt(mintPoolCoin, expectedMintPoolCoinAmtBasedB).Sub(sdk.MinInt(mintPoolCoin, expectedMintPoolCoinAmtBasedA)).Quo(mintPoolCoin).GT(errorRateThreshold) {
		panic("invariant check fails due to incorrect amount of pool coins")
	}
}

// DepositInvariant checks after deposit amounts.
func DepositInvariant(lastReserveCoinA, lastReserveCoinB, depositCoinA, depositCoinB, afterReserveCoinA, afterReserveCoinB, refundedCoinA, refundedCoinB math.Int) {
	depositCoinA = depositCoinA.Sub(refundedCoinA)
	depositCoinB = depositCoinB.Sub(refundedCoinB)

	depositCoinRatio := depositCoinA.Quo(depositCoinB)
	lastReserveRatio := lastReserveCoinA.Quo(lastReserveCoinB)
	afterReserveRatio := afterReserveCoinA.Quo(afterReserveCoinB)

	// AfterDepositReserveCoinA = LastReserveCoinA + AfterRefundedDepositCoinA
	// AfterDepositReserveCoinB = LastReserveCoinB + AfterRefundedDepositCoinA
	if !afterReserveCoinA.Equal(lastReserveCoinA.Add(depositCoinA)) ||
		!afterReserveCoinB.Equal(lastReserveCoinB.Add(depositCoinB)) {
		panic("invariant check fails due to incorrect deposit amounts")
	}

	if depositCoinA.GTE(coinAmountThreshold) && depositCoinB.GTE(coinAmountThreshold) &&
		lastReserveCoinA.GTE(coinAmountThreshold) && lastReserveCoinB.GTE(coinAmountThreshold) {
		// AfterRefundedDepositCoinA / AfterRefundedDepositCoinA = LastReserveCoinA / LastReserveCoinB
		if errorRate(lastReserveRatio, depositCoinRatio).GT(errorRateThreshold) {
			panic("invariant check fails due to incorrect deposit ratio")
		}
		// LastReserveCoinA / LastReserveCoinB = AfterDepositReserveCoinA / AfterDepositReserveCoinB
		if errorRate(lastReserveRatio, afterReserveRatio).GT(errorRateThreshold) {
			panic("invariant check fails due to incorrect pool price ratio")
		}
	}
}

// BurningPoolCoinsInvariant checks the correct burning amount of pool coins.
func BurningPoolCoinsInvariant(burnedPoolCoin, withdrawCoinA, withdrawCoinB, reserveCoinA, reserveCoinB, lastPoolCoinSupply math.Int, withdrawFeeCoins sdk.Coins) {
	burningPoolCoinRatio := burnedPoolCoin.Quo(lastPoolCoinSupply)
	if burningPoolCoinRatio.Equal(sdk.OneInt()) {
		return
	}

	withdrawCoinARatio := withdrawCoinA.Add(withdrawFeeCoins[0].Amount).Quo(reserveCoinA)
	withdrawCoinBRatio := withdrawCoinB.Add(withdrawFeeCoins[1].Amount).Quo(reserveCoinB)

	// BurnedPoolCoinAmount / LastPoolCoinSupply >= (WithdrawCoinA+WithdrawFeeCoinA) / LastReserveCoinA
	// BurnedPoolCoinAmount / LastPoolCoinSupply >= (WithdrawCoinB+WithdrawFeeCoinB) / LastReserveCoinB
	if withdrawCoinARatio.GT(burningPoolCoinRatio) || withdrawCoinBRatio.GT(burningPoolCoinRatio) {
		panic("invariant check fails due to incorrect ratio of burning pool coins")
	}

	expectedBurningPoolCoinBasedA := lastPoolCoinSupply.Mul(withdrawCoinARatio)
	expectedBurningPoolCoinBasedB := lastPoolCoinSupply.Mul(withdrawCoinBRatio)

	if burnedPoolCoin.GTE(coinAmountThreshold) &&
		(sdk.MaxInt(burnedPoolCoin, expectedBurningPoolCoinBasedA).Sub(sdk.MinInt(burnedPoolCoin, expectedBurningPoolCoinBasedA))).Quo(burnedPoolCoin).GT(errorRateThreshold) ||
		sdk.MaxInt(burnedPoolCoin, expectedBurningPoolCoinBasedB).Sub(sdk.MinInt(burnedPoolCoin, expectedBurningPoolCoinBasedB)).Quo(burnedPoolCoin).GT(errorRateThreshold) {
		panic("invariant check fails due to incorrect amount of burning pool coins")
	}
}

// WithdrawReserveCoinsInvariant checks the after withdraw amounts.
func WithdrawReserveCoinsInvariant(withdrawCoinA, withdrawCoinB, reserveCoinA, reserveCoinB,
	afterReserveCoinA, afterReserveCoinB, afterPoolCoinTotalSupply, lastPoolCoinSupply, burnedPoolCoin math.Int) {
	// AfterWithdrawReserveCoinA = LastReserveCoinA - WithdrawCoinA
	if !afterReserveCoinA.Equal(reserveCoinA.Sub(withdrawCoinA)) {
		panic("invariant check fails due to incorrect withdraw coin A amount")
	}

	// AfterWithdrawReserveCoinB = LastReserveCoinB - WithdrawCoinB
	if !afterReserveCoinB.Equal(reserveCoinB.Sub(withdrawCoinB)) {
		panic("invariant check fails due to incorrect withdraw coin B amount")
	}

	// AfterWithdrawPoolCoinSupply = LastPoolCoinSupply - BurnedPoolCoinAmount
	if !afterPoolCoinTotalSupply.Equal(lastPoolCoinSupply.Sub(burnedPoolCoin)) {
		panic("invariant check fails due to incorrect total supply")
	}
}

// WithdrawAmountInvariant checks the correct ratio of withdraw coin amounts.
func WithdrawAmountInvariant(withdrawCoinA, withdrawCoinB, reserveCoinA, reserveCoinB, burnedPoolCoin, poolCoinSupply math.Int, withdrawFeeRate math.Int) {
	ratio := burnedPoolCoin.Quo(poolCoinSupply).Mul(math.OneInt().Sub(withdrawFeeRate))
	idealWithdrawCoinA := reserveCoinA.Mul(ratio)
	idealWithdrawCoinB := reserveCoinB.Mul(ratio)
	diffA := idealWithdrawCoinA.Sub(withdrawCoinA).Abs()
	diffB := idealWithdrawCoinB.Sub(withdrawCoinB).Abs()
	if !burnedPoolCoin.Equal(poolCoinSupply) {
		if diffA.GTE(math.OneInt()) {
			panic(fmt.Sprintf("withdraw coin amount %v differs too much from %v", withdrawCoinA, idealWithdrawCoinA))
		}
		if diffB.GTE(math.OneInt()) {
			panic(fmt.Sprintf("withdraw coin amount %v differs too much from %v", withdrawCoinB, idealWithdrawCoinB))
		}
	}
}

// ImmutablePoolPriceAfterWithdrawInvariant checks the immutable pool price after withdrawing coins.
func ImmutablePoolPriceAfterWithdrawInvariant(reserveCoinA, reserveCoinB, withdrawCoinA, withdrawCoinB, afterReserveCoinA, afterReserveCoinB math.Int) {
	// TestReinitializePool tests a scenario where after reserve coins are zero
	if !afterReserveCoinA.IsZero() && !afterReserveCoinB.IsZero() {
		reserveCoinA = reserveCoinA.Sub(withdrawCoinA)
		reserveCoinB = reserveCoinB.Sub(withdrawCoinB)

		reserveCoinRatio := reserveCoinA.Quo(reserveCoinB)
		afterReserveCoinRatio := afterReserveCoinA.Quo(afterReserveCoinB)

		// LastReserveCoinA / LastReserveCoinB = AfterWithdrawReserveCoinA / AfterWithdrawReserveCoinB
		if reserveCoinA.GTE(coinAmountThreshold) && reserveCoinB.GTE(coinAmountThreshold) &&
			withdrawCoinA.GTE(coinAmountThreshold) && withdrawCoinB.GTE(coinAmountThreshold) &&
			errorRate(reserveCoinRatio, afterReserveCoinRatio).GT(errorRateThreshold) {
			panic("invariant check fails due to incorrect pool price ratio")
		}
	}
}
