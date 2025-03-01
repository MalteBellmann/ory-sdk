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
pub struct CreateVerifiableCredentialRequestBody {
    #[serde(rename = "format", skip_serializing_if = "Option::is_none")]
    pub format: Option<String>,
    #[serde(rename = "proof", skip_serializing_if = "Option::is_none")]
    pub proof: Option<Box<crate::models::VerifiableCredentialProof>>,
    #[serde(rename = "types", skip_serializing_if = "Option::is_none")]
    pub types: Option<Vec<String>>,
}

impl Default for CreateVerifiableCredentialRequestBody {
    fn default() -> Self {
        Self::new()
    }
}

impl CreateVerifiableCredentialRequestBody {
    pub fn new() -> CreateVerifiableCredentialRequestBody {
        CreateVerifiableCredentialRequestBody {
                format: None,
                proof: None,
                types: None,
        }
    }
}


