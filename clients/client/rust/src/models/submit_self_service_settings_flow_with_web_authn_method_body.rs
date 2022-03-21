/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.144
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody {
    /// CSRFToken is the anti-CSRF token
    #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
    pub csrf_token: Option<String>,
    /// Method  Should be set to \"webauthn\" when trying to add, update, or remove a webAuthn pairing.
    #[serde(rename = "method")]
    pub method: String,
    /// Register a WebAuthn Security Key  It is expected that the JSON returned by the WebAuthn registration process is included here.
    #[serde(rename = "webauthn_register", skip_serializing_if = "Option::is_none")]
    pub webauthn_register: Option<String>,
    /// Name of the WebAuthn Security Key to be Added  A human-readable name for the security key which will be added.
    #[serde(rename = "webauthn_register_displayname", skip_serializing_if = "Option::is_none")]
    pub webauthn_register_displayname: Option<String>,
    /// Remove a WebAuthn Security Key  This must contain the ID of the WebAuthN connection.
    #[serde(rename = "webauthn_remove", skip_serializing_if = "Option::is_none")]
    pub webauthn_remove: Option<String>,
}

impl SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody {
    pub fn new(method: String) -> SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody {
        SubmitSelfServiceSettingsFlowWithWebAuthnMethodBody {
            csrf_token: None,
            method,
            webauthn_register: None,
            webauthn_register_displayname: None,
            webauthn_remove: None,
        }
    }
}


