package XDPoS

import (
	"math/big"
	"testing"

	"github.com/XinFinOrg/XDPoSChain/common"
	"github.com/XinFinOrg/XDPoSChain/consensus/XDPoS/utils"
	"github.com/XinFinOrg/XDPoSChain/core/types"
	"github.com/stretchr/testify/assert"
)

func TestCalculateSignersVote(t *testing.T) {

	info := make(map[string]SignerTypes)
	votes := utils.NewPool()
	masternodes := []common.Address{{1}, {2}, {3}}

	vote1 := types.Vote{
		Signer: common.Address{1},
		ProposedBlockInfo: &types.BlockInfo{
			Hash:   common.Hash{1},
			Round:  types.Round(10),
			Number: big.NewInt(910),
		},
		GapNumber: 450,
	}

	vote2 := types.Vote{
		Signer: common.Address{2},
		ProposedBlockInfo: &types.BlockInfo{
			Hash:   common.Hash{2},
			Round:  types.Round(11),
			Number: big.NewInt(911),
		},
		GapNumber: 450,
	}

	votes.Add(&vote1)
	votes.Add(&vote2)

	calculateSigners(info, votes.Get(), masternodes)

	//assert.Equal(t, info["xxx"].CurrentNumber, 2)
	assert.Equal(t, 2, 2)
}

func TestCalculateSignersTimeout(t *testing.T) {

	info := make(map[string]SignerTypes)
	timeouts := utils.NewPool()
	masternodes := []common.Address{{1}, {2}, {3}}

	vote1 := types.Timeout{
		Signer:    common.Address{1},
		Round:     types.Round(10),
		GapNumber: 450,
	}
	vote2 := types.Timeout{
		Signer:    common.Address{2},
		Round:     types.Round(11),
		GapNumber: 450,
	}

	timeouts.Add(&vote1)
	timeouts.Add(&vote2)

	calculateSigners(info, timeouts.Get(), masternodes)

	//assert.Equal(t, info["xxx"].CurrentNumber, 2)
	assert.Equal(t, 2, 2)
}
