package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

// TODO: remove redundant function
func randRange(r *rand.Rand, min, max int) sdk.Int {
	return sdk.NewInt(int64(r.Intn(max-min) + min))
}

// TODO: remove redundant function
func randFloats(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// TODO: remove redundant function
func GetRandomOrders(denomX, denomY string, X, Y sdk.Int, r *rand.Rand) (XtoY, YtoX []BatchPoolSwapMsg) {
	currentPrice := X.ToDec().Quo(Y.ToDec())

	XtoYnewSize := int(r.Int31n(20)) // 0~19
	YtoXnewSize := int(r.Int31n(20)) // 0~19

	for i := 0; i < XtoYnewSize; i++ {
		randFloats(0.1, 0.9)
		orderPrice := currentPrice.Mul(sdk.NewDecFromIntWithPrec(randRange(r, 991, 1009), 3))
		orderAmt := X.ToDec().Mul(sdk.NewDecFromIntWithPrec(randRange(r, 1, 100), 4))
		orderCoin := sdk.NewCoin(denomX, orderAmt.RoundInt())

		XtoY = append(XtoY, BatchPoolSwapMsg{
			Msg: &MsgSwap{
				OfferCoin:       orderCoin,
				DemandCoinDenom: denomY,
				OrderPrice:      orderPrice,
			},
		})
	}

	for i := 0; i < YtoXnewSize; i++ {
		orderPrice := currentPrice.Mul(sdk.NewDecFromIntWithPrec(randRange(r, 991, 1009), 3))
		orderAmt := Y.ToDec().Mul(sdk.NewDecFromIntWithPrec(randRange(r, 1, 100), 4))
		orderCoin := sdk.NewCoin(denomY, orderAmt.RoundInt())

		YtoX = append(YtoX, BatchPoolSwapMsg{
			Msg: &MsgSwap{
				OfferCoin:       orderCoin,
				DemandCoinDenom: denomX,
				OrderPrice:      orderPrice,
			},
		})
	}
	return
}

func TestGetOrderMap(t *testing.T) {
	//var msgs []BatchPoolSwapMsg
	X := sdk.NewInt(100000000)
	Y := sdk.NewInt(50000000)
	//currentYPriceOverX := X.Quo(Y)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	XtoY, YtoX := GetRandomOrders("denomX", "denomY", X, Y, r)
	fmt.Println(XtoY)
	fmt.Println(YtoX)
}

func TestOrderBookSort(t *testing.T) {
	orderMap := make(OrderMap)
	a, _ := sdk.NewDecFromStr("0.1")
	b, _ := sdk.NewDecFromStr("0.2")
	c, _ := sdk.NewDecFromStr("0.3")
	orderMap[a.String()] = OrderByPrice{
		OrderPrice:   a,
		BuyOrderAmt:  sdk.ZeroInt(),
		SellOrderAmt: sdk.ZeroInt(),
	}
	orderMap[b.String()] = OrderByPrice{
		OrderPrice:   b,
		BuyOrderAmt:  sdk.ZeroInt(),
		SellOrderAmt: sdk.ZeroInt(),
	}
	orderMap[c.String()] = OrderByPrice{
		OrderPrice:   c,
		BuyOrderAmt:  sdk.ZeroInt(),
		SellOrderAmt: sdk.ZeroInt(),
	}
	// make orderbook to sort orderMap
	orderBook := orderMap.SortOrderBook()
	fmt.Println(orderBook)

	res := orderBook.Less(0, 1)
	require.True(t, res)
	res = orderBook.Less(1, 2)
	require.True(t, res)
	res = orderBook.Less(2, 1)
	require.False(t, res)

	orderBook.Swap(1, 2)
	fmt.Println(orderBook)
	require.Equal(t, c, orderBook[1].OrderPrice)
	require.Equal(t, b, orderBook[2].OrderPrice)

	orderBook.Sort()
	fmt.Println(orderBook)
	require.Equal(t, a, orderBook[0].OrderPrice)
	require.Equal(t, b, orderBook[1].OrderPrice)
	require.Equal(t, c, orderBook[2].OrderPrice)

	orderBook.Reverse()
	fmt.Println(orderBook)
	require.Equal(t, a, orderBook[2].OrderPrice)
	require.Equal(t, b, orderBook[1].OrderPrice)
	require.Equal(t, c, orderBook[0].OrderPrice)

}

func TestMinMaxDec(t *testing.T) {
	a, _ := sdk.NewDecFromStr("0.1")
	b, _ := sdk.NewDecFromStr("0.2")
	c, _ := sdk.NewDecFromStr("0.3")

	require.Equal(t, a, MinDec(a, b))
	require.Equal(t, a, MinDec(a, c))
	require.Equal(t, b, MaxDec(a, b))
	require.Equal(t, c, MaxDec(a, c))
	require.Equal(t, a, MaxDec(a, a))
	require.Equal(t, a, MinDec(a, a))
}

func TestGetExecutableAmt(t *testing.T) {
	orderMap := make(OrderMap)
	a, _ := sdk.NewDecFromStr("0.1")
	b, _ := sdk.NewDecFromStr("0.2")
	c, _ := sdk.NewDecFromStr("0.3")
	orderMap[a.String()] = OrderByPrice{
		OrderPrice:   a,
		BuyOrderAmt:  sdk.ZeroInt(),
		SellOrderAmt: sdk.NewInt(30000000),
	}
	orderMap[b.String()] = OrderByPrice{
		OrderPrice:   b,
		BuyOrderAmt:  sdk.NewInt(90000000),
		SellOrderAmt: sdk.ZeroInt(),
	}
	orderMap[c.String()] = OrderByPrice{
		OrderPrice:   c,
		BuyOrderAmt:  sdk.NewInt(50000000),
		SellOrderAmt: sdk.ZeroInt(),
	}
	// make orderbook to sort orderMap
	orderBook := orderMap.SortOrderBook()

	executableBuyAmtX, executableSellAmtY := GetExecutableAmt(b, orderBook)
	require.Equal(t, sdk.NewInt(140000000), executableBuyAmtX)
	require.Equal(t, sdk.NewInt(30000000), executableSellAmtY)
}

// TODO: WIP
func TestGetPriceDirection(t *testing.T) {

	// decrease case
	orderMap := make(OrderMap)
	a, _ := sdk.NewDecFromStr("0.1")
	b, _ := sdk.NewDecFromStr("0.2")
	c, _ := sdk.NewDecFromStr("0.3")
	orderMap[a.String()] = OrderByPrice{
		OrderPrice:   a,
		BuyOrderAmt:  sdk.ZeroInt(),
		SellOrderAmt: sdk.NewInt(30000000),
	}
	orderMap[b.String()] = OrderByPrice{
		OrderPrice:   b,
		BuyOrderAmt:  sdk.NewInt(90000000),
		SellOrderAmt: sdk.ZeroInt(),
	}
	orderMap[c.String()] = OrderByPrice{
		OrderPrice:   c,
		BuyOrderAmt:  sdk.NewInt(50000000),
		SellOrderAmt: sdk.ZeroInt(),
	}
	// make orderbook to sort orderMap
	orderBook := orderMap.SortOrderBook()

	// increase
	X := sdk.NewInt(10000000).ToDec()
	Y := sdk.NewInt(50000000).ToDec()
	currentYPriceOverX := X.Quo(Y)
	require.Equal(t, currentYPriceOverX, b)
	result := GetPriceDirection(currentYPriceOverX, orderBook)
	require.Equal(t, Increase, result)

	// decrease
	X = sdk.NewInt(100000000).ToDec()
	Y = sdk.NewInt(50000000).ToDec()
	currentYPriceOverX = X.Quo(Y)
	result = GetPriceDirection(currentYPriceOverX, orderBook)
	require.Equal(t, Decrease, result)

	// TODO: stay case
}

// TODO: WIP
func TestComputePriceDirection(t *testing.T) {

	// decrease case
	orderMap := make(OrderMap)
	a, _ := sdk.NewDecFromStr("2.0")
	b, _ := sdk.NewDecFromStr("2.1")
	c, _ := sdk.NewDecFromStr("1.9")
	orderMap[a.String()] = OrderByPrice{
		OrderPrice:   a,
		BuyOrderAmt:  sdk.ZeroInt(),
		SellOrderAmt: sdk.NewInt(3000000),
	}
	orderMap[b.String()] = OrderByPrice{
		OrderPrice:   b,
		BuyOrderAmt:  sdk.NewInt(9000000),
		SellOrderAmt: sdk.ZeroInt(),
	}
	orderMap[c.String()] = OrderByPrice{
		OrderPrice:   c,
		BuyOrderAmt:  sdk.NewInt(5000000),
		SellOrderAmt: sdk.ZeroInt(),
	}
	// make orderbook to sort orderMap
	orderBook := orderMap.SortOrderBook()

	X := sdk.NewInt(100000000).ToDec()
	Y := sdk.NewInt(50000000).ToDec()
	currentYPriceOverX := X.Quo(Y)
	result := ComputePriceDirection(X, Y, currentYPriceOverX, orderBook)

	fmt.Println(X, Y, currentYPriceOverX)
	fmt.Println(result)

	// increase case
	orderMap[c.String()] = OrderByPrice{
		OrderPrice:   c,
		BuyOrderAmt:  sdk.ZeroInt(),
		SellOrderAmt: sdk.NewInt(1000000),
	}
	orderMap[b.String()] = OrderByPrice{
		OrderPrice:   b,
		BuyOrderAmt:  sdk.NewInt(4000000),
		SellOrderAmt: sdk.ZeroInt(),
	}
	orderMap[a.String()] = OrderByPrice{
		OrderPrice:   a,
		BuyOrderAmt:  sdk.NewInt(7000000),
		SellOrderAmt: sdk.ZeroInt(),
	}
	// make orderbook to sort orderMap
	orderBook = orderMap.SortOrderBook()

	X = sdk.NewInt(100000000).ToDec()
	Y = sdk.NewInt(50000000).ToDec()
	currentYPriceOverX = X.Quo(Y)
	result = ComputePriceDirection(X, Y, currentYPriceOverX, orderBook)

	fmt.Println(X, Y, currentYPriceOverX)
	fmt.Println(result)
}
