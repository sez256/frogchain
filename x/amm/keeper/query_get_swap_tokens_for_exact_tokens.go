package keeper

import (
	"context"

	"frogchain/x/amm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetSwapTokensForExactTokens(goCtx context.Context, req *types.QueryGetSwapTokensForExactTokensRequest) (*types.QueryGetSwapTokensForExactTokensResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	tokenInAmount, _, err := k.SwapExactAmountOut(ctx, req.PoolId, sdk.NewDec(int64(req.AmountOut)), req.Path)
	if err != nil {
		return nil, err
	}

	return &types.QueryGetSwapTokensForExactTokensResponse{
		AmountIn: tokenInAmount.RoundInt().Uint64(),
	}, nil
}
