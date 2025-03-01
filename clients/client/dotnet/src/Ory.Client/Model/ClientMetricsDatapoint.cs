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
    /// Represents a single datapoint/bucket of a time series
    /// </summary>
    [DataContract(Name = "metricsDatapoint")]
    public partial class ClientMetricsDatapoint : IEquatable<ClientMetricsDatapoint>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientMetricsDatapoint" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected ClientMetricsDatapoint()
        {
            this.AdditionalProperties = new Dictionary<string, object>();
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientMetricsDatapoint" /> class.
        /// </summary>
        /// <param name="count">The count of events that occured in this time (required).</param>
        /// <param name="time">The time of the bucket (required).</param>
        public ClientMetricsDatapoint(long count = default(long), DateTime time = default(DateTime))
        {
            this.Count = count;
            this.Time = time;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// The count of events that occured in this time
        /// </summary>
        /// <value>The count of events that occured in this time</value>
        [DataMember(Name = "count", IsRequired = true, EmitDefaultValue = false)]
        public long Count { get; set; }

        /// <summary>
        /// The time of the bucket
        /// </summary>
        /// <value>The time of the bucket</value>
        [DataMember(Name = "time", IsRequired = true, EmitDefaultValue = false)]
        public DateTime Time { get; set; }

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
            sb.Append("class ClientMetricsDatapoint {\n");
            sb.Append("  Count: ").Append(Count).Append("\n");
            sb.Append("  Time: ").Append(Time).Append("\n");
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
            return this.Equals(input as ClientMetricsDatapoint);
        }

        /// <summary>
        /// Returns true if ClientMetricsDatapoint instances are equal
        /// </summary>
        /// <param name="input">Instance of ClientMetricsDatapoint to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ClientMetricsDatapoint input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.Count == input.Count ||
                    this.Count.Equals(input.Count)
                ) && 
                (
                    this.Time == input.Time ||
                    (this.Time != null &&
                    this.Time.Equals(input.Time))
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
                hashCode = (hashCode * 59) + this.Count.GetHashCode();
                if (this.Time != null)
                {
                    hashCode = (hashCode * 59) + this.Time.GetHashCode();
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
