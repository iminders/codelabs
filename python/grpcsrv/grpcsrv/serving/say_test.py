# -*- coding:utf-8 -*-

import logging
import unittest

from grpcsrv.serving.say import say


class TestSay(unittest.TestCase):
    @classmethod
    def setUpClass(self):
        self.msg = "Hi"

    def test_say(self):
        self.assertEqual("Hello world!", say(self.msg))
        self.assertEqual("Hello world!", say(None))


if __name__ == '__main__':
    logging.root.setLevel(logging.ERROR)
    unittest.main()
