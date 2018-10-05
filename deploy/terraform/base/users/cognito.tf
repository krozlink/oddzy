resource "aws_cognito_user_pool" "users" {
    name = "${var.application_name}-${var.application_stage}"

    alias_attributes = ["email", "preferred_username"]
    auto_verified_attributes = ["email"]

    admin_create_user_config {
        allow_admin_create_user_only = false
        unused_account_validity_days = 7
    }

    password_policy {
        minimum_length = 8
        require_lowercase = true
        require_numbers = true
        require_uppercase = true
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
        },
        {
            attribute_data_type = "String"
            developer_only_attribute = false
            mutable = true
            name = "given_name"
            required = true
        },
        {
            attribute_data_type = "String"
            developer_only_attribute = false
            mutable = true
            name = "family_name"
            required = true
        },
        {
            attribute_data_type = "String"
            developer_only_attribute = false
            mutable = true
            name = "preferred_username"
            required = false
        },
        {
            attribute_data_type = "String"
            developer_only_attribute = false
            mutable = true
            name = "phone_number"
            required = true
        },
        {
            attribute_data_type = "String"
            developer_only_attribute = false
            mutable = true
            name = "gender"
            required = true
        },
        {
            attribute_data_type = "String"
            developer_only_attribute = false
            mutable = true
            name = "address"
            required = true
        },
        {
            attribute_data_type = "String"
            mutable = true
            name = "birthdate"
            required = true
        }
    ]
}