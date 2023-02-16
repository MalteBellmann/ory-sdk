/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.1.13
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// IdentityWithCredentialsOidc : Create Identity and Import Social Sign In Credentials



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct IdentityWithCredentialsOidc {
    #[serde(rename = "config", skip_serializing_if = "Option::is_none")]
    pub config: Option<Box<crate::models::IdentityWithCredentialsOidcConfig>>,
}

impl Default for IdentityWithCredentialsOidc {
    fn default() -> Self {
        Self::new()
    }
}

impl IdentityWithCredentialsOidc {
    /// Create Identity and Import Social Sign In Credentials
    pub fn new() -> IdentityWithCredentialsOidc {
        IdentityWithCredentialsOidc {
                config: None,
        }
    }
}


