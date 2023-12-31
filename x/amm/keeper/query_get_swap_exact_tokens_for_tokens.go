package keeper

import (
	"context"

	"frogchain/x/amm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetSwapExactTokensForTokens(goCtx context.Context, req *types.QueryGetSwapExactTokensForTokensRequest) (*types.QueryGetSwapExactTokensForTokensResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	tokenOutAmout, _, err := k.SwapExactAmountIn(ctx, req.PoolId, sdk.NewDec(int64(req.AmountIn)), req.Path)
	if err != nil {
		return nil, err
	}

	return &types.QueryGetSwapExactTokensForTokensResponse{
		AmountOut: tokenOutAmout.TruncateInt().Uint64(),
	}, nil
}
