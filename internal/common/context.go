package common

import "context"

// TxID is a context key type for transaction IDs.
type TxID string

// ContextWithTxID adds a transaction ID to the context.
func ContextWithTxID(parent context.Context, value int64) context.Context {
	return context.WithValue(parent, TxID("tx"), value)
}

// GetTxIDFromContext extracts the transaction ID from the context.
func GetTxIDFromContext(ctx context.Context) int64 {
	return ctx.Value(TxID("tx")).(int64)
}
