set -e
cd "$(dirname "$0")"

export BAZEL_RUNID=$RANDOM

bazel --batch \
      --output_user_root=/root/cache/bazel \
      test --package_path=/root/pylab \
           --spawn_strategy=standalone \
           --genrule_strategy=standalone \
           --action_env="TUSHARE_TOKEN=$TUSHARE_TOKEN" \
           --test_output=errors \
           -c opt \
           -- //...
