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
from ory_client.model.continue_with_recovery_ui import ContinueWithRecoveryUi
from ory_client.model.continue_with_recovery_ui_flow import ContinueWithRecoveryUiFlow
from ory_client.model.continue_with_set_ory_session_token import ContinueWithSetOrySessionToken
from ory_client.model.continue_with_settings_ui import ContinueWithSettingsUi
from ory_client.model.continue_with_verification_ui import ContinueWithVerificationUi
globals()['ContinueWithRecoveryUi'] = ContinueWithRecoveryUi
globals()['ContinueWithRecoveryUiFlow'] = ContinueWithRecoveryUiFlow
globals()['ContinueWithSetOrySessionToken'] = ContinueWithSetOrySessionToken
globals()['ContinueWithSettingsUi'] = ContinueWithSettingsUi
globals()['ContinueWithVerificationUi'] = ContinueWithVerificationUi
from ory_client.model.continue_with import ContinueWith


class TestContinueWith(unittest.TestCase):
    """ContinueWith unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testContinueWith(self):
        """Test ContinueWith"""
        # FIXME: construct object with mandatory attributes with example values
        # model = ContinueWith()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
