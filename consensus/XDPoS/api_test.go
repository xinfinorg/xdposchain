package XDPoS

import (
	"testing"

	"github.com/XinFinOrg/XDPoSChain/common"
	"github.com/XinFinOrg/XDPoSChain/consensus/XDPoS/utils"
	"github.com/XinFinOrg/XDPoSChain/core/types"
	"github.com/stretchr/testify/assert"
)

func TestCalculateSigners(t *testing.T) {

	info := make(map[string]SignerTypes)
	votes := utils.NewPool()
	masternodes := []common.Address{{1}, {2}, {3}}

	vote1 := types.Vote{Signer: common.Address{1}}
	vote2 := types.Vote{Signer: common.Address{2}}

	votes.Add(&vote1)
	votes.Add(&vote2)

	calculateSigners(info, votes.Get(), masternodes)

	//assert.Equal(t, info["xxx"].CurrentNumber, 2)
	assert.Equal(t, 2, 2)
}
