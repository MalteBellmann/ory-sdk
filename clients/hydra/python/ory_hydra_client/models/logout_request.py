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


class LogoutRequest(object):
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
        'request_url': 'str',
        'rp_initiated': 'bool',
        'sid': 'str',
        'subject': 'str'
    }

    attribute_map = {
        'request_url': 'request_url',
        'rp_initiated': 'rp_initiated',
        'sid': 'sid',
        'subject': 'subject'
    }

    def __init__(self, request_url=None, rp_initiated=None, sid=None, subject=None, local_vars_configuration=None):  # noqa: E501
        """LogoutRequest - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._request_url = None
        self._rp_initiated = None
        self._sid = None
        self._subject = None
        self.discriminator = None

        if request_url is not None:
            self.request_url = request_url
        if rp_initiated is not None:
            self.rp_initiated = rp_initiated
        if sid is not None:
            self.sid = sid
        if subject is not None:
            self.subject = subject

    @property
    def request_url(self):
        """Gets the request_url of this LogoutRequest.  # noqa: E501

        RequestURL is the original Logout URL requested.  # noqa: E501

        :return: The request_url of this LogoutRequest.  # noqa: E501
        :rtype: str
        """
        return self._request_url

    @request_url.setter
    def request_url(self, request_url):
        """Sets the request_url of this LogoutRequest.

        RequestURL is the original Logout URL requested.  # noqa: E501

        :param request_url: The request_url of this LogoutRequest.  # noqa: E501
        :type: str
        """

        self._request_url = request_url

    @property
    def rp_initiated(self):
        """Gets the rp_initiated of this LogoutRequest.  # noqa: E501

        RPInitiated is set to true if the request was initiated by a Relying Party (RP), also known as an OAuth 2.0 Client.  # noqa: E501

        :return: The rp_initiated of this LogoutRequest.  # noqa: E501
        :rtype: bool
        """
        return self._rp_initiated

    @rp_initiated.setter
    def rp_initiated(self, rp_initiated):
        """Sets the rp_initiated of this LogoutRequest.

        RPInitiated is set to true if the request was initiated by a Relying Party (RP), also known as an OAuth 2.0 Client.  # noqa: E501

        :param rp_initiated: The rp_initiated of this LogoutRequest.  # noqa: E501
        :type: bool
        """

        self._rp_initiated = rp_initiated

    @property
    def sid(self):
        """Gets the sid of this LogoutRequest.  # noqa: E501

        SessionID is the login session ID that was requested to log out.  # noqa: E501

        :return: The sid of this LogoutRequest.  # noqa: E501
        :rtype: str
        """
        return self._sid

    @sid.setter
    def sid(self, sid):
        """Sets the sid of this LogoutRequest.

        SessionID is the login session ID that was requested to log out.  # noqa: E501

        :param sid: The sid of this LogoutRequest.  # noqa: E501
        :type: str
        """

        self._sid = sid

    @property
    def subject(self):
        """Gets the subject of this LogoutRequest.  # noqa: E501

        Subject is the user for whom the logout was request.  # noqa: E501

        :return: The subject of this LogoutRequest.  # noqa: E501
        :rtype: str
        """
        return self._subject

    @subject.setter
    def subject(self, subject):
        """Sets the subject of this LogoutRequest.

        Subject is the user for whom the logout was request.  # noqa: E501

        :param subject: The subject of this LogoutRequest.  # noqa: E501
        :type: str
        """

        self._subject = subject

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
        if not isinstance(other, LogoutRequest):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, LogoutRequest):
            return True

        return self.to_dict() != other.to_dict()
