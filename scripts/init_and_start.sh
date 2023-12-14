#!/bin/bash
set -eux

# The script is for auto deployer, node_id could be as moniker
COMMAND=$1
MONIKER=$2
CHAINID=$3
PEERS=$4
GENESIS_FILE=$5
DATA_DIR=/data

if [ ! -f "$DATA_DIR/config/config.toml" ]; then
	# Initialize node's configuration files.
	$COMMAND init $MONIKER --chain-id $CHAINID --home $DATA_DIR

	# Modify the persistent_peers field of config.toml
	sed -i.bak "s/persistent_peers = \"\"/persistent_peers = \"${PEERS}\"/" $DATA_DIR/config/config.toml

	# Download genesis file
	curl -L -o $DATA_DIR/config/genesis.json $GENESIS_FILE

	# Get the IBC_TOKEN_DENOM from genesis.json
	IBC_TOKEN_DENOM=`jq -r '.app_state.crisis.constant_fee.denom' $DATA_DIR/config/genesis.json`

	# Modify the minimum-gas-price field of app.toml
	sed -i.bak "s#minimum-gas-prices = \"0aotto\"#minimum-gas-prices = \"20000000000${IBC_TOKEN_DENOM}\"#" $DATA_DIR/config/app.toml

	# Copy cosmovisor folder to data directory
	cp -R /root/cosmovisor $DATA_DIR/

	# Create a symbolic link for the current version
	ln -s $DATA_DIR/cosmovisor/genesis $DATA_DIR/cosmovisor/current
fi

# Copy cosmovisor folder to data directory
cp -R /root/cosmovisor $DATA_DIR/

cosmovisor run start --home $DATA_DIR --rpc.laddr="tcp://0.0.0.0:26657"