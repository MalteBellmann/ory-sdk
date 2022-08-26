/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.2.0-alpha.20
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct TrustedOAuth2JwtGrantJsonWebKey {
    /// The \"key_id\" is key unique identifier (same as kid header in jws/jwt).
    #[serde(rename = "kid", skip_serializing_if = "Option::is_none")]
    pub kid: Option<String>,
    /// The \"set\" is basically a name for a group(set) of keys. Will be the same as \"issuer\" in grant.
    #[serde(rename = "set", skip_serializing_if = "Option::is_none")]
    pub set: Option<String>,
}

impl Default for TrustedOAuth2JwtGrantJsonWebKey {
    fn default() -> Self {
        Self::new()
    }
}

impl TrustedOAuth2JwtGrantJsonWebKey {
    pub fn new() -> TrustedOAuth2JwtGrantJsonWebKey {
        TrustedOAuth2JwtGrantJsonWebKey {
                kid: None,
                set: None,
        }
    }
}


