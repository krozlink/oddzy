resource "aws_cognito_user_pool" "users" {
    name = "${var.application_name}-${var.application_stage}"

    alias_attributes = ["email", "preferred_username"]
    auto_verified_attributes = ["email"]
    
    lifecycle {
        prevent_destroy = true
    }

    admin_create_user_config {
        allow_admin_create_user_only = false
        unused_account_validity_days = 7
    }

    password_policy {
        minimum_length = 8
        require_lowercase = false
        require_numbers = false
        require_uppercase = false
        require_symbols = false
    }

    lambda_config { 
        pre_sign_up = "${aws_lambda_function.pre_signup.arn}"
    }

    verification_message_template {
        default_email_option = "CONFIRM_WITH_LINK"
        email_message_by_link = "Please click the link below to verify your email address. {##Verify Email##}"
        email_subject_by_link = "Oddzy - Your verification link"
    }

    schema = [
         {
             attribute_data_type = "String"
             developer_only_attribute = false
             mutable = true
             name = "email"
             required = true
 
            string_attribute_constraints {
                min_length = 1
                max_length = 255
            }
         },
        {
            attribute_data_type = "String"
            developer_only_attribute = false
            mutable = true
            name = "given_name"
            required = true
 
            string_attribute_constraints {
                min_length = 1
                max_length = 255
            }
        },
        {
            attribute_data_type = "String"
            developer_only_attribute = false
            mutable = true
            name = "family_name"
            required = true
 
            string_attribute_constraints {
                min_length = 1
                max_length = 255
            }
        },
        {
            attribute_data_type = "String"
            developer_only_attribute = false
            mutable = true
            name = "preferred_username"
            required = false
 
            string_attribute_constraints {
                min_length = 1
                max_length = 255
            }
        },
        {
            attribute_data_type = "String"
            developer_only_attribute = false
            mutable = true
            name = "phone_number"
            required = true
 
            string_attribute_constraints {
                min_length = 1
                max_length = 255
            }
        },
        {
            attribute_data_type = "String"
            developer_only_attribute = false
            mutable = true
            name = "address"
            required = true
 
            string_attribute_constraints {
                min_length = 1
                max_length = 255
            }
        },
        {
            attribute_data_type = "String"
            mutable = true
            name = "birthdate"
            required = true
 
            string_attribute_constraints {
                min_length = 1
                max_length = 255
            }
        }
    ]
}

resource "aws_cognito_user_pool_client" "users" {
    name = "${var.application_name}-${var.application_stage}-users"
    user_pool_id = "${aws_cognito_user_pool.users.id}"
    refresh_token_validity = 30
    generate_secret = false

    lifecycle {
        prevent_destroy = true
    }
}

resource "aws_cognito_user_pool_domain" "users" {
    domain = "oddzy"
    user_pool_id = "${aws_cognito_user_pool.users.id}"
}

resource "aws_cognito_identity_pool" "users" {
    identity_pool_name = "${var.application_name} ${var.application_stage}"
    allow_unauthenticated_identities = true,

    cognito_identity_providers {
        client_id = "${aws_cognito_user_pool_client.users.id}"
        provider_name = "${aws_cognito_user_pool.users.endpoint}"
        server_side_token_check = false
    }
}


resource "aws_iam_role" "authenticated" {
  name = "cognito_authenticated"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Federated": "cognito-identity.amazonaws.com"
      },
      "Action": "sts:AssumeRoleWithWebIdentity",
      "Condition": {
        "StringEquals": {
          "cognito-identity.amazonaws.com:aud": "${aws_cognito_identity_pool.users.id}"
        },
        "ForAnyValue:StringLike": {
          "cognito-identity.amazonaws.com:amr": "authenticated"
        }
      }
    }
  ]
}
EOF
}

resource "aws_iam_role" "unauthenticated" {
  name = "cognito_unauthenticated"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
       "Principal": {
         "Federated": "cognito-identity.amazonaws.com"
      },
      "Action": "sts:AssumeRoleWithWebIdentity",
      "Condition": {
        "StringEquals": {
          "cognito-identity.amazonaws.com:aud": "${aws_cognito_identity_pool.users.id}"
        },
        "ForAnyValue:StringLike": {
          "cognito-identity.amazonaws.com:amr": "unauthenticated"
        }
      }
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "authenticated" {
  name = "authenticated_policy"
  role = "${aws_iam_role.authenticated.id}"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "firehose:PutRecordBatch"
      ],
      "Resource": [
        "arn:aws:firehose:${var.region}:${data.aws_caller_identity.current.account_id}:deliverystream/${var.application_name}-${var.application_stage}-tracking"
      ]
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "unauthenticated" {
  name = "unauthenticated_policy"
  role = "${aws_iam_role.unauthenticated.id}"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "firehose:PutRecordBatch"
      ],
      "Resource": [
        "arn:aws:firehose:${var.region}:${data.aws_caller_identity.current.account_id}:deliverystream/${var.application_name}-${var.application_stage}-tracking"
      ]
    }
  ]
}
EOF
}

resource "aws_cognito_identity_pool_roles_attachment" "main" {
  identity_pool_id = "${aws_cognito_identity_pool.users.id}"

  roles {
    "authenticated" = "${aws_iam_role.authenticated.arn}"
    "unauthenticated" = "${aws_iam_role.unauthenticated.arn}"
  }
}