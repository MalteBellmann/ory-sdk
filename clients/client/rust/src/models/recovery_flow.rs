/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.6.1
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// RecoveryFlow : This request is used when an identity wants to recover their account.  We recommend reading the [Account Recovery Documentation](../self-service/flows/password-reset-account-recovery)



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct RecoveryFlow {
    /// Active, if set, contains the recovery method that is being used. It is initially not set.
    #[serde(rename = "active", skip_serializing_if = "Option::is_none")]
    pub active: Option<String>,
    /// Contains possible actions that could follow this flow
    #[serde(rename = "continue_with", skip_serializing_if = "Option::is_none")]
    pub continue_with: Option<Vec<crate::models::ContinueWith>>,
    /// ExpiresAt is the time (UTC) when the request expires. If the user still wishes to update the setting, a new request has to be initiated.
    #[serde(rename = "expires_at")]
    pub expires_at: String,
    /// ID represents the request's unique ID. When performing the recovery flow, this represents the id in the recovery ui's query parameter: http://<selfservice.flows.recovery.ui_url>?request=<id>
    #[serde(rename = "id")]
    pub id: String,
    /// IssuedAt is the time (UTC) when the request occurred.
    #[serde(rename = "issued_at")]
    pub issued_at: String,
    /// RequestURL is the initial URL that was requested from Ory Kratos. It can be used to forward information contained in the URL's path or query for example.
    #[serde(rename = "request_url")]
    pub request_url: String,
    /// ReturnTo contains the requested return_to URL.
    #[serde(rename = "return_to", skip_serializing_if = "Option::is_none")]
    pub return_to: Option<String>,
    /// State represents the state of this request:  choose_method: ask the user to choose a method (e.g. recover account via email) sent_email: the email has been sent to the user passed_challenge: the request was successful and the recovery challenge was passed.
    #[serde(rename = "state")]
    pub state: Option<serde_json::Value>,
    /// The flow type can either be `api` or `browser`.
    #[serde(rename = "type")]
    pub _type: String,
    #[serde(rename = "ui")]
    pub ui: Box<crate::models::UiContainer>,
}


impl RecoveryFlow {
    /// This request is used when an identity wants to recover their account.  We recommend reading the [Account Recovery Documentation](../self-service/flows/password-reset-account-recovery)
    pub fn new(expires_at: String, id: String, issued_at: String, request_url: String, state: Option<serde_json::Value>, _type: String, ui: crate::models::UiContainer) -> RecoveryFlow {
        RecoveryFlow {
                active: None,
                continue_with: None,
                expires_at,
                id,
                issued_at,
                request_url,
                return_to: None,
                state,
                _type,
                ui: Box::new(ui),
        }
    }
}


