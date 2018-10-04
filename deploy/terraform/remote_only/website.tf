provider "archive" {}

// AWS SSM Command which removes the existing website and replaces it with the latest version from S3
resource "aws_ssm_document" "update_website" {
  name          = "${var.application_name}-${var.application_stage}-update-website"
  document_type = "Command"

  content = <<DOC
  {
    "schemaVersion": "2.2",
    "description": "Redeploy the website using the latest version from S3",
    "mainSteps": [
    {
        "action":"aws:runShellScript",
         "name":"deployWebsite",
         "inputs":{
            "runCommand":[
              "aws s3 sync s3://oddzy/web/dist /mnt/efs/website/oddzy --delete"
            ]
         }
    }]
  }
DOC
}