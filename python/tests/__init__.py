import unittest
import test_message

def orca_test_suite():
    loader = unittest.TestLoader()
    suite = loader.loadTestsFromModule(test_message)
    return suite
