package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cognize/axon/x/agent/keeper"
	"github.com/cognize/axon/x/agent/types"
	"github.com/stretchr/testify/require"
)

func TestEscrowCreate(t *testing.T) {
	_, ctx, _, k := setupTestKeeper(t)

	seller := "cognize1seller123456789012345678901234567890"
	buyer := "cognize1buyer123456789012345678901234567"

	k.SetAgent(ctx, buyer, types.Agent{
		Address:      buyer,
		StakeAmount:  sdk.NewInt64Coin("cognize", 100),
		Reputation:   50,
		Status:       "online",
	})

	escrowID, err := k.CreateEscrow(ctx, seller, buyer, "50", "service-1", "", "test escrow")
	require.NoError(t, err)
	require.NotEmpty(t, escrowID)

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyEscrow(escrowID))
	require.NotNil(t, bz)
	require.Contains(t, string(bz), "service-1")
}

func TestEscrowCreateInvalidAmount(t *testing.T) {
	_, ctx, _, k := setupTestKeeper(t)

	seller := "cognize1seller123456789012345678901234567890"
	buyer := "cognize1buyer123456789012345678901234567"

	k.SetAgent(ctx, buyer, types.Agent{
		Address:      buyer,
		StakeAmount:  sdk.NewInt64Coin("cognize", 100),
		Reputation:   50,
		Status:       "online",
	})

	_, err := k.CreateEscrow(ctx, seller, buyer, "0", "service-1", "", "")
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrInvalidAmount)

	_, err = k.CreateEscrow(ctx, seller, buyer, "-10", "service-1", "", "")
	require.Error(t, err)
}

func TestEscrowCreateBuyerNotAgent(t *testing.T) {
	_, ctx, _, k := setupTestKeeper(t)

	seller := "cognize1seller123456789012345678901234567890"
	buyer := "cognize1notagent12345678901234567890123"

	_, err := k.CreateEscrow(ctx, seller, buyer, "50", "service-1", "", "")
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrAgentNotFound)
}

func TestEscrowCreateInsufficientFunds(t *testing.T) {
	_, ctx, _, k := setupTestKeeper(t)

	seller := "cognize1seller123456789012345678901234567890"
	buyer := "cognize1buyer123456789012345678901234567"

	k.SetAgent(ctx, buyer, types.Agent{
		Address:      buyer,
		StakeAmount:  sdk.NewInt64Coin("cognize", 10),
		Reputation:   50,
		Status:       "online",
	})

	_, err := k.CreateEscrow(ctx, seller, buyer, "50", "service-1", "", "")
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrInsufficientFunds)
}

func TestEscrowComplete(t *testing.T) {
	_, ctx, _, k := setupTestKeeper(t)

	seller := "cognize1seller123456789012345678901234567890"
	buyer := "cognize1buyer123456789012345678901234567"

	k.SetAgent(ctx, buyer, types.Agent{
		Address:      buyer,
		StakeAmount:  sdk.NewInt64Coin("cognize", 100),
		Reputation:   50,
		Status:       "online",
	})

	escrowID, err := k.CreateEscrow(ctx, seller, buyer, "50", "service-1", "", "")
	require.NoError(t, err)

	err = k.ConfirmDelivery(ctx, escrowID, seller)
	require.NoError(t, err)
}