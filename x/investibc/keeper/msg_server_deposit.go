package keeper

import (
	"context"

	"frogchain/x/investibc/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get message creator
	creator := sdk.MustAccAddressFromBech32(msg.Creator)

	receiveToken := msg.Amount

	if receiveToken.Denom != k.DepositDenom(ctx) {
		return nil, types.ErrInvalidDenom
	}

	// send asset tokens from creator to module
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, sdk.NewCoins(receiveToken))
	if sdkError != nil {
		return nil, sdkError
	}

	// set share token data
	moduleToken := sdk.Coin{
		Denom:  types.ModuleToken,
		Amount: receiveToken.Amount,
	}

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(moduleToken)); err != nil {
		return nil, err
	}

	// send module_token to the depositor
	sdkError = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(moduleToken))
	if sdkError != nil {
		return nil, sdkError
	}

	valArr := k.stakingKeeper.GetAllValidators(ctx)
	if len(valArr) < 1 {
		return nil, types.ErrNoValidator
	}

	validator := valArr[0]
	k.stakingKeeper.Delegate(ctx, creator, moduleToken.Amount, stakingtypes.Unbonded, validator, true)

	return &types.MsgDepositResponse{}, nil
}
