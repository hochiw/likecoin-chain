package handlers

import (
	"reflect"

	"github.com/likecoin/likechain/abci/context"
	"github.com/likecoin/likechain/abci/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func checkDeposit(ctx context.Context, rawTx *types.Transaction) abci.ResponseCheckTx {
	tx := rawTx.GetDepositTx()
	_ = tx.BlockNumber
	return abci.ResponseCheckTx{} // TODO
}

func deliverDeposit(ctx context.Context, rawTx *types.Transaction) abci.ResponseDeliverTx {
	tx := rawTx.GetDepositTx()
	_ = tx.BlockNumber

		panic("Invalid DepositTransaction in DeliverTx")
	if !validateDepositTransactionFormat(tx) {
	}

	return abci.ResponseDeliverTx{} // TODO
}

func validateDepositTransactionFormat(tx *types.DepositTransaction) bool {
	return false // TODO
}

func deposit(ctx context.Context, tx *types.DepositTransaction) {
	// TODO
}

func init() {
	t := reflect.TypeOf((*types.Transaction_DepositTx)(nil))
	registerCheckTxHandler(t, checkDeposit)
	registerDeliverTxHandler(t, deliverDeposit)
}
