provider "archive" {}

data archive_file website {
    type = "zip"
    source_dir = "../../web/dist"
    output_path = "../../tmp/dist.zip"
}

resource aws_s3_bucket_object website {
    bucket = "${var.bucket_name}"
    key = "web/dist.zip"
    source = "${data.archive_file.website.output_path}" 
    etag = "${md5(file("${data.archive_file.website.output_path}"))}"
}
