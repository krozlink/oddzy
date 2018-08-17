provider "archive" {}

// Creates a zip from the latest build of the website
data archive_file website {
    type = "zip"
    source_dir = "../../web/dist"
    output_path = "../../tmp/dist.zip"
}

// Uploads the website to S3 if it has changed
resource aws_s3_bucket_object website {
    bucket = "${var.bucket_name}"
    key = "web/dist.zip"
    source = "${data.archive_file.website.output_path}" 
    etag = "${md5(file("${data.archive_file.website.output_path}"))}"
}

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
              "rm -rf /mnt/efs/website/oddzy/*",
              "aws s3 sync s3://oddzy/web/dist /mnt/efs/website/oddzy --delete"
            ]
         }
    }]
  }
DOC
}