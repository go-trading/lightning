package core

type Strategy interface {
	Service
	Positions() []Position
	ClosePositions()
}