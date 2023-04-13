
##################FF_x64快速验证(build)
# go build ./edge/cmd/edgecore/
# ERR:verifying github.com/docker/distribution@v2.8.0+incompatible: checksum mismatch

# env
env; go version
cat /etc/passwd |grep root
cat /etc/group  |grep root

export GOSUMDB=off #noEffect
# go clean -modcache
# rm go.sum
# go mod tidy

# handBuildErr DO:
# GOSUMDB=off go get github.com/docker/docker/errdefs@v20.10.7+incompatible
# GOSUMDB=off go get k8s.io/kubernetes/pkg/kubelet/images@v1.23.15

# X64
# export CGO_ENABLED=0 #sqlite
GOSUMDB=off go build -o edgecore-x64 -v -ldflags "-s -w $flags" ./edge/cmd/edgecore/
ls -lh edgecore-*
# exit 0
# ARM64
GOARM="" # need to clear the value since golang compiler doesn't allow this env when building the binary for ARMv8.
# CC=aarch64-linux-gnu-gcc 
GOSUMDB=off GOARCH=arm64 GOOS="linux" CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc go build -o edgecore-arm64 -ldflags "$flags" ./edge/cmd/edgecore/

tar -zcf edgecore-x64.tar.gz edgecore-x64
tar -zcf edgecore-arm64.tar.gz edgecore-arm64
ls -lh edgecore-*
exit 0


##################hack/lib/golang.sh
kubeedge::version::get_version_info() {
  GIT_COMMIT=$(git rev-parse "HEAD^{commit}" 2>/dev/null)
  if git_status=$(git status --porcelain 2>/dev/null) && [[ -z ${git_status} ]]; then
    GIT_TREE_STATE="clean"
  else
    GIT_TREE_STATE="dirty"
  fi
  GIT_VERSION=$(git describe --tags --abbrev=14 "${GIT_COMMIT}^{commit}" 2>/dev/null)

  # This translates the "git describe" to an actual semver.org
  # compatible semantic version that looks something like this:
  #   v1.1.0-alpha.0.6+84c76d1142ea4d
  #
  # TODO: We continue calling this "git version" because so many
  # downstream consumers are expecting it there.
  #
  # These regexes are painful enough in sed...
  # We don't want to do them in pure shell, so disable SC2001
  # shellcheck disable=SC2001
  DASHES_IN_VERSION=$(echo "${GIT_VERSION}" | sed "s/[^-]//g")
  if [[ "${DASHES_IN_VERSION}" == "---" ]] ; then
    # shellcheck disable=SC2001
    # We have distance to subversion (v1.1.0-subversion-1-gCommitHash)
    GIT_VERSION=$(echo "${GIT_VERSION}" | sed "s/-\([0-9]\{1,\}\)-g\([0-9a-f]\{14\}\)$/.\1\+\2/")
  elif [[ "${DASHES_IN_VERSION}" == "--" ]] ; then
      # shellcheck disable=SC2001
      # We have distance to base tag (v1.1.0-1-gCommitHash)
      GIT_VERSION=$(echo "${GIT_VERSION}" | sed "s/-g\([0-9a-f]\{14\}\)$/+\1/")
  fi

  if [[ "${GIT_TREE_STATE}" == "dirty" ]]; then
    # git describe --dirty only considers changes to existing files, but
    # that is problematic since new untracked .go files affect the build,
    # so use our idea of "dirty" from git status instead.
    GIT_VERSION+="-dirty"
  fi

  # Try to match the "git describe" output to a regex to try to extract
  # the "major" and "minor" versions and whether this is the exact tagged
  # version or whether the tree is between two tagged versions.
  if [[ "${GIT_VERSION}" =~ ^v([0-9]+)\.([0-9]+)(\.[0-9]+)?([-].*)?([+].*)?$ ]]; then
    GIT_MAJOR=${BASH_REMATCH[1]}
    GIT_MINOR=${BASH_REMATCH[2]}
    if [[ -n "${BASH_REMATCH[4]}" ]]; then
      GIT_MINOR+="+"
    fi
  fi

  # If GIT_VERSION is not a valid Semantic Version, then refuse to build.
  if ! [[ "${GIT_VERSION}" =~ ^v([0-9]+)\.([0-9]+)(\.[0-9]+)?(-[0-9A-Za-z.-]+)?(\+[0-9A-Za-z.-]+)?$ ]]; then
      echo "GIT_VERSION should be a valid Semantic Version. Current value: ${GIT_VERSION}"
      echo "Please see more details here: https://semver.org"
      exit 1
  fi
}

