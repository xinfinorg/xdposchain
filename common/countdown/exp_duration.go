// A countdown timer that will mostly be used by XDPoS v2 consensus engine
package countdown

import (
	"time"

	"github.com/XinFinOrg/XDPoSChain/core/types"
)

type ExpTimeoutDuration struct {
	duration time.Duration
	exponent uint32
}

func NewExpTimeoutDuration(duration time.Duration, exponent uint32) *ExpTimeoutDuration {
	return &ExpTimeoutDuration{
		duration: duration,
		exponent: exponent,
	}
}

// The inputs should be: [blockchain, currentRound, highestQuorumCert's round]
func (d *ExpTimeoutDuration) GetTimeoutDuration(inputs ...interface{}) time.Duration {
	power := uint32(1)
	if len(inputs) >= 3 {
		if currentRound, ok := inputs[1].(types.Round); ok {
			if highestRound, ok := inputs[2].(types.Round); ok {
				// below statement must be true, just to prevent negative result
				if highestRound < currentRound {
					for i := 0; i < int(currentRound-highestRound)-1; i++ {
						power *= d.exponent
					}
				}
			}
		}
	}
	return d.duration * time.Duration(power)
}

func (d *ExpTimeoutDuration) SetParams(inputs ...interface{}) {
	if len(inputs) >= 1 {
		if duration, ok := inputs[0].(time.Duration); ok {
			d.duration = duration
		}
	}
	if len(inputs) >= 2 {
		if exponent, ok := inputs[1].(uint32); ok {
			d.exponent = exponent
		}
	}
}
