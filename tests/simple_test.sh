#!/bin/bash
set -eux

CHAIN_BIN=ottod
CHAINID=otto_9100-1
MONIKER=dev
HOME_DIR=~/.$MONIKER-node
KEY=$MONIKER
TEST_ADDR=otto1jgp2w8t5zausx4vgnc4w3ywh22xwke2lljk0n3
TEST_AMOUNT=2000000000000000000aotto

# Before: Get balance
$CHAIN_BIN query bank balances $TEST_ADDR --chain-id $CHAINID

# Test transfer
$CHAIN_BIN tx bank send \
$($CHAIN_BIN keys show $KEY -a --home $HOME_DIR --keyring-backend test) \
$TEST_ADDR \
$TEST_AMOUNT \
--chain-id $CHAINID \
--home $HOME_DIR \
--keyring-backend test \
--trace -y

sleep 15s

# After: Get balance
$CHAIN_BIN query bank balances $TEST_ADDR --chain-id $CHAINID