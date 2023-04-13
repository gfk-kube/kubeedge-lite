#!/usr/bin/env bash

###
#Copyright 2022 The KubeEdge Authors.
#
#Licensed under the Apache License, Version 2.0 (the "License");
#you may not use this file except in compliance with the License.
#You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
#Unless required by applicable law or agreed to in writing, software
#distributed under the License is distributed on an "AS IS" BASIS,
#WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#See the License for the specific language governing permissions and
#limitations under the License.
###

# set -o errexit
# set -o nounset
# set -o pipefail

KUBEEDGE_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/" && pwd -P)" #/../..
MOUNTPATH="${MOUNTPATH:-/kubeedge}"
KUBEEDGE_BUILD_IMAGE=${KUBEEDGE_BUILD_IMAGE:-"kubeedge/build-tools:1.17.13-ke1"}
DOCKER_GID="${DOCKER_GID:-$(grep '^docker:' /etc/group | cut -f3 -d:)}"
CONTAINER_RUN_OPTIONS="${CONTAINER_RUN_OPTIONS:--it}"
# 
DOCKER_GID=0 #root
UID=0 #默认的: mkdir /go 无权限;

echo "start building inside container"
echo docker run --rm ${CONTAINER_RUN_OPTIONS} -u "${UID}:${DOCKER_GID}" \
    --init \
    --sig-proxy=true \
    -e XDG_CACHE_HOME=/tmp/.cache \
    -v /_ext/ke_go:/go \
    -v ${KUBEEDGE_ROOT}:${MOUNTPATH} \
    -w ${MOUNTPATH} ${KUBEEDGE_BUILD_IMAGE} "$@"
# docker run --rm -it -u 0:0 --init --sig-proxy=true -e XDG_CACHE_HOME=/tmp/.cache -v /_ext/working/_ct/fk-kubeedge-bfg-b1:/kubeedge -w /kubeedge kubeedge/build-tools:1.17.13-ke1 ./build.sh
# # headless @ mac23-199 in .../_ct/fk-kubeedge-bfg-b1 |21:23:42  |fix-offline-pods ?:6 ✗| 
# $ git pull; docker run --rm -it -u 0:0 --init --sig-proxy=true -e XDG_CACHE_HOME=/tmp/.cache -v /_ext/ke_go:/go -v /_ext/working/_ct/fk-kubeedge-bfg-b1:/kubeedge -w /kubeedge kubeedge/build-tools:1.17.13-ke1 bash

# exit 0

docker run --rm ${CONTAINER_RUN_OPTIONS} -u "${UID}:${DOCKER_GID}" \
    --init \
    --sig-proxy=true \
    -e XDG_CACHE_HOME=/tmp/.cache \
    -v /_ext/ke_go:/go \
    -v ${KUBEEDGE_ROOT}:${MOUNTPATH} \
    -w ${MOUNTPATH} ${KUBEEDGE_BUILD_IMAGE} "$@"
