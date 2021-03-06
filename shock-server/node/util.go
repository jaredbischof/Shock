package node

import (
	"sync"
)

type mappy map[string]bool

func IsInMappy(item string, mp mappy) bool {
	if _, ok := mp[item]; ok {
		return true
	}
	return false
}

var virtIdx = mappy{"size": true}

var (
	LockMgr = NewLocker()
)

type Locker struct {
	sync.Mutex
	nLock map[string]*NodeLock
}

type NodeLock struct {
	sync.Mutex
}

func NewLocker() *Locker {
	return &Locker{
		nLock: map[string]*NodeLock{},
	}
}

func (l *Locker) LockNode(id string) {
	// add to map if missing
	l.Lock()
	if _, ok := l.nLock[id]; !ok {
		l.nLock[id] = new(NodeLock)
	}
	l.nLock[id].Lock()
	l.Unlock()
}

func (l *Locker) UnlockNode(id string) {
	if _, ok := l.nLock[id]; ok {
		l.nLock[id].Unlock()
	}
}

func (l *Locker) RemoveNode(id string) {
	delete(l.nLock, id)
}
