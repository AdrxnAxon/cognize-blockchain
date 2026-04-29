package keeper_test

import (
	"testing"

	"github.com/cognize/axon/x/privacy/keeper"
	"github.com/cognize/axon/x/privacy/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupPrivacyTestKeeper(t *testing.T) (keeper.Keeper, sdk.Context) {
	k, ctx := newPrivacyTestKeeper(t)
	return k, ctx
}

func TestPrivacyCommitmentTree(t *testing.T) {
	_, ctx, pk := setupPrivacyTestKeeper(t)

	size := pk.GetTreeSize(ctx)
	require.Equal(t, uint64(0), size)

	pk.setTreeSize(ctx, 10)
	size = pk.GetTreeSize(ctx)
	require.Equal(t, uint64(10), size)
}

func TestPrivacyCommitmentInsert(t *testing.T) {
	_, ctx, pk := setupPrivacyTestKeeper(t)

	commitment := [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}

	err := pk.InsertCommitment(ctx, commitment)
	require.NoError(t, err)

	size := pk.GetTreeSize(ctx)
	require.Equal(t, uint64(1), size)
}

func TestPrivacyViewingKey(t *testing.T) {
	_, ctx, pk := setupPrivacyTestKeeper(t)

	addr := "cognize1test1234567890123456789012345678"
	keyID, key, err := pk.GenerateViewingKey(ctx, addr, "test-service", 1)
	require.NoError(t, err)
	require.NotEmpty(t, keyID)
	require.NotEmpty(t, key)
}

func TestPrivacyMixerCreate(t *testing.T) {
	_, ctx, pk := setupPrivacyTestKeeper(t)

	depositAmt := "1000"
	participants := uint64(10)

	poolID, err := pk.CreateMixerPool(ctx, depositAmt, participants, "cognize")
	require.NoError(t, err)
	require.NotEmpty(t, poolID)
}