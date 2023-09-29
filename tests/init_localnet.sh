#!/bin/bash
set -eux

CHAIN_BIN=ottod
CHAIN_ID=otto_9100-1
# set denom
DENOM="aotto"
MONIKER=validator
KEY=$MONIKER
TEST_DIR=$(pwd)/tests
FAUCET_KEY=faucet
FAUCET_SEED=$TEST_DIR/faucet_seed.txt
DENOM_METADAT_SECTION=$(pwd)/tests/denom_metadata.json
PEERS=$(cat $TEST_DIR/persistent_peers.txt)
BUILD_DIR=$(pwd)/build
DATA_DIR=$BUILD_DIR/validatornode01
CONF_DIR=$DATA_DIR/config
GENTX_DIR=$CONF_DIR/gentxs
GENESIS=$CONF_DIR/genesis.json
TEMP_GENESIS=$CONF_DIR/tmp_genesis.json

# clean
rm -rf $BUILD_DIR

echo "Init validatornodes and import keys with chain-id=$CHAIN_ID"
for i in $(ls $TEST_DIR | grep 'validator'); do
    echo "Init $i under $BUILD_DIR/$i"
    $CHAIN_BIN init $MONIKER --chain-id $CHAIN_ID --home $BUILD_DIR/$i
    echo "Import $i key"
    cat $TEST_DIR/$i/seed.txt | $CHAIN_BIN keys add $KEY --home $BUILD_DIR/$i --no-backup --keyring-backend test --recover
done

echo "Init fullnodes with chain-id=$CHAIN_ID"
for i in $(ls $TEST_DIR | grep 'fullnode'); do
    echo "Init $i under $BUILD_DIR/$i"
    $CHAIN_BIN init $MONIKER --chain-id $CHAIN_ID --home $BUILD_DIR/$i
done

echo "Set the validators"
for i in $(ls $BUILD_DIR | grep 'validator'); do
    echo "- Set the priv_validator_key.json file"
    cp $TEST_DIR/$i/priv_validator_key.json $BUILD_DIR/$i/config/priv_validator_key.json
    echo "- Set the node_key.json file"
    cp $TEST_DIR/$i/node_key.json $BUILD_DIR/$i/config/node_key.json
    echo "- Modify the persistent_peers field of config.toml"
    sed -i "s/persistent_peers = \"\"/persistent_peers = \"${PEERS}\"/" $BUILD_DIR/$i/config/config.toml
done

echo "Set the fullnodes"
for i in $(ls $BUILD_DIR | grep 'fullnode');do
    echo "- Set the node_key.json files"
    cp $TEST_DIR/$i/node_key.json $BUILD_DIR/$i/config/node_key.json
    echo "- Modify the pruning field of app.toml"
    sed -i "s/pruning = \"default\"/pruning = \"nothing\"/" $BUILD_DIR/$i/config/app.toml
    echo "- Modify the persistent_peers field of config.toml"
    sed -i "s/persistent_peers = \"\"/persistent_peers = \"${PEERS}\"/" $BUILD_DIR/$i/config/config.toml
done

echo "Prepare genesis..."
echo "- Set gas limit in genesis"
jq '.consensus_params["block"]["max_gas"]="10000000"' "$GENESIS" > "$TEMP_GENESIS" && mv "$TEMP_GENESIS" "$GENESIS"

echo "- Set bank.denom_metadata with the content of $DENOM_METADAT_SECTION"
jq -s '.[0].app_state.bank.denom_metadata = .[1] | .[0]' $GENESIS $DENOM_METADAT_SECTION > $CONF_DIR/genesis_denom_metadata.json
mv $CONF_DIR/genesis_denom_metadata.json $GENESIS

echo "- Set $DENOM as denom"
sed -i.bak "s/aevmos/$DENOM/g" $GENESIS
sed -i.bak "s/stake/$DENOM/g" $GENESIS

echo "- Set no_base_fee"
sed -i.bak 's/"no_base_fee": false/"no_base_fee": true/g' "$GENESIS"

# # Change proposal periods to pass within a reasonable time for local testing
# sed -i 's/"max_deposit_period": "172800s"/"max_deposit_period": "30s"/g' "$GENESIS"
# sed -i 's/"voting_period": "172800s"/"voting_period": "30s"/g' "$GENESIS"
# # Change proposal required quorum to 15%, so with the orchestrator vote the proposals pass 
# sed -i 's/"quorum": "0.334000000000000000"/"quorum": "0.150000000000000000"/g' "$GENESIS"

echo "Import $FAUCET_KEY key"
cat $FAUCET_SEED | $CHAIN_BIN keys add $FAUCET_KEY --home $DATA_DIR --no-backup --keyring-backend test --recover

echo "- Allocate the faucet account"
$CHAIN_BIN add-genesis-account \
"$($CHAIN_BIN keys show $FAUCET_KEY -a --home $DATA_DIR --keyring-backend test)" 100000000000000000000000000$DENOM \
--home $DATA_DIR --keyring-backend test

echo "- Allocate the validators genesis accounts"
for i in $(ls $BUILD_DIR | grep 'validator'); do
    address=$($CHAIN_BIN keys show $KEY -a --home $BUILD_DIR/$i --keyring-backend test)
    $CHAIN_BIN add-genesis-account \
    "$address" 100000000000000000000000000$DENOM --home $DATA_DIR --keyring-backend test
done

echo "- Distribute genesis.json to all validators and fullnodes"
for i in $(ls $BUILD_DIR | grep 'node'); do
    if [ "$i" != "validatornode01" ]; then
        cp $GENESIS $BUILD_DIR/$i/config/genesis.json
    fi
done

echo "- Sign genesis transaction"
mkdir $GENTX_DIR
for i in $(ls $BUILD_DIR | grep 'validator'); do
    $CHAIN_BIN gentx $KEY 10000000000000000000000$DENOM --home $BUILD_DIR/$i --keyring-backend test --output-document $CONF_DIR/gentxs/gentx-$i.json 
done

echo "- Collect genesis tx"
$CHAIN_BIN collect-gentxs --gentx-dir $CONF_DIR/gentxs --home $DATA_DIR

echo "- Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
$CHAIN_BIN validate-genesis --home $DATA_DIR

echo "- Distribute final genesis.json to all validators and fullnodes"
for i in $(ls $BUILD_DIR | grep 'node'); do
    if [ "$i" != "validatornode01" ]; then
        cp $GENESIS $BUILD_DIR/$i/config/genesis.json
    fi
done