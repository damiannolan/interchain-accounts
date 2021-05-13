variable "do_token" {
    description = "A valid DigitalOcean API token"
}

variable "mnemonic" {
    description = "Seed phrase used to recover existing key"
}

variable "pvt_key" {
    description = "Private ssh key path location"
}

module "ica-node-1" {
    source      = "../../modules/ica-node"
    
    image       = "ubuntu-18-04-x64"
    chain_id    = "ica-1"
    region      = "fra1"
    size        = "s-1vcpu-1gb"

    do_token    = var.do_token
    mnemonic    = var.mnemonic
    pvt_key     = var.pvt_key
}
