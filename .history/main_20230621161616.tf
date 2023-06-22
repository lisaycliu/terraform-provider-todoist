terraform {
  required_providers {
    todoist = {
      version = "~> 0.3.1"
      "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
      source  = "todoist"
    }
  }
}