# coding: utf-8

"""
    ORY Hydra

    Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here.  # noqa: E501

    The version of the OpenAPI document: v1.8.0-pre.1
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from ory_hydra_client.configuration import Configuration


class OAuth2TokenIntrospection(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'active': 'bool',
        'aud': 'list[str]',
        'client_id': 'str',
        'exp': 'int',
        'ext': 'object',
        'iat': 'int',
        'iss': 'str',
        'nbf': 'int',
        'obfuscated_subject': 'str',
        'scope': 'str',
        'sub': 'str',
        'token_type': 'str',
        'username': 'str'
    }

    attribute_map = {
        'active': 'active',
        'aud': 'aud',
        'client_id': 'client_id',
        'exp': 'exp',
        'ext': 'ext',
        'iat': 'iat',
        'iss': 'iss',
        'nbf': 'nbf',
        'obfuscated_subject': 'obfuscated_subject',
        'scope': 'scope',
        'sub': 'sub',
        'token_type': 'token_type',
        'username': 'username'
    }

    def __init__(self, active=None, aud=None, client_id=None, exp=None, ext=None, iat=None, iss=None, nbf=None, obfuscated_subject=None, scope=None, sub=None, token_type=None, username=None, local_vars_configuration=None):  # noqa: E501
        """OAuth2TokenIntrospection - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._active = None
        self._aud = None
        self._client_id = None
        self._exp = None
        self._ext = None
        self._iat = None
        self._iss = None
        self._nbf = None
        self._obfuscated_subject = None
        self._scope = None
        self._sub = None
        self._token_type = None
        self._username = None
        self.discriminator = None

        self.active = active
        if aud is not None:
            self.aud = aud
        if client_id is not None:
            self.client_id = client_id
        if exp is not None:
            self.exp = exp
        if ext is not None:
            self.ext = ext
        if iat is not None:
            self.iat = iat
        if iss is not None:
            self.iss = iss
        if nbf is not None:
            self.nbf = nbf
        if obfuscated_subject is not None:
            self.obfuscated_subject = obfuscated_subject
        if scope is not None:
            self.scope = scope
        if sub is not None:
            self.sub = sub
        if token_type is not None:
            self.token_type = token_type
        if username is not None:
            self.username = username

    @property
    def active(self):
        """Gets the active of this OAuth2TokenIntrospection.  # noqa: E501

        Active is a boolean indicator of whether or not the presented token is currently active.  The specifics of a token's \"active\" state will vary depending on the implementation of the authorization server and the information it keeps about its tokens, but a \"true\" value return for the \"active\" property will generally indicate that a given token has been issued by this authorization server, has not been revoked by the resource owner, and is within its given time window of validity (e.g., after its issuance time and before its expiration time).  # noqa: E501

        :return: The active of this OAuth2TokenIntrospection.  # noqa: E501
        :rtype: bool
        """
        return self._active

    @active.setter
    def active(self, active):
        """Sets the active of this OAuth2TokenIntrospection.

        Active is a boolean indicator of whether or not the presented token is currently active.  The specifics of a token's \"active\" state will vary depending on the implementation of the authorization server and the information it keeps about its tokens, but a \"true\" value return for the \"active\" property will generally indicate that a given token has been issued by this authorization server, has not been revoked by the resource owner, and is within its given time window of validity (e.g., after its issuance time and before its expiration time).  # noqa: E501

        :param active: The active of this OAuth2TokenIntrospection.  # noqa: E501
        :type: bool
        """
        if self.local_vars_configuration.client_side_validation and active is None:  # noqa: E501
            raise ValueError("Invalid value for `active`, must not be `None`")  # noqa: E501

        self._active = active

    @property
    def aud(self):
        """Gets the aud of this OAuth2TokenIntrospection.  # noqa: E501

        Audience contains a list of the token's intended audiences.  # noqa: E501

        :return: The aud of this OAuth2TokenIntrospection.  # noqa: E501
        :rtype: list[str]
        """
        return self._aud

    @aud.setter
    def aud(self, aud):
        """Sets the aud of this OAuth2TokenIntrospection.

        Audience contains a list of the token's intended audiences.  # noqa: E501

        :param aud: The aud of this OAuth2TokenIntrospection.  # noqa: E501
        :type: list[str]
        """

        self._aud = aud

    @property
    def client_id(self):
        """Gets the client_id of this OAuth2TokenIntrospection.  # noqa: E501

        ID is aclient identifier for the OAuth 2.0 client that requested this token.  # noqa: E501

        :return: The client_id of this OAuth2TokenIntrospection.  # noqa: E501
        :rtype: str
        """
        return self._client_id

    @client_id.setter
    def client_id(self, client_id):
        """Sets the client_id of this OAuth2TokenIntrospection.

        ID is aclient identifier for the OAuth 2.0 client that requested this token.  # noqa: E501

        :param client_id: The client_id of this OAuth2TokenIntrospection.  # noqa: E501
        :type: str
        """

        self._client_id = client_id

    @property
    def exp(self):
        """Gets the exp of this OAuth2TokenIntrospection.  # noqa: E501

        Expires at is an integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this token will expire.  # noqa: E501

        :return: The exp of this OAuth2TokenIntrospection.  # noqa: E501
        :rtype: int
        """
        return self._exp

    @exp.setter
    def exp(self, exp):
        """Sets the exp of this OAuth2TokenIntrospection.

        Expires at is an integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this token will expire.  # noqa: E501

        :param exp: The exp of this OAuth2TokenIntrospection.  # noqa: E501
        :type: int
        """

        self._exp = exp

    @property
    def ext(self):
        """Gets the ext of this OAuth2TokenIntrospection.  # noqa: E501

        Extra is arbitrary data set by the session.  # noqa: E501

        :return: The ext of this OAuth2TokenIntrospection.  # noqa: E501
        :rtype: object
        """
        return self._ext

    @ext.setter
    def ext(self, ext):
        """Sets the ext of this OAuth2TokenIntrospection.

        Extra is arbitrary data set by the session.  # noqa: E501

        :param ext: The ext of this OAuth2TokenIntrospection.  # noqa: E501
        :type: object
        """

        self._ext = ext

    @property
    def iat(self):
        """Gets the iat of this OAuth2TokenIntrospection.  # noqa: E501

        Issued at is an integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this token was originally issued.  # noqa: E501

        :return: The iat of this OAuth2TokenIntrospection.  # noqa: E501
        :rtype: int
        """
        return self._iat

    @iat.setter
    def iat(self, iat):
        """Sets the iat of this OAuth2TokenIntrospection.

        Issued at is an integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this token was originally issued.  # noqa: E501

        :param iat: The iat of this OAuth2TokenIntrospection.  # noqa: E501
        :type: int
        """

        self._iat = iat

    @property
    def iss(self):
        """Gets the iss of this OAuth2TokenIntrospection.  # noqa: E501

        IssuerURL is a string representing the issuer of this token  # noqa: E501

        :return: The iss of this OAuth2TokenIntrospection.  # noqa: E501
        :rtype: str
        """
        return self._iss

    @iss.setter
    def iss(self, iss):
        """Sets the iss of this OAuth2TokenIntrospection.

        IssuerURL is a string representing the issuer of this token  # noqa: E501

        :param iss: The iss of this OAuth2TokenIntrospection.  # noqa: E501
        :type: str
        """

        self._iss = iss

    @property
    def nbf(self):
        """Gets the nbf of this OAuth2TokenIntrospection.  # noqa: E501

        NotBefore is an integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this token is not to be used before.  # noqa: E501

        :return: The nbf of this OAuth2TokenIntrospection.  # noqa: E501
        :rtype: int
        """
        return self._nbf

    @nbf.setter
    def nbf(self, nbf):
        """Sets the nbf of this OAuth2TokenIntrospection.

        NotBefore is an integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this token is not to be used before.  # noqa: E501

        :param nbf: The nbf of this OAuth2TokenIntrospection.  # noqa: E501
        :type: int
        """

        self._nbf = nbf

    @property
    def obfuscated_subject(self):
        """Gets the obfuscated_subject of this OAuth2TokenIntrospection.  # noqa: E501

        ObfuscatedSubject is set when the subject identifier algorithm was set to \"pairwise\" during authorization. It is the `sub` value of the ID Token that was issued.  # noqa: E501

        :return: The obfuscated_subject of this OAuth2TokenIntrospection.  # noqa: E501
        :rtype: str
        """
        return self._obfuscated_subject

    @obfuscated_subject.setter
    def obfuscated_subject(self, obfuscated_subject):
        """Sets the obfuscated_subject of this OAuth2TokenIntrospection.

        ObfuscatedSubject is set when the subject identifier algorithm was set to \"pairwise\" during authorization. It is the `sub` value of the ID Token that was issued.  # noqa: E501

        :param obfuscated_subject: The obfuscated_subject of this OAuth2TokenIntrospection.  # noqa: E501
        :type: str
        """

        self._obfuscated_subject = obfuscated_subject

    @property
    def scope(self):
        """Gets the scope of this OAuth2TokenIntrospection.  # noqa: E501

        Scope is a JSON string containing a space-separated list of scopes associated with this token.  # noqa: E501

        :return: The scope of this OAuth2TokenIntrospection.  # noqa: E501
        :rtype: str
        """
        return self._scope

    @scope.setter
    def scope(self, scope):
        """Sets the scope of this OAuth2TokenIntrospection.

        Scope is a JSON string containing a space-separated list of scopes associated with this token.  # noqa: E501

        :param scope: The scope of this OAuth2TokenIntrospection.  # noqa: E501
        :type: str
        """

        self._scope = scope

    @property
    def sub(self):
        """Gets the sub of this OAuth2TokenIntrospection.  # noqa: E501

        Subject of the token, as defined in JWT [RFC7519]. Usually a machine-readable identifier of the resource owner who authorized this token.  # noqa: E501

        :return: The sub of this OAuth2TokenIntrospection.  # noqa: E501
        :rtype: str
        """
        return self._sub

    @sub.setter
    def sub(self, sub):
        """Sets the sub of this OAuth2TokenIntrospection.

        Subject of the token, as defined in JWT [RFC7519]. Usually a machine-readable identifier of the resource owner who authorized this token.  # noqa: E501

        :param sub: The sub of this OAuth2TokenIntrospection.  # noqa: E501
        :type: str
        """

        self._sub = sub

    @property
    def token_type(self):
        """Gets the token_type of this OAuth2TokenIntrospection.  # noqa: E501

        TokenType is the introspected token's type, for example `access_token` or `refresh_token`.  # noqa: E501

        :return: The token_type of this OAuth2TokenIntrospection.  # noqa: E501
        :rtype: str
        """
        return self._token_type

    @token_type.setter
    def token_type(self, token_type):
        """Sets the token_type of this OAuth2TokenIntrospection.

        TokenType is the introspected token's type, for example `access_token` or `refresh_token`.  # noqa: E501

        :param token_type: The token_type of this OAuth2TokenIntrospection.  # noqa: E501
        :type: str
        """

        self._token_type = token_type

    @property
    def username(self):
        """Gets the username of this OAuth2TokenIntrospection.  # noqa: E501

        Username is a human-readable identifier for the resource owner who authorized this token.  # noqa: E501

        :return: The username of this OAuth2TokenIntrospection.  # noqa: E501
        :rtype: str
        """
        return self._username

    @username.setter
    def username(self, username):
        """Sets the username of this OAuth2TokenIntrospection.

        Username is a human-readable identifier for the resource owner who authorized this token.  # noqa: E501

        :param username: The username of this OAuth2TokenIntrospection.  # noqa: E501
        :type: str
        """

        self._username = username

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, OAuth2TokenIntrospection):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, OAuth2TokenIntrospection):
            return True

        return self.to_dict() != other.to_dict()
