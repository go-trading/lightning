package core

import (
	"sync"
	"time"
)

type SymbolsCache struct {
	sync.RWMutex
	time.Time
	cache map[string]*Symbol
}

func (sc *SymbolsCache) Get(symbolName string) *Symbol {
	sc.RLock()
	defer sc.RUnlock()
	return sc.cache[symbolName]
}

func (sc *SymbolsCache) IsCacheInvalid(now time.Time) bool {
	sc.RLock()
	defer sc.RUnlock()
	return len(sc.cache) == 0 || now.Sub(sc.Time) > 24*time.Hour
}

func (sc *SymbolsCache) Add(symbols map[string]*Symbol, now time.Time) {
	sc.Lock()
	defer sc.Unlock()
	sc.Time = now
	//only add new Symbol. Don't change pointer to old Symbol
	for _, s := range symbols {
		if _, ok := sc.cache[s.Name()]; !ok {
			if sc.cache == nil {
				sc.cache = make(map[string]*Symbol)
			}
			sc.cache[s.Name()] = s
		}
	}
}
