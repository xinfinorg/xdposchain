package countdown

import (
	"testing"
	"time"

	"github.com/XinFinOrg/XDPoSChain/core/types"
	"github.com/stretchr/testify/assert"
)

func TestExpDuration(t *testing.T) {
	exponent := uint32(2)
	base := time.Second * 59
	duration := NewExpTimeoutDuration(base, exponent)
	// round 10 = 9+1, normal case, should be base
	currentRound := types.Round(10)
	highestQCRound := types.Round(9)
	result := duration.GetTimeoutDuration(nil, currentRound, highestQCRound)
	assert.Equal(t, base, result)

	// round 11 = 9+2, already 1 round timeout, should be base*exponent
	currentRound++
	result = duration.GetTimeoutDuration(nil, currentRound, highestQCRound)
	assert.Equal(t, base*time.Duration(exponent), result)

	// round 12 = 9+3, already 2 rounds timeout, should be base*exponent^2
	currentRound++
	result = duration.GetTimeoutDuration(nil, currentRound, highestQCRound)
	assert.Equal(t, base*time.Duration(exponent)*time.Duration(exponent), result)

	// test compatible with const timeout user calling it
	result = duration.GetTimeoutDuration(nil)
	assert.Equal(t, base, result)

	// test SetParams
	base++
	exponent++
	duration.SetParams(base, exponent)
	result = duration.GetTimeoutDuration(nil, currentRound, highestQCRound)
	assert.Equal(t, base*time.Duration(exponent)*time.Duration(exponent), result)
	base++
	duration.SetParams(base)
	result = duration.GetTimeoutDuration(nil, currentRound, highestQCRound)
	assert.Equal(t, base*time.Duration(exponent)*time.Duration(exponent), result)
}
