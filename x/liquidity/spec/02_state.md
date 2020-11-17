<!--
order: 2
-->

# State

## LiquidityPool

`LiquidityPool` stores static information of a liquidity pool

```go
type LiquidityPool struct {
	PoolID             uint64         // index of this liquidity pool
	PoolTypeIndex      uint32         // pool type of this liquidity pool
	ReserveCoinDenoms  []string       // list of reserve coin denoms for this liquidity pool
	ReserveAccount     sdk.AccAddress // module account address for this liquidity pool to store reserve coins
	PoolCoinDenom      string         // denom of pool coin for this liquidity pool
}
```

LiquidityPool: `0x11 | LiquidityPoolID -> amino(LiquidityPool)`

LiquidityPoolByReserveAccIndex: `0x12 | ReserveAcc -> nil`


## LiquidityPoolBatch

```go
type LiquidityPoolBatch struct {
	PoolID                  uint64                     // id of target liquidity pool
	BatchIndex              uint64                     // index of this batch
	BeginHeight             uint64                     // height where this batch is begun
	DepositMsgIndex         uint64                     // last index of BatchPoolDepositMsgs	
	WithdrawMsgIndex        uint64                     // last index of BatchPoolWithdrawMsgs	
	SwapMsgIndex            uint64                     // last index of BatchPoolSwapMsgs	
	ExecutionStatus         bool                       // true if executed, false if not executed yet
}


type BatchPoolDepositMsg struct {
	MsgHeight uint64 // height where this message is appended to the batch
	MsgIndex  uint64 // index of this deposit message in this batch
	Msg       MsgDepositToLiquidityPool
}

type BatchPoolWithdrawMsg struct {
	MsgHeight uint64 // height where this message is appended to the batch
	MsgIndex  uint64 // index of this withdraw message in this batch
	Msg       MsgWithdrawFromLiquidityPool
}

type BatchPoolSwapMsg struct {
	MsgHeight    uint64 // height where this message is appended to the batch
	MsgIndex     uint64 // index of this swap message in this batch
	CancelHeight uint32 // swap orders are cancelled when current height is equal or higher than CancelHeight
	Msg          MsgSwap
}

```

LiquidityPoolBatchIndex: `0x21 | PoolID -> amino(int64)`

LiquidityPoolBatch: `0x22 | PoolID -> amino(LiquidityPoolBatch)`

LiquidityPoolBatchDepositMsgIndex: `0x31 | PoolID -> nil`

LiquidityPoolBatchDepositMsgs: `0x31 | PoolID | MsgIndex -> amino(BatchPoolDepositMsg)`

LiquidityPoolBatchWithdrawMsgIndex: `0x32 | PoolID -> nil`

LiquidityPoolBatchWithdrawMsgs: `0x32 | PoolID | MsgIndex -> amino(BatchPoolWithdrawMsg)`

LiquidityPoolBatchSwapMsgIndex: `0x33 | PoolID -> nil`

LiquidityPoolBatchSwapMsgs: `0x33 | PoolID | MsgIndex -> amino(BatchPoolSwapMsg)`
