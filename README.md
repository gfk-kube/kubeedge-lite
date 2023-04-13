

**build-X64**

- `git pull; docker run --rm -it -u 0:0 --init --sig-proxy=true -e XDG_CACHE_HOME=/tmp/.cache -v /_ext/ke_go:/go -v /_ext/working/_ct/fk-kubeedge-bfg-b1:/kubeedge -w /kubeedge kubeedge/build-tools:1.17.13-ke1 bash`

```bash
# gosum不对>> 手动删；


# go build err:
/go/pkg/mod/github.com/kubeedge/kubernetes@v1.23.15-kubeedge1/pkg/util/parsers/parsers.go:26:2: missing go.sum entry for module providing package github.com/docker/distribution/reference (imported by k8s.io/kubernetes/pkg/kubelet/images); to add:
	go get k8s.io/kubernetes/pkg/kubelet/images@v1.23.15
/go/pkg/mod/github.com/docker/docker@v20.10.7+incompatible/errdefs/http_helpers.go:8:2: missing go.sum entry for module providing package github.com/docker/distribution/registry/api/errcode (imported by github.com/docker/docker/errdefs); to add:
	go get github.com/docker/docker/errdefs@v20.10.7+incompatible


# go get err:
root@199a722a2423:/kubeedge# go get k8s.io/kubernetes/pkg/kubelet/images@v1.23.15
github.com/docker/distribution@v2.8.0+incompatible: verifying module: checksum mismatch
	downloaded: h1:u9vuu6qqG7nN9a735Noed0ahoUm30iipVRlhgh72N0M=
	sum.golang.org: h1:l9EaZDICImO1ngI+uTifW+ZYvvz7fKISBAKpg+MbWbY=

SECURITY ERROR
This download does NOT match the one reported by the checksum server.
The bits may have been replaced on the origin server, or an attacker may
have intercepted the download attempt.

For more information, see 'go help module-auth'.
root@199a722a2423:/kubeedge# go get github.com/docker/docker/errdefs@v20.10.7+incompatible
github.com/docker/distribution@v2.8.0+incompatible: verifying module: checksum mismatch
	downloaded: h1:u9vuu6qqG7nN9a735Noed0ahoUm30iipVRlhgh72N0M=
	sum.golang.org: h1:l9EaZDICImO1ngI+uTifW+ZYvvz7fKISBAKpg+MbWbY=

SECURITY ERROR
This download does NOT match the one reported by the checksum server.
The bits may have been replaced on the origin server, or an attacker may
have intercepted the download attempt.

For more information, see 'go help module-auth'.

# GOSUMDB=off 
root@199a722a2423:/kubeedge# GOSUMDB=off go get github.com/docker/docker/errdefs@v20.10.7+incompatible
go: module github.com/golang/protobuf is deprecated: Use the "google.golang.org/protobuf" module instead.
root@199a722a2423:/kubeedge# GOSUMDB=off go get k8s.io/kubernetes/pkg/kubelet/images@v1.23.15
go: module github.com/golang/protobuf is deprecated: Use the "google.golang.org/protobuf" module instead.
root@199a722a2423:/kubeedge# 

# 手动buildOK
```

**build-arm64**

```bash
# 同23.199一样错：
# root@199a722a2423:/kubeedge# GOARCH=arm64 GOOS="linux" CGO_ENABLED=1 go build -o edgecore-arm64 -ldflags "$flags" ./edge/cmd/edgecore/
# runtime/cgo
gcc_arm64.S: Assembler messages:
gcc_arm64.S:28: Error: no such instruction: `stp x29,x30,[sp,'
gcc_arm64.S:32: Error: too many memory references for `mov'
gcc_arm64.S:34: Error: no such instruction: `stp x19,x20,[sp,'
gcc_arm64.S:37: Error: no such instruction: `stp x21,x22,[sp,'
gcc_arm64.S:40: Error: no such instruction: `stp x23,x24,[sp,'
gcc_arm64.S:43: Error: no such instruction: `stp x25,x26,[sp,'
gcc_arm64.S:46: Error: no such instruction: `stp x27,x28,[sp,'
gcc_arm64.S:50: Error: too many memory references for `mov'
gcc_arm64.S:51: Error: too many memory references for `mov'
gcc_arm64.S:52: Error: too many memory references for `mov'
gcc_arm64.S:54: Error: no such instruction: `blr x20'
gcc_arm64.S:55: Error: no such instruction: `blr x19'
gcc_arm64.S:57: Error: no such instruction: `ldp x27,x28,[sp,'
gcc_arm64.S:60: Error: no such instruction: `ldp x25,x26,[sp,'
gcc_arm64.S:63: Error: no such instruction: `ldp x23,x24,[sp,'
gcc_arm64.S:66: Error: no such instruction: `ldp x21,x22,[sp,'
gcc_arm64.S:69: Error: no such instruction: `ldp x19,x20,[sp,'
gcc_arm64.S:72: Error: no such instruction: `ldp x29,x30,[sp],'
root@199a722a2423:/kubeedge# 
root@199a722a2423:/kubeedge# 

# https://blog.csdn.net/sun007700/article/details/120487881
# 使用以下解决
# CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc go build -o baseline_client-arm64  baseline_client.go  

# root@199a722a2423:/kubeedge# GOARCH=arm64 GOOS="linux" CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc go build -o edgecore-arm64 -ldflags "$flags" ./edge/cmd/edgecore/
root@199a722a2423:/kubeedge# ls -lh edgecore-*
-rwxr-xr-x 1 root root 95M Apr 13 13:38 edgecore-arm64
-rw-r--r-- 1 1000 1000  45 Apr 13 13:09 edgecore-arm64.tar.gz
-rwxr-xr-x 1 root root 70M Apr 13 13:31 edgecore-x64
-rw-r--r-- 1 1000 1000 20M Apr 13 13:09 edgecore-x64.tar.gz
```