package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreatePool = "create_pool"

var _ sdk.Msg = &MsgCreatePool{}

func NewMsgCreatePool(creator string, poolParam *PoolParam, poolAssets []*PoolToken, assetAmounts []uint64) *MsgCreatePool {
	return &MsgCreatePool{
		Creator:      creator,
		PoolParam:    poolParam,
		PoolAssets:   poolAssets,
		AssetAmounts: assetAmounts,
	}
}

func (msg *MsgCreatePool) Route() string {
	return RouterKey
}

func (msg *MsgCreatePool) Type() string {
	return TypeMsgCreatePool
}

func (msg *MsgCreatePool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidAddress, "create | invalid creator address (%s)", err)
	}

	swapFeeAmount := msg.PoolParam.SwapFee
	if swapFeeAmount > TOTALPERCENT {
		return sdkerrors.Wrapf(ErrFeeOverflow, ErrFeeOverflow.Error(), fmt.Sprint(swapFeeAmount), "Swap Fee")
	}

	exitFeeAmount := msg.PoolParam.SwapFee
	if exitFeeAmount > TOTALPERCENT {
		return sdkerrors.Wrapf(ErrFeeOverflow, ErrFeeOverflow.Error(), fmt.Sprint(exitFeeAmount), "Exit Fee")
	}

	for _, poolAsset := range msg.PoolAssets {
		weight := poolAsset.TokenWeight
		if weight == 0 {
			return ErrWeightZero
		}
	}
	return nil
}
