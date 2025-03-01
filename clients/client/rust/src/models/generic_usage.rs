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
pub struct GenericUsage {
    /// AdditionalPrice is the price per-unit in cent exceeding IncludedUsage. A price of 0 means that no other items can be consumed.
    #[serde(rename = "additional_price")]
    pub additional_price: i64,
    /// IncludedUsage is the number of included items.
    #[serde(rename = "included_usage")]
    pub included_usage: i64,
}


impl GenericUsage {
    pub fn new(additional_price: i64, included_usage: i64) -> GenericUsage {
        GenericUsage {
                additional_price,
                included_usage,
        }
    }
}


