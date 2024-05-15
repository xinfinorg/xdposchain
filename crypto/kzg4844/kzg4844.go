// Copyright 2023 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package kzg4844 implements the KZG crypto for EIP-4844.
package kzg4844

import (
	"embed"
	"github.com/XinFinOrg/XDPoSChain/common/hexutil"
	"hash"
	"reflect"
	"sync/atomic"

)

//go:embed trusted_setup.json
var content embed.FS

var (
	blobT       = reflect.TypeOf(Blob{})
	commitmentT = reflect.TypeOf(Commitment{})
	proofT      = reflect.TypeOf(Proof{})
)

// Blob represents a 4844 data blob.
type Blob [131072]byte

// UnmarshalJSON parses a blob in hex syntax.
func (b *Blob) UnmarshalJSON(input []byte) error {
	return hexutil.UnmarshalFixedJSON(blobT, input, b[:])
}

// MarshalText returns the hex representation of b.
func (b Blob) MarshalText() ([]byte, error) {
	return hexutil.Bytes(b[:]).MarshalText()
}

// Commitment is a serialized commitment to a polynomial.
type Commitment [48]byte

// UnmarshalJSON parses a commitment in hex syntax.
func (c *Commitment) UnmarshalJSON(input []byte) error {
	return hexutil.UnmarshalFixedJSON(commitmentT, input, c[:])
}

// MarshalText returns the hex representation of c.
func (c Commitment) MarshalText() ([]byte, error) {
	return hexutil.Bytes(c[:]).MarshalText()
}

// Proof is a serialized commitment to the quotient polynomial.
type Proof [48]byte

// UnmarshalJSON parses a proof in hex syntax.
func (p *Proof) UnmarshalJSON(input []byte) error {
	return hexutil.UnmarshalFixedJSON(proofT, input, p[:])
}

// MarshalText returns the hex representation of p.
func (p Proof) MarshalText() ([]byte, error) {
	return hexutil.Bytes(p[:]).MarshalText()
}

// Point is a BLS field element.
type Point [32]byte

// Claim is a claimed evaluation value in a specific point.
type Claim [32]byte

// useCKZG controls whether the cryptography should use the Go or C backend.
var useCKZG atomic.Bool

// CalcBlobHashV1 calculates the 'versioned blob hash' of a commitment.
// The given hasher must be a sha256 hash instance, otherwise the result will be invalid!
func CalcBlobHashV1(hasher hash.Hash, commit *Commitment) (vh [32]byte) {
	if hasher.Size() != 32 {
		panic("wrong hash size")
	}
	hasher.Reset()
	hasher.Write(commit[:])
	hasher.Sum(vh[:0])
	vh[0] = 0x01 // version
	return vh
}
