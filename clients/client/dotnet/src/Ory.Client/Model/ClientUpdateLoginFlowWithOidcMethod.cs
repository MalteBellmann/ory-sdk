/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.6.1
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
    /// Update Login Flow with OpenID Connect Method
    /// </summary>
    [DataContract(Name = "updateLoginFlowWithOidcMethod")]
    public partial class ClientUpdateLoginFlowWithOidcMethod : IEquatable<ClientUpdateLoginFlowWithOidcMethod>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientUpdateLoginFlowWithOidcMethod" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected ClientUpdateLoginFlowWithOidcMethod()
        {
            this.AdditionalProperties = new Dictionary<string, object>();
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientUpdateLoginFlowWithOidcMethod" /> class.
        /// </summary>
        /// <param name="csrfToken">The CSRF Token.</param>
        /// <param name="idToken">IDToken is an optional id token provided by an OIDC provider  If submitted, it is verified using the OIDC provider&#39;s public key set and the claims are used to populate the OIDC credentials of the identity. If the OIDC provider does not store additional claims (such as name, etc.) in the IDToken itself, you can use the &#x60;traits&#x60; field to populate the identity&#39;s traits. Note, that Apple only includes the users email in the IDToken.  Supported providers are Apple.</param>
        /// <param name="idTokenNonce">IDTokenNonce is the nonce, used when generating the IDToken. If the provider supports nonce validation, the nonce will be validated against this value and required..</param>
        /// <param name="method">Method to use  This field must be set to &#x60;oidc&#x60; when using the oidc method. (required).</param>
        /// <param name="provider">The provider to register with (required).</param>
        /// <param name="traits">The identity traits. This is a placeholder for the registration flow..</param>
        /// <param name="upstreamParameters">UpstreamParameters are the parameters that are passed to the upstream identity provider.  These parameters are optional and depend on what the upstream identity provider supports. Supported parameters are: &#x60;login_hint&#x60; (string): The &#x60;login_hint&#x60; parameter suppresses the account chooser and either pre-fills the email box on the sign-in form, or selects the proper session. &#x60;hd&#x60; (string): The &#x60;hd&#x60; parameter limits the login/registration process to a Google Organization, e.g. &#x60;mycollege.edu&#x60;. &#x60;prompt&#x60; (string): The &#x60;prompt&#x60; specifies whether the Authorization Server prompts the End-User for reauthentication and consent, e.g. &#x60;select_account&#x60;..</param>
        public ClientUpdateLoginFlowWithOidcMethod(string csrfToken = default(string), string idToken = default(string), string idTokenNonce = default(string), string method = default(string), string provider = default(string), Object traits = default(Object), Object upstreamParameters = default(Object))
        {
            // to ensure "method" is required (not null)
            if (method == null) {
                throw new ArgumentNullException("method is a required property for ClientUpdateLoginFlowWithOidcMethod and cannot be null");
            }
            this.Method = method;
            // to ensure "provider" is required (not null)
            if (provider == null) {
                throw new ArgumentNullException("provider is a required property for ClientUpdateLoginFlowWithOidcMethod and cannot be null");
            }
            this.Provider = provider;
            this.CsrfToken = csrfToken;
            this.IdToken = idToken;
            this.IdTokenNonce = idTokenNonce;
            this.Traits = traits;
            this.UpstreamParameters = upstreamParameters;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// The CSRF Token
        /// </summary>
        /// <value>The CSRF Token</value>
        [DataMember(Name = "csrf_token", EmitDefaultValue = false)]
        public string CsrfToken { get; set; }

        /// <summary>
        /// IDToken is an optional id token provided by an OIDC provider  If submitted, it is verified using the OIDC provider&#39;s public key set and the claims are used to populate the OIDC credentials of the identity. If the OIDC provider does not store additional claims (such as name, etc.) in the IDToken itself, you can use the &#x60;traits&#x60; field to populate the identity&#39;s traits. Note, that Apple only includes the users email in the IDToken.  Supported providers are Apple
        /// </summary>
        /// <value>IDToken is an optional id token provided by an OIDC provider  If submitted, it is verified using the OIDC provider&#39;s public key set and the claims are used to populate the OIDC credentials of the identity. If the OIDC provider does not store additional claims (such as name, etc.) in the IDToken itself, you can use the &#x60;traits&#x60; field to populate the identity&#39;s traits. Note, that Apple only includes the users email in the IDToken.  Supported providers are Apple</value>
        [DataMember(Name = "id_token", EmitDefaultValue = false)]
        public string IdToken { get; set; }

        /// <summary>
        /// IDTokenNonce is the nonce, used when generating the IDToken. If the provider supports nonce validation, the nonce will be validated against this value and required.
        /// </summary>
        /// <value>IDTokenNonce is the nonce, used when generating the IDToken. If the provider supports nonce validation, the nonce will be validated against this value and required.</value>
        [DataMember(Name = "id_token_nonce", EmitDefaultValue = false)]
        public string IdTokenNonce { get; set; }

        /// <summary>
        /// Method to use  This field must be set to &#x60;oidc&#x60; when using the oidc method.
        /// </summary>
        /// <value>Method to use  This field must be set to &#x60;oidc&#x60; when using the oidc method.</value>
        [DataMember(Name = "method", IsRequired = true, EmitDefaultValue = false)]
        public string Method { get; set; }

        /// <summary>
        /// The provider to register with
        /// </summary>
        /// <value>The provider to register with</value>
        [DataMember(Name = "provider", IsRequired = true, EmitDefaultValue = false)]
        public string Provider { get; set; }

        /// <summary>
        /// The identity traits. This is a placeholder for the registration flow.
        /// </summary>
        /// <value>The identity traits. This is a placeholder for the registration flow.</value>
        [DataMember(Name = "traits", EmitDefaultValue = false)]
        public Object Traits { get; set; }

        /// <summary>
        /// UpstreamParameters are the parameters that are passed to the upstream identity provider.  These parameters are optional and depend on what the upstream identity provider supports. Supported parameters are: &#x60;login_hint&#x60; (string): The &#x60;login_hint&#x60; parameter suppresses the account chooser and either pre-fills the email box on the sign-in form, or selects the proper session. &#x60;hd&#x60; (string): The &#x60;hd&#x60; parameter limits the login/registration process to a Google Organization, e.g. &#x60;mycollege.edu&#x60;. &#x60;prompt&#x60; (string): The &#x60;prompt&#x60; specifies whether the Authorization Server prompts the End-User for reauthentication and consent, e.g. &#x60;select_account&#x60;.
        /// </summary>
        /// <value>UpstreamParameters are the parameters that are passed to the upstream identity provider.  These parameters are optional and depend on what the upstream identity provider supports. Supported parameters are: &#x60;login_hint&#x60; (string): The &#x60;login_hint&#x60; parameter suppresses the account chooser and either pre-fills the email box on the sign-in form, or selects the proper session. &#x60;hd&#x60; (string): The &#x60;hd&#x60; parameter limits the login/registration process to a Google Organization, e.g. &#x60;mycollege.edu&#x60;. &#x60;prompt&#x60; (string): The &#x60;prompt&#x60; specifies whether the Authorization Server prompts the End-User for reauthentication and consent, e.g. &#x60;select_account&#x60;.</value>
        [DataMember(Name = "upstream_parameters", EmitDefaultValue = false)]
        public Object UpstreamParameters { get; set; }

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
            sb.Append("class ClientUpdateLoginFlowWithOidcMethod {\n");
            sb.Append("  CsrfToken: ").Append(CsrfToken).Append("\n");
            sb.Append("  IdToken: ").Append(IdToken).Append("\n");
            sb.Append("  IdTokenNonce: ").Append(IdTokenNonce).Append("\n");
            sb.Append("  Method: ").Append(Method).Append("\n");
            sb.Append("  Provider: ").Append(Provider).Append("\n");
            sb.Append("  Traits: ").Append(Traits).Append("\n");
            sb.Append("  UpstreamParameters: ").Append(UpstreamParameters).Append("\n");
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
            return this.Equals(input as ClientUpdateLoginFlowWithOidcMethod);
        }

        /// <summary>
        /// Returns true if ClientUpdateLoginFlowWithOidcMethod instances are equal
        /// </summary>
        /// <param name="input">Instance of ClientUpdateLoginFlowWithOidcMethod to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ClientUpdateLoginFlowWithOidcMethod input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.CsrfToken == input.CsrfToken ||
                    (this.CsrfToken != null &&
                    this.CsrfToken.Equals(input.CsrfToken))
                ) && 
                (
                    this.IdToken == input.IdToken ||
                    (this.IdToken != null &&
                    this.IdToken.Equals(input.IdToken))
                ) && 
                (
                    this.IdTokenNonce == input.IdTokenNonce ||
                    (this.IdTokenNonce != null &&
                    this.IdTokenNonce.Equals(input.IdTokenNonce))
                ) && 
                (
                    this.Method == input.Method ||
                    (this.Method != null &&
                    this.Method.Equals(input.Method))
                ) && 
                (
                    this.Provider == input.Provider ||
                    (this.Provider != null &&
                    this.Provider.Equals(input.Provider))
                ) && 
                (
                    this.Traits == input.Traits ||
                    (this.Traits != null &&
                    this.Traits.Equals(input.Traits))
                ) && 
                (
                    this.UpstreamParameters == input.UpstreamParameters ||
                    (this.UpstreamParameters != null &&
                    this.UpstreamParameters.Equals(input.UpstreamParameters))
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
                if (this.CsrfToken != null)
                {
                    hashCode = (hashCode * 59) + this.CsrfToken.GetHashCode();
                }
                if (this.IdToken != null)
                {
                    hashCode = (hashCode * 59) + this.IdToken.GetHashCode();
                }
                if (this.IdTokenNonce != null)
                {
                    hashCode = (hashCode * 59) + this.IdTokenNonce.GetHashCode();
                }
                if (this.Method != null)
                {
                    hashCode = (hashCode * 59) + this.Method.GetHashCode();
                }
                if (this.Provider != null)
                {
                    hashCode = (hashCode * 59) + this.Provider.GetHashCode();
                }
                if (this.Traits != null)
                {
                    hashCode = (hashCode * 59) + this.Traits.GetHashCode();
                }
                if (this.UpstreamParameters != null)
                {
                    hashCode = (hashCode * 59) + this.UpstreamParameters.GetHashCode();
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
