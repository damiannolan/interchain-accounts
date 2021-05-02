#!/bin/sh
set -e

echo "Initializing ica..."
icad init ica --chain-id ica

echo "Adding genesis account..."
if [ -z "$MNEMONIC" ]
then 
    echo "y" | icad keys add validator --keyring-backend test
else 
    echo "$MNEMONIC" | icad keys add validator --recover --keyring-backend test
fi

icad add-genesis-account $(icad keys show validator --keyring-backend test -a) 100000000000stake

echo "Creating and collecting gentx..."
icad gentx validator 7000000000stake --chain-id ica --keyring-backend test
icad collect-gentxs

icad start --pruning=nothing
