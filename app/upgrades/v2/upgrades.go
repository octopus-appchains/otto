// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package v2

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/cosmos/interchain-security/v3/x/ccv/consumer"
	ccvconsumertypes "github.com/cosmos/interchain-security/v3/x/ccv/consumer/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	cdc codec.JSONCodec,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		logger := ctx.Logger().With("upgrade", UpgradeName)
		vm[ccvconsumertypes.ModuleName] = consumer.AppModule{}.ConsensusVersion()
		ccvModule := mm.Modules[ccvconsumertypes.ModuleName].(module.HasGenesis)
		ccvModule.InitGenesis(ctx, cdc, json.RawMessage(genesisJson))

		// Leave modules are as-is to avoid running InitGenesis.
		logger.Debug("running module migrations ...")
		return mm.RunMigrations(ctx, configurator, vm)
	}
}
