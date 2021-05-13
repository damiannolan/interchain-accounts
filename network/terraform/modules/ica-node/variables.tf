variable "do_token" {
  description = "A valid DigitalOcean API token"
}

variable "pvt_key" {
  description = "Private ssh key path location"
}

variable "image" {
  description = "The virtual machine image name"
}

variable "chain_id" {
  description = "The chain id, also used a droplet name"
}

variable "mnemonic" {
    description = "Seed phrase used to recover existing key"
}

variable "region" {
  description = "The DigitalOcean region in which the droplet will be deployed"
}

variable "size" {
  description = "The CPU and RAM resources assigned to the droplet"
}
