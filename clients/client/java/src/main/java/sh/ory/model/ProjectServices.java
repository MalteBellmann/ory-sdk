/*
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.6.1
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
import sh.ory.model.ProjectServiceIdentity;
import sh.ory.model.ProjectServiceOAuth2;
import sh.ory.model.ProjectServicePermission;

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
 * ProjectServices
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2024-02-05T19:48:39.584840582Z[Etc/UTC]")
public class ProjectServices {
  public static final String SERIALIZED_NAME_IDENTITY = "identity";
  @SerializedName(SERIALIZED_NAME_IDENTITY)
  private ProjectServiceIdentity identity;

  public static final String SERIALIZED_NAME_OAUTH2 = "oauth2";
  @SerializedName(SERIALIZED_NAME_OAUTH2)
  private ProjectServiceOAuth2 oauth2;

  public static final String SERIALIZED_NAME_PERMISSION = "permission";
  @SerializedName(SERIALIZED_NAME_PERMISSION)
  private ProjectServicePermission permission;

  public ProjectServices() {
  }

  public ProjectServices identity(ProjectServiceIdentity identity) {
    
    this.identity = identity;
    return this;
  }

   /**
   * Get identity
   * @return identity
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ProjectServiceIdentity getIdentity() {
    return identity;
  }


  public void setIdentity(ProjectServiceIdentity identity) {
    this.identity = identity;
  }


  public ProjectServices oauth2(ProjectServiceOAuth2 oauth2) {
    
    this.oauth2 = oauth2;
    return this;
  }

   /**
   * Get oauth2
   * @return oauth2
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ProjectServiceOAuth2 getOauth2() {
    return oauth2;
  }


  public void setOauth2(ProjectServiceOAuth2 oauth2) {
    this.oauth2 = oauth2;
  }


  public ProjectServices permission(ProjectServicePermission permission) {
    
    this.permission = permission;
    return this;
  }

   /**
   * Get permission
   * @return permission
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ProjectServicePermission getPermission() {
    return permission;
  }


  public void setPermission(ProjectServicePermission permission) {
    this.permission = permission;
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
   * @return the ProjectServices instance itself
   */
  public ProjectServices putAdditionalProperty(String key, Object value) {
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
    ProjectServices projectServices = (ProjectServices) o;
    return Objects.equals(this.identity, projectServices.identity) &&
        Objects.equals(this.oauth2, projectServices.oauth2) &&
        Objects.equals(this.permission, projectServices.permission)&&
        Objects.equals(this.additionalProperties, projectServices.additionalProperties);
  }

  @Override
  public int hashCode() {
    return Objects.hash(identity, oauth2, permission, additionalProperties);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ProjectServices {\n");
    sb.append("    identity: ").append(toIndentedString(identity)).append("\n");
    sb.append("    oauth2: ").append(toIndentedString(oauth2)).append("\n");
    sb.append("    permission: ").append(toIndentedString(permission)).append("\n");
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
    openapiFields.add("identity");
    openapiFields.add("oauth2");
    openapiFields.add("permission");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to ProjectServices
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!ProjectServices.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in ProjectServices is not found in the empty JSON string", ProjectServices.openapiRequiredFields.toString()));
        }
      }
      // validate the optional field `identity`
      if (jsonObj.get("identity") != null && !jsonObj.get("identity").isJsonNull()) {
        ProjectServiceIdentity.validateJsonObject(jsonObj.getAsJsonObject("identity"));
      }
      // validate the optional field `oauth2`
      if (jsonObj.get("oauth2") != null && !jsonObj.get("oauth2").isJsonNull()) {
        ProjectServiceOAuth2.validateJsonObject(jsonObj.getAsJsonObject("oauth2"));
      }
      // validate the optional field `permission`
      if (jsonObj.get("permission") != null && !jsonObj.get("permission").isJsonNull()) {
        ProjectServicePermission.validateJsonObject(jsonObj.getAsJsonObject("permission"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!ProjectServices.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'ProjectServices' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<ProjectServices> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(ProjectServices.class));

       return (TypeAdapter<T>) new TypeAdapter<ProjectServices>() {
           @Override
           public void write(JsonWriter out, ProjectServices value) throws IOException {
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
           public ProjectServices read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             // store additional fields in the deserialized instance
             ProjectServices instance = thisAdapter.fromJsonTree(jsonObj);
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
  * Create an instance of ProjectServices given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of ProjectServices
  * @throws IOException if the JSON string is invalid with respect to ProjectServices
  */
  public static ProjectServices fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, ProjectServices.class);
  }

 /**
  * Convert an instance of ProjectServices to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

