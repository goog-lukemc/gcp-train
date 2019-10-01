package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Badway
	BadWay()

	// A one better
	//BetterOne()

	// Look at the build in sync map type
	time.Sleep(time.Second * 10)

}

func BetterOne() {
	c := newProtectedMap()
	for i := 0; i < 50; i++ {
		go c.mutate("a", "1")
		go c.mutate("a", "2")
	}
}

// A custom map that contains a protected map
type protectedMap struct {
	cache map[string]string
	mux   sync.Mutex
}

// A mutex blocks the execution of this specific code
func (p *protectedMap) mutate(k string, v string) {
	p.mux.Lock()
	p.cache[k] = v
	fmt.Println(p.cache[k])
	p.mux.Unlock()
}

// We need to initialize the map when we new up the type
func newProtectedMap() (p protectedMap) {
	p.cache = make(map[string]string)
	return p
}

func BadWay() {
	thing := make(map[string]string)
	for i := 0; i < 1; i++ {
		go mutate(thing, "a", "1")
		go mutate(thing, "a", "2")
	}

}

func mutate(p map[string]string, k string, v string) {
	for {
		p[k] = v
	}
}
