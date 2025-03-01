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
pub struct NormalizedProjectRevisionTokenizerTemplate {
    /// Claims mapper URL
    #[serde(rename = "claims_mapper_url", skip_serializing_if = "Option::is_none")]
    pub claims_mapper_url: Option<String>,
    /// The Project's Revision Creation Date
    #[serde(rename = "created_at", skip_serializing_if = "Option::is_none")]
    pub created_at: Option<String>,
    /// The revision ID.
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<String>,
    /// JSON Web Key URL
    #[serde(rename = "jwks_url", skip_serializing_if = "Option::is_none")]
    pub jwks_url: Option<String>,
    /// The unique key of the template
    #[serde(rename = "key", skip_serializing_if = "Option::is_none")]
    pub key: Option<String>,
    /// The Revision's ID this schema belongs to
    #[serde(rename = "project_revision_id", skip_serializing_if = "Option::is_none")]
    pub project_revision_id: Option<String>,
    /// Token time to live
    #[serde(rename = "ttl", skip_serializing_if = "Option::is_none")]
    pub ttl: Option<String>,
    /// Last Time Project's Revision was Updated
    #[serde(rename = "updated_at", skip_serializing_if = "Option::is_none")]
    pub updated_at: Option<String>,
}

impl Default for NormalizedProjectRevisionTokenizerTemplate {
    fn default() -> Self {
        Self::new()
    }
}

impl NormalizedProjectRevisionTokenizerTemplate {
    pub fn new() -> NormalizedProjectRevisionTokenizerTemplate {
        NormalizedProjectRevisionTokenizerTemplate {
                claims_mapper_url: None,
                created_at: None,
                id: None,
                jwks_url: None,
                key: None,
                project_revision_id: None,
                ttl: None,
                updated_at: None,
        }
    }
}


