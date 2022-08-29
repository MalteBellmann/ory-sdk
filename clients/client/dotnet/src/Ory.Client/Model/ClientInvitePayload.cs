/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.2.0-alpha.22
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
    /// ClientInvitePayload
    /// </summary>
    [DataContract(Name = "InvitePayload")]
    public partial class ClientInvitePayload : IEquatable<ClientInvitePayload>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientInvitePayload" /> class.
        /// </summary>
        /// <param name="inviteeEmail">inviteeEmail.</param>
        public ClientInvitePayload(string inviteeEmail = default(string))
        {
            this.InviteeEmail = inviteeEmail;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Gets or Sets InviteeEmail
        /// </summary>
        [DataMember(Name = "invitee_email", EmitDefaultValue = false)]
        public string InviteeEmail { get; set; }

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
            StringBuilder sb = new StringBuilder();
            sb.Append("class ClientInvitePayload {\n");
            sb.Append("  InviteeEmail: ").Append(InviteeEmail).Append("\n");
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
            return this.Equals(input as ClientInvitePayload);
        }

        /// <summary>
        /// Returns true if ClientInvitePayload instances are equal
        /// </summary>
        /// <param name="input">Instance of ClientInvitePayload to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ClientInvitePayload input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.InviteeEmail == input.InviteeEmail ||
                    (this.InviteeEmail != null &&
                    this.InviteeEmail.Equals(input.InviteeEmail))
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
                if (this.InviteeEmail != null)
                {
                    hashCode = (hashCode * 59) + this.InviteeEmail.GetHashCode();
                }
                if (this.AdditionalProperties != null)
                {
                    hashCode = (hashCode * 59) + this.AdditionalProperties.GetHashCode();
                }
                return hashCode;
            }
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        public IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}
