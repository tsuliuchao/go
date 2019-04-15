//fatal error: concurrent map read and map write
package main

import "sync"

func main() {
	sm := newSafeMap()

	for i := 0; i < 100000; i++ {
		go sm.writeMap( i, i)
		go sm.readMap( i)
	}

}
type safeMap struct {
	sync.RWMutex
	Map map[int]int
}

func newSafeMap() *safeMap{
	sm := new(safeMap)
	sm.Map = make(map[int]int)
	return sm
}

func (sm *safeMap)readMap(key int) int {
	sm.Lock()
	value := sm.Map[key]
	sm.Unlock()
	return value
}

func (sm *safeMap)writeMap(key int, value int) {
	sm.Lock()
	sm.Map[key] = value
	sm.Unlock()
}