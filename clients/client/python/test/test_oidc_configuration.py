"""
    Ory APIs

    Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.   # noqa: E501

    The version of the OpenAPI document: v1.6.1
    Contact: support@ory.sh
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import ory_client
from ory_client.model.credential_supported_draft00 import CredentialSupportedDraft00
globals()['CredentialSupportedDraft00'] = CredentialSupportedDraft00
from ory_client.model.oidc_configuration import OidcConfiguration


class TestOidcConfiguration(unittest.TestCase):
    """OidcConfiguration unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testOidcConfiguration(self):
        """Test OidcConfiguration"""
        # FIXME: construct object with mandatory attributes with example values
        # model = OidcConfiguration()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
