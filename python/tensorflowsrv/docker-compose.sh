set -e
cd "$(dirname "$0")"

export TEST_RUNID=$RANDOM
echo "test run id: $TEST_RUNID"

cd /root/tensorflowsrv
pip install -e .
# 忽略warnings
python -W ignore:ImportWarning -m pytest