# Get the value that needs to be passed to the -ldflags parameter of go build
kubeedge::version::ldflags() {
  kubeedge::version::get_version_info

  local -a ldflags
  function add_ldflag() {
    local key=${1}
    local val=${2}
    # If you update these, also update the list pkg/version/def.bzl.
    ldflags+=(
      "-X ${KUBEEDGE_GO_PACKAGE}/pkg/version.${key}=${val}"
    )
  }

  add_ldflag "buildDate" "$(date ${SOURCE_DATE_EPOCH:+"--date=@${SOURCE_DATE_EPOCH}"} -u +'%Y-%m-%dT%H:%M:%SZ')"
  if [[ -n ${GIT_COMMIT-} ]]; then
    add_ldflag "gitCommit" "${GIT_COMMIT}"
    add_ldflag "gitTreeState" "${GIT_TREE_STATE}"
  fi

  if [[ -n ${GIT_VERSION-} ]]; then
    add_ldflag "gitVersion" "${GIT_VERSION}"
  fi

  if [[ -n ${GIT_MAJOR-} && -n ${GIT_MINOR-} ]]; then
    add_ldflag "gitMajor" "${GIT_MAJOR}"
    add_ldflag "gitMinor" "${GIT_MINOR}"
  fi

  # The -ldflags parameter takes a single string, so join the output.
  echo "${ldflags[*]-}"
}

##kubeedge::golang::build_binaries() {  ###x64构建
local goldflags gogcflags
# If GOLDFLAGS is unset, then set it to the a default of "-s -w".
goldflags="${GOLDFLAGS=-s -w -buildid=} $(kubeedge::version::ldflags)"
gogcflags="${GOGCFLAGS:-}"
echo "building $bin"
local name="${bin##*/}"
set -x
go build -o ${KUBEEDGE_OUTPUT_BINPATH}/${name} -gcflags="${gogcflags:-}" -ldflags "${goldflags:-}" $bin
set +x


##kubeedge::golang::cross_build_place_binaries() {  #arm,arm64构建
local ldflags
read -r ldflags <<< "$(kubeedge::version::ldflags)"
echo "cross building $bin GOARM${goarm}"
local name="${bin##*/}"
# mark-c2
set -x
GOARM="" # need to clear the value since golang compiler doesn't allow this env when building the binary for ARMv8.
GOARCH=arm64 GOOS="linux" CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc go build -o ${KUBEEDGE_OUTPUT_BINPATH}/${name} -ldflags "$ldflags" $bin
set +x


##kubeedge::golang::small_build_place_binaries() { ##small, upx压缩版
local ldflags
read -r ldflags <<< "$(kubeedge::version::ldflags)"
set -x
go build -o ${KUBEEDGE_OUTPUT_BINPATH}/${name} -ldflags "-w -s -extldflags -static $ldflags" $bin
upx-ucl -9 ${KUBEEDGE_OUTPUT_BINPATH}/${name}
set +x

#IMG
GO_LDFLAGS="$(${KUBEEDGE_ROOT}/hack/make-rules/version.sh)"
set -x
docker build --build-arg GO_LDFLAGS="${GO_LDFLAGS}" -t kubeedge/${IMAGE_NAME}:${IMAGE_TAG} -f ${DOCKERFILE_PATH} .
set +x

#IMG-cross
IMAGE_NAME="$(get_imagename_by_target ${arg})"
DOCKERFILE_PATH="$(get_dockerfile_by_target ${arg})"
set -x
# If there's any issues when using buildx, can refer to the issue below
# https://github.com/docker/buildx/issues/495
# https://github.com/multiarch/qemu-user-static/issues/100
# docker run --rm --privileged multiarch/qemu-user-static --reset -p yes
docker buildx build --build-arg GO_LDFLAGS="${GO_LDFLAGS}" -t ${IMAGE_REPO_NAME}/${IMAGE_NAME}:${IMAGE_TAG} -f ${DOCKERFILE_PATH} --platform linux/amd64,linux/arm64,linux/arm/v7 --push .
set +x



