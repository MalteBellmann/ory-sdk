/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.1.13
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct Pagination {
    /// Pagination Page  This value is currently an integer, but it is not sequential. The value is not the page number, but a reference. The next page can be any number and some numbers might return an empty list.  For example, page 2 might not follow after page 1. And even if page 3 and 5 exist, but page 4 might not exist.
    #[serde(rename = "page", skip_serializing_if = "Option::is_none")]
    pub page: Option<i64>,
    /// Items per page  This is the number of items per page to return. For details on pagination please head over to the [pagination documentation](https://www.ory.sh/docs/ecosystem/api-design#pagination).
    #[serde(rename = "page_size", skip_serializing_if = "Option::is_none")]
    pub page_size: Option<i64>,
    /// Next Page Token  The next page token. For details on pagination please head over to the [pagination documentation](https://www.ory.sh/docs/ecosystem/api-design#pagination).
    #[serde(rename = "page_token", skip_serializing_if = "Option::is_none")]
    pub page_token: Option<String>,
    /// Items per Page  This is the number of items per page.
    #[serde(rename = "per_page", skip_serializing_if = "Option::is_none")]
    pub per_page: Option<i64>,
}

impl Default for Pagination {
    fn default() -> Self {
        Self::new()
    }
}

impl Pagination {
    pub fn new() -> Pagination {
        Pagination {
                page: None,
                page_size: None,
                page_token: None,
                per_page: None,
        }
    }
}


