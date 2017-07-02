# domain-socket-tester

Build `domain-socket-tester-M.m.P-I.x86_64.rpm`
and   `domain-socket-tester_M.m.P-I_amd64.deb`
where "M.m.P-I" is Major.minor.Patch-Iteration.

## Usage

A program that listens on a Unix Domain Socket

### Invocation

```console
domain-socket-tester --socket-file /var/run/xyz.sock
```

## Development

### Dependencies

#### Set environment variables

```console
export GOPATH="${HOME}/go"
export PATH="${PATH}:${GOPATH}/bin:/usr/local/go/bin"
export PROJECT_DIR=${GOPATH}/src/github.com/docktermj
```

#### Download project

```console
mkdir -p ${PROJECT_DIR}
cd ${PROJECT_DIR}
git clone git@github.com:docktermj/domain-socket-tester.git
```

#### Download dependencies

```console
cd ${PROJECT_DIR}/domain-socket-tester
make dependencies
```

### Build

#### Local build

```console
cd ${PROJECT_DIR}/domain-socket-tester
make build-local
```

The results will be in the `${GOPATH}/bin` directory.

#### Docker build

```console
cd ${PROJECT_DIR}/domain-socket-tester
make build
```

The results will be in the `.../target` directory.

### Test

```console
cd ${PROJECT_DIR}/domain-socket-tester
make test-local
```

### Install

#### RPM-based

Example distributions: openSUSE, Fedora, CentOS, Mandrake

##### RPM Install

Example:

```console
sudo rpm -ivh domain-socket-tester-M.m.P-I.x86_64.rpm
```

##### RPM Update

Example: 

```console
sudo rpm -Uvh domain-socket-tester-M.m.P-I.x86_64.rpm
```

#### Debian

Example distributions: Ubuntu

##### Debian Install / Update

Example:

```console
sudo dpkg -i domain-socket-tester_M.m.P-I_amd64.deb
```

### Cleanup

```console
make clean
```
