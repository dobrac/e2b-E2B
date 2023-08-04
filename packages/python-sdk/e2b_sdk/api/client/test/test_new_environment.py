# coding: utf-8

"""
    Devbook

    Devbook API  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


import unittest
import datetime

import client
from client.models.new_environment import NewEnvironment  # noqa: E501
from client.rest import ApiException


class TestNewEnvironment(unittest.TestCase):
    """NewEnvironment unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test NewEnvironment
        include_option is a boolean, when False only required
        params are included, when True both required and
        optional params are included"""
        # uncomment below to create an instance of `NewEnvironment`
        """
        model = client.models.new_environment.NewEnvironment()  # noqa: E501
        if include_optional :
            return NewEnvironment(
                title = '', 
                template = ''
            )
        else :
            return NewEnvironment(
                template = '',
        )
        """

    def testNewEnvironment(self):
        """Test NewEnvironment"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == "__main__":
    unittest.main()