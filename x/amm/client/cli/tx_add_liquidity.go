package cli

import (
	"strconv"

	"frogchain/x/amm/types"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddLiquidity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-liquidity [pool-id] [desired-amounts] [min-amounts]",
		Short: "Broadcast message add-liquidity",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			// get pool id
			argPoolId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			// get desired token amounts to add liquidity
			argCastDesiredAmounts := strings.Split(args[1], listSeparator)
			argDesiredAmounts := make([]sdk.Dec, len(argCastDesiredAmounts))
			for i, arg := range argCastDesiredAmounts {
				value, err := sdk.NewDecFromStr(arg)
				if err != nil {
					return err
				}
				argDesiredAmounts[i] = value
			}

			// get minimum token amounts to add liquidity
			argCastMinAmounts := strings.Split(args[2], listSeparator)
			argMinAmounts := make([]sdk.Dec, len(argCastMinAmounts))
			for i, arg := range argCastMinAmounts {
				value, err := sdk.NewDecFromStr(arg)
				if err != nil {
					return err
				}
				argMinAmounts[i] = value
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddLiquidity(
				clientCtx.GetFromAddress().String(),
				argPoolId,
				argDesiredAmounts,
				argMinAmounts,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
