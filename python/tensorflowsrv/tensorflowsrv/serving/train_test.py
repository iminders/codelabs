# -*- coding:utf-8 -*-

import logging
import unittest


class TestNone(unittest.TestCase):
    @classmethod
    def setUpClass(self):
        self.msg = "Hi"

    def test_nothing(self):
        self.assertEqual(1, 1)


if __name__ == '__main__':
    logging.root.setLevel(logging.ERROR)
    unittest.main()
