package privacy

import (
	"crypto/elliptic"
	"math/big"

	"github.com/coinbase/kryptology/pkg/core/curves"
)

// This file implements hashing to scalar and hashing to point (on curve) using coinbase/kryptology package.
// Usage ref: https://asecuritysite.com/kryptology/hash2curve

// The curve used in ring_ct.go and bulletproof.go is secp256k1 curve but it has different implementation type than the hashing library.

var kryptoCurve = curves.K256()

func hashToScalar(msg []byte, c elliptic.Curve) (*big.Int) {
	return kryptoCurve.Scalar.Hash(msg[:]).BigInt()
}

func hashToPoint(msg []byte, c elliptic.Curve) (ECPoint) {
	pointHash := kryptoCurve.Point.Hash(msg[:])
	x := pointHash.ToAffineUncompressed()[1:33]
	y := pointHash.ToAffineUncompressed()[33:]
	bigX := new(big.Int).SetBytes(x)
	bigY := new(big.Int).SetBytes(y)

	return ECPoint{bigX, bigY}
}
