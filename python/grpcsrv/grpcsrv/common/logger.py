import logging
import os

logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s %(filename)s[%(lineno)d] %(levelname)s %(message)s')

logger = logging.getLogger()

logger_dir = os.path.join("/tmp", "codelab")
if not os.path.exists(logger_dir):
    os.makedirs(logger_dir)

handler = logging.FileHandler(os.path.join(logger_dir, "grpcsrv.log"))
handler.setLevel(logging.INFO)
formatter = logging.Formatter(
    '%(asctime)s %(filename)s[%(lineno)d] %(levelname)s %(message)s')
handler.setFormatter(formatter)
logger.addHandler(handler)
