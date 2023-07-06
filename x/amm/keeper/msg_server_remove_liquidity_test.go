package keeper_test

import (
	"context"
	"fmt"
	"testing"

	keepertest "frogchain/testutil/keeper"
	"frogchain/x/amm"
	"frogchain/x/amm/keeper"
	"frogchain/x/amm/types"

	"frogchain/x/amm/testutil"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func setupMsgRemoveLiquidity(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context,
	*gomock.Controller, *testutil.MockBankKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	k, ctx := keepertest.AmmKeeperWithMocks(t, bankMock)
	amm.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)

	bankMock.ExpectAny(context)

	createNPool(k, ctx, 1)

	return server, *k, context, ctrl, bankMock
}

func TestMsgRemoveLiquidityNoKey(t *testing.T) {
	ms, _, context, ctrl, _ := setupMsgRemoveLiquidity(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	removeResponse, err := ms.RemoveLiquidity(ctx, &types.MsgRemoveLiquidity{
		Creator:    alice,
		PoolId:     1,
		Liquidity:  10,
		MinAmounts: []uint64{10, 10},
	})

	require.Nil(t, removeResponse)
	require.Equal(t,
		"key 1 doesn't exist: key not found",
		err.Error())
}

func TestMsgRemoveLiquidity(t *testing.T) {
	ms, _, context, ctrl, _ := setupMsgRemoveLiquidity(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	removeResponse, err := ms.RemoveLiquidity(ctx, &types.MsgRemoveLiquidity{
		Creator:    alice,
		PoolId:     0,
		Liquidity:  10,
		MinAmounts: []uint64{10, 10},
	})

	fmt.Print(err)

	response := &types.MsgRemoveLiquidityResponse{
		ReceivedTokens: []sdk.Coin{
			sdk.NewCoin("foocoin", sdk.NewInt(10)),
			sdk.NewCoin("token", sdk.NewInt(10)),
		},
	}

	require.EqualValues(t, response, removeResponse)
}
