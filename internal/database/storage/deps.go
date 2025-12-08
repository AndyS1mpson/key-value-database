package storage

import "context"

//go:generate mockgen -source=deps.go -destination=./mocks/mock.go

type engine interface {
	Set(context.Context, string, string)
	Get(context.Context, string) (string, bool)
	Del(context.Context, string)
}
