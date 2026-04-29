package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cognize/axon/x/agent/keeper"
	"github.com/cognize/axon/x/agent/types"
	"github.com/stretchr/testify/require"
)

func TestGovernanceCreateProposal(t *testing.T) {
	_, ctx, _, k := setupTestKeeper(t)

	proposer := "cognize1proposer123456789012345678901234"

	k.SetAgent(ctx, proposer, types.Agent{
		Address:      proposer,
		StakeAmount:  sdk.NewInt64Coin("cognize", 15000),
		Reputation:   30,
		Status:       "online",
	})

	proposalID, err := k.SubmitProposal(ctx, proposer, "test proposal", "text", "description", "100")
	require.NoError(t, err)
	require.Equal(t, int64(1), proposalID)
}

func TestGovernanceCreateProposalInsufficientStake(t *testing.T) {
	_, ctx, _, k := setupTestKeeper(t)

	proposer := "cognize1proposer123456789012345678901234"

	k.SetAgent(ctx, proposer, types.Agent{
		Address:      proposer,
		StakeAmount:  sdk.NewInt64Coin("cognize", 5000),
		Reputation:   30,
		Status:       "online",
	})

	_, err := k.SubmitProposal(ctx, proposer, "test proposal", "text", "description", "100")
	require.Error(t, err)
}

func TestGovernanceCreateProposalInsufficientReputation(t *testing.T) {
	_, ctx, _, k := setupTestKeeper(t)

	proposer := "cognize1proposer123456789012345678901234"

	k.SetAgent(ctx, proposer, types.Agent{
		Address:      proposer,
		StakeAmount:  sdk.NewInt64Coin("cognize", 15000),
		Reputation:   10,
		Status:       "online",
	})

	_, err := k.SubmitProposal(ctx, proposer, "test proposal", "text", "description", "100")
	require.Error(t, err)
}

func TestGovernanceVote(t *testing.T) {
	_, ctx, _, k := setupTestKeeper(t)

	proposer := "cognize1proposer123456789012345678901234"
	voter := "cognize1voter12345678901234567890123456"

	k.SetAgent(ctx, proposer, types.Agent{
		Address:      proposer,
		StakeAmount:  sdk.NewInt64Coin("cognize", 15000),
		Reputation:   30,
		Status:       "online",
	})

	k.SetAgent(ctx, voter, types.Agent{
		Address:      voter,
		StakeAmount:  sdk.NewInt64Coin("cognize", 10000),
		Reputation:   25,
		Status:       "online",
	})

	proposalID, err := k.SubmitProposal(ctx, proposer, "test proposal", "text", "description", "100")
	require.NoError(t, err)

	err = k.Vote(ctx, voter, proposalID, "yes", "")
	require.NoError(t, err)
}

func TestGovernanceVoteInvalidProposal(t *testing.T) {
	_, ctx, _, k := setupTestKeeper(t)

	voter := "cognize1voter12345678901234567890123456"

	k.SetAgent(ctx, voter, types.Agent{
		Address:      voter,
		StakeAmount:  sdk.NewInt64Coin("cognize", 10000),
		Reputation:   25,
		Status:       "online",
	})

	err := k.Vote(ctx, voter, 999, "yes", "")
	require.Error(t, err)
}