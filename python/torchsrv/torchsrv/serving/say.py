# -*- coding:utf-8 -*-

from grpcsrv.common.logger import logger


def say(utterance):
    msg = "Hello world!"
    logger.info(msg)
    return msg
