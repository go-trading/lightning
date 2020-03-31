package core

/*
func (s *Symbol) Status() SymbolStatus {
	return SymbolStatus(atomic.LoadInt32((*int32)(&s.status)))
}

func (s *Symbol) SetStatus(status SymbolStatus) {
	atomic.StoreInt32((*int32)(&s.status), int32(status))
	s.statusDispatch(status)
}

func (s *Symbol) OnStatus(fn func(*Symbol, SymbolStatus)) int {
	s.statusSubscribersMutex.Lock()
	defer s.statusSubscribersMutex.Unlock()

	prevStatusSubscribersExist := existsNotNil(s.statusSubscribers)
	defer func() {
		if !prevStatusSubscribersExist && existsNotNil(s.statusSubscribers) {
			s.Exchange().SubscribeSymbolStatus(s)
		}
	}()

	//try insert in place of removed function
	for i := 0; i < len(s.statusSubscribers); i++ {
		if s.statusSubscribers[i] == nil {
			s.statusSubscribers[i] = fn
			return i
		}
	}
	// insert fn in new place
	s.statusSubscribers = append(s.statusSubscribers, fn)
	return len(s.statusSubscribers) - 1
}

func (s *Symbol) RemoveOnStatus(id int) {
	s.statusSubscribersMutex.Lock()
	defer s.statusSubscribersMutex.Unlock()

	prevStatusSubscribersExist := existsNotNil(s.statusSubscribers)
	s.statusSubscribers[id] = nil
	if prevStatusSubscribersExist && !existsNotNil(s.statusSubscribers) {
		s.Exchange().UnsubscribeSymbolStatus(s)
	}
}

func (s *Symbol) statusDispatch(newStatus SymbolStatus) {
	s.statusSubscribersMutex.RLock()
	defer s.statusSubscribersMutex.RUnlock()
	for _, fn := range s.statusSubscribers {
		if fn != nil {
			go fn(s, newStatus)
		}
	}
}
*/
