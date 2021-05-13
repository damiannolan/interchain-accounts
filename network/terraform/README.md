# Terraform IAC

## Prerequisites

1. Install the [Terraform CLI](https://www.terraform.io/downloads.html)

2. Install the [Digital Ocean CLI](https://github.com/digitalocean/doctl) 

3. Obtain a valid access token via the DigitalOcean administration dashboard. Export the access token to a environment variable.
```
export DIGITALOCEAN_TOKEN="$YOUR_PERSONAL_ACCESS_TOKEN"
```

3. Configure the `doctl` CLI by running the following:
```
doctl auth init --access-token $DIGITALOCEAN_TOKEN
```

4. Add an SSH key via the DigitalOcean administration dashboard for Terraform: Accounts -> Settings -> Security -> SSH Keys


## Setup and Usage

```
# Create a workspace
terraform workspace new ica-dev-eu-central

# Initialize Terraform to install module dependencies and prepare local terraform state
terraform -chdir=environments/zone-a init

# Prepare a dry-run execution plan, detailing the resources to be deployed
terraform -chdir=environments/zone-a plan -var "do_token=${DIGITALOCEAN_TOKEN}" -var "pvt_key=$HOME/.ssh/tf/id_rsa" -var "mnemonic=${MNEMONIC_A}" -out zone-a.out

# Apply the environment execution plan
terraform -chdir=environments/zone-a apply zone-a.out

# Teardown the environment
terraform -chdir=environments/zone-a plan --destroy -var "do_token=${DIGITALOCEAN_TOKEN}" -var "pvt_key=$HOME/.ssh/tf/id_rsa" -var "mnemonic=${MNEMONIC_A}" -out zone-a.out

# Apply the environment execution plan
terraform -chdir=environments/zone-a apply zone-a.out

# Initialize Terraform to install module dependencies and prepare local terraform state
terraform -chdir=environments/zone-b init

# Prepare a dry-run execution plan, detailing the resources to be deployed
terraform -chdir=environments/zone-b plan -var "do_token=${DIGITALOCEAN_TOKEN}" -var "pvt_key=$HOME/.ssh/tf/id_rsa" -var "mnemonic=${MNEMONIC_B}" -out zone-b.out

# Apply the environment execution plan
terraform -chdir=environments/zone-b apply zone-b.out

# Teardown the environment
terraform -chdir=environments/zone-b plan --destroy -var "do_token=${DIGITALOCEAN_TOKEN}" -var "pvt_key=$HOME/.ssh/tf/id_rsa" -var "mnemonic=${MNEMONIC_B}" -out zone-b.out

# Apply the environment execution plan
terraform -chdir=environments/zone-b apply zone-b.out
```

## WIP notes
create ssh key and add to digital ocean

ssh-keygen
cat ~/.ssh/tf/id_rsa

If you added a passphrase to your ssh key
Error: Failed to parse ssh private key: ssh: this private key is passphrase protected

Ensure ssh-agent is running:
ssh-agent

Open the ssh config file, create if not exists
vim ~/.ssh/config

Add the following:
Host *
  AddKeysToAgent yes
  UseKeychain yes
  IdentityFile ~/.ssh/id_rsa

Host *
  AddKeysToAgent yes
  UseKeychain yes
  IdentityFile ~/.ssh/tf/id_rsa

Now run:
ssh-add -K ~/.ssh/id_rsa
ssh-add -K ~/.ssh/tf/id_rsa

SSH into node:
ssh -i ~/.ssh/tf/id_rsa root@ipv4_addr

## Hermes - Observations

- hermes start will only relay packets over a single channel, must use hermes start-multi
- Start the relayer in multi-paths mode. Handles packet relaying across all open channels between all chains in the config.

# icad start --pruning=nothing > ~/.ica/ica_0.log 2>&1 &
# icad start --pruning=nothing > ~/.ica/ica_1.log 2>&1 &

export MNEMONIC_A="alley afraid soup fall idea toss can goose become valve initial strong forward bright dish figure check leopard decide warfare hub unusual join cart"

export MNEMONIC_B="record gift you once hip style during joke field prize dust unique length more pencil transfer quit train device arrive energy sort steak upset"

## Hermes Relayer

### Restore Keys

```
hermes keys restore ica-0 -m "alley afraid soup fall idea toss can goose become valve initial strong forward bright dish figure check leopard decide warfare hub unusual join cart"

hermes keys restore ica-1 -m "record gift you once hip style during joke field prize dust unique length more pencil transfer quit train device arrive energy sort steak upset"
```

### Configuring Clients

- Create a client where - (dest: ica-0) (src: ica-1)
```
hermes tx raw create-client ica-0 ica-1
hermes query client state ica-0 07-tendermint-0
```

- Create a second client where - (dest: ica-1) (src: ica-0)
```
hermes tx raw create-client ica-1 ica-0
hermes query client state ica-1 07-tendermint-0
```

- Update client states by sending an update-client transaction:
```
hermes tx raw update-client ica-0 07-tendermint-0
hermes tx raw update-client ica-1 07-tendermint-0
```

### Connection Handshake

- Initialise a new ibc connection
```
# conn-init
hermes tx raw conn-init ica-0 ica-1 07-tendermint-0 07-tendermint-0

# conn-try
hermes tx raw conn-try ica-1 ica-0 07-tendermint-0 07-tendermint-0 -s connection-0

# conn-ack
hermes tx raw conn-ack ica-0 ica-1 07-tendermint-0 07-tendermint-0 -d connection-0 -s connection-0

# conn-confirm
hermes tx raw conn-confirm ica-1 ica-0 07-tendermint-0 07-tendermint-0 -d connection-0 -s connection-0

# query the connection
hermes query connection end ica-0 connection-0
hermes query connection end ica-1 connection-0
```

- Create an unordered ics-20 transfer channel
```
# chan-open-init
hermes tx raw chan-open-init ica-0 ica-1 connection-0 transfer transfer -o UNORDERED

# chan-open-try
hermes tx raw chan-open-try ica-1 ica-0 connection-0 transfer transfer -s channel-0

# chan-open-ack
hermes tx raw chan-open-ack ica-0 ica-1 connection-0 transfer transfer -d channel-0 -s channel-0

# chan-open-confirm
hermes tx raw chan-open-confirm ica-1 ica-0 connection-0 transfer transfer -d channel-0 -s channel-0

# query the channel
hermes query channel end ica-0 transfer channel-0
hermes query channel end ica-1 transfer channel-0
```

- Create an ordered ics-27 ibcaccount channel
```
# chan-open-init
hermes tx raw chan-open-init ica-0 ica-1 connection-0 ibcaccount ibcaccount -o ORDERED

# chan-open-try
hermes tx raw chan-open-try ica-1 ica-0 connection-0 ibcaccount ibcaccount -s channel-1

# chan-open-ack
hermes tx raw chan-open-ack ica-0 ica-1 connection-0 ibcaccount ibcaccount -d channel-1 -s channel-1

# chan-open-confirm
hermes tx raw chan-open-confirm ica-1 ica-0 connection-0 ibcaccount ibcaccount -d channel-1 -s channel-1

# query the channel
hermes query channel end ica-0 ibcaccount channel-1
hermes query channel end ica-0 ibcaccount channel-1

# start the relayer
hermes start-multi
```

--- 

# Register an IBC Account on chain test-2 
icad tx intertx register --from validator --source-port ibcaccount --source-channel channel-1 --chain-id ica-0 --gas 90000 --node tcp://159.65.121.50:26657

hermes tx raw packet-recv ica-0 ica-1 ibcaccount channel-1
hermes tx raw packet-ack ica-1 ica-0 ibcaccount channel-1

icad query intertx ibcaccount cosmos1mjk79fjjgpplak5wq838w0yd982gzkyfrk07am  ibcaccount channel-0 --node tcp://134.209.238.196:26657

terraform -chdir=environments/zone-b apply zone-b.out