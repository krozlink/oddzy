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
    allow_unauthenticated_identities = false

    cognito_identity_providers {
        client_id = "${aws_cognito_user_pool_client.users.id}"
        provider_name = "${aws_cognito_user_pool.users.endpoint}"
        server_side_token_check = false
    }
}