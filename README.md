# rpminfo
It is a tool to extract the package name, version, etc. from the RPM file name.

## build

```sh
$ go build ./cmd/rpminfo
```

## run

```sh
$ rpminfo openssl-1.1.1c-2.el8.x86_64
```

```sh
$ rpminfo -newline openssl-1.1.1c-2.el8.x86_64
```

```sh
$ rpminfo -field name openssl-1.1.1c-2.el8.x86_64
```