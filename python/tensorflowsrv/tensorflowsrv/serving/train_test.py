# -*- coding:utf-8 -*-

import logging
import unittest

from tensorflowsrv.serving.train import MODEL_DIR


class TestTrain(unittest.TestCase):
    @classmethod
    def setUpClass(self):
        self.msg = "Hi"

    def test_nothing(self):
        self.assertEqual(MODEL_DIR, "testdata/faishon")


if __name__ == '__main__':
    logging.root.setLevel(logging.ERROR)
    unittest.main()
