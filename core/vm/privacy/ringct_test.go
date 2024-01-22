package privacy

import (
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSign(t *testing.T) {
	/*for i := 14; i < 15; i++ {
	for j := 14; j < 15; j++ {
		for k := 0; k <= j; k++ {*/
	numRing := 5
	ringSize := 10
	s := 9
	fmt.Println("Generate random ring parameter ")
	rings, privkeys, m, err := GenerateMultiRingParams(numRing, ringSize, s)

	fmt.Println("numRing  ", numRing)
	fmt.Println("ringSize  ", ringSize)
	fmt.Println("index of real one  ", s)

	fmt.Println("Ring  ", rings)
	fmt.Println("privkeys  ", privkeys)
	fmt.Println("m  ", m)

	ringSignature, err := Sign(m, rings, privkeys, s)
	if err != nil {
		t.Error("Failed to create Ring signature")
	}

	sig, err := ringSignature.Serialize()
	if err != nil {
		t.Error("Failed to Serialize input Ring signature")
	}

	deserializedSig, err := Deserialize(sig)
	if err != nil {
		t.Error("Failed to Deserialize Ring signature")
	}
	verified := Verify(deserializedSig, false)

	if !verified {
		t.Error("Failed to verify Ring signature")
	}

}

func TestDeserialize(t *testing.T) {
	numRing := 5
	ringSize := 10
	s := 5
	rings, privkeys, m, err := GenerateMultiRingParams(numRing, ringSize, s)

	ringSignature, err := Sign(m, rings, privkeys, s)
	if err != nil {
		t.Error("Failed to create Ring signature")
	}

	// A normal signature.
	sig, err := ringSignature.Serialize()
	if err != nil {
		t.Error("Failed to Serialize input Ring signature")
	}

	// Modify the serialized signature s.t.
	// the new signature passes the length check
	// but triggers buffer overflow in Deserialize().
	// ringSize: 10 -> 56759212534490939
	// len(sig): 3495 -> 3804
	// 80 + 5 * (56759212534490939*65 + 33) = 18446744073709551616 + 3804
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, 56759212534490939)
	for i := 0; i < 8; i++ {
		sig[i+8] = bs[i]
	}
	tail := make([]byte, 3804-len(sig))
	sig = append(sig, tail...)

	_, err = Deserialize(sig)
	assert.EqualError(t, err, "incorrect ring size, len r: 3804, sig.NumRing: 5 sig.Size: 56759212534490939")
}
