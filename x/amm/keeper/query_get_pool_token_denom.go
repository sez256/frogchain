package keeper

import (
	"context"

	"frogchain/x/amm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetPoolTokenDenom(goCtx context.Context, req *types.QueryGetPoolTokenDenomRequest) (*types.QueryGetPoolTokenDenomResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	tokenAsset, err := k.GetPoolTokenForId(ctx, req.Id, req.AssetId)
	if err != nil {
		return nil, err
	}

	return &types.QueryGetPoolTokenDenomResponse{
		TokenDenom: tokenAsset.TokenDenom,
	}, nil
}