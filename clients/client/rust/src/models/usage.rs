/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.6.1
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct Usage {
    #[serde(rename = "GenericUsage", skip_serializing_if = "Option::is_none")]
    pub generic_usage: Option<Box<crate::models::GenericUsage>>,
}

impl Default for Usage {
    fn default() -> Self {
        Self::new()
    }
}

impl Usage {
    pub fn new() -> Usage {
        Usage {
                generic_usage: None,
        }
    }
}


