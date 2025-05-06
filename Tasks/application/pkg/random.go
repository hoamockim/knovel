package pkg

import (
	"crypto/rand"
	"fmt"
	"math/big"
	mathRand "math/rand"
	"sync"
)

var Math *mathRand.Rand

type lockedSource struct {
	lk  sync.Mutex
	src mathRand.Source
}

func init() {
	maxInt63 := new(big.Int).SetUint64(1 << 63)
	var seed *big.Int
	var err error
	if seed, err = rand.Int(rand.Reader, maxInt63); err != nil {
		panic(err)
	}
	Math = mathRand.New(&lockedSource{src: mathRand.NewSource(seed.Int64())})
}

func (r *lockedSource) Int63() (n int64) {
	r.lk.Lock()
	n = r.src.Int63()
	r.lk.Unlock()
	return
}

func (r *lockedSource) Seed(seed int64) {
	r.lk.Lock()
	r.src.Seed(seed)
	r.lk.Unlock()
}

func Bytes(n int) []byte {
	data := make([]byte, n)
	return data
}

func UUID() string {
	id := Bytes(16)
	id[6] &= 0x0F // clear version
	id[6] |= 0x40 // set version to 4 (random uuid)
	id[8] &= 0x3F // clear variant
	id[8] |= 0x80 // set to IETF variant
	return fmt.Sprintf("%x-%x-%x-%x-%x", id[0:4], id[4:6], id[6:8], id[8:10], id[10:])
}

func UUID8() string {
	id := Bytes(12)
	return fmt.Sprintf("%x", id[0:])
}
