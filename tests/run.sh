#!/bin/sh

# tests are supposed to be located in the same directory as this file

DIR=$(readlink -f $(dirname $0))

export TESTING_HOST=${TESTING_HOST:="mender-reporting:8080"}

export PYTHONDONTWRITEBYTECODE=1

pip3 install --quiet --force-reinstall -r /testing/requirements.txt

# if we're running in a container, wait a little before starting tests
[ $$ -eq 1 ] && sleep 10

py.test -vv -s --tb=short --verbose \
        --junitxml=$DIR/results.xml \
        $DIR/tests "$@"
