<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [tendermint/liquidity/v1beta1/tx.proto](#tendermint/liquidity/v1beta1/tx.proto)
    - [MsgCreatePool](#tendermint.liquidity.v1beta1.MsgCreatePool)
    - [MsgCreatePoolResponse](#tendermint.liquidity.v1beta1.MsgCreatePoolResponse)
    - [MsgDepositWithinBatch](#tendermint.liquidity.v1beta1.MsgDepositWithinBatch)
    - [MsgDepositWithinBatchResponse](#tendermint.liquidity.v1beta1.MsgDepositWithinBatchResponse)
    - [MsgSwapWithinBatch](#tendermint.liquidity.v1beta1.MsgSwapWithinBatch)
    - [MsgSwapWithinBatchResponse](#tendermint.liquidity.v1beta1.MsgSwapWithinBatchResponse)
    - [MsgWithdrawWithinBatch](#tendermint.liquidity.v1beta1.MsgWithdrawWithinBatch)
    - [MsgWithdrawWithinBatchResponse](#tendermint.liquidity.v1beta1.MsgWithdrawWithinBatchResponse)
  
    - [Msg](#tendermint.liquidity.v1beta1.Msg)
  
- [tendermint/liquidity/v1beta1/liquidity.proto](#tendermint/liquidity/v1beta1/liquidity.proto)
    - [DepositMsgState](#tendermint.liquidity.v1beta1.DepositMsgState)
    - [Params](#tendermint.liquidity.v1beta1.Params)
    - [Pool](#tendermint.liquidity.v1beta1.Pool)
    - [PoolBatch](#tendermint.liquidity.v1beta1.PoolBatch)
    - [PoolMetadata](#tendermint.liquidity.v1beta1.PoolMetadata)
    - [PoolType](#tendermint.liquidity.v1beta1.PoolType)
    - [SwapMsgState](#tendermint.liquidity.v1beta1.SwapMsgState)
    - [WithdrawMsgState](#tendermint.liquidity.v1beta1.WithdrawMsgState)
  
- [tendermint/liquidity/v1beta1/genesis.proto](#tendermint/liquidity/v1beta1/genesis.proto)
    - [GenesisState](#tendermint.liquidity.v1beta1.GenesisState)
    - [PoolRecord](#tendermint.liquidity.v1beta1.PoolRecord)
  
- [tendermint/liquidity/v1beta1/query.proto](#tendermint/liquidity/v1beta1/query.proto)
    - [QueryLiquidityPoolBatchRequest](#tendermint.liquidity.v1beta1.QueryLiquidityPoolBatchRequest)
    - [QueryLiquidityPoolBatchResponse](#tendermint.liquidity.v1beta1.QueryLiquidityPoolBatchResponse)
    - [QueryLiquidityPoolByPoolCoinDenomRequest](#tendermint.liquidity.v1beta1.QueryLiquidityPoolByPoolCoinDenomRequest)
    - [QueryLiquidityPoolByReserveAccRequest](#tendermint.liquidity.v1beta1.QueryLiquidityPoolByReserveAccRequest)
    - [QueryLiquidityPoolRequest](#tendermint.liquidity.v1beta1.QueryLiquidityPoolRequest)
    - [QueryLiquidityPoolResponse](#tendermint.liquidity.v1beta1.QueryLiquidityPoolResponse)
    - [QueryLiquidityPoolsRequest](#tendermint.liquidity.v1beta1.QueryLiquidityPoolsRequest)
    - [QueryLiquidityPoolsResponse](#tendermint.liquidity.v1beta1.QueryLiquidityPoolsResponse)
    - [QueryParamsRequest](#tendermint.liquidity.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#tendermint.liquidity.v1beta1.QueryParamsResponse)
    - [QueryPoolBatchDepositMsgRequest](#tendermint.liquidity.v1beta1.QueryPoolBatchDepositMsgRequest)
    - [QueryPoolBatchDepositMsgResponse](#tendermint.liquidity.v1beta1.QueryPoolBatchDepositMsgResponse)
    - [QueryPoolBatchDepositMsgsRequest](#tendermint.liquidity.v1beta1.QueryPoolBatchDepositMsgsRequest)
    - [QueryPoolBatchDepositMsgsResponse](#tendermint.liquidity.v1beta1.QueryPoolBatchDepositMsgsResponse)
    - [QueryPoolBatchSwapMsgRequest](#tendermint.liquidity.v1beta1.QueryPoolBatchSwapMsgRequest)
    - [QueryPoolBatchSwapMsgResponse](#tendermint.liquidity.v1beta1.QueryPoolBatchSwapMsgResponse)
    - [QueryPoolBatchSwapMsgsRequest](#tendermint.liquidity.v1beta1.QueryPoolBatchSwapMsgsRequest)
    - [QueryPoolBatchSwapMsgsResponse](#tendermint.liquidity.v1beta1.QueryPoolBatchSwapMsgsResponse)
    - [QueryPoolBatchWithdrawMsgRequest](#tendermint.liquidity.v1beta1.QueryPoolBatchWithdrawMsgRequest)
    - [QueryPoolBatchWithdrawMsgResponse](#tendermint.liquidity.v1beta1.QueryPoolBatchWithdrawMsgResponse)
    - [QueryPoolBatchWithdrawMsgsRequest](#tendermint.liquidity.v1beta1.QueryPoolBatchWithdrawMsgsRequest)
    - [QueryPoolBatchWithdrawMsgsResponse](#tendermint.liquidity.v1beta1.QueryPoolBatchWithdrawMsgsResponse)
  
    - [Query](#tendermint.liquidity.v1beta1.Query)
  
- [Scalar Value Types](#scalar-value-types)



<a name="tendermint/liquidity/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## tendermint/liquidity/v1beta1/tx.proto



<a name="tendermint.liquidity.v1beta1.MsgCreatePool"></a>

### MsgCreatePool
MsgCreatePool defines an sdk.Msg type that supports submitting a create liquidity pool tx.

See: https://github.com/gravity-devs/liquidity/blob/develop/x/liquidity/spec/04_messages.md


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_creator_address` | [string](#string) |  |  |
| `pool_type_id` | [uint32](#uint32) |  | id of the target pool type, must match the value in the pool. Only pool-type-id 1 is supported. |
| `deposit_coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | reserve coin pair of the pool to deposit. |






<a name="tendermint.liquidity.v1beta1.MsgCreatePoolResponse"></a>

### MsgCreatePoolResponse
MsgCreatePoolResponse defines the Msg/CreatePool response type.






<a name="tendermint.liquidity.v1beta1.MsgDepositWithinBatch"></a>

### MsgDepositWithinBatch
`MsgDepositWithinBatch defines` an `sdk.Msg` type that supports submitting 
a deposit request to the batch of the liquidity pool.
Deposit is submitted to the batch of the Liquidity pool with the specified 
`pool_id`, `deposit_coins` for reserve.
This request is stacked in the batch of the liquidity pool, is not processed 
immediately, and is processed in the `endblock` at the same time as other requests.

See: https://github.com/gravity-devs/liquidity/blob/develop/x/liquidity/spec/04_messages.md


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor_address` | [string](#string) |  |  |
| `pool_id` | [uint64](#uint64) |  | id of the target pool |
| `deposit_coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | reserve coin pair of the pool to deposit |






<a name="tendermint.liquidity.v1beta1.MsgDepositWithinBatchResponse"></a>

### MsgDepositWithinBatchResponse
MsgDepositWithinBatchResponse defines the Msg/DepositWithinBatch response type.






<a name="tendermint.liquidity.v1beta1.MsgSwapWithinBatch"></a>

### MsgSwapWithinBatch
`MsgSwapWithinBatch` defines an sdk.Msg type that supports submitting a swap offer request to the batch of the liquidity pool.
Submit swap offer to the liquidity pool batch with the specified the `pool_id`, `swap_type_id`,
`demand_coin_denom` with the coin and the price you're offering
and `offer_coin_fee` must be half of offer coin amount * current `params.swap_fee_rate` and ceil for reservation to pay fees.
This request is stacked in the batch of the liquidity pool, is not processed 
immediately, and is processed in the `endblock` at the same time as other requests.
You must request the same fields as the pool.
Only the default `swap_type_id` 1 is supported.

See: https://github.com/gravity-devs/liquidity/tree/develop/doc
https://github.com/gravity-devs/liquidity/blob/develop/x/liquidity/spec/04_messages.md


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `swap_requester_address` | [string](#string) |  | address of swap requester |
| `pool_id` | [uint64](#uint64) |  | id of swap type, must match the value in the pool. Only `swap_type_id` 1 is supported. |
| `swap_type_id` | [uint32](#uint32) |  | id of swap type. Must match the value in the pool. |
| `offer_coin` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | offer sdk.coin for the swap request, must match the denom in the pool. |
| `demand_coin_denom` | [string](#string) |  | denom of demand coin to be exchanged on the swap request, must match the denom in the pool. |
| `offer_coin_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | half of offer coin amount * params.swap_fee_rate and ceil for reservation to pay fees. |
| `order_price` | [string](#string) |  | limit order price for the order, the price is the exchange ratio of X/Y where X is the amount of the first coin and Y is the amount of the second coin when their denoms are sorted alphabetically. |






<a name="tendermint.liquidity.v1beta1.MsgSwapWithinBatchResponse"></a>

### MsgSwapWithinBatchResponse
MsgSwapWithinBatchResponse defines the Msg/Swap response type.






<a name="tendermint.liquidity.v1beta1.MsgWithdrawWithinBatch"></a>

### MsgWithdrawWithinBatch
`MsgWithdrawWithinBatch` defines an `sdk.Msg` type that supports submitting 
a withdraw request to the batch of the liquidity pool.
Withdraw is submitted to the batch from the Liquidity pool with the 
specified `pool_id`, `pool_coin` of the pool.
This request is stacked in the batch of the liquidity pool, is not processed 
immediately, and is processed in the `endblock` at the same time as other requests.

See: https://github.com/gravity-devs/liquidity/blob/develop/x/liquidity/spec/04_messages.md


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `withdrawer_address` | [string](#string) |  |  |
| `pool_id` | [uint64](#uint64) |  | id of the target pool |
| `pool_coin` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="tendermint.liquidity.v1beta1.MsgWithdrawWithinBatchResponse"></a>

### MsgWithdrawWithinBatchResponse
MsgWithdrawWithinBatchResponse defines the Msg/WithdrawWithinBatch response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="tendermint.liquidity.v1beta1.Msg"></a>

### Msg
Msg defines the liquidity Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreatePool` | [MsgCreatePool](#tendermint.liquidity.v1beta1.MsgCreatePool) | [MsgCreatePoolResponse](#tendermint.liquidity.v1beta1.MsgCreatePoolResponse) | Submit a create liquidity pool message. | |
| `DepositWithinBatch` | [MsgDepositWithinBatch](#tendermint.liquidity.v1beta1.MsgDepositWithinBatch) | [MsgDepositWithinBatchResponse](#tendermint.liquidity.v1beta1.MsgDepositWithinBatchResponse) | Submit a deposit to the liquidity pool batch. | |
| `WithdrawWithinBatch` | [MsgWithdrawWithinBatch](#tendermint.liquidity.v1beta1.MsgWithdrawWithinBatch) | [MsgWithdrawWithinBatchResponse](#tendermint.liquidity.v1beta1.MsgWithdrawWithinBatchResponse) | Submit a withdraw from the liquidity pool batch. | |
| `Swap` | [MsgSwapWithinBatch](#tendermint.liquidity.v1beta1.MsgSwapWithinBatch) | [MsgSwapWithinBatchResponse](#tendermint.liquidity.v1beta1.MsgSwapWithinBatchResponse) | Submit a swap to the liquidity pool batch. | |

 <!-- end services -->



<a name="tendermint/liquidity/v1beta1/liquidity.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## tendermint/liquidity/v1beta1/liquidity.proto



<a name="tendermint.liquidity.v1beta1.DepositMsgState"></a>

### DepositMsgState
DepositMsgState defines the state of deposit message that contains state information as it is processed in the next batch or batches.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `msg_height` | [int64](#int64) |  | height where this message is appended to the batch |
| `msg_index` | [uint64](#uint64) |  | index of this deposit message in this liquidity pool |
| `executed` | [bool](#bool) |  | true if executed on this batch, false if not executed |
| `succeeded` | [bool](#bool) |  | true if executed successfully on this batch, false if failed |
| `to_be_deleted` | [bool](#bool) |  | true if ready to be deleted on kvstore, false if not ready to be deleted |
| `msg` | [MsgDepositWithinBatch](#tendermint.liquidity.v1beta1.MsgDepositWithinBatch) |  | MsgDepositWithinBatch |






<a name="tendermint.liquidity.v1beta1.Params"></a>

### Params
Params defines the parameters for the liquidity module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_types` | [PoolType](#tendermint.liquidity.v1beta1.PoolType) | repeated | list of available pool types |
| `min_init_deposit_amount` | [string](#string) |  | Minimum number of coins to be deposited to the liquidity pool on pool creation. |
| `init_pool_coin_mint_amount` | [string](#string) |  | Initial mint amount of pool coins upon pool creation. |
| `max_reserve_coin_amount` | [string](#string) |  | Limit the size of each liquidity pool to minimize risk. In development, set to 0 for no limit. In production, set a limit. |
| `pool_creation_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | Fee paid to create a Liquidity Pool. Set a fee to prevent spamming. |
| `swap_fee_rate` | [string](#string) |  | Swap fee rate for every executed swap. |
| `withdraw_fee_rate` | [string](#string) |  | Reserve coin withdrawal with less proportion by withdrawFeeRate. |
| `max_order_amount_ratio` | [string](#string) |  | Maximum ratio of reserve coins that can be ordered at a swap order. |
| `unit_batch_height` | [uint32](#uint32) |  | The smallest unit batch height for every liquidity pool. |
| `circuit_breaker_enabled` | [bool](#bool) |  | Circuit breaker enables or disables transaction messages in liquidity module. |






<a name="tendermint.liquidity.v1beta1.Pool"></a>

### Pool
Pool defines the liquidity pool that contains pool information.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  | id of the pool |
| `type_id` | [uint32](#uint32) |  | id of the pool_type |
| `reserve_coin_denoms` | [string](#string) | repeated | denoms of reserve coin pair of the pool |
| `reserve_account_address` | [string](#string) |  | reserve account address of the pool |
| `pool_coin_denom` | [string](#string) |  | denom of pool coin of the pool |






<a name="tendermint.liquidity.v1beta1.PoolBatch"></a>

### PoolBatch
PoolBatch defines the batch or batches of a given liquidity pool that contains indexes of deposit, withdraw, and swap messages. 
Index param increments by 1 if the pool id is same.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_id` | [uint64](#uint64) |  | id of the pool |
| `index` | [uint64](#uint64) |  | index of this batch |
| `begin_height` | [int64](#int64) |  | height where this batch is started |
| `deposit_msg_index` | [uint64](#uint64) |  | last index of DepositMsgStates |
| `withdraw_msg_index` | [uint64](#uint64) |  | last index of WithdrawMsgStates |
| `swap_msg_index` | [uint64](#uint64) |  | last index of SwapMsgStates |
| `executed` | [bool](#bool) |  | true if executed, false if not executed |






<a name="tendermint.liquidity.v1beta1.PoolMetadata"></a>

### PoolMetadata
Metadata for the state of each pool for invariant checking after genesis export or import.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_id` | [uint64](#uint64) |  | id of the pool |
| `pool_coin_total_supply` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | pool coin issued at the pool |
| `reserve_coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | reserve coins deposited in the pool |






<a name="tendermint.liquidity.v1beta1.PoolType"></a>

### PoolType
Structure for the pool type to distinguish the characteristics of the reserve pools.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint32](#uint32) |  | This is the id of the pool_type that is used as pool_type_id for pool creation. In this version, only pool-type-id 1 is supported. {"id":1,"name":"ConstantProductLiquidityPool","min_reserve_coin_num":2,"max_reserve_coin_num":2,"description":""} |
| `name` | [string](#string) |  | name of the pool type. |
| `min_reserve_coin_num` | [uint32](#uint32) |  | minimum number of reserveCoins for LiquidityPoolType, only 2 reserve coins are supported. |
| `max_reserve_coin_num` | [uint32](#uint32) |  | maximum number of reserveCoins for LiquidityPoolType, only 2 reserve coins are supported. |
| `description` | [string](#string) |  | description of the pool type. |






<a name="tendermint.liquidity.v1beta1.SwapMsgState"></a>

### SwapMsgState
SwapMsgState defines the state of the swap message that contains state information as the message is processed in the next batch or batches.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `msg_height` | [int64](#int64) |  | height where this message is appended to the batch |
| `msg_index` | [uint64](#uint64) |  | index of this swap message in this liquidity pool |
| `executed` | [bool](#bool) |  | true if executed on this batch, false if not executed |
| `succeeded` | [bool](#bool) |  | true if executed successfully on this batch, false if failed |
| `to_be_deleted` | [bool](#bool) |  | true if ready to be deleted on kvstore, false if not ready to be deleted |
| `order_expiry_height` | [int64](#int64) |  | swap orders are cancelled when current height is equal to or higher than ExpiryHeight |
| `exchanged_offer_coin` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | offer coin exchanged until now |
| `remaining_offer_coin` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | offer coin currently remaining to be exchanged |
| `reserved_offer_coin_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | reserve fee for pays fee in half offer coin |
| `msg` | [MsgSwapWithinBatch](#tendermint.liquidity.v1beta1.MsgSwapWithinBatch) |  | MsgSwapWithinBatch |






<a name="tendermint.liquidity.v1beta1.WithdrawMsgState"></a>

### WithdrawMsgState
WithdrawMsgState defines the state of the withdraw message that contains state information as the message is processed in the next batch or batches.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `msg_height` | [int64](#int64) |  | height where this message is appended to the batch |
| `msg_index` | [uint64](#uint64) |  | index of this withdraw message in this liquidity pool |
| `executed` | [bool](#bool) |  | true if executed on this batch, false if not executed |
| `succeeded` | [bool](#bool) |  | true if executed successfully on this batch, false if failed |
| `to_be_deleted` | [bool](#bool) |  | true if ready to be deleted on kvstore, false if not ready to be deleted |
| `msg` | [MsgWithdrawWithinBatch](#tendermint.liquidity.v1beta1.MsgWithdrawWithinBatch) |  | MsgWithdrawWithinBatch |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="tendermint/liquidity/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## tendermint/liquidity/v1beta1/genesis.proto



<a name="tendermint.liquidity.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the liquidity module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#tendermint.liquidity.v1beta1.Params) |  | params defines all the parameters for the liquidity module. |
| `pool_records` | [PoolRecord](#tendermint.liquidity.v1beta1.PoolRecord) | repeated |  |






<a name="tendermint.liquidity.v1beta1.PoolRecord"></a>

### PoolRecord
records the state of each pool after genesis export or import, used to check variables


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool` | [Pool](#tendermint.liquidity.v1beta1.Pool) |  |  |
| `pool_metadata` | [PoolMetadata](#tendermint.liquidity.v1beta1.PoolMetadata) |  |  |
| `pool_batch` | [PoolBatch](#tendermint.liquidity.v1beta1.PoolBatch) |  |  |
| `deposit_msg_states` | [DepositMsgState](#tendermint.liquidity.v1beta1.DepositMsgState) | repeated |  |
| `withdraw_msg_states` | [WithdrawMsgState](#tendermint.liquidity.v1beta1.WithdrawMsgState) | repeated |  |
| `swap_msg_states` | [SwapMsgState](#tendermint.liquidity.v1beta1.SwapMsgState) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="tendermint/liquidity/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## tendermint/liquidity/v1beta1/query.proto



<a name="tendermint.liquidity.v1beta1.QueryLiquidityPoolBatchRequest"></a>

### QueryLiquidityPoolBatchRequest
the request type for the QueryLiquidityPoolBatch RPC method. requestable including specified pool_id.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_id` | [uint64](#uint64) |  | id of the target pool for query |






<a name="tendermint.liquidity.v1beta1.QueryLiquidityPoolBatchResponse"></a>

### QueryLiquidityPoolBatchResponse
the response type for the QueryLiquidityPoolBatchResponse RPC method. Returns the liquidity pool batch that corresponds to the requested pool_id.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `batch` | [PoolBatch](#tendermint.liquidity.v1beta1.PoolBatch) |  |  |






<a name="tendermint.liquidity.v1beta1.QueryLiquidityPoolByPoolCoinDenomRequest"></a>

### QueryLiquidityPoolByPoolCoinDenomRequest
the request type for the QueryLiquidityByPoolCoinDenomPool RPC method. Requestable specified pool_coin_denom.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_coin_denom` | [string](#string) |  |  |






<a name="tendermint.liquidity.v1beta1.QueryLiquidityPoolByReserveAccRequest"></a>

### QueryLiquidityPoolByReserveAccRequest
the request type for the QueryLiquidityByReserveAcc RPC method. Requestable specified reserve_acc.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `reserve_acc` | [string](#string) |  |  |






<a name="tendermint.liquidity.v1beta1.QueryLiquidityPoolRequest"></a>

### QueryLiquidityPoolRequest
the request type for the QueryLiquidityPool RPC method. requestable specified pool_id.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_id` | [uint64](#uint64) |  |  |






<a name="tendermint.liquidity.v1beta1.QueryLiquidityPoolResponse"></a>

### QueryLiquidityPoolResponse
the response type for the QueryLiquidityPoolResponse RPC method. Returns the liquidity pool that corresponds to the requested pool_id.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool` | [Pool](#tendermint.liquidity.v1beta1.Pool) |  |  |






<a name="tendermint.liquidity.v1beta1.QueryLiquidityPoolsRequest"></a>

### QueryLiquidityPoolsRequest
the request type for the QueryLiquidityPools RPC method. Requestable including pagination offset, limit, key.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="tendermint.liquidity.v1beta1.QueryLiquidityPoolsResponse"></a>

### QueryLiquidityPoolsResponse
the response type for the QueryLiquidityPoolsResponse RPC method. This includes a list of all existing liquidity pools and paging results that contain next_key and total count.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pools` | [Pool](#tendermint.liquidity.v1beta1.Pool) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. not working on this version. |






<a name="tendermint.liquidity.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the QueryParams RPC method.






<a name="tendermint.liquidity.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
the response type for the QueryParamsResponse RPC method. This includes current parameter of the liquidity module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#tendermint.liquidity.v1beta1.Params) |  | params holds all the parameters of this module. |






<a name="tendermint.liquidity.v1beta1.QueryPoolBatchDepositMsgRequest"></a>

### QueryPoolBatchDepositMsgRequest
the request type for the QueryPoolBatchDeposit RPC method. requestable including specified pool_id and msg_index.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_id` | [uint64](#uint64) |  | id of the target pool for query |
| `msg_index` | [uint64](#uint64) |  | target msg_index of the pool |






<a name="tendermint.liquidity.v1beta1.QueryPoolBatchDepositMsgResponse"></a>

### QueryPoolBatchDepositMsgResponse
the response type for the QueryPoolBatchDepositMsg RPC method. This includes a batch swap message of the batch.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposit` | [DepositMsgState](#tendermint.liquidity.v1beta1.DepositMsgState) |  |  |






<a name="tendermint.liquidity.v1beta1.QueryPoolBatchDepositMsgsRequest"></a>

### QueryPoolBatchDepositMsgsRequest
the request type for the QueryPoolBatchDeposit RPC method. Requestable including specified pool_id and pagination offset, limit, key.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_id` | [uint64](#uint64) |  | id of the target pool for query |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="tendermint.liquidity.v1beta1.QueryPoolBatchDepositMsgsResponse"></a>

### QueryPoolBatchDepositMsgsResponse
the response type for the QueryPoolBatchDeposit RPC method. This includes a list of all currently existing deposit messages of the batch and paging results that contain next_key and total count.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposits` | [DepositMsgState](#tendermint.liquidity.v1beta1.DepositMsgState) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. not working on this version. |






<a name="tendermint.liquidity.v1beta1.QueryPoolBatchSwapMsgRequest"></a>

### QueryPoolBatchSwapMsgRequest
the request type for the QueryPoolBatchSwap RPC method. Requestable including specified pool_id and msg_index.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_id` | [uint64](#uint64) |  | id of the target pool for query |
| `msg_index` | [uint64](#uint64) |  | target msg_index of the pool |






<a name="tendermint.liquidity.v1beta1.QueryPoolBatchSwapMsgResponse"></a>

### QueryPoolBatchSwapMsgResponse
the response type for the QueryPoolBatchSwapMsg RPC method. This includes a batch swap message of the batch.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `swap` | [SwapMsgState](#tendermint.liquidity.v1beta1.SwapMsgState) |  |  |






<a name="tendermint.liquidity.v1beta1.QueryPoolBatchSwapMsgsRequest"></a>

### QueryPoolBatchSwapMsgsRequest
the request type for the QueryPoolBatchSwapMsgs RPC method. Requestable including specified pool_id and pagination offset, limit, key.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_id` | [uint64](#uint64) |  | id of the target pool for query |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="tendermint.liquidity.v1beta1.QueryPoolBatchSwapMsgsResponse"></a>

### QueryPoolBatchSwapMsgsResponse
the response type for the QueryPoolBatchSwapMsgs RPC method. This includes list of all currently existing swap messages of the batch and paging results that contain next_key and total count.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `swaps` | [SwapMsgState](#tendermint.liquidity.v1beta1.SwapMsgState) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. not working on this version. |






<a name="tendermint.liquidity.v1beta1.QueryPoolBatchWithdrawMsgRequest"></a>

### QueryPoolBatchWithdrawMsgRequest
the request type for the QueryPoolBatchWithdraw RPC method. requestable including specified pool_id and msg_index.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_id` | [uint64](#uint64) |  | id of the target pool for query |
| `msg_index` | [uint64](#uint64) |  | target msg_index of the pool |






<a name="tendermint.liquidity.v1beta1.QueryPoolBatchWithdrawMsgResponse"></a>

### QueryPoolBatchWithdrawMsgResponse
the response type for the QueryPoolBatchWithdrawMsg RPC method. This includes a batch swap message of the batch.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `withdraw` | [WithdrawMsgState](#tendermint.liquidity.v1beta1.WithdrawMsgState) |  |  |






<a name="tendermint.liquidity.v1beta1.QueryPoolBatchWithdrawMsgsRequest"></a>

### QueryPoolBatchWithdrawMsgsRequest
the request type for the QueryPoolBatchWithdraw RPC method. Requestable including specified pool_id and pagination offset, limit, key.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_id` | [uint64](#uint64) |  | id of the target pool for query |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="tendermint.liquidity.v1beta1.QueryPoolBatchWithdrawMsgsResponse"></a>

### QueryPoolBatchWithdrawMsgsResponse
the response type for the QueryPoolBatchWithdraw RPC method. This includes a list of all currently existing withdraw messages of the batch and paging results that contain next_key and total count.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `withdraws` | [WithdrawMsgState](#tendermint.liquidity.v1beta1.WithdrawMsgState) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. Not supported on this version. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="tendermint.liquidity.v1beta1.Query"></a>

### Query
Query defines the gRPC query service for the liquidity module.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `LiquidityPools` | [QueryLiquidityPoolsRequest](#tendermint.liquidity.v1beta1.QueryLiquidityPoolsRequest) | [QueryLiquidityPoolsResponse](#tendermint.liquidity.v1beta1.QueryLiquidityPoolsResponse) | Get existing liquidity pools. | GET|/cosmos/liquidity/v1beta1/pools|
| `LiquidityPool` | [QueryLiquidityPoolRequest](#tendermint.liquidity.v1beta1.QueryLiquidityPoolRequest) | [QueryLiquidityPoolResponse](#tendermint.liquidity.v1beta1.QueryLiquidityPoolResponse) | Get specific liquidity pool. | GET|/cosmos/liquidity/v1beta1/pools/{pool_id}|
| `LiquidityPoolByPoolCoinDenom` | [QueryLiquidityPoolByPoolCoinDenomRequest](#tendermint.liquidity.v1beta1.QueryLiquidityPoolByPoolCoinDenomRequest) | [QueryLiquidityPoolResponse](#tendermint.liquidity.v1beta1.QueryLiquidityPoolResponse) | Get specific liquidity pool corresponding to the pool_coin_denom. | GET|/cosmos/liquidity/v1beta1/pools/pool_coin_denom/{pool_coin_denom}|
| `LiquidityPoolByReserveAcc` | [QueryLiquidityPoolByReserveAccRequest](#tendermint.liquidity.v1beta1.QueryLiquidityPoolByReserveAccRequest) | [QueryLiquidityPoolResponse](#tendermint.liquidity.v1beta1.QueryLiquidityPoolResponse) | Get specific liquidity pool corresponding to the reserve account. | GET|/cosmos/liquidity/v1beta1/pools/reserve_acc/{reserve_acc}|
| `LiquidityPoolBatch` | [QueryLiquidityPoolBatchRequest](#tendermint.liquidity.v1beta1.QueryLiquidityPoolBatchRequest) | [QueryLiquidityPoolBatchResponse](#tendermint.liquidity.v1beta1.QueryLiquidityPoolBatchResponse) | Get the pool's current batch. | GET|/cosmos/liquidity/v1beta1/pools/{pool_id}/batch|
| `PoolBatchSwapMsgs` | [QueryPoolBatchSwapMsgsRequest](#tendermint.liquidity.v1beta1.QueryPoolBatchSwapMsgsRequest) | [QueryPoolBatchSwapMsgsResponse](#tendermint.liquidity.v1beta1.QueryPoolBatchSwapMsgsResponse) | Get all swap messages in the pool's current batch. | GET|/cosmos/liquidity/v1beta1/pools/{pool_id}/batch/swaps|
| `PoolBatchSwapMsg` | [QueryPoolBatchSwapMsgRequest](#tendermint.liquidity.v1beta1.QueryPoolBatchSwapMsgRequest) | [QueryPoolBatchSwapMsgResponse](#tendermint.liquidity.v1beta1.QueryPoolBatchSwapMsgResponse) | Get a specific swap message in the pool's current batch. | GET|/cosmos/liquidity/v1beta1/pools/{pool_id}/batch/swaps/{msg_index}|
| `PoolBatchDepositMsgs` | [QueryPoolBatchDepositMsgsRequest](#tendermint.liquidity.v1beta1.QueryPoolBatchDepositMsgsRequest) | [QueryPoolBatchDepositMsgsResponse](#tendermint.liquidity.v1beta1.QueryPoolBatchDepositMsgsResponse) | Get all deposit messages in the pool's current batch. | GET|/cosmos/liquidity/v1beta1/pools/{pool_id}/batch/deposits|
| `PoolBatchDepositMsg` | [QueryPoolBatchDepositMsgRequest](#tendermint.liquidity.v1beta1.QueryPoolBatchDepositMsgRequest) | [QueryPoolBatchDepositMsgResponse](#tendermint.liquidity.v1beta1.QueryPoolBatchDepositMsgResponse) | Get a specific deposit message in the pool's current batch. | GET|/cosmos/liquidity/v1beta1/pools/{pool_id}/batch/deposits/{msg_index}|
| `PoolBatchWithdrawMsgs` | [QueryPoolBatchWithdrawMsgsRequest](#tendermint.liquidity.v1beta1.QueryPoolBatchWithdrawMsgsRequest) | [QueryPoolBatchWithdrawMsgsResponse](#tendermint.liquidity.v1beta1.QueryPoolBatchWithdrawMsgsResponse) | Get all withdraw messages in the pool's current batch. | GET|/cosmos/liquidity/v1beta1/pools/{pool_id}/batch/withdraws|
| `PoolBatchWithdrawMsg` | [QueryPoolBatchWithdrawMsgRequest](#tendermint.liquidity.v1beta1.QueryPoolBatchWithdrawMsgRequest) | [QueryPoolBatchWithdrawMsgResponse](#tendermint.liquidity.v1beta1.QueryPoolBatchWithdrawMsgResponse) | Get a specific withdraw message in the pool's current batch. | GET|/cosmos/liquidity/v1beta1/pools/{pool_id}/batch/withdraws/{msg_index}|
| `Params` | [QueryParamsRequest](#tendermint.liquidity.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#tendermint.liquidity.v1beta1.QueryParamsResponse) | Get all parameters of the liquidity module. | GET|/cosmos/liquidity/v1beta1/params|

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |
