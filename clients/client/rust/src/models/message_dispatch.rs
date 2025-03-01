/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.6.1
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// MessageDispatch : MessageDispatch represents an attempt of sending a courier message It contains the status of the attempt (failed or successful) and the error if any occured



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct MessageDispatch {
    /// CreatedAt is a helper struct field for gobuffalo.pop.
    #[serde(rename = "created_at")]
    pub created_at: String,
    #[serde(rename = "error", skip_serializing_if = "Option::is_none")]
    pub error: Option<serde_json::Value>,
    /// The ID of this message dispatch
    #[serde(rename = "id")]
    pub id: String,
    /// The ID of the message being dispatched
    #[serde(rename = "message_id")]
    pub message_id: String,
    /// The status of this dispatch Either \"failed\" or \"success\" failed CourierMessageDispatchStatusFailed success CourierMessageDispatchStatusSuccess
    #[serde(rename = "status")]
    pub status: StatusEnum,
    /// UpdatedAt is a helper struct field for gobuffalo.pop.
    #[serde(rename = "updated_at")]
    pub updated_at: String,
}


impl MessageDispatch {
    /// MessageDispatch represents an attempt of sending a courier message It contains the status of the attempt (failed or successful) and the error if any occured
    pub fn new(created_at: String, id: String, message_id: String, status: StatusEnum, updated_at: String) -> MessageDispatch {
        MessageDispatch {
                created_at,
                error: None,
                id,
                message_id,
                status,
                updated_at,
        }
    }
}

/// The status of this dispatch Either \"failed\" or \"success\" failed CourierMessageDispatchStatusFailed success CourierMessageDispatchStatusSuccess
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum StatusEnum {
    #[serde(rename = "failed")]
    Failed,
    #[serde(rename = "success")]
    Success,
}

