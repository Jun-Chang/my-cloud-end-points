resource "google_storage_bucket" "storage" {
  project                     = var.gcp_project_id
  location                    = local.region
  name                        = "my-cloud-deploy"
  uniform_bucket_level_access = true
  public_access_prevention    = "enforced"

  lifecycle_rule {
    action {
      type = "Delete"
    }
    condition {
      age = 7
    }
  }
}

resource "google_clouddeploy_target" "target" {
  for_each = toset(local.envs)

  project          = var.gcp_project_id
  location         = local.region
  name             = "my-cloud-deploy-${each.value}"
  require_approval = false

  execution_configs {
    usages           = ["RENDER", "DEPLOY"]
    artifact_storage = "gs://${google_storage_bucket.storage.name}/artifacts"
  }

  run {
    location = "projects/${var.gcp_project_id}/locations/${local.region}"
  }

  deploy_parameters = {
    service_name = "my-cloud-endpoint-backend-${each.value}"
  }
}

resource "google_clouddeploy_delivery_pipeline" "pipeline-app" {
  project     = var.gcp_project_id
  location    = local.region
  name        = "my-cloud-deploy-pipeline"
  description = "Delivery Pipeline for my-cloud-endpoint"

  serial_pipeline {
    stages {
      target_id = google_clouddeploy_target.target["dev"].name
    }

    stages {
      target_id = google_clouddeploy_target.target["stg"].name
    }

    stages {
      target_id = google_clouddeploy_target.target["prd"].name
    }
  }
}
