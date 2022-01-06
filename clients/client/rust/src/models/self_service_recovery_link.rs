/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.39
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct SelfServiceRecoveryLink {
    /// Recovery Link Expires At  The timestamp when the recovery link expires.
    #[serde(rename = "expires_at", skip_serializing_if = "Option::is_none")]
    pub expires_at: Option<String>,
    /// Recovery Link  This link can be used to recover the account.
    #[serde(rename = "recovery_link")]
    pub recovery_link: String,
}

impl SelfServiceRecoveryLink {
    pub fn new(recovery_link: String) -> SelfServiceRecoveryLink {
        SelfServiceRecoveryLink {
            expires_at: None,
            recovery_link,
        }
    }
}


