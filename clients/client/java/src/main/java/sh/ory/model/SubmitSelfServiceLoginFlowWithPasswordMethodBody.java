/*
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.2.0-alpha.20
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

/**
 * SubmitSelfServiceLoginFlowWithPasswordMethodBody
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-08-26T08:30:03.017335438Z[Etc/UTC]")
public class SubmitSelfServiceLoginFlowWithPasswordMethodBody {
  public static final String SERIALIZED_NAME_CSRF_TOKEN = "csrf_token";
  @SerializedName(SERIALIZED_NAME_CSRF_TOKEN)
  private String csrfToken;

  public static final String SERIALIZED_NAME_IDENTIFIER = "identifier";
  @SerializedName(SERIALIZED_NAME_IDENTIFIER)
  private String identifier;

  public static final String SERIALIZED_NAME_METHOD = "method";
  @SerializedName(SERIALIZED_NAME_METHOD)
  private String method;

  public static final String SERIALIZED_NAME_PASSWORD = "password";
  @SerializedName(SERIALIZED_NAME_PASSWORD)
  private String password;

  public static final String SERIALIZED_NAME_PASSWORD_IDENTIFIER = "password_identifier";
  @SerializedName(SERIALIZED_NAME_PASSWORD_IDENTIFIER)
  private String passwordIdentifier;

  public SubmitSelfServiceLoginFlowWithPasswordMethodBody() { 
  }

  public SubmitSelfServiceLoginFlowWithPasswordMethodBody csrfToken(String csrfToken) {
    
    this.csrfToken = csrfToken;
    return this;
  }

   /**
   * Sending the anti-csrf token is only required for browser login flows.
   * @return csrfToken
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Sending the anti-csrf token is only required for browser login flows.")

  public String getCsrfToken() {
    return csrfToken;
  }


  public void setCsrfToken(String csrfToken) {
    this.csrfToken = csrfToken;
  }


  public SubmitSelfServiceLoginFlowWithPasswordMethodBody identifier(String identifier) {
    
    this.identifier = identifier;
    return this;
  }

   /**
   * Identifier is the email or username of the user trying to log in.
   * @return identifier
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Identifier is the email or username of the user trying to log in.")

  public String getIdentifier() {
    return identifier;
  }


  public void setIdentifier(String identifier) {
    this.identifier = identifier;
  }


  public SubmitSelfServiceLoginFlowWithPasswordMethodBody method(String method) {
    
    this.method = method;
    return this;
  }

   /**
   * Method should be set to \&quot;password\&quot; when logging in using the identifier and password strategy.
   * @return method
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Method should be set to \"password\" when logging in using the identifier and password strategy.")

  public String getMethod() {
    return method;
  }


  public void setMethod(String method) {
    this.method = method;
  }


  public SubmitSelfServiceLoginFlowWithPasswordMethodBody password(String password) {
    
    this.password = password;
    return this;
  }

   /**
   * The user&#39;s password.
   * @return password
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "The user's password.")

  public String getPassword() {
    return password;
  }


  public void setPassword(String password) {
    this.password = password;
  }


  public SubmitSelfServiceLoginFlowWithPasswordMethodBody passwordIdentifier(String passwordIdentifier) {
    
    this.passwordIdentifier = passwordIdentifier;
    return this;
  }

   /**
   * Identifier is the email or username of the user trying to log in. This field is deprecated!
   * @return passwordIdentifier
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Identifier is the email or username of the user trying to log in. This field is deprecated!")

  public String getPasswordIdentifier() {
    return passwordIdentifier;
  }


  public void setPasswordIdentifier(String passwordIdentifier) {
    this.passwordIdentifier = passwordIdentifier;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SubmitSelfServiceLoginFlowWithPasswordMethodBody submitSelfServiceLoginFlowWithPasswordMethodBody = (SubmitSelfServiceLoginFlowWithPasswordMethodBody) o;
    return Objects.equals(this.csrfToken, submitSelfServiceLoginFlowWithPasswordMethodBody.csrfToken) &&
        Objects.equals(this.identifier, submitSelfServiceLoginFlowWithPasswordMethodBody.identifier) &&
        Objects.equals(this.method, submitSelfServiceLoginFlowWithPasswordMethodBody.method) &&
        Objects.equals(this.password, submitSelfServiceLoginFlowWithPasswordMethodBody.password) &&
        Objects.equals(this.passwordIdentifier, submitSelfServiceLoginFlowWithPasswordMethodBody.passwordIdentifier);
  }

  @Override
  public int hashCode() {
    return Objects.hash(csrfToken, identifier, method, password, passwordIdentifier);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SubmitSelfServiceLoginFlowWithPasswordMethodBody {\n");
    sb.append("    csrfToken: ").append(toIndentedString(csrfToken)).append("\n");
    sb.append("    identifier: ").append(toIndentedString(identifier)).append("\n");
    sb.append("    method: ").append(toIndentedString(method)).append("\n");
    sb.append("    password: ").append(toIndentedString(password)).append("\n");
    sb.append("    passwordIdentifier: ").append(toIndentedString(passwordIdentifier)).append("\n");
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

}

