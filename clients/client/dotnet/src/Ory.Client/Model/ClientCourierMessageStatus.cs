/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.2.8
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
    /// A Message&#39;s Status
    /// </summary>
    /// <value>A Message&#39;s Status</value>
    [JsonConverter(typeof(StringEnumConverter))]
    public enum ClientCourierMessageStatus
    {
        /// <summary>
        /// Enum Queued for value: queued
        /// </summary>
        [EnumMember(Value = "queued")]
        Queued = 1,

        /// <summary>
        /// Enum Sent for value: sent
        /// </summary>
        [EnumMember(Value = "sent")]
        Sent = 2,

        /// <summary>
        /// Enum Processing for value: processing
        /// </summary>
        [EnumMember(Value = "processing")]
        Processing = 3,

        /// <summary>
        /// Enum Abandoned for value: abandoned
        /// </summary>
        [EnumMember(Value = "abandoned")]
        Abandoned = 4

    }

}
