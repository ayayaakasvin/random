package randomtool

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func RandomInt(min, max int64) (int64, error) {
	if max <= min {
		return 0, fmt.Errorf("min must be less than max")
	}

    nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max+1-min)))
	if err != nil {
		return -1, err
	}

	n := nBig.Int64()
    return n + min, nil
}