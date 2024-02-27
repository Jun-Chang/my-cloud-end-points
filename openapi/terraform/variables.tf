variable "gcp_project_id" {
  type = string
}

locals {
  region = "asia-northeast1"
  envs   = ["dev", "stg", "prd"]
}
