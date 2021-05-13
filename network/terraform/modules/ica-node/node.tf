resource "digitalocean_droplet" "ica_node" {
    image   = var.image
    name    = var.chain_id
    region  = var.region
    size    = var.size
    
    private_networking = true
    
    ssh_keys = [
      data.digitalocean_ssh_key.terraform.id
    ]

    connection {
      host        = self.ipv4_address
      user        = "root"
      type        = "ssh"
      private_key = file(var.pvt_key)
      timeout     = "2m"
    }

    provisioner "remote-exec" {
      # TODO: Much of this boilerplate should be moved to scripts
      # Terraform documents remote-exec as last-resort
      inline = [
        "export PATH=$PATH:/usr/bin",
        "sudo apt-get update",
        "sudo add-apt-repository ppa:longsleep/golang-backports -y",
        "sudo apt install golang-go -y",
        "export PATH=$PATH:$(go env GOPATH)/bin",
        "git clone https://github.com/cosmos/interchain-accounts.git && cd interchain-accounts",
        "make install && which icad",
        "icad init ica --chain-id ${var.chain_id}",
        "echo '${var.mnemonic}' | icad keys add validator --recover --keyring-backend test",
        "icad add-genesis-account $(icad keys show validator --keyring-backend test -a) 100000000000stake",
        "icad gentx validator 7000000000stake --chain-id ${var.chain_id} --keyring-backend test",
        "icad collect-gentxs",
        "sed -i -e 's#\"tcp://127.0.0.1:26657\"#\"tcp://0.0.0.0:'\"26657\"'\"#g' ~/.ica/config/config.toml",
        "icad start --pruning=nothing > ~/.ica/chain.log 2>&1 &"
      ]
    }
}

# resource "digitalocean_loadbalancer" "public" {
#   name   = "loadbalancer-1"
#   region = "fra1"

#   forwarding_rule {
#     entry_port     = 80
#     entry_protocol = "http"

#     target_port     = 80
#     target_protocol = "http"
#   }

#   forwarding_rule {
#     entry_port     = 26657
#     entry_protocol = "tcp"

#     target_port     = 26657
#     target_protocol = "tcp"
#   }

#   forwarding_rule {
#     entry_port     = 9090
#     entry_protocol = "http2"

#     target_port     = 9090
#     target_protocol = "http2"

#     tls_passthrough = true
#   }
  
#   healthcheck {
#     port     = 22
#     protocol = "tcp"
#   }

#   droplet_ids = [digitalocean_droplet.ica-node.id]
# }