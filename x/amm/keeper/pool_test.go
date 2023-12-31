package keeper_test

import (
	"testing"

	keepertest "frogchain/testutil/keeper"
	"frogchain/testutil/nullify"
	"frogchain/x/amm/keeper"
	"frogchain/x/amm/types"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNPool(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Pool {
	items := make([]types.Pool, n)
	for i := range items {
		items[i].PoolAssets = []sdk.Coin{
			sdk.NewCoin("foocoin", math.NewInt(10000)),
			sdk.NewCoin("token", math.NewInt(10000)),
		}

		items[i].PoolParam = types.PoolParam{
			SwapFee:      sdk.NewDec(12),
			ExitFee:      sdk.NewDec(15),
			FeeCollector: "user1",
		}

		items[i].ShareToken = sdk.NewCoin(types.ShareTokenIndex(uint64(i)), sdk.NewInt(50))

		items[i].AssetWeights = []sdk.Dec{sdk.NewDec(1), sdk.NewDec(1)}

		items[i].MinimumLiquidity = sdk.NewDec(types.MINIMUM_LIQUIDITY)

		items[i].Id = keeper.GetPoolCount(ctx)
		keeper.SetPool(ctx, items[i])
		keeper.SetPoolCount(ctx, items[i].Id+1)
	}
	return items
}

func TestPoolGet(t *testing.T) {
	keeper, ctx := keepertest.AmmKeeper(t)
	items := createNPool(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetPool(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestPoolRemove(t *testing.T) {
	keeper, ctx := keepertest.AmmKeeper(t)
	items := createNPool(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePool(ctx, item.Id)
		_, found := keeper.GetPool(ctx, item.Id)
		require.False(t, found)
	}
}

func TestPoolGetAll(t *testing.T) {
	keeper, ctx := keepertest.AmmKeeper(t)
	items := createNPool(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPool(ctx)),
	)
}

func TestPoolCount(t *testing.T) {
	keeper, ctx := keepertest.AmmKeeper(t)
	items := createNPool(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetPoolCount(ctx))
}
