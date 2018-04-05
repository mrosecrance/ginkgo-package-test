#!/bin/bash
set -e

cd /var/vcap/packages/cf-redis-smoke-tests/src/github.com/pivotal-cf/cf-redis-smoke-tests
/var/vcap/packages/ginkgo -r

