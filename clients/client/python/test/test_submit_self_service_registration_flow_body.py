"""
    Ory APIs

    Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.   # noqa: E501

    The version of the OpenAPI document: v0.2.0-alpha.20
    Contact: support@ory.sh
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import ory_client
from ory_client.model.submit_self_service_registration_flow_with_oidc_method_body import SubmitSelfServiceRegistrationFlowWithOidcMethodBody
from ory_client.model.submit_self_service_registration_flow_with_password_method_body import SubmitSelfServiceRegistrationFlowWithPasswordMethodBody
from ory_client.model.submit_self_service_registration_flow_with_web_authn_method_body import SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody
globals()['SubmitSelfServiceRegistrationFlowWithOidcMethodBody'] = SubmitSelfServiceRegistrationFlowWithOidcMethodBody
globals()['SubmitSelfServiceRegistrationFlowWithPasswordMethodBody'] = SubmitSelfServiceRegistrationFlowWithPasswordMethodBody
globals()['SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody'] = SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody
from ory_client.model.submit_self_service_registration_flow_body import SubmitSelfServiceRegistrationFlowBody


class TestSubmitSelfServiceRegistrationFlowBody(unittest.TestCase):
    """SubmitSelfServiceRegistrationFlowBody unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testSubmitSelfServiceRegistrationFlowBody(self):
        """Test SubmitSelfServiceRegistrationFlowBody"""
        # FIXME: construct object with mandatory attributes with example values
        # model = SubmitSelfServiceRegistrationFlowBody()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
