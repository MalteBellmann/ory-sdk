/*
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.1.13
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
import java.util.ArrayList;
import java.util.List;
import sh.ory.model.OAuth2Client;
import sh.ory.model.OAuth2ConsentRequestOpenIDConnectContext;

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
 * OAuth2LoginRequest
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-02-16T15:37:49.887299015Z[Etc/UTC]")
public class OAuth2LoginRequest {
  public static final String SERIALIZED_NAME_CHALLENGE = "challenge";
  @SerializedName(SERIALIZED_NAME_CHALLENGE)
  private String challenge;

  public static final String SERIALIZED_NAME_CLIENT = "client";
  @SerializedName(SERIALIZED_NAME_CLIENT)
  private OAuth2Client client;

  public static final String SERIALIZED_NAME_OIDC_CONTEXT = "oidc_context";
  @SerializedName(SERIALIZED_NAME_OIDC_CONTEXT)
  private OAuth2ConsentRequestOpenIDConnectContext oidcContext;

  public static final String SERIALIZED_NAME_REQUEST_URL = "request_url";
  @SerializedName(SERIALIZED_NAME_REQUEST_URL)
  private String requestUrl;

  public static final String SERIALIZED_NAME_REQUESTED_ACCESS_TOKEN_AUDIENCE = "requested_access_token_audience";
  @SerializedName(SERIALIZED_NAME_REQUESTED_ACCESS_TOKEN_AUDIENCE)
  private List<String> requestedAccessTokenAudience = new ArrayList<>();

  public static final String SERIALIZED_NAME_REQUESTED_SCOPE = "requested_scope";
  @SerializedName(SERIALIZED_NAME_REQUESTED_SCOPE)
  private List<String> requestedScope = new ArrayList<>();

  public static final String SERIALIZED_NAME_SESSION_ID = "session_id";
  @SerializedName(SERIALIZED_NAME_SESSION_ID)
  private String sessionId;

  public static final String SERIALIZED_NAME_SKIP = "skip";
  @SerializedName(SERIALIZED_NAME_SKIP)
  private Boolean skip;

  public static final String SERIALIZED_NAME_SUBJECT = "subject";
  @SerializedName(SERIALIZED_NAME_SUBJECT)
  private String subject;

  public OAuth2LoginRequest() {
  }

  public OAuth2LoginRequest challenge(String challenge) {
    
    this.challenge = challenge;
    return this;
  }

   /**
   * ID is the identifier (\&quot;login challenge\&quot;) of the login request. It is used to identify the session.
   * @return challenge
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "ID is the identifier (\"login challenge\") of the login request. It is used to identify the session.")

  public String getChallenge() {
    return challenge;
  }


  public void setChallenge(String challenge) {
    this.challenge = challenge;
  }


  public OAuth2LoginRequest client(OAuth2Client client) {
    
    this.client = client;
    return this;
  }

   /**
   * Get client
   * @return client
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")

  public OAuth2Client getClient() {
    return client;
  }


  public void setClient(OAuth2Client client) {
    this.client = client;
  }


  public OAuth2LoginRequest oidcContext(OAuth2ConsentRequestOpenIDConnectContext oidcContext) {
    
    this.oidcContext = oidcContext;
    return this;
  }

   /**
   * Get oidcContext
   * @return oidcContext
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OAuth2ConsentRequestOpenIDConnectContext getOidcContext() {
    return oidcContext;
  }


  public void setOidcContext(OAuth2ConsentRequestOpenIDConnectContext oidcContext) {
    this.oidcContext = oidcContext;
  }


  public OAuth2LoginRequest requestUrl(String requestUrl) {
    
    this.requestUrl = requestUrl;
    return this;
  }

   /**
   * RequestURL is the original OAuth 2.0 Authorization URL requested by the OAuth 2.0 client. It is the URL which initiates the OAuth 2.0 Authorization Code or OAuth 2.0 Implicit flow. This URL is typically not needed, but might come in handy if you want to deal with additional request parameters.
   * @return requestUrl
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "RequestURL is the original OAuth 2.0 Authorization URL requested by the OAuth 2.0 client. It is the URL which initiates the OAuth 2.0 Authorization Code or OAuth 2.0 Implicit flow. This URL is typically not needed, but might come in handy if you want to deal with additional request parameters.")

  public String getRequestUrl() {
    return requestUrl;
  }


  public void setRequestUrl(String requestUrl) {
    this.requestUrl = requestUrl;
  }


  public OAuth2LoginRequest requestedAccessTokenAudience(List<String> requestedAccessTokenAudience) {
    
    this.requestedAccessTokenAudience = requestedAccessTokenAudience;
    return this;
  }

  public OAuth2LoginRequest addRequestedAccessTokenAudienceItem(String requestedAccessTokenAudienceItem) {
    this.requestedAccessTokenAudience.add(requestedAccessTokenAudienceItem);
    return this;
  }

   /**
   * Get requestedAccessTokenAudience
   * @return requestedAccessTokenAudience
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")

  public List<String> getRequestedAccessTokenAudience() {
    return requestedAccessTokenAudience;
  }


  public void setRequestedAccessTokenAudience(List<String> requestedAccessTokenAudience) {
    this.requestedAccessTokenAudience = requestedAccessTokenAudience;
  }


  public OAuth2LoginRequest requestedScope(List<String> requestedScope) {
    
    this.requestedScope = requestedScope;
    return this;
  }

  public OAuth2LoginRequest addRequestedScopeItem(String requestedScopeItem) {
    this.requestedScope.add(requestedScopeItem);
    return this;
  }

   /**
   * Get requestedScope
   * @return requestedScope
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "")

  public List<String> getRequestedScope() {
    return requestedScope;
  }


  public void setRequestedScope(List<String> requestedScope) {
    this.requestedScope = requestedScope;
  }


  public OAuth2LoginRequest sessionId(String sessionId) {
    
    this.sessionId = sessionId;
    return this;
  }

   /**
   * SessionID is the login session ID. If the user-agent reuses a login session (via cookie / remember flag) this ID will remain the same. If the user-agent did not have an existing authentication session (e.g. remember is false) this will be a new random value. This value is used as the \&quot;sid\&quot; parameter in the ID Token and in OIDC Front-/Back- channel logout. It&#39;s value can generally be used to associate consecutive login requests by a certain user.
   * @return sessionId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "SessionID is the login session ID. If the user-agent reuses a login session (via cookie / remember flag) this ID will remain the same. If the user-agent did not have an existing authentication session (e.g. remember is false) this will be a new random value. This value is used as the \"sid\" parameter in the ID Token and in OIDC Front-/Back- channel logout. It's value can generally be used to associate consecutive login requests by a certain user.")

  public String getSessionId() {
    return sessionId;
  }


  public void setSessionId(String sessionId) {
    this.sessionId = sessionId;
  }


  public OAuth2LoginRequest skip(Boolean skip) {
    
    this.skip = skip;
    return this;
  }

   /**
   * Skip, if true, implies that the client has requested the same scopes from the same user previously. If true, you can skip asking the user to grant the requested scopes, and simply forward the user to the redirect URL.  This feature allows you to update / set session information.
   * @return skip
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Skip, if true, implies that the client has requested the same scopes from the same user previously. If true, you can skip asking the user to grant the requested scopes, and simply forward the user to the redirect URL.  This feature allows you to update / set session information.")

  public Boolean getSkip() {
    return skip;
  }


  public void setSkip(Boolean skip) {
    this.skip = skip;
  }


  public OAuth2LoginRequest subject(String subject) {
    
    this.subject = subject;
    return this;
  }

   /**
   * Subject is the user ID of the end-user that authenticated. Now, that end user needs to grant or deny the scope requested by the OAuth 2.0 client. If this value is set and &#x60;skip&#x60; is true, you MUST include this subject type when accepting the login request, or the request will fail.
   * @return subject
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Subject is the user ID of the end-user that authenticated. Now, that end user needs to grant or deny the scope requested by the OAuth 2.0 client. If this value is set and `skip` is true, you MUST include this subject type when accepting the login request, or the request will fail.")

  public String getSubject() {
    return subject;
  }


  public void setSubject(String subject) {
    this.subject = subject;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    OAuth2LoginRequest oAuth2LoginRequest = (OAuth2LoginRequest) o;
    return Objects.equals(this.challenge, oAuth2LoginRequest.challenge) &&
        Objects.equals(this.client, oAuth2LoginRequest.client) &&
        Objects.equals(this.oidcContext, oAuth2LoginRequest.oidcContext) &&
        Objects.equals(this.requestUrl, oAuth2LoginRequest.requestUrl) &&
        Objects.equals(this.requestedAccessTokenAudience, oAuth2LoginRequest.requestedAccessTokenAudience) &&
        Objects.equals(this.requestedScope, oAuth2LoginRequest.requestedScope) &&
        Objects.equals(this.sessionId, oAuth2LoginRequest.sessionId) &&
        Objects.equals(this.skip, oAuth2LoginRequest.skip) &&
        Objects.equals(this.subject, oAuth2LoginRequest.subject);
  }

  @Override
  public int hashCode() {
    return Objects.hash(challenge, client, oidcContext, requestUrl, requestedAccessTokenAudience, requestedScope, sessionId, skip, subject);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class OAuth2LoginRequest {\n");
    sb.append("    challenge: ").append(toIndentedString(challenge)).append("\n");
    sb.append("    client: ").append(toIndentedString(client)).append("\n");
    sb.append("    oidcContext: ").append(toIndentedString(oidcContext)).append("\n");
    sb.append("    requestUrl: ").append(toIndentedString(requestUrl)).append("\n");
    sb.append("    requestedAccessTokenAudience: ").append(toIndentedString(requestedAccessTokenAudience)).append("\n");
    sb.append("    requestedScope: ").append(toIndentedString(requestedScope)).append("\n");
    sb.append("    sessionId: ").append(toIndentedString(sessionId)).append("\n");
    sb.append("    skip: ").append(toIndentedString(skip)).append("\n");
    sb.append("    subject: ").append(toIndentedString(subject)).append("\n");
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
    openapiFields.add("challenge");
    openapiFields.add("client");
    openapiFields.add("oidc_context");
    openapiFields.add("request_url");
    openapiFields.add("requested_access_token_audience");
    openapiFields.add("requested_scope");
    openapiFields.add("session_id");
    openapiFields.add("skip");
    openapiFields.add("subject");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("challenge");
    openapiRequiredFields.add("client");
    openapiRequiredFields.add("request_url");
    openapiRequiredFields.add("requested_access_token_audience");
    openapiRequiredFields.add("requested_scope");
    openapiRequiredFields.add("skip");
    openapiRequiredFields.add("subject");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to OAuth2LoginRequest
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!OAuth2LoginRequest.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in OAuth2LoginRequest is not found in the empty JSON string", OAuth2LoginRequest.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!OAuth2LoginRequest.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `OAuth2LoginRequest` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : OAuth2LoginRequest.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if (!jsonObj.get("challenge").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `challenge` to be a primitive type in the JSON string but got `%s`", jsonObj.get("challenge").toString()));
      }
      // validate the required field `client`
      OAuth2Client.validateJsonObject(jsonObj.getAsJsonObject("client"));
      // validate the optional field `oidc_context`
      if (jsonObj.get("oidc_context") != null && !jsonObj.get("oidc_context").isJsonNull()) {
        OAuth2ConsentRequestOpenIDConnectContext.validateJsonObject(jsonObj.getAsJsonObject("oidc_context"));
      }
      if (!jsonObj.get("request_url").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `request_url` to be a primitive type in the JSON string but got `%s`", jsonObj.get("request_url").toString()));
      }
      // ensure the required json array is present
      if (jsonObj.get("requested_access_token_audience") == null) {
        throw new IllegalArgumentException("Expected the field `linkedContent` to be an array in the JSON string but got `null`");
      } else if (!jsonObj.get("requested_access_token_audience").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `requested_access_token_audience` to be an array in the JSON string but got `%s`", jsonObj.get("requested_access_token_audience").toString()));
      }
      // ensure the required json array is present
      if (jsonObj.get("requested_scope") == null) {
        throw new IllegalArgumentException("Expected the field `linkedContent` to be an array in the JSON string but got `null`");
      } else if (!jsonObj.get("requested_scope").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `requested_scope` to be an array in the JSON string but got `%s`", jsonObj.get("requested_scope").toString()));
      }
      if ((jsonObj.get("session_id") != null && !jsonObj.get("session_id").isJsonNull()) && !jsonObj.get("session_id").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `session_id` to be a primitive type in the JSON string but got `%s`", jsonObj.get("session_id").toString()));
      }
      if (!jsonObj.get("subject").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `subject` to be a primitive type in the JSON string but got `%s`", jsonObj.get("subject").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!OAuth2LoginRequest.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'OAuth2LoginRequest' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<OAuth2LoginRequest> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(OAuth2LoginRequest.class));

       return (TypeAdapter<T>) new TypeAdapter<OAuth2LoginRequest>() {
           @Override
           public void write(JsonWriter out, OAuth2LoginRequest value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public OAuth2LoginRequest read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of OAuth2LoginRequest given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of OAuth2LoginRequest
  * @throws IOException if the JSON string is invalid with respect to OAuth2LoginRequest
  */
  public static OAuth2LoginRequest fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, OAuth2LoginRequest.class);
  }

 /**
  * Convert an instance of OAuth2LoginRequest to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

