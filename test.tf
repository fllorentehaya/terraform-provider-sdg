terraform {
  required_providers {
    sdg = {
      source  = "fer.com/fll/sdg"
      version = "0.1.2"
    }
  }
}

  provider "sdg" {}

resource "sdg_batch_template" "my-server" {
    location = "West Europe"
    client_id = "7101e195-2f4d-44b3-a29a-f2acf4e047ca"
    client_secret = "xWMT0iRnwzCDPUY9GJzhNdVODTFDTIwg.5"
    subscription = "14f69348-8506-41f1-ac5f-830f0c4ffceb"
    tenant = "cd149d37-c0df-4919-a5de-b2294e066d62"
    entorno="produccion"
}

