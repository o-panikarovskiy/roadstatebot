package commands

import (
	"crypto/rand"
	"math/big"
	"unicode/utf8"
)

func min(a uint64, b uint64) uint64 {
	if a > b {
		return b
	}
	return a
}

func max(a uint64, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func getRandValueInArr(arr []string) string {
	index, err := rand.Int(rand.Reader, big.NewInt(int64(len(arr))))

	var idx int64
	if err == nil {
		idx = index.Int64()
	} else {
		idx = 0
	}

	return arr[idx]
}
