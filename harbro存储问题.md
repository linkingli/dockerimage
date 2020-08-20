```
err.detail= filesystem mkdir /registry/docker structure needs cleaning
```
```
shenqian (Cloud)	@lijun 大佬 我问一下 平台这边 docker的存储是本地划了一块？还是用外部存储？
	16:08
zhangzhaoyang (Cloud)	无标题[root@e5103-d036-s2-vm ~]# docker info
Containers: 148
 Running: 139
 Paused: 0
 Stopped: 9
Images: 261
Server Version: 1.13.1
Storage Driver: overlay2
 Backing Filesystem: extfs
 Supports d_type: true
 Native Overlay Diff: true
Logging Driver: json-file
Cgroup Driver: cgroupfs
Plugins:
 Volume: local
 Network: bridge host macvlan null overlay
Swarm: inactive
Runtimes: docker-runc runc
Default Runtime: docker-runc
Init Binary: /usr/libexec/docker/docker-init-current
containerd version:  (expected: aa8187dbd3b7ad67d8e5e3a15115d3eef43a7ed1)
runc version: 66aedde759f33c190954815fb765eedc1d782dd9 (expected: 9df8b306d01f59d3a8029be411de015b7304dd8f)
init version: fec3683b971d9c3ef73f284f176672c44b448662 (expected: 949e6facb77383876aeff8a6944dde66b3089574)
Security Options:
 seccomp
  WARNING: You're not using the default seccomp profile
  Profile: /etc/docker/seccomp.json
 selinux
Kernel Version: 4.14.131-generic
Operating System: CentOS Linux 7 (Core)
OSType: linux
Architecture: x86_64
Number of Docker Hooks: 3
CPUs: 32
Total Memory: 125.7 GiB
Name: e5103-d036-s2-vm
ID: ID7A:FYBX:6D4X:IO5Y:OGDX:RG3E:XQDZ:RZRV:6CXL:35VZ:SS6C:DWS2
Docker Root Dir: /var/lib/docker
Debug Mode (client): false
Debug Mode (server): false
Registry: https://index.docker.io/v1/
WARNING: No swap limit support
WARNING: bridge-nf-call-iptables is disabled
WARNING: bridge-nf-call-ip6tables is disabled
Experimental: false
Insecure Registries:
 registry.cluster.local:9999
 127.0.0.0/8
Live Restore Enabled: true
Registries: docker.io (secure) 
```
