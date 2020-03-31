package core

import "context"

type Exchange interface {
	Service
	Symbol(string) *Symbol

	SubscribeTrades(context.Context, *Trades)
}
