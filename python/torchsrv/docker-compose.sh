set -e
cd "$(dirname "$0")"

export TEST_RUNID=$RANDOM
echo "test run id: $TEST_RUNID"

cd /root/grpcsrv
pip install -e .
python -m pytest
