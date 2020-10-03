/* 
 * ORY Hydra
 *
 * Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here.
 *
 * The version of the OpenAPI document: v1.8.0-pre.1
 * 
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using System;
using System.Linq;
using System.IO;
using System.Text;
using System.Text.RegularExpressions;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Runtime.Serialization;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using System.ComponentModel.DataAnnotations;
using OpenAPIDateConverter = Ory.Hydra.Client.Client.OpenAPIDateConverter;

namespace Ory.Hydra.Client.Model
{
    /// <summary>
    /// HydraJsonWebKeySetGeneratorRequest
    /// </summary>
    [DataContract]
    public partial class HydraJsonWebKeySetGeneratorRequest :  IEquatable<HydraJsonWebKeySetGeneratorRequest>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="HydraJsonWebKeySetGeneratorRequest" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected HydraJsonWebKeySetGeneratorRequest() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="HydraJsonWebKeySetGeneratorRequest" /> class.
        /// </summary>
        /// <param name="alg">The algorithm to be used for creating the key. Supports \&quot;RS256\&quot;, \&quot;ES512\&quot;, \&quot;HS512\&quot;, and \&quot;HS256\&quot; (required).</param>
        /// <param name="kid">The kid of the key to be created (required).</param>
        /// <param name="use">The \&quot;use\&quot; (public key use) parameter identifies the intended use of the public key. The \&quot;use\&quot; parameter is employed to indicate whether a public key is used for encrypting data or verifying the signature on data. Valid values are \&quot;enc\&quot; and \&quot;sig\&quot;. (required).</param>
        public HydraJsonWebKeySetGeneratorRequest(string alg = default(string), string kid = default(string), string use = default(string))
        {
            // to ensure "alg" is required (not null)
            this.Alg = alg ?? throw new ArgumentNullException("alg is a required property for HydraJsonWebKeySetGeneratorRequest and cannot be null");
            // to ensure "kid" is required (not null)
            this.Kid = kid ?? throw new ArgumentNullException("kid is a required property for HydraJsonWebKeySetGeneratorRequest and cannot be null");
            // to ensure "use" is required (not null)
            this.Use = use ?? throw new ArgumentNullException("use is a required property for HydraJsonWebKeySetGeneratorRequest and cannot be null");
        }
        
        /// <summary>
        /// The algorithm to be used for creating the key. Supports \&quot;RS256\&quot;, \&quot;ES512\&quot;, \&quot;HS512\&quot;, and \&quot;HS256\&quot;
        /// </summary>
        /// <value>The algorithm to be used for creating the key. Supports \&quot;RS256\&quot;, \&quot;ES512\&quot;, \&quot;HS512\&quot;, and \&quot;HS256\&quot;</value>
        [DataMember(Name="alg", EmitDefaultValue=false)]
        public string Alg { get; set; }

        /// <summary>
        /// The kid of the key to be created
        /// </summary>
        /// <value>The kid of the key to be created</value>
        [DataMember(Name="kid", EmitDefaultValue=false)]
        public string Kid { get; set; }

        /// <summary>
        /// The \&quot;use\&quot; (public key use) parameter identifies the intended use of the public key. The \&quot;use\&quot; parameter is employed to indicate whether a public key is used for encrypting data or verifying the signature on data. Valid values are \&quot;enc\&quot; and \&quot;sig\&quot;.
        /// </summary>
        /// <value>The \&quot;use\&quot; (public key use) parameter identifies the intended use of the public key. The \&quot;use\&quot; parameter is employed to indicate whether a public key is used for encrypting data or verifying the signature on data. Valid values are \&quot;enc\&quot; and \&quot;sig\&quot;.</value>
        [DataMember(Name="use", EmitDefaultValue=false)]
        public string Use { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class HydraJsonWebKeySetGeneratorRequest {\n");
            sb.Append("  Alg: ").Append(Alg).Append("\n");
            sb.Append("  Kid: ").Append(Kid).Append("\n");
            sb.Append("  Use: ").Append(Use).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }
  
        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public virtual string ToJson()
        {
            return JsonConvert.SerializeObject(this, Formatting.Indented);
        }

        /// <summary>
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="input">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object input)
        {
            return this.Equals(input as HydraJsonWebKeySetGeneratorRequest);
        }

        /// <summary>
        /// Returns true if HydraJsonWebKeySetGeneratorRequest instances are equal
        /// </summary>
        /// <param name="input">Instance of HydraJsonWebKeySetGeneratorRequest to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(HydraJsonWebKeySetGeneratorRequest input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.Alg == input.Alg ||
                    (this.Alg != null &&
                    this.Alg.Equals(input.Alg))
                ) && 
                (
                    this.Kid == input.Kid ||
                    (this.Kid != null &&
                    this.Kid.Equals(input.Kid))
                ) && 
                (
                    this.Use == input.Use ||
                    (this.Use != null &&
                    this.Use.Equals(input.Use))
                );
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
                if (this.Alg != null)
                    hashCode = hashCode * 59 + this.Alg.GetHashCode();
                if (this.Kid != null)
                    hashCode = hashCode * 59 + this.Kid.GetHashCode();
                if (this.Use != null)
                    hashCode = hashCode * 59 + this.Use.GetHashCode();
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
