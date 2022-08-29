/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.2.0-alpha.22
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct OAuth2ConsentSessionExpiresAt {
    #[serde(rename = "access_token", skip_serializing_if = "Option::is_none")]
    pub access_token: Option<String>,
    #[serde(rename = "authorize_code", skip_serializing_if = "Option::is_none")]
    pub authorize_code: Option<String>,
    #[serde(rename = "id_token", skip_serializing_if = "Option::is_none")]
    pub id_token: Option<String>,
    #[serde(rename = "par_context", skip_serializing_if = "Option::is_none")]
    pub par_context: Option<String>,
    #[serde(rename = "refresh_token", skip_serializing_if = "Option::is_none")]
    pub refresh_token: Option<String>,
}

impl Default for OAuth2ConsentSessionExpiresAt {
    fn default() -> Self {
        Self::new()
    }
}

impl OAuth2ConsentSessionExpiresAt {
    pub fn new() -> OAuth2ConsentSessionExpiresAt {
        OAuth2ConsentSessionExpiresAt {
                access_token: None,
                authorize_code: None,
                id_token: None,
                par_context: None,
                refresh_token: None,
        }
    }
}


