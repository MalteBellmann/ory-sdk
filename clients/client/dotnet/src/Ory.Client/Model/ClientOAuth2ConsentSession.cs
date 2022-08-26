/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.2.0-alpha.20
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
    /// ClientOAuth2ConsentSession
    /// </summary>
    [DataContract(Name = "oAuth2ConsentSession")]
    public partial class ClientOAuth2ConsentSession : IEquatable<ClientOAuth2ConsentSession>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientOAuth2ConsentSession" /> class.
        /// </summary>
        /// <param name="allowedTopLevelClaims">allowedTopLevelClaims.</param>
        /// <param name="clientId">clientId.</param>
        /// <param name="consentChallenge">consentChallenge.</param>
        /// <param name="excludeNotBeforeClaim">excludeNotBeforeClaim.</param>
        /// <param name="expiresAt">expiresAt.</param>
        /// <param name="extra">extra.</param>
        /// <param name="headers">headers.</param>
        /// <param name="idTokenClaims">idTokenClaims.</param>
        /// <param name="kid">kid.</param>
        /// <param name="subject">subject.</param>
        /// <param name="username">username.</param>
        public ClientOAuth2ConsentSession(List<string> allowedTopLevelClaims = default(List<string>), string clientId = default(string), string consentChallenge = default(string), bool excludeNotBeforeClaim = default(bool), ClientOAuth2ConsentSessionExpiresAt expiresAt = default(ClientOAuth2ConsentSessionExpiresAt), Dictionary<string, Object> extra = default(Dictionary<string, Object>), ClientHeaders headers = default(ClientHeaders), ClientIDTokenClaims idTokenClaims = default(ClientIDTokenClaims), string kid = default(string), string subject = default(string), string username = default(string))
        {
            this.AllowedTopLevelClaims = allowedTopLevelClaims;
            this.ClientId = clientId;
            this.ConsentChallenge = consentChallenge;
            this.ExcludeNotBeforeClaim = excludeNotBeforeClaim;
            this.ExpiresAt = expiresAt;
            this.Extra = extra;
            this.Headers = headers;
            this.IdTokenClaims = idTokenClaims;
            this.Kid = kid;
            this.Subject = subject;
            this.Username = username;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Gets or Sets AllowedTopLevelClaims
        /// </summary>
        [DataMember(Name = "allowed_top_level_claims", EmitDefaultValue = false)]
        public List<string> AllowedTopLevelClaims { get; set; }

        /// <summary>
        /// Gets or Sets ClientId
        /// </summary>
        [DataMember(Name = "client_id", EmitDefaultValue = false)]
        public string ClientId { get; set; }

        /// <summary>
        /// Gets or Sets ConsentChallenge
        /// </summary>
        [DataMember(Name = "consent_challenge", EmitDefaultValue = false)]
        public string ConsentChallenge { get; set; }

        /// <summary>
        /// Gets or Sets ExcludeNotBeforeClaim
        /// </summary>
        [DataMember(Name = "exclude_not_before_claim", EmitDefaultValue = true)]
        public bool ExcludeNotBeforeClaim { get; set; }

        /// <summary>
        /// Gets or Sets ExpiresAt
        /// </summary>
        [DataMember(Name = "expires_at", EmitDefaultValue = false)]
        public ClientOAuth2ConsentSessionExpiresAt ExpiresAt { get; set; }

        /// <summary>
        /// Gets or Sets Extra
        /// </summary>
        [DataMember(Name = "extra", EmitDefaultValue = false)]
        public Dictionary<string, Object> Extra { get; set; }

        /// <summary>
        /// Gets or Sets Headers
        /// </summary>
        [DataMember(Name = "headers", EmitDefaultValue = false)]
        public ClientHeaders Headers { get; set; }

        /// <summary>
        /// Gets or Sets IdTokenClaims
        /// </summary>
        [DataMember(Name = "id_token_claims", EmitDefaultValue = false)]
        public ClientIDTokenClaims IdTokenClaims { get; set; }

        /// <summary>
        /// Gets or Sets Kid
        /// </summary>
        [DataMember(Name = "kid", EmitDefaultValue = false)]
        public string Kid { get; set; }

        /// <summary>
        /// Gets or Sets Subject
        /// </summary>
        [DataMember(Name = "subject", EmitDefaultValue = false)]
        public string Subject { get; set; }

        /// <summary>
        /// Gets or Sets Username
        /// </summary>
        [DataMember(Name = "username", EmitDefaultValue = false)]
        public string Username { get; set; }

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
            sb.Append("class ClientOAuth2ConsentSession {\n");
            sb.Append("  AllowedTopLevelClaims: ").Append(AllowedTopLevelClaims).Append("\n");
            sb.Append("  ClientId: ").Append(ClientId).Append("\n");
            sb.Append("  ConsentChallenge: ").Append(ConsentChallenge).Append("\n");
            sb.Append("  ExcludeNotBeforeClaim: ").Append(ExcludeNotBeforeClaim).Append("\n");
            sb.Append("  ExpiresAt: ").Append(ExpiresAt).Append("\n");
            sb.Append("  Extra: ").Append(Extra).Append("\n");
            sb.Append("  Headers: ").Append(Headers).Append("\n");
            sb.Append("  IdTokenClaims: ").Append(IdTokenClaims).Append("\n");
            sb.Append("  Kid: ").Append(Kid).Append("\n");
            sb.Append("  Subject: ").Append(Subject).Append("\n");
            sb.Append("  Username: ").Append(Username).Append("\n");
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
            return this.Equals(input as ClientOAuth2ConsentSession);
        }

        /// <summary>
        /// Returns true if ClientOAuth2ConsentSession instances are equal
        /// </summary>
        /// <param name="input">Instance of ClientOAuth2ConsentSession to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ClientOAuth2ConsentSession input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.AllowedTopLevelClaims == input.AllowedTopLevelClaims ||
                    this.AllowedTopLevelClaims != null &&
                    input.AllowedTopLevelClaims != null &&
                    this.AllowedTopLevelClaims.SequenceEqual(input.AllowedTopLevelClaims)
                ) && 
                (
                    this.ClientId == input.ClientId ||
                    (this.ClientId != null &&
                    this.ClientId.Equals(input.ClientId))
                ) && 
                (
                    this.ConsentChallenge == input.ConsentChallenge ||
                    (this.ConsentChallenge != null &&
                    this.ConsentChallenge.Equals(input.ConsentChallenge))
                ) && 
                (
                    this.ExcludeNotBeforeClaim == input.ExcludeNotBeforeClaim ||
                    this.ExcludeNotBeforeClaim.Equals(input.ExcludeNotBeforeClaim)
                ) && 
                (
                    this.ExpiresAt == input.ExpiresAt ||
                    (this.ExpiresAt != null &&
                    this.ExpiresAt.Equals(input.ExpiresAt))
                ) && 
                (
                    this.Extra == input.Extra ||
                    this.Extra != null &&
                    input.Extra != null &&
                    this.Extra.SequenceEqual(input.Extra)
                ) && 
                (
                    this.Headers == input.Headers ||
                    (this.Headers != null &&
                    this.Headers.Equals(input.Headers))
                ) && 
                (
                    this.IdTokenClaims == input.IdTokenClaims ||
                    (this.IdTokenClaims != null &&
                    this.IdTokenClaims.Equals(input.IdTokenClaims))
                ) && 
                (
                    this.Kid == input.Kid ||
                    (this.Kid != null &&
                    this.Kid.Equals(input.Kid))
                ) && 
                (
                    this.Subject == input.Subject ||
                    (this.Subject != null &&
                    this.Subject.Equals(input.Subject))
                ) && 
                (
                    this.Username == input.Username ||
                    (this.Username != null &&
                    this.Username.Equals(input.Username))
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
                if (this.AllowedTopLevelClaims != null)
                {
                    hashCode = (hashCode * 59) + this.AllowedTopLevelClaims.GetHashCode();
                }
                if (this.ClientId != null)
                {
                    hashCode = (hashCode * 59) + this.ClientId.GetHashCode();
                }
                if (this.ConsentChallenge != null)
                {
                    hashCode = (hashCode * 59) + this.ConsentChallenge.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.ExcludeNotBeforeClaim.GetHashCode();
                if (this.ExpiresAt != null)
                {
                    hashCode = (hashCode * 59) + this.ExpiresAt.GetHashCode();
                }
                if (this.Extra != null)
                {
                    hashCode = (hashCode * 59) + this.Extra.GetHashCode();
                }
                if (this.Headers != null)
                {
                    hashCode = (hashCode * 59) + this.Headers.GetHashCode();
                }
                if (this.IdTokenClaims != null)
                {
                    hashCode = (hashCode * 59) + this.IdTokenClaims.GetHashCode();
                }
                if (this.Kid != null)
                {
                    hashCode = (hashCode * 59) + this.Kid.GetHashCode();
                }
                if (this.Subject != null)
                {
                    hashCode = (hashCode * 59) + this.Subject.GetHashCode();
                }
                if (this.Username != null)
                {
                    hashCode = (hashCode * 59) + this.Username.GetHashCode();
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
