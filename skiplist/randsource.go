package skiplist

import "math/rand"

type defaultRandSource struct{}

func (r defaultRandSource) Int63() int64 {
	return rand.Int63()
}

func (r defaultRandSource) Seed(seed int64) {
	rand.Seed(seed)
}
