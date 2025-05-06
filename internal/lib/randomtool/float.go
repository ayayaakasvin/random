package randomtool

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math"
)

func RandomFloat(min, max float64, precision uint) (float64, error) {
	if min >= max {
		return 0, fmt.Errorf("min (%f) must be less than max (%f)", min, max)
	}
	if math.IsNaN(min) || math.IsNaN(max) || math.IsInf(min, 0) || math.IsInf(max, 0) {
		return 0, fmt.Errorf("min and max must be finite numbers")
	}

	var randomFloat float64
	var buf [8]byte
	_, err := rand.Read(buf[:])
	if err != nil {
		return 0, fmt.Errorf("failed to generate random number: %v", err)
	}
	randBits := binary.BigEndian.Uint64(buf[:]) >> 12
	randomFloat = float64(randBits) / (1 << 52)

	scaled := min + (max-min)*randomFloat

	if precision < math.MaxUint32 {
		scale := math.Pow10(int(precision))
		scaled = math.Round(scaled*scale) / scale
	}

	return scaled, nil
}