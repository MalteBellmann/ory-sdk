# coding: utf-8

"""
    ORY Hydra

    Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here.  # noqa: E501

    The version of the OpenAPI document: v1.9.0-alpha.3
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from ory_hydra_client.configuration import Configuration


class PluginMount(object):
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
        'description': 'str',
        'destination': 'str',
        'name': 'str',
        'options': 'list[str]',
        'settable': 'list[str]',
        'source': 'str',
        'type': 'str'
    }

    attribute_map = {
        'description': 'Description',
        'destination': 'Destination',
        'name': 'Name',
        'options': 'Options',
        'settable': 'Settable',
        'source': 'Source',
        'type': 'Type'
    }

    def __init__(self, description=None, destination=None, name=None, options=None, settable=None, source=None, type=None, local_vars_configuration=None):  # noqa: E501
        """PluginMount - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._description = None
        self._destination = None
        self._name = None
        self._options = None
        self._settable = None
        self._source = None
        self._type = None
        self.discriminator = None

        self.description = description
        self.destination = destination
        self.name = name
        self.options = options
        self.settable = settable
        self.source = source
        self.type = type

    @property
    def description(self):
        """Gets the description of this PluginMount.  # noqa: E501

        description  # noqa: E501

        :return: The description of this PluginMount.  # noqa: E501
        :rtype: str
        """
        return self._description

    @description.setter
    def description(self, description):
        """Sets the description of this PluginMount.

        description  # noqa: E501

        :param description: The description of this PluginMount.  # noqa: E501
        :type: str
        """
        if self.local_vars_configuration.client_side_validation and description is None:  # noqa: E501
            raise ValueError("Invalid value for `description`, must not be `None`")  # noqa: E501

        self._description = description

    @property
    def destination(self):
        """Gets the destination of this PluginMount.  # noqa: E501

        destination  # noqa: E501

        :return: The destination of this PluginMount.  # noqa: E501
        :rtype: str
        """
        return self._destination

    @destination.setter
    def destination(self, destination):
        """Sets the destination of this PluginMount.

        destination  # noqa: E501

        :param destination: The destination of this PluginMount.  # noqa: E501
        :type: str
        """
        if self.local_vars_configuration.client_side_validation and destination is None:  # noqa: E501
            raise ValueError("Invalid value for `destination`, must not be `None`")  # noqa: E501

        self._destination = destination

    @property
    def name(self):
        """Gets the name of this PluginMount.  # noqa: E501

        name  # noqa: E501

        :return: The name of this PluginMount.  # noqa: E501
        :rtype: str
        """
        return self._name

    @name.setter
    def name(self, name):
        """Sets the name of this PluginMount.

        name  # noqa: E501

        :param name: The name of this PluginMount.  # noqa: E501
        :type: str
        """
        if self.local_vars_configuration.client_side_validation and name is None:  # noqa: E501
            raise ValueError("Invalid value for `name`, must not be `None`")  # noqa: E501

        self._name = name

    @property
    def options(self):
        """Gets the options of this PluginMount.  # noqa: E501

        options  # noqa: E501

        :return: The options of this PluginMount.  # noqa: E501
        :rtype: list[str]
        """
        return self._options

    @options.setter
    def options(self, options):
        """Sets the options of this PluginMount.

        options  # noqa: E501

        :param options: The options of this PluginMount.  # noqa: E501
        :type: list[str]
        """
        if self.local_vars_configuration.client_side_validation and options is None:  # noqa: E501
            raise ValueError("Invalid value for `options`, must not be `None`")  # noqa: E501

        self._options = options

    @property
    def settable(self):
        """Gets the settable of this PluginMount.  # noqa: E501

        settable  # noqa: E501

        :return: The settable of this PluginMount.  # noqa: E501
        :rtype: list[str]
        """
        return self._settable

    @settable.setter
    def settable(self, settable):
        """Sets the settable of this PluginMount.

        settable  # noqa: E501

        :param settable: The settable of this PluginMount.  # noqa: E501
        :type: list[str]
        """
        if self.local_vars_configuration.client_side_validation and settable is None:  # noqa: E501
            raise ValueError("Invalid value for `settable`, must not be `None`")  # noqa: E501

        self._settable = settable

    @property
    def source(self):
        """Gets the source of this PluginMount.  # noqa: E501

        source  # noqa: E501

        :return: The source of this PluginMount.  # noqa: E501
        :rtype: str
        """
        return self._source

    @source.setter
    def source(self, source):
        """Sets the source of this PluginMount.

        source  # noqa: E501

        :param source: The source of this PluginMount.  # noqa: E501
        :type: str
        """
        if self.local_vars_configuration.client_side_validation and source is None:  # noqa: E501
            raise ValueError("Invalid value for `source`, must not be `None`")  # noqa: E501

        self._source = source

    @property
    def type(self):
        """Gets the type of this PluginMount.  # noqa: E501

        type  # noqa: E501

        :return: The type of this PluginMount.  # noqa: E501
        :rtype: str
        """
        return self._type

    @type.setter
    def type(self, type):
        """Sets the type of this PluginMount.

        type  # noqa: E501

        :param type: The type of this PluginMount.  # noqa: E501
        :type: str
        """
        if self.local_vars_configuration.client_side_validation and type is None:  # noqa: E501
            raise ValueError("Invalid value for `type`, must not be `None`")  # noqa: E501

        self._type = type

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
        if not isinstance(other, PluginMount):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, PluginMount):
            return True

        return self.to_dict() != other.to_dict()
