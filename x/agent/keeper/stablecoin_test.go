package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cognize/axon/x/agent/keeper"
	"github.com/cognize/axon/x/agent/types"
	"github.com/stretchr/testify/require"
)

func TestStablecoinDeposit(t *testing.T) {
	_, ctx, _, k := setupTestKeeper(t)

	depositor := "cognize1depositor12345678901234567890123456"

	k.SetAgent(ctx, depositor, types.Agent{
		Address:      depositor,
		StakeAmount:  sdk.NewInt64Coin("cognize", 1000),
		Reputation:   50,
		Status:       "online",
	})

	result, err := k.ProcessStablecoinDeposit(ctx, depositor, "200", "cusd")
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestStablecoinDepositBelowMinimum(t *testing.T) {
	_, ctx, _, k := setupTestKeeper(t)

	depositor := "cognize1depositor12345678901234567890123456"

	_, err := k.ProcessStablecoinDeposit(ctx, depositor, "50", "cusd")
	require.Error(t, err)
}

func TestStablecoinWithdrawal(t *testing.T) {
	_, ctx, _, k := setupTestKeeper(t)

	depositor := "cognize1depositor12345678901234567890123456"

	k.SetAgent(ctx, depositor, types.Agent{
		Address:      depositor,
		StakeAmount:  sdk.NewInt64Coin("cognize", 1000),
		Reputation:   50,
		Status:       "online",
	})

	_, err := k.ProcessStablecoinDeposit(ctx, depositor, "200", "cusd")
	require.NoError(t, err)

	err = k.ProcessStablecoinWithdrawal(ctx, depositor, "100", "cusd")
	require.NoError(t, err)
}