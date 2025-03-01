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
from ory_client.model.workspace import Workspace
globals()['Workspace'] = Workspace
from ory_client.model.workspaces import Workspaces


class TestWorkspaces(unittest.TestCase):
    """Workspaces unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testWorkspaces(self):
        """Test Workspaces"""
        # FIXME: construct object with mandatory attributes with example values
        # model = Workspaces()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
