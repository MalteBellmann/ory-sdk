/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.39
 * Contact: support@ory.sh
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Runtime.Serialization;
using System.Text;
using System.Text.RegularExpressions;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;
using System.ComponentModel.DataAnnotations;
using OpenAPIDateConverter = Ory.Client.Client.OpenAPIDateConverter;

namespace Ory.Client.Model
{
    /// <summary>
    /// ClientProjectVerificationConfig
    /// </summary>
    [DataContract(Name = "projectVerificationConfig")]
    public partial class ClientProjectVerificationConfig : IEquatable<ClientProjectVerificationConfig>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientProjectVerificationConfig" /> class.
        /// </summary>
        /// <param name="enabled">Set to true to enable email / phone verification..</param>
        /// <param name="requireVerifiedAddress">Set to true to only allow users with a verified address to sign in..</param>
        public ClientProjectVerificationConfig(bool enabled = default(bool), bool requireVerifiedAddress = default(bool))
        {
            this.Enabled = enabled;
            this.RequireVerifiedAddress = requireVerifiedAddress;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Set to true to enable email / phone verification.
        /// </summary>
        /// <value>Set to true to enable email / phone verification.</value>
        [DataMember(Name = "enabled", EmitDefaultValue = true)]
        public bool Enabled { get; set; }

        /// <summary>
        /// Set to true to only allow users with a verified address to sign in.
        /// </summary>
        /// <value>Set to true to only allow users with a verified address to sign in.</value>
        [DataMember(Name = "require_verified_address", EmitDefaultValue = true)]
        public bool RequireVerifiedAddress { get; set; }

        /// <summary>
        /// Gets or Sets additional properties
        /// </summary>
        [JsonExtensionData]
        public IDictionary<string, object> AdditionalProperties { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class ClientProjectVerificationConfig {\n");
            sb.Append("  Enabled: ").Append(Enabled).Append("\n");
            sb.Append("  RequireVerifiedAddress: ").Append(RequireVerifiedAddress).Append("\n");
            sb.Append("  AdditionalProperties: ").Append(AdditionalProperties).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public virtual string ToJson()
        {
            return Newtonsoft.Json.JsonConvert.SerializeObject(this, Newtonsoft.Json.Formatting.Indented);
        }

        /// <summary>
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="input">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object input)
        {
            return this.Equals(input as ClientProjectVerificationConfig);
        }

        /// <summary>
        /// Returns true if ClientProjectVerificationConfig instances are equal
        /// </summary>
        /// <param name="input">Instance of ClientProjectVerificationConfig to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ClientProjectVerificationConfig input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.Enabled == input.Enabled ||
                    this.Enabled.Equals(input.Enabled)
                ) && 
                (
                    this.RequireVerifiedAddress == input.RequireVerifiedAddress ||
                    this.RequireVerifiedAddress.Equals(input.RequireVerifiedAddress)
                )
                && (this.AdditionalProperties.Count == input.AdditionalProperties.Count && !this.AdditionalProperties.Except(input.AdditionalProperties).Any());
        }

        /// <summary>
        /// Gets the hash code
        /// </summary>
        /// <returns>Hash code</returns>
        public override int GetHashCode()
        {
            unchecked // Overflow is fine, just wrap
            {
                int hashCode = 41;
                hashCode = hashCode * 59 + this.Enabled.GetHashCode();
                hashCode = hashCode * 59 + this.RequireVerifiedAddress.GetHashCode();
                if (this.AdditionalProperties != null)
                    hashCode = hashCode * 59 + this.AdditionalProperties.GetHashCode();
                return hashCode;
            }
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> IValidatableObject.Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}
