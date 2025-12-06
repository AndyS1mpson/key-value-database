package common

import "context"

type TxID string

// ContextWithTxID fill context transaction id
func ContextWithTxID(parent context.Context, value int64) context.Context {
	return context.WithValue(parent, TxID("tx"), value)
}

// GetTxIDFromContext get transaction id from context
func GetTxIDFromContext(ctx context.Context) int64 {
	return ctx.Value(TxID("tx")).(int64)
}
