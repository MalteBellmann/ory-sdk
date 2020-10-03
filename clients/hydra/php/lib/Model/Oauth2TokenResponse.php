<?php
/**
 * Oauth2TokenResponse
 *
 * PHP version 5
 *
 * @category Class
 * @package  Ory\Hydra\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * ORY Hydra
 *
 * Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here.
 *
 * The version of the OpenAPI document: v1.8.0-pre.1
 * 
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 4.3.1
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace Ory\Hydra\Client\Model;

use \ArrayAccess;
use \Ory\Hydra\Client\ObjectSerializer;

/**
 * Oauth2TokenResponse Class Doc Comment
 *
 * @category Class
 * @description The Access Token Response
 * @package  Ory\Hydra\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */
class Oauth2TokenResponse implements ModelInterface, ArrayAccess
{
    const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'oauth2TokenResponse';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'accessToken' => 'string',
        'expiresIn' => 'int',
        'idToken' => 'string',
        'refreshToken' => 'string',
        'scope' => 'string',
        'tokenType' => 'string'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPIFormats = [
        'accessToken' => null,
        'expiresIn' => 'int64',
        'idToken' => null,
        'refreshToken' => null,
        'scope' => null,
        'tokenType' => null
    ];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPITypes()
    {
        return self::$openAPITypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPIFormats()
    {
        return self::$openAPIFormats;
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'accessToken' => 'access_token',
        'expiresIn' => 'expires_in',
        'idToken' => 'id_token',
        'refreshToken' => 'refresh_token',
        'scope' => 'scope',
        'tokenType' => 'token_type'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'accessToken' => 'setAccessToken',
        'expiresIn' => 'setExpiresIn',
        'idToken' => 'setIdToken',
        'refreshToken' => 'setRefreshToken',
        'scope' => 'setScope',
        'tokenType' => 'setTokenType'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'accessToken' => 'getAccessToken',
        'expiresIn' => 'getExpiresIn',
        'idToken' => 'getIdToken',
        'refreshToken' => 'getRefreshToken',
        'scope' => 'getScope',
        'tokenType' => 'getTokenType'
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * The original name of the model.
     *
     * @return string
     */
    public function getModelName()
    {
        return self::$openAPIModelName;
    }

    

    

    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     *                      initializing the model
     */
    public function __construct(array $data = null)
    {
        $this->container['accessToken'] = isset($data['accessToken']) ? $data['accessToken'] : null;
        $this->container['expiresIn'] = isset($data['expiresIn']) ? $data['expiresIn'] : null;
        $this->container['idToken'] = isset($data['idToken']) ? $data['idToken'] : null;
        $this->container['refreshToken'] = isset($data['refreshToken']) ? $data['refreshToken'] : null;
        $this->container['scope'] = isset($data['scope']) ? $data['scope'] : null;
        $this->container['tokenType'] = isset($data['tokenType']) ? $data['tokenType'] : null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }


    /**
     * Gets accessToken
     *
     * @return string|null
     */
    public function getAccessToken()
    {
        return $this->container['accessToken'];
    }

    /**
     * Sets accessToken
     *
     * @param string|null $accessToken accessToken
     *
     * @return $this
     */
    public function setAccessToken($accessToken)
    {
        $this->container['accessToken'] = $accessToken;

        return $this;
    }

    /**
     * Gets expiresIn
     *
     * @return int|null
     */
    public function getExpiresIn()
    {
        return $this->container['expiresIn'];
    }

    /**
     * Sets expiresIn
     *
     * @param int|null $expiresIn expiresIn
     *
     * @return $this
     */
    public function setExpiresIn($expiresIn)
    {
        $this->container['expiresIn'] = $expiresIn;

        return $this;
    }

    /**
     * Gets idToken
     *
     * @return string|null
     */
    public function getIdToken()
    {
        return $this->container['idToken'];
    }

    /**
     * Sets idToken
     *
     * @param string|null $idToken idToken
     *
     * @return $this
     */
    public function setIdToken($idToken)
    {
        $this->container['idToken'] = $idToken;

        return $this;
    }

    /**
     * Gets refreshToken
     *
     * @return string|null
     */
    public function getRefreshToken()
    {
        return $this->container['refreshToken'];
    }

    /**
     * Sets refreshToken
     *
     * @param string|null $refreshToken refreshToken
     *
     * @return $this
     */
    public function setRefreshToken($refreshToken)
    {
        $this->container['refreshToken'] = $refreshToken;

        return $this;
    }

    /**
     * Gets scope
     *
     * @return string|null
     */
    public function getScope()
    {
        return $this->container['scope'];
    }

    /**
     * Sets scope
     *
     * @param string|null $scope scope
     *
     * @return $this
     */
    public function setScope($scope)
    {
        $this->container['scope'] = $scope;

        return $this;
    }

    /**
     * Gets tokenType
     *
     * @return string|null
     */
    public function getTokenType()
    {
        return $this->container['tokenType'];
    }

    /**
     * Sets tokenType
     *
     * @param string|null $tokenType tokenType
     *
     * @return $this
     */
    public function setTokenType($tokenType)
    {
        $this->container['tokenType'] = $tokenType;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param integer $offset Offset
     *
     * @return mixed
     */
    public function offsetGet($offset)
    {
        return isset($this->container[$offset]) ? $this->container[$offset] : null;
    }

    /**
     * Sets value based on offset.
     *
     * @param integer $offset Offset
     * @param mixed   $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param integer $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }

    /**
     * Gets the string presentation of the object
     *
     * @return string
     */
    public function __toString()
    {
        return json_encode(
            ObjectSerializer::sanitizeForSerialization($this),
            JSON_PRETTY_PRINT
        );
    }

    /**
     * Gets a header-safe presentation of the object
     *
     * @return string
     */
    public function toHeaderValue()
    {
        return json_encode(ObjectSerializer::sanitizeForSerialization($this));
    }
}


