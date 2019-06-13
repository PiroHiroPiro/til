# Docker

## namespace

```
$ docker pull busybox:latest
$ docker create --name busybox-rootfs busybox:latest
$ docker export busybox-rootfs > rootfs.tar
$ mkdir rootfs
$ tar xf rootfs.tar -C rootfs
$ unshare -rmpfiun chroot rootfs sh
# ls /
# mount -t proc proc /proc
# ps a
# id -u
```

```
$ ps -o euid,cmd a | grep unshare | grep -v grep
```

## cgroups

```
$ docker run --rm -it ubuntu
# cat /sys/fs/cgroup/devices/devices.list
# mknod /dev/fb0 c 29 0
# cat /dev/fb0 > /dev/null
# ls /sys/fs/cgroup
$ ls -l /sys/fs/cgroup/devices | grep docker
$ cat /sys/fs/cgroup/devices/[ID].list
```

## overlays

```
$ docker save ubuntu:18.04 > image.tar
$ mkdir image
$ tar xf image.tar -C image
$ mkdir layer
$ tar xf image/[ID]/layer.tar -C layer
$ mkdir upper
$ mkdir lower
$ echo "upper A" > ./upper/a
$ echo "lower A" > ./lower/a
$ echo "lower B" > ./lower/b
$ mkdir work merged
$ sudo mount -t overlay overlay -olowerdir=lower,upperdir=upper,workdir=work merged
$ tree merged/
$ cat merged/a
$ cat merged/b
$ echo "Rewrite" > merged/b
$ cat upper/b
```
