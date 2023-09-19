/*
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.2.8
 * Contact: support@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonArray;
import com.google.gson.JsonDeserializationContext;
import com.google.gson.JsonDeserializer;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonParseException;
import com.google.gson.TypeAdapterFactory;
import com.google.gson.reflect.TypeToken;

import java.lang.reflect.Type;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Map.Entry;
import java.util.Set;

import sh.ory.JSON;

/**
 * Ory Identity Schema Validation Result
 */
@ApiModel(description = "Ory Identity Schema Validation Result")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-09-19T13:27:41.914746600Z[Etc/UTC]")
public class ManagedIdentitySchemaValidationResult {
  public static final String SERIALIZED_NAME_MESSAGE = "message";
  @SerializedName(SERIALIZED_NAME_MESSAGE)
  private String message;

  public static final String SERIALIZED_NAME_VALID = "valid";
  @SerializedName(SERIALIZED_NAME_VALID)
  private Boolean valid;

  public ManagedIdentitySchemaValidationResult() {
  }

  public ManagedIdentitySchemaValidationResult message(String message) {
    
    this.message = message;
    return this;
  }

   /**
   * Get message
   * @return message
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMessage() {
    return message;
  }


  public void setMessage(String message) {
    this.message = message;
  }


  public ManagedIdentitySchemaValidationResult valid(Boolean valid) {
    
    this.valid = valid;
    return this;
  }

   /**
   * Get valid
   * @return valid
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getValid() {
    return valid;
  }


  public void setValid(Boolean valid) {
    this.valid = valid;
  }

  /**
   * A container for additional, undeclared properties.
   * This is a holder for any undeclared properties as specified with
   * the 'additionalProperties' keyword in the OAS document.
   */
  private Map<String, Object> additionalProperties;

  /**
   * Set the additional (undeclared) property with the specified name and value.
   * If the property does not already exist, create it otherwise replace it.
   *
   * @param key name of the property
   * @param value value of the property
   * @return the ManagedIdentitySchemaValidationResult instance itself
   */
  public ManagedIdentitySchemaValidationResult putAdditionalProperty(String key, Object value) {
    if (this.additionalProperties == null) {
        this.additionalProperties = new HashMap<String, Object>();
    }
    this.additionalProperties.put(key, value);
    return this;
  }

  /**
   * Return the additional (undeclared) property.
   *
   * @return a map of objects
   */
  public Map<String, Object> getAdditionalProperties() {
    return additionalProperties;
  }

  /**
   * Return the additional (undeclared) property with the specified name.
   *
   * @param key name of the property
   * @return an object
   */
  public Object getAdditionalProperty(String key) {
    if (this.additionalProperties == null) {
        return null;
    }
    return this.additionalProperties.get(key);
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ManagedIdentitySchemaValidationResult managedIdentitySchemaValidationResult = (ManagedIdentitySchemaValidationResult) o;
    return Objects.equals(this.message, managedIdentitySchemaValidationResult.message) &&
        Objects.equals(this.valid, managedIdentitySchemaValidationResult.valid)&&
        Objects.equals(this.additionalProperties, managedIdentitySchemaValidationResult.additionalProperties);
  }

  @Override
  public int hashCode() {
    return Objects.hash(message, valid, additionalProperties);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ManagedIdentitySchemaValidationResult {\n");
    sb.append("    message: ").append(toIndentedString(message)).append("\n");
    sb.append("    valid: ").append(toIndentedString(valid)).append("\n");
    sb.append("    additionalProperties: ").append(toIndentedString(additionalProperties)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }


  public static HashSet<String> openapiFields;
  public static HashSet<String> openapiRequiredFields;

  static {
    // a set of all properties/fields (JSON key names)
    openapiFields = new HashSet<String>();
    openapiFields.add("message");
    openapiFields.add("valid");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to ManagedIdentitySchemaValidationResult
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!ManagedIdentitySchemaValidationResult.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in ManagedIdentitySchemaValidationResult is not found in the empty JSON string", ManagedIdentitySchemaValidationResult.openapiRequiredFields.toString()));
        }
      }
      if ((jsonObj.get("message") != null && !jsonObj.get("message").isJsonNull()) && !jsonObj.get("message").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `message` to be a primitive type in the JSON string but got `%s`", jsonObj.get("message").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!ManagedIdentitySchemaValidationResult.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'ManagedIdentitySchemaValidationResult' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<ManagedIdentitySchemaValidationResult> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(ManagedIdentitySchemaValidationResult.class));

       return (TypeAdapter<T>) new TypeAdapter<ManagedIdentitySchemaValidationResult>() {
           @Override
           public void write(JsonWriter out, ManagedIdentitySchemaValidationResult value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             obj.remove("additionalProperties");
             // serialize additonal properties
             if (value.getAdditionalProperties() != null) {
               for (Map.Entry<String, Object> entry : value.getAdditionalProperties().entrySet()) {
                 if (entry.getValue() instanceof String)
                   obj.addProperty(entry.getKey(), (String) entry.getValue());
                 else if (entry.getValue() instanceof Number)
                   obj.addProperty(entry.getKey(), (Number) entry.getValue());
                 else if (entry.getValue() instanceof Boolean)
                   obj.addProperty(entry.getKey(), (Boolean) entry.getValue());
                 else if (entry.getValue() instanceof Character)
                   obj.addProperty(entry.getKey(), (Character) entry.getValue());
                 else {
                   obj.add(entry.getKey(), gson.toJsonTree(entry.getValue()).getAsJsonObject());
                 }
               }
             }
             elementAdapter.write(out, obj);
           }

           @Override
           public ManagedIdentitySchemaValidationResult read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             // store additional fields in the deserialized instance
             ManagedIdentitySchemaValidationResult instance = thisAdapter.fromJsonTree(jsonObj);
             for (Map.Entry<String, JsonElement> entry : jsonObj.entrySet()) {
               if (!openapiFields.contains(entry.getKey())) {
                 if (entry.getValue().isJsonPrimitive()) { // primitive type
                   if (entry.getValue().getAsJsonPrimitive().isString())
                     instance.putAdditionalProperty(entry.getKey(), entry.getValue().getAsString());
                   else if (entry.getValue().getAsJsonPrimitive().isNumber())
                     instance.putAdditionalProperty(entry.getKey(), entry.getValue().getAsNumber());
                   else if (entry.getValue().getAsJsonPrimitive().isBoolean())
                     instance.putAdditionalProperty(entry.getKey(), entry.getValue().getAsBoolean());
                   else
                     throw new IllegalArgumentException(String.format("The field `%s` has unknown primitive type. Value: %s", entry.getKey(), entry.getValue().toString()));
                 } else if (entry.getValue().isJsonArray()) {
                     instance.putAdditionalProperty(entry.getKey(), gson.fromJson(entry.getValue(), List.class));
                 } else { // JSON object
                     instance.putAdditionalProperty(entry.getKey(), gson.fromJson(entry.getValue(), HashMap.class));
                 }
               }
             }
             return instance;
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of ManagedIdentitySchemaValidationResult given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of ManagedIdentitySchemaValidationResult
  * @throws IOException if the JSON string is invalid with respect to ManagedIdentitySchemaValidationResult
  */
  public static ManagedIdentitySchemaValidationResult fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, ManagedIdentitySchemaValidationResult.class);
  }

 /**
  * Convert an instance of ManagedIdentitySchemaValidationResult to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

