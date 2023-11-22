package app

import (
	"fmt"

	admintypes "github.com/cosmos/admin-module/x/adminmodule/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	erc20cli "github.com/evmos/evmos/v14/x/erc20/client/cli"
	"github.com/evmos/evmos/v14/x/erc20/types"
	"github.com/spf13/cobra"
)

// NewRegisterCoinProposalCmd implements the command to submit a community-pool-spend proposal
func NewRegisterCoinProposalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-coin METADATA_FILE",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a register coin proposal",
		Long:  `Submit a proposal to register a Cosmos coin to the erc20 along with an initial deposit. The proposal details must be supplied via a JSON file.`,
		Example: fmt.Sprintf(`$ %s tx adminmodule submit-proposal register-coin metadata.json --from=<key_or_address>

Where metadata.json contains (example):

{
  "metadata": [
    {
			"description": "The native staking and governance token of the Osmosis chain",
			"denom_units": [
				{
						"denom": "ibc/<HASH>",
						"exponent": 0,
						"aliases": ["ibcuosmo"]
				},
				{
						"denom": "OSMO",
						"exponent": 6
				}
			],
			"base": "ibc/<HASH>",
			"display": "OSMO",
			"name": "Osmo",
			"symbol": "OSMO"
		}
	]
}`, version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			title, err := cmd.Flags().GetString(cli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(cli.FlagDescription) //nolint:staticcheck,nolintlint
			if err != nil {
				return err
			}

			metadata, err := erc20cli.ParseMetadata(clientCtx.Codec, args[0])
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			content := types.NewRegisterCoinProposal(title, description, metadata...)

			msg, err := admintypes.NewMsgSubmitProposal(content, from)
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal") //nolint:staticcheck,nolintlint
	if err := cmd.MarkFlagRequired(cli.FlagTitle); err != nil {
		panic(err)
	}
	if err := cmd.MarkFlagRequired(cli.FlagDescription); err != nil { //nolint:staticcheck,nolintlint
		panic(err)
	}
	return cmd
}

// NewRegisterERC20ProposalCmd implements the command to submit a community-pool-spend proposal
func NewRegisterERC20ProposalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "register-erc20 ERC20_ADDRESS...",
		Args:    cobra.MinimumNArgs(1),
		Short:   "Submit a proposal to register ERC20 token",
		Long:    "Submit a proposal to register ERC20 tokens along with an initial deposit. To register multiple tokens in one proposal pass them after each other e.g. `register-erc20 <contract-address1> <contract-address2>` ",
		Example: fmt.Sprintf("$ %s tx adminmodule submit-proposal register-erc20 <contract-address> --from=<key_or_address>", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			title, err := cmd.Flags().GetString(cli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(cli.FlagDescription) //nolint:staticcheck
			if err != nil {
				return err
			}

			erc20Addresses := args
			from := clientCtx.GetFromAddress()
			content := types.NewRegisterERC20Proposal(title, description, erc20Addresses...)

			msg, err := admintypes.NewMsgSubmitProposal(content, from)
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal") //nolint:staticcheck
	if err := cmd.MarkFlagRequired(cli.FlagTitle); err != nil {
		panic(err)
	}
	if err := cmd.MarkFlagRequired(cli.FlagDescription); err != nil { //nolint:staticcheck
		panic(err)
	}
	return cmd
}

// NewToggleTokenConversionProposalCmd implements the command to submit a community-pool-spend proposal
func NewToggleTokenConversionProposalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "toggle-token-conversion TOKEN",
		Args:    cobra.ExactArgs(1),
		Short:   "Submit a toggle token conversion proposal",
		Long:    "Submit a proposal to toggle the conversion of a token pair along with an initial deposit.",
		Example: fmt.Sprintf("$ %s tx adminmodule submit-proposal toggle-token-conversion DENOM_OR_CONTRACT --from=<key_or_address>", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			title, err := cmd.Flags().GetString(cli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(cli.FlagDescription) //nolint:staticcheck
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			token := args[0]
			content := types.NewToggleTokenConversionProposal(title, description, token)

			msg, err := admintypes.NewMsgSubmitProposal(content, from)
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal") //nolint:staticcheck
	if err := cmd.MarkFlagRequired(cli.FlagTitle); err != nil {
		panic(err)
	}
	if err := cmd.MarkFlagRequired(cli.FlagDescription); err != nil { //nolint:staticcheck
		panic(err)
	}
	return cmd
}
