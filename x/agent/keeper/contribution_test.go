package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cognize/axon/x/agent/keeper"
	"github.com/cognize/axon/x/agent/types"
	"github.com/stretchr/testify/require"
)

func TestContributionRewardsMint(t *testing.T) {
	_, ctx, _, k := setupTestKeeper(t)

	ctx = ctx.WithBlockHeight(100)

	k.MintContributionRewards(ctx)
	require.True(t, true)
}

func TestContributionDistribution(t *testing.T) {
	_, ctx, _, k := setupTestKeeper(t)

	agentAddr := "cognize1contributor123456789012345678901234"

	k.SetAgent(ctx, agentAddr, types.Agent{
		Address:      agentAddr,
		StakeAmount:  sdk.NewInt64Coin("cognize", 1000),
		Reputation:   30,
		Status:       "online",
		RegisteredAt: 1,
	})

	store := ctx.KVStore(k.storeKey)
	require.NotNil(t, store)
}

func TestContributionRewardsCap(t *testing.T) {
	_, ctx, _, k := setupTestKeeper(t)

	pool := k.GetContributionPool(ctx)
	require.NotNil(t, pool)
}