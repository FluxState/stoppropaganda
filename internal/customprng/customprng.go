package customprng

import (
	"time"

	"github.com/valyala/fastrand"
)

/*
THIS CODE IS NOT THREAD-SAFE
*/

type RNG struct {
	rng    fastrand.RNG
	buff   []byte
	number uint32
}

func New(maxStringLength int) *RNG {
	r := &RNG{
		rng:  fastrand.RNG{},
		buff: make([]byte, maxStringLength),
	}
	r.rng.Seed(uint32(time.Now().UnixNano()))
	return r
}

var randomRunesList = []byte("abcdefghijklmnopqrstuvwxyz_")
var randomRunesListLength = uint32(len(randomRunesList))

func (r *RNG) Uint32n(min, max uint32) uint32 {
    return r.rng.Uint32n(max-min) + min;
}

func (r *RNG) slice(min, max uint32) []byte {
	r.number = r.Uint32n(min, max)
	r.buff = r.buff[:r.number]
	for r.number > 0 {
		r.number--
		idx := r.rng.Uint32n(randomRunesListLength)
		r.buff[r.number] = randomRunesList[idx]
	}
	return r.buff
}

func (r *RNG) String(min, max uint32) string {
	return string(r.slice(min, max))
}

func (r *RNG) StringSuffix(min, max uint32, suffix string) string {
	return string(append(r.slice(min, max), []byte(suffix)...))
}
