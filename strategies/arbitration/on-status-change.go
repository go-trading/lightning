package arbitration

import "github.com/go-trading/lightning/core"

func (b *Bot) onStatusChange(symbol *core.Symbol, status core.SymbolStatus) {
	log.Errorf("TODO new status for %v is %v", symbol.Name(), status)
	return
}
