/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.2.8
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// GetManagedIdentitySchemaLocation : Ory Identity Schema Location



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct GetManagedIdentitySchemaLocation {
    #[serde(rename = "location", skip_serializing_if = "Option::is_none")]
    pub location: Option<String>,
}

impl Default for GetManagedIdentitySchemaLocation {
    fn default() -> Self {
        Self::new()
    }
}

impl GetManagedIdentitySchemaLocation {
    /// Ory Identity Schema Location
    pub fn new() -> GetManagedIdentitySchemaLocation {
        GetManagedIdentitySchemaLocation {
                location: None,
        }
    }
}


