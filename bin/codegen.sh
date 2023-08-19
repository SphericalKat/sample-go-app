#!/bin/bash

script_dir=`(cd $(dirname $0); pwd)`

set -eux
cd ${script_dir}/../spec/
bash bin/validate.sh
bash bin/build.sh

cd ${script_dir}/../api/gen
oapi-codegen --config server.cfg.yaml ../../spec/build/sample/index.yaml
oapi-codegen --config types.cfg.yaml ../../spec/build/sample/index.yaml

cd ${script_dir}/..
# go generate --tags wireinject ./...