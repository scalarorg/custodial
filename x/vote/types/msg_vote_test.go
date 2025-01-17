package types

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/stretchr/testify/assert"

	"github.com/scalarorg/scalar-core/testutils/rand"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	"github.com/scalarorg/scalar-core/x/vote/exported"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
)

func TestVoteRequest_ValidateBasic(t *testing.T) {
	t.Run("no sender", func(t *testing.T) {
		vote := NewVoteRequest(nil, exported.PollID(rand.PosI64()), chainsTypes.NewVoteEvents(nexus.ChainName(rand.NormalizedStr(3))))
		assert.Error(t, vote.ValidateBasic())
	})

	t.Run("no vote", func(t *testing.T) {
		vote := &VoteRequest{
			Sender: rand.AccAddr(),
			PollID: exported.PollID(rand.PosI64()),
			Vote:   nil,
		}
		assert.Error(t, vote.ValidateBasic())
	})

	t.Run("faulty vote type", func(t *testing.T) {
		vote := &VoteRequest{
			Sender: rand.AccAddr(),
			PollID: exported.PollID(rand.PosI64()),
			Vote:   types.UnsafePackAny("result"),
		}
		assert.Error(t, vote.ValidateBasic())
	})

	t.Run("correct vote", func(t *testing.T) {
		vote := NewVoteRequest(rand.AccAddr(), exported.PollID(rand.PosI64()), chainsTypes.NewVoteEvents(nexus.ChainName(rand.NormalizedStr(3))))
		assert.NoError(t, vote.ValidateBasic())
	})
}
